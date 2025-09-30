package service

import (
	"clinic-vet-api/app/modules/auth/infrastructure/token"
	"clinic-vet-api/app/modules/core/domain/entity/customer"
	"clinic-vet-api/app/modules/core/domain/entity/employee"
	"clinic-vet-api/app/modules/core/domain/entity/notification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	service "clinic-vet-api/app/modules/notifications/application"
	commondto "clinic-vet-api/app/shared/dto"
	"context"
)

type UserAccountService struct {
	userRepository      repository.UserRepository
	customerRepository  repository.CustomerRepository
	employeeRepository  repository.EmployeeRepository
	notificationService service.NotificationService
	tokenManager        token.TokenManager
}

func NewUserAccountService(
	userRepo repository.UserRepository,
	customerRepo repository.CustomerRepository,
	employeeRepo repository.EmployeeRepository,
	tokenManager token.TokenManager,
	notificationSvc service.NotificationService,
) *UserAccountService {
	return &UserAccountService{
		userRepository:      userRepo,
		customerRepository:  customerRepo,
		employeeRepository:  employeeRepo,
		notificationService: notificationSvc,
		tokenManager:        tokenManager,
	}
}

func (s *UserAccountService) CreateCustomer(ctx context.Context, userID valueobject.UserID, personalData commondto.PersonalData) (valueobject.CustomerID, error) {
	cust := customer.NewCustomerBuilder().
		WithIsActive(true).
		WithUserID(&userID).
		WithDateOfBirth(personalData.DateOfBirth).
		WithName(personalData.Name).
		WithGender(personalData.Gender).
		Build()

	if err := s.customerRepository.Save(ctx, cust); err != nil {
		return valueobject.CustomerID{}, err
	}

	return cust.ID(), nil
}

func (s *UserAccountService) AttachEmployeeToUser(ctx context.Context, userID valueobject.UserID, employee employee.Employee) error {
	if err := employee.AssignUser(ctx, userID); err != nil {
		return err
	}

	if err := s.employeeRepository.Save(ctx, &employee); err != nil {
		return err
	}

	return nil
}

func (s *UserAccountService) SendActivationEmail(ctx context.Context, userID, email, name string) error {
	token, err := s.tokenManager.GenerateToken(token.ActivationToken, token.TokenConfig{
		UserID: userID,
	})
	if err != nil {
		return err
	}

	notif := notification.NewActivationEmail(userID, email, name, token)
	return s.notificationService.Send(ctx, notif)
}

func (s *UserAccountService) SendWelcomeEmail(ctx context.Context, email, name string) error {
	notif := notification.NewWelcomeEmail(email, name)
	return s.notificationService.Send(ctx, notif)
}
