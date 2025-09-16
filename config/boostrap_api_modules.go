package config

import (
	"fmt"
	"log"

	customerAPI "clinic-vet-api/app/modules/customer/presentation"
	vetAPI "clinic-vet-api/app/modules/employee/presentation"
	medHistoryAPI "clinic-vet-api/app/modules/medical/presentation"
	petRepo "clinic-vet-api/app/modules/pets/infrastructure/repository"
	petAPI "clinic-vet-api/app/modules/pets/presentation"
	userAPI "clinic-vet-api/app/modules/users/presentation"
	"clinic-vet-api/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
)

func BootstrapAPIModules(
	router *gin.Engine,
	queries *sqlc.Queries,
	db *pgxpool.Pool,
	validator *validator.Validate,
) error {
	petRepository := petRepo.NewSqlcPetRepository(queries)

	// Bootstrap Employee Module
	vetModule := vetAPI.NewEmployeeModule(&vetAPI.EmployeeAPIConfig{
		Router:        router,
		DB:            db,
		Queries:       queries,
		DataValidator: validator,
	})

	if err := vetModule.Bootstrap(); err != nil {
		return fmt.Errorf("failed to bootstrap vet module: %w", err)
	}

	// Bootstrap Customer Module
	customerModule := customerAPI.NewCustomerAPIModule(&customerAPI.CustomerAPIConfig{
		Router:    router,
		Queries:   queries,
		Validator: validator,
		PetRepo:   petRepository,
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
		Router:       router,
		Queries:      queries,
		Validator:    validator,
		CustomerRepo: customerRepo,
	})

	if err := petModule.Bootstrap(); err != nil {
		return fmt.Errorf("failed to bootstrap pet module: %w", err)
	}

	vetRepo, err := vetModule.GetRepository()
	if err != nil {
		return fmt.Errorf("failed to get vet repository: %w", err)
	}

	// Bootstrap Medical History Module
	medHistoryModule := medHistoryAPI.NewMedicalHistoryModule(&medHistoryAPI.MedicalHistoryModuleConfig{
		Router:       router,
		Queries:      queries,
		Validator:    validator,
		CustomerRepo: &customerRepo,
		EmployeeRepo: &vetRepo,
		PetRepo:      &petRepository,
	})

	if err := medHistoryModule.Bootstrap(); err != nil {
		return fmt.Errorf("failed to bootstrap medical history module: %w", err)
	}

	if err := userAPI.NewUserAPIModule(userAPI.UserAPIConfig{
		Router:        router,
		Queries:       queries,
		DataValidator: validator,
	}).Bootstrap(); err != nil {
		return fmt.Errorf("failed to bootstrap user API module: %w", err)
	}

	log.Println("modules bootstrapped successfully")
	return nil
}
