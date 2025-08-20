package config

import (
	"fmt"
	"log"

	medHistoryAPI "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/infrastructure/api"
	ownerAPI "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/infrastructure/api"
	petAPI "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/infrastructure/api"
	userAPI "github.com/alexisTrejo11/Clinic-Vet-API/app/users/infrastructure/api"
	vetAPI "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/infrastructure/api"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BootstrapAPIModules(
	router *gin.Engine,
	queries *sqlc.Queries,
	validator *validator.Validate,
) error {

	// Bootstrap Vet Module
	vetModule := vetAPI.NewVeterinarianModule(&vetAPI.VeterinarianAPIConfig{
		Router:        router,
		Queries:       queries,
		DataValidator: validator,
	})

	if err := vetModule.Bootstrap(); err != nil {
		return fmt.Errorf("failed to bootstrap vet module: %w", err)
	}

	// Bootstrap Owner Module
	ownerModule := ownerAPI.NewOwnerAPIModule(&ownerAPI.OwnerAPIConfig{
		Router:    router,
		Queries:   queries,
		Validator: validator,
	})

	if err := ownerModule.Bootstrap(); err != nil {
		return fmt.Errorf("failed to bootstrap owner module: %w", err)
	}

	ownerRepo, err := ownerModule.GetRepository()
	if err != nil {
		return fmt.Errorf("failed to get owner repository: %w", err)
	}

	// Bootstrap Pet Module
	petModule := petAPI.NewPetModule(&petAPI.PetModuleConfig{
		Router:    router,
		Queries:   queries,
		Validator: validator,
		OwnerRepo: ownerRepo,
	})

	if err := petModule.Bootstrap(); err != nil {
		return fmt.Errorf("failed to bootstrap pet module: %w", err)
	}

	// Get pet repository for medical history module
	petRepo, err := petModule.GetRepository()
	if err != nil {
		return fmt.Errorf("failed to get pet repository: %w", err)
	}

	vetRepo, err := vetModule.GetRepository()
	if err != nil {
		return fmt.Errorf("failed to get vet repository: %w", err)
	}

	// Bootstrap Medical History Module
	medHistoryModule := medHistoryAPI.NewMedicalHistoryModule(&medHistoryAPI.MedicalHistoryModuleConfig{
		Router:    router,
		Queries:   queries,
		Validator: validator,
		OwnerRepo: ownerRepo,
		VetRepo:   vetRepo,
		PetRepo:   petRepo,
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
