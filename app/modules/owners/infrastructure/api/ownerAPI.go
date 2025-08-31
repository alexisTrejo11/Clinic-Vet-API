package ownerAPI

import (
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/owners/application/usecase"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/owners/infrastructure/api/controller"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/owners/infrastructure/api/routes"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/owners/infrastructure/persistence"
	appError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/application"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type OwnerAPIConfig struct {
	Router    *gin.Engine
	Validator *validator.Validate
	Queries   *sqlc.Queries
	PetRepo   repository.PetRepository
}

type OwnerAPIComponents struct {
	Repository repository.OwnerRepository
	UseCase    usecase.OwnerServiceFacade
	Controller *controller.OwnerController
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
	ownerRepo := persistence.NewSqlcOwnerRepository(f.config.Queries, f.config.PetRepo)

	// Create use cases
	useCases := f.createUseCases(ownerRepo)

	// Create controller
	controller := controller.NewOwnerController(f.config.Validator, useCases)

	// Register routes
	routes.OwnerRoutes(f.config.Router, controller)

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
func (f *OwnerAPIModule) createUseCases(repo repository.OwnerRepository) usecase.OwnerServiceFacade {
	getUseCase := usecase.NewGetOwnerByIDUseCase(repo)
	listUseCase := usecase.NewListOwnersUseCase(repo)
	createUseCase := usecase.NewCreateOwnerUseCase(repo)
	updateUseCase := usecase.NewUpdateOwnerUseCase(repo)
	deleteUseCase := usecase.NewSoftDeleteOwnerUseCase(repo)

	return usecase.NewOwnerUseCases(getUseCase, listUseCase, createUseCase, updateUseCase, deleteUseCase)
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

func (f *OwnerAPIModule) GetRepository() (repository.OwnerRepository, error) {
	components, err := f.GetComponents()
	if err != nil {
		return nil, err
	}
	return components.Repository, nil
}

func (f *OwnerAPIModule) GetUseCase() (usecase.OwnerServiceFacade, error) {
	components, err := f.GetComponents()
	if err != nil {
		return nil, err
	}
	return components.UseCase, nil
}

func (f *OwnerAPIModule) GetController() (*controller.OwnerController, error) {
	components, err := f.GetComponents()
	if err != nil {
		return nil, err
	}
	return components.Controller, nil
}
