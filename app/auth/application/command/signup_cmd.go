package authCmd

import (
	"context"
	"errors"
	"time"

	session "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/domain"
	sessionRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/domain/repositories"
	ownerRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/repository"
	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	userApplication "github.com/alexisTrejo11/Clinic-Vet-API/app/users/application"
	userCommand "github.com/alexisTrejo11/Clinic-Vet-API/app/users/application/command"
	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
	userRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
	vetRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/repositories"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
)

type SignupCommand struct {
	// User credentials
	Email       *string `json:"email"`
	PhoneNumber *string `json:"phone_number"`
	Password    string  `json:"password"`

	// Personal details
	FirstName      string        `json:"first_name"`
	LastName       string        `json:"last_name"`
	Gender         user.Gender   `json:"gender"`
	DateOfBirth    time.Time     `json:"date_of_birth"`
	Location       string        `json:"location"`
	Address        string        `json:"address"`
	Role           user.UserRole `json:"role"`
	ProfilePicture string        `json:"profile_picture"`
	Bio            string        `json:"bio"`

	// Veterinarian details
	LicenseNumber   *string
	Specialty       *vetDomain.VetSpecialty
	YearsExperience *int

	CTX context.Context `json:"-"`

	// Metadata
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	Source    string `json:"source"`
}

type SignupHandler interface {
	Handle(cmd *SignupCommand) shared.CommandResult
}

type SignupCommandHandler struct {
	userRepo    userRepository.UserRepository
	dispatcher  *userApplication.CommandDispatcher
	sessionRepo sessionRepo.SessionRepository
	ownerRepo   ownerRepository.OwnerRepository
	vetRepo     vetRepo.VeterinarianRepository
}

func NewSignupCommandHandler(
	userRepo userRepository.UserRepository,
	dispatcher *userApplication.CommandDispatcher,
	ownerRepo ownerRepository.OwnerRepository,
	sessionRepo sessionRepo.SessionRepository,
	vetRepo vetRepo.VeterinarianRepository,

) *SignupCommandHandler {
	return &SignupCommandHandler{
		userRepo:    userRepo,
		dispatcher:  dispatcher,
		sessionRepo: sessionRepo,
		ownerRepo:   ownerRepo,
		vetRepo:     vetRepo,
	}
}

func (h *SignupCommandHandler) Handle(cmd *SignupCommand) shared.CommandResult {
	if err := h.validateCredentials(cmd); err != nil {
		return shared.FailureResult("an conflict ocurred", err)
	}

	userId, err := h.createUser(cmd)
	if err != nil {
		return shared.FailureResult("an error ocurred while creating user", err)
	}

	// TODO: Produce Event

	return shared.SuccesResult(userId.String(), "User created successfully")

}

func (h *SignupCommandHandler) validateCredentials(cmd *SignupCommand) error {
	if cmd.Email == nil && cmd.PhoneNumber == nil {
		return errors.New("email or phone number must be provided")
	}

	if err := h.validateUniqueCredentials(cmd); err != nil {
		return err
	}

	if err := user.ValidatePassword(cmd.Password); err != nil {
		return err
	}

	return nil
}

func (h *SignupCommandHandler) validateUniqueCredentials(cmd *SignupCommand) error {
	if cmd.Email != nil {
		exists, err := h.userRepo.ExistsByEmail(cmd.CTX, *cmd.Email)
		if err != nil {
			return err
		}
		if exists {
			return errors.New("email already exists")
		}
	}

	if cmd.PhoneNumber != nil {
		exists, err := h.userRepo.ExistsByPhone(cmd.CTX, *cmd.PhoneNumber)
		if err != nil {
			return err
		}
		if exists {
			return errors.New("phone number already exists")
		}
	}

	return nil
}

func (h *SignupCommandHandler) createUser(cmd *SignupCommand) (user.UserId, error) {
	userCommand := ToCreteUserCommand(*cmd)

	result := h.dispatcher.Dispatch(userCommand)
	if !result.IsSuccess {
		return user.UserId{}, errors.New(result.Message)
	}

	userId, err := user.NewUserId(result.Id)
	if err != nil {
		return user.UserId{}, err
	}

	if err := h.createOwner(*cmd, userId); err != nil {
		return user.UserId{}, err
	}

	if err := h.createVet(*cmd, userId); err != nil {
		return user.UserId{}, err
	}

	return userId, nil
}

func (h *SignupCommandHandler) createSession(userId user.UserId, cmd SignupCommand) error {
	session := session.Session{
		UserId: userId.String(),
	}

	if err := h.sessionRepo.Create(cmd.CTX, &session); err != nil {
		return err
	}

	return nil
}

func ToCreteUserCommand(cmd SignupCommand) userCommand.CreateUserCommand {
	command := userCommand.CreateUserCommand{
		Password: cmd.Password,
		Address:  cmd.Address,
		Role:     cmd.Role.String(),
		Status:   user.UserStatusPending,
	}

	if cmd.Email != nil {
		command.Email = *cmd.Email
	}
	if cmd.PhoneNumber != nil {
		command.PhoneNumber = *cmd.PhoneNumber
	}

	return command
}

func (h *SignupCommandHandler) createOwner(cmd SignupCommand, userId user.UserId) error {
	if cmd.Role != user.UserRoleOwner {
		return nil
	}

	name, err := user.NewPersonName(cmd.FirstName, cmd.LastName)
	if err != nil {
		return err
	}

	newUserOwner := &ownerDomain.Owner{
		Photo:       cmd.ProfilePicture,
		FullName:    name,
		Gender:      cmd.Gender,
		DateOfBirth: cmd.DateOfBirth,
		UserId:      func(v int) *int { return &v }(userId.GetValue()),
		IsActive:    true,
		Address:     &cmd.Address,
	}

	if err := h.ownerRepo.Save(cmd.CTX, newUserOwner); err != nil {
		return err
	}

	return nil
}

func (h *SignupCommandHandler) createVet(cmd SignupCommand, userId user.UserId) error {
	if cmd.Role != user.UserRoleVeterinarian {
		return nil
	}

	name, err := user.NewPersonName(cmd.FirstName, cmd.LastName)
	if err != nil {
		return err
	}

	vet := vetDomain.Veterinarian{
		Name:     name,
		Photo:    cmd.ProfilePicture,
		IsActive: true,
		UserID:   func(v int) *int { return &v }(userId.GetValue()),
	}

	if cmd.LicenseNumber != nil {
		vet.LicenseNumber = *cmd.LicenseNumber
	}
	if cmd.Specialty != nil {
		vet.Specialty = *cmd.Specialty
	}
	if cmd.YearsExperience != nil {
		vet.YearsExperience = *cmd.YearsExperience
	}

	if err := h.vetRepo.Save(cmd.CTX, &vet); err != nil {
		return err
	}
	return nil
}
