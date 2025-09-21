package config

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/sqlc"
	"fmt"
	"log"

	api "clinic-vet-api/app/modules/appointment/presentation"
	authAPI "clinic-vet-api/app/modules/auth/presentation"
	customerAPI "clinic-vet-api/app/modules/customer/presentation"
	vetAPI "clinic-vet-api/app/modules/employee/presentation"
	medSessionAPI "clinic-vet-api/app/modules/medical/presentation"
	paymentAPI "clinic-vet-api/app/modules/payments/presentation"
	petAPI "clinic-vet-api/app/modules/pets/presentation"
	userAPI "clinic-vet-api/app/modules/users/presentation"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func BootstrapAPIModules(
	routerGroup *gin.RouterGroup,
	queries *sqlc.Queries,
	db *pgxpool.Pool,
	validator *validator.Validate,
	redis *redis.Client,
	jwtSecret string,
) error {
	userModule := userAPI.NewUserAPIModule(userAPI.UserAPIConfig{
		Router:        routerGroup,
		Queries:       queries,
		DataValidator: validator,
		DB:            db,
	})

	authMiddleware := middleware.NewAuthMiddleware(jwtSecret, queries)
	userModule.SetAuthMiddleware(authMiddleware)

	if err := userModule.Bootstrap(); err != nil {
		return fmt.Errorf("failed to bootstrap user API module: %w", err)
	}

	// Bootstrap Employee Module
	vetModule := vetAPI.NewEmployeeModule(&vetAPI.EmployeeAPIConfig{
		Router:         routerGroup,
		DB:             db,
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
	authAPI.SetupAuthModule(routerGroup, validator, vetRepo, customerRepo, redis, queries, jwtSecret, authMiddleware)

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

	apptModule := api.NewAppointmentAPIBuilder(&api.AppointmentAPIConfig{
		Router:         routerGroup,
		Validator:      validator,
		Queries:        queries,
		AuthMiddleware: authMiddleware,
		CustomerRepo:   customerRepo,
	})

	if err := apptModule.Build(); err != nil {
		return fmt.Errorf("failed to bootstrap appointment API module: %w", err)
	}

	log.Println("modules bootstrapped successfully")
	return nil
}
