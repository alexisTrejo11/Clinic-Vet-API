package ownerAPI

import (
	ownerUsecase "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/usecase"
	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	ownerController "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/infrastructure/api/controller"
	ownerRoutes "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/infrastructure/api/routes"
	sqlcOwnerRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/infrastructure/persistence"
	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
	appError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/application"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type OwnerAPIConfig struct {
	Router    *gin.Engine
	Validator *validator.Validate
	Queries   *sqlc.Queries
	PetRepo   petDomain.PetRepository
}

type OwnerAPIComponents struct {
	Repository ownerDomain.OwnerRepository
	UseCase    ownerUsecase.OwnerServiceFacade
	Controller *ownerController.OwnerController
}

type OwnerAPIModule struct {
	config     *OwnerAPIConfig
	components *OwnerAPIComponents
	isBuilt    bool
}

func NewOwnerAPIModule(config *OwnerAPIConfig) *OwnerAPIModule {
	return &OwnerAPIModule{
		config:  config,
		isBuilt: false,
	}
}

func (f *OwnerAPIModule) Bootstrap() error {
	if f.isBuilt {
		return nil
	}

	if err := f.validateConfig(); err != nil {
		return err
	}

	// Create repository
	ownerRepo := sqlcOwnerRepository.NewSqlcOwnerRepository(f.config.Queries, f.config.PetRepo)

	// Create use cases
	useCases := f.createUseCases(ownerRepo)

	// Create controller
	controller := ownerController.NewOwnerController(f.config.Validator, useCases)

	// Register routes
	ownerRoutes.OwnerRoutes(f.config.Router, controller)

	// Store components
	f.components = &OwnerAPIComponents{
		Repository: ownerRepo,
		UseCase:    useCases,
		Controller: controller,
	}

	f.isBuilt = true
	return nil
}

// createUseCases creates and wires all use cases
func (f *OwnerAPIModule) createUseCases(repo ownerDomain.OwnerRepository) ownerUsecase.OwnerServiceFacade {
	getUseCase := ownerUsecase.NewGetOwnerByIdUseCase(repo)
	listUseCase := ownerUsecase.NewListOwnersUseCase(repo)
	createUseCase := ownerUsecase.NewCreateOwnerUseCase(repo)
	updateUseCase := ownerUsecase.NewUpdateOwnerUseCase(repo)
	deleteUseCase := ownerUsecase.NewSoftDeleteOwnerUseCase(repo)

	return ownerUsecase.NewOwnerUseCases(getUseCase, listUseCase, createUseCase, updateUseCase, deleteUseCase)
}

// validateConfig validates the Module configuration
func (f *OwnerAPIModule) validateConfig() error {
	if f.config == nil {
		return appError.NewConflictError("INVALID_CONFIG", "configuration cannot be nil")
	}
	if f.config.Router == nil {
		return appError.NewConflictError("INVALID_CONFIG", "router cannot be nil")
	}
	if f.config.Validator == nil {
		return appError.NewConflictError("INVALID_CONFIG", "validator cannot be nil")
	}
	if f.config.Queries == nil {
		return appError.NewConflictError("INVALID_CONFIG", "queries cannot be nil")
	}
	if f.config.PetRepo == nil {
		return appError.NewConflictError("INVALID_CONFIG", "pet repository cannot be nil")
	}
	return nil
}

func (f *OwnerAPIModule) GetComponents() (*OwnerAPIComponents, error) {
	if !f.isBuilt {
		if err := f.Bootstrap(); err != nil {
			return nil, err
		}
	}
	return f.components, nil
}

func (f *OwnerAPIModule) GetRepository() (ownerDomain.OwnerRepository, error) {
	components, err := f.GetComponents()
	if err != nil {
		return nil, err
	}
	return components.Repository, nil
}

func (f *OwnerAPIModule) GetUseCase() (ownerUsecase.OwnerServiceFacade, error) {
	components, err := f.GetComponents()
	if err != nil {
		return nil, err
	}
	return components.UseCase, nil
}

func (f *OwnerAPIModule) GetController() (*ownerController.OwnerController, error) {
	components, err := f.GetComponents()
	if err != nil {
		return nil, err
	}
	return components.Controller, nil
}
