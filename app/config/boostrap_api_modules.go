package config

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/sqlc"
	"fmt"
	"log"

	authAPI "clinic-vet-api/app/modules/account/auth"
	userAPI "clinic-vet-api/app/modules/account/user/presentation"
	apptApi "clinic-vet-api/app/modules/appointment/presentation"
	"clinic-vet-api/app/modules/core/service"
	customerAPI "clinic-vet-api/app/modules/customer/presentation"
	vetAPI "clinic-vet-api/app/modules/employee/presentation"
	dewormApi "clinic-vet-api/app/modules/medical/deworm/presentation"
	medSessionAPI "clinic-vet-api/app/modules/medical/session/presentation"
	api "clinic-vet-api/app/modules/medical/vaccination/presentation"
	paymentAPI "clinic-vet-api/app/modules/payment/presentation"
	petAPI "clinic-vet-api/app/modules/pet/presentation"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
)

func BootstrapAPIModules(
	routerGroup *gin.RouterGroup,
	queries *sqlc.Queries,
	notificationService service.NotificationService,
	validator *validator.Validate,
	redis *redis.Client,
	jwtSecret string,
) error {
	userModule := userAPI.NewUserAPIModule(userAPI.UserAPIConfig{
		Router:        routerGroup,
		Queries:       queries,
		DataValidator: validator,
	})

	authMiddleware := middleware.NewAuthMiddleware(jwtSecret, queries)
	userModule.SetAuthMiddleware(authMiddleware)

	if err := userModule.Bootstrap(); err != nil {
		return fmt.Errorf("failed to bootstrap user API module: %w", err)
	}

	// Bootstrap Employee Module
	vetModule := vetAPI.NewEmployeeModule(&vetAPI.EmployeeAPIConfig{
		Router:         routerGroup,
		Queries:        queries,
		DataValidator:  validator,
		AuthMiddleware: authMiddleware,
	})

	if err := vetModule.Bootstrap(); err != nil {
		return fmt.Errorf("failed to bootstrap vet module: %w", err)
	}

	// Bootstrap Customer Module
	customerModule := customerAPI.NewCustomerAPIModule(&customerAPI.CustomerAPIConfig{
		Router:         routerGroup,
		Queries:        queries,
		Validator:      validator,
		AuthMiddleware: authMiddleware,
	})

	if err := customerModule.Bootstrap(); err != nil {
		return fmt.Errorf("failed to bootstrap customer module: %w", err)
	}

	customerRepo, err := customerModule.GetRepository()
	if err != nil {
		return fmt.Errorf("failed to get customer repository: %w", err)
	}

	// Bootstrap Pet Module
	petModule := petAPI.NewPetModule(&petAPI.PetModuleConfig{
		Router:         routerGroup,
		Queries:        queries,
		Validator:      validator,
		CustomerRepo:   customerRepo,
		AuthMiddleware: authMiddleware,
	})

	if err := petModule.Bootstrap(); err != nil {
		return fmt.Errorf("failed to bootstrap pet module: %w", err)
	}

	vetRepo, err := vetModule.GetRepository()
	if err != nil {
		return fmt.Errorf("failed to get vet repository: %w", err)
	}

	petRepository, err := petModule.GetRepository()
	if err != nil {
		return fmt.Errorf("failed to get pet repository: %w", err)
	}

	// Bootstrap Medical Session Module
	medSessionModule := medSessionAPI.NewMedicalSessionModule(&medSessionAPI.MedicalSessionModuleConfig{
		Router:         routerGroup,
		Queries:        queries,
		Validator:      validator,
		CustomerRepo:   &customerRepo,
		EmployeeRepo:   &vetRepo,
		PetRepo:        &petRepository,
		AuthMiddleware: authMiddleware,
	})

	if err := medSessionModule.Bootstrap(); err != nil {
		return fmt.Errorf("failed to bootstrap medical history module: %w", err)
	}

	customerRepo, err = customerModule.GetRepository()
	if err != nil {
		return fmt.Errorf("failed to get customer repository: %w", err)
	}

	vetRepo, err = vetModule.GetRepository()
	if err != nil {
		return fmt.Errorf("failed to get vet repository: %w", err)
	}

	// Bootstrap Auth Module
	authModule := authAPI.NewAuthAPI(&authAPI.AuthModuleConfig{
		RouterGroup:         routerGroup,
		Validator:           validator,
		Queries:             queries,
		AuthMiddleware:      authMiddleware,
		EmployeeRepo:        vetRepo,
		CustomerRepo:        customerRepo,
		Client:              redis,
		SecretKey:           jwtSecret,
		NotificationService: notificationService,
	})

	if err := authModule.Boostrap(); err != nil {
		return fmt.Errorf("failed to bootstrap auth API module: %w", err)
	}

	// Bootstrap Payment Module
	paymentModule := paymentAPI.NewPaymentAPIBuilder(&paymentAPI.PaymentAPIConfig{
		Router:         routerGroup,
		Validator:      validator,
		Queries:        queries,
		AuthMiddleware: authMiddleware,
	})

	if err := paymentModule.Build(); err != nil {
		return fmt.Errorf("failed to bootstrap payment API module: %w", err)
	}

	employeeRepo, err := vetModule.GetRepository()
	if err != nil {
		return fmt.Errorf("failed to get employee repository: %w", err)
	}

	apptModule := apptApi.NewAppointmentAPIBuilder(&apptApi.AppointmentAPIConfig{
		Router:         routerGroup,
		Validator:      validator,
		Queries:        queries,
		AuthMiddleware: authMiddleware,
		CustomerRepo:   customerRepo,
		EmployeeRepo:   employeeRepo,
	})

	if err := apptModule.Build(); err != nil {
		return fmt.Errorf("failed to bootstrap appointment API module: %w", err)
	}

	dewormModule := dewormApi.NewDewormAPIModule(&dewormApi.DewormAPIConfig{
		RouterGroup:    routerGroup,
		Queries:        queries,
		Validator:      validator,
		AuthMiddleware: authMiddleware,
		PetRepo:        petRepository,
		EmployeeRepo:   vetRepo,
		CustomerRepo:   customerRepo,
	})

	if err := dewormModule.Bootstrap(); err != nil {
		return fmt.Errorf("failed to bootstrap deworm API module: %w", err)
	}

	vaccinationModule := api.NewVaccinationAPIModule(&api.VaccinationConfig{
		Router:         routerGroup,
		Queries:        queries,
		Validator:      validator,
		AuthMiddleware: authMiddleware,
		PetRepo:        petRepository,
		EmployeeRepo:   vetRepo,
		CustomerRepo:   customerRepo,
	})

	if err := vaccinationModule.Bootstrap(); err != nil {
		return fmt.Errorf("failed to bootstrap vaccination API module: %w", err)
	}

	log.Println("modules bootstrapped successfully")
	return nil
}
