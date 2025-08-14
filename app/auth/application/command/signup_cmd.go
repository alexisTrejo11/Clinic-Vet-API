package authCmd

import (
	"context"
	"errors"
	"strconv"
	"time"

	session "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/domain"
	sessionRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/domain/repositories"
	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/valueObjects"
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
	UserId      int     `json:"user_id"`

	// Personal details
	FirstName      string              `json:"first_name"`
	LastName       string              `json:"last_name"`
	Gender         valueObjects.Gender `json:"gender"`
	DateOfBirth    time.Time           `json:"date_of_birth"`
	Location       string              `json:"location"`
	Address        string              `json:"address"`
	Role           user.UserRole       `json:"role"`
	ProfilePicture string              `json:"profile_picture"`
	Bio            string              `json:"bio"`

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
	ownerRepo   ownerDomain.OwnerRepository
	vetRepo     vetRepo.VeterinarianRepository
}

func NewSignupCommandHandler(
	userRepo userRepository.UserRepository,
	dispatcher *userApplication.CommandDispatcher,
	ownerRepo ownerDomain.OwnerRepository,
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

	return shared.SuccesResult(strconv.Itoa(userId), "User created successfully")

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

func (h *SignupCommandHandler) createUser(cmd *SignupCommand) (int, error) {
	userCommand := ToCreateUserCommand(*cmd)

	result := h.dispatcher.Dispatch(userCommand)
	if !result.IsSuccess {
		return 0, errors.New(result.Message)
	}

	if err := h.createOwner(*cmd, cmd.UserId); err != nil {
		return 0, err
	}

	if err := h.createVet(*cmd); err != nil {
		return 0, err
	}

	return cmd.UserId, nil
}

func (h *SignupCommandHandler) CreateSession(cmd SignupCommand) error {
	session := session.Session{
		UserId: strconv.Itoa(cmd.UserId),
	}

	if err := h.sessionRepo.Create(cmd.CTX, &session); err != nil {
		return err
	}

	return nil
}

func ToCreateUserCommand(cmd SignupCommand) userCommand.CreateUserCommand {
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

func (h *SignupCommandHandler) createOwner(cmd SignupCommand, userId int) error {
	if cmd.Role != user.UserRoleOwner {
		return nil
	}

	name, err := valueObjects.NewPersonName(cmd.FirstName, cmd.LastName)
	if err != nil {
		return err
	}

	newUserOwner := &ownerDomain.Owner{}
	newUserOwner.SetFullName(name)
	newUserOwner.SetUserId(userId)
	newUserOwner.SetPhoto(cmd.ProfilePicture)
	newUserOwner.SetGender(cmd.Gender)
	newUserOwner.SetDateOfBirth(cmd.DateOfBirth)
	newUserOwner.SetAddress(cmd.Address)

	if err := h.ownerRepo.Save(cmd.CTX, newUserOwner); err != nil {
		return err
	}

	return nil
}

func (h *SignupCommandHandler) createVet(cmd SignupCommand) error {
	name, err := valueObjects.NewPersonName(cmd.FirstName, cmd.LastName)
	if err != nil {
		return err
	}

	builder := vetDomain.NewVeterinarianBuilder().
		WithName(name).
		WithPhoto(cmd.ProfilePicture).
		WithIsActive(true).
		WithUserID(func(v int) *int { return &v }(cmd.UserId))

	if cmd.LicenseNumber != nil {
		builder.WithLicenseNumber(*cmd.LicenseNumber)
	}
	if cmd.Specialty != nil {
		builder.WithSpecialty(*cmd.Specialty)
	}
	if cmd.YearsExperience != nil {
		builder.WithYearsExperience(*cmd.YearsExperience)
	}

	vet := builder.Build()

	if err := h.vetRepo.Save(cmd.CTX, vet); err != nil {
		return err
	}

	return nil
}
