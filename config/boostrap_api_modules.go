package config

import (
	"fmt"
	"log"

	medHistoryAPI "clinic-vet-api/app/modules/medical/infrastructure/api"
	ownerAPI "clinic-vet-api/app/modules/owners/infrastructure/api"
	petAPI "clinic-vet-api/app/modules/pets/infrastructure/api"
	sqlcPetRepository "clinic-vet-api/app/modules/pets/infrastructure/persistence"
	userAPI "clinic-vet-api/app/modules/users/infrastructure/api"
	vetAPI "clinic-vet-api/app/modules/veterinarians/infrastructure/api"
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
	petRepository := sqlcPetRepository.NewSqlcPetRepository(queries)

	// Bootstrap Vet Module
	vetModule := vetAPI.NewVeterinarianModule(&vetAPI.VeterinarianAPIConfig{
		Router:        router,
		DB:            db,
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
		PetRepo:   petRepository,
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

	vetRepo, err := vetModule.GetRepository()
	if err != nil {
		return fmt.Errorf("failed to get vet repository: %w", err)
	}

	// Bootstrap Medical History Module
	medHistoryModule := medHistoryAPI.NewMedicalHistoryModule(&medHistoryAPI.MedicalHistoryModuleConfig{
		Router:    router,
		Queries:   queries,
		Validator: validator,
		OwnerRepo: &ownerRepo,
		VetRepo:   &vetRepo,
		PetRepo:   &petRepository,
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
