package service

import (
	"clinic-vet-api/app/modules/core/domain/entity/customer"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	commondto "clinic-vet-api/app/shared/dto"
	"context"
	"fmt"
)

type UserAccountService interface {
	CreateCustomer(ctx context.Context, userID valueobject.UserID, personalData commondto.PersonalData) (valueobject.CustomerID, error)
	AttachCustomerToUser(ctx context.Context, userID valueobject.UserID, customerID valueobject.CustomerID) error
	AttachEmployeeToUser(ctx context.Context, userID valueobject.UserID, employeeID valueobject.EmployeeID) error
	SendActivationEmail(ctx context.Context, email valueobject.Email, name string) error
	SendWelcomeEmail(ctx context.Context, email valueobject.Email, name string) error
}

type userAccountService struct {
	userRepository     repository.UserRepository
	profileRepository  repository.ProfileRepository
	customerRepository repository.CustomerRepository
	employeeRepository repository.EmployeeRepository
	emailService       EmailService
}

func NewUserAccountService(
	userRepo repository.UserRepository,
	profileRepo repository.ProfileRepository,
	customerRepo repository.CustomerRepository,
	employeeRepo repository.EmployeeRepository,
	emailSvc EmailService,
) UserAccountService {
	return &userAccountService{
		userRepository:     userRepo,
		profileRepository:  profileRepo,
		customerRepository: customerRepo,
		employeeRepository: employeeRepo,
		emailService:       emailSvc,
	}
}

func (s *userAccountService) CreateCustomer(ctx context.Context, userID valueobject.UserID, personalData commondto.PersonalData) (valueobject.CustomerID, error) {
	cust, err := customer.CreateCustomer(
		ctx,
		customer.WithIsActive(true),
		customer.WithUserID(&userID),
		customer.WithDateOfBirth(personalData.DateOfBirth),
		customer.WithFullName(personalData.Name),
		customer.WithGender(personalData.Gender),
	)
	if err != nil {
		return valueobject.CustomerID{}, err
	}

	if err := s.customerRepository.Save(ctx, cust); err != nil {
		return valueobject.CustomerID{}, err
	}

	return cust.ID(), nil
}

func (s *userAccountService) AttachCustomerToUser(ctx context.Context, userID valueobject.UserID, customerID valueobject.CustomerID) error {
	user, err := s.userRepository.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	customer, err := s.customerRepository.FindByID(ctx, customerID)
	if err != nil {
		return err
	}

	if err := user.AssignCustomer(customer.ID()); err != nil {
		return err
	}

	if err := customer.AssignUser(user.ID()); err != nil {
		return err
	}

	if err := s.userRepository.Save(ctx, &user); err != nil {
		return err
	}

	if err := s.customerRepository.Save(ctx, &customer); err != nil {
		return err
	}

	return nil
}

func (s *userAccountService) AttachEmployeeToUser(ctx context.Context, userID valueobject.UserID, employeeID valueobject.EmployeeID) error {
	user, err := s.userRepository.FindByID(ctx, userID)
	if err != nil {
		return err
	}
	employee, err := s.employeeRepository.FindByID(ctx, employeeID)
	if err != nil {
		return err
	}

	if err := user.AssignEmployee(employee.ID()); err != nil {
		return err
	}

	if err := employee.AssignUser(ctx, user.ID()); err != nil {
		return err
	}

	if err := s.userRepository.Save(ctx, &user); err != nil {
		return err
	}

	if err := s.employeeRepository.Save(ctx, &employee); err != nil {
		return err
	}

	return nil
}

func (s *userAccountService) SendActivationEmail(ctx context.Context, email valueobject.Email, name string) error {
	return s.emailService.SendEmail(email.String(), "Activate Your Account", fmt.Sprintf("Hello %s,\n\nPlease activate your account by clicking the link below.\n\nBest regards,\nClinic Vet Team", name))
}

func (s *userAccountService) SendWelcomeEmail(ctx context.Context, email valueobject.Email, name string) error {
	return s.emailService.SendEmail(email.String(), "Welcome to Clinic Vet", fmt.Sprintf("Hello %s,\n\nWelcome to Clinic Vet! We're excited to have you on board.\n\nBest regards,\nClinic Vet Team", name))
}
