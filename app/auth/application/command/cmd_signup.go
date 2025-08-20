package authCmd

import (
	"context"
	"errors"
	"strconv"
	"sync"
	"time"

	session "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/domain"
	notificationService "github.com/alexisTrejo11/Clinic-Vet-API/app/notifications/application"
	notificationDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/notifications/domain"
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

type signupHandler struct {
	userRepo            userRepository.UserRepository
	dispatcher          *userApplication.CommandDispatcher
	sessionRepo         session.SessionRepository
	ownerRepo           ownerDomain.OwnerRepository
	notificationService notificationService.NotificationService
	vetRepo             vetRepo.VeterinarianRepository
}

func NewSignupCommandHandler(
	userRepo userRepository.UserRepository,
	dispatcher *userApplication.CommandDispatcher,
	ownerRepo ownerDomain.OwnerRepository,
	sessionRepo session.SessionRepository,
	notificationService notificationService.NotificationService,
	vetRepo vetRepo.VeterinarianRepository,

) *signupHandler {
	return &signupHandler{
		userRepo:            userRepo,
		dispatcher:          dispatcher,
		sessionRepo:         sessionRepo,
		ownerRepo:           ownerRepo,
		vetRepo:             vetRepo,
		notificationService: notificationService,
	}
}

func (h *signupHandler) Handle(cmd SignupCommand) AuthCommandResult {
	if err := h.validateCredentials(&cmd); err != nil {
		return AuthCommandResult{CommandResult: shared.FailureResult("an conflict ocurred", err)}
	}

	userId, err := h.createUser(&cmd)
	if err != nil {
		return AuthCommandResult{CommandResult: shared.FailureResult("an error ocurred while creating user", err)}
	}

	go h.SendActivationEmail(cmd.CTX, strconv.Itoa(userId), *cmd.Email)

	return AuthCommandResult{CommandResult: shared.SuccessResult(strconv.Itoa(userId), "User created successfully")}

}

func (h *signupHandler) validateCredentials(cmd *SignupCommand) error {
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

func (h *signupHandler) validateUniqueCredentials(cmd *SignupCommand) error {
	var wg sync.WaitGroup
	errChan := make(chan error, 2)

	if cmd.Email != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			exists, err := h.userRepo.ExistsByEmail(cmd.CTX, *cmd.Email)
			if err != nil {
				errChan <- err
				return
			}
			if exists {
				errChan <- errors.New("email already exists")
			}
		}()
	}

	if cmd.PhoneNumber != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			exists, err := h.userRepo.ExistsByPhone(cmd.CTX, *cmd.PhoneNumber)
			if err != nil {
				errChan <- err
				return
			}
			if exists {
				errChan <- errors.New("phone number already exists")
			}
		}()
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *signupHandler) createUser(cmd *SignupCommand) (int, error) {
	var wg sync.WaitGroup
	errChan := make(chan error, 3)

	wg.Add(1)
	go func() {
		defer wg.Done()
		userCreateCommand := toCreateUserCommand(*cmd)
		result := h.dispatcher.Dispatch(userCreateCommand)
		if !result.IsSuccess {
			errChan <- errors.New(result.Message)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := h.createOwner(*cmd, cmd.UserId); err != nil {
			errChan <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := h.createVet(*cmd); err != nil {
			errChan <- err
		}
	}()

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return 0, err
		}
	}

	return cmd.UserId, nil
}

func (h *signupHandler) createOwner(cmd SignupCommand, userId int) error {
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

func (h *signupHandler) createVet(cmd SignupCommand) error {
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

func toCreateUserCommand(cmd SignupCommand) userCommand.CreateUserCommand {
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

func (h *signupHandler) SendActivationEmail(ctx context.Context, userId, email string) error {
	notification := notificationDomain.NewActivateAccountNotification(userId, email)
	return h.notificationService.SendNotification(ctx, &notification)
}
