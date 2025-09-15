package api

import (
	"fmt"

	repository "clinic-vet-api/app/core/repositories"
	"clinic-vet-api/app/modules/pets/application/usecase"
	"clinic-vet-api/app/modules/pets/infrastructure/api/controller"
	"clinic-vet-api/app/modules/pets/infrastructure/api/routes"
	"clinic-vet-api/app/modules/pets/infrastructure/persistence"
	"clinic-vet-api/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PetModuleConfig struct {
	Router    *gin.Engine
	Queries   *sqlc.Queries
	Validator *validator.Validate
	OwnerRepo repository.OwnerRepository
}

type PetModuleComponents struct {
	Repository repository.PetRepository
	UseCases   usecase.PetUseCasesFacade
	Controller *controller.PetController
}

type PetModule struct {
	config     *PetModuleConfig
	components *PetModuleComponents
	isBuilt    bool
}

func NewPetModule(config *PetModuleConfig) *PetModule {
	return &PetModule{
		config:  config,
		isBuilt: false,
	}
}

func (m *PetModule) Bootstrap() error {
	if m.isBuilt {
		return nil
	}

	if err := m.validateConfig(); err != nil {
		return err
	}

	repository := m.createRepository()

	useCases := m.createUseCases(repository)

	controller := m.createController(useCases)

	m.registerRoutes(controller)

	m.components = &PetModuleComponents{
		Repository: repository,
		UseCases:   useCases,
		Controller: controller,
	}

	m.isBuilt = true
	return nil
}

func (m *PetModule) createRepository() repository.PetRepository {
	return persistence.NewSqlcPetRepository(m.config.Queries)
}

func (m *PetModule) createUseCases(repository repository.PetRepository) usecase.PetUseCasesFacade {
	return usecase.NewPetUseCasesFacade(repository, m.config.OwnerRepo)
}

func (m *PetModule) createController(useCases usecase.PetUseCasesFacade) *controller.PetController {
	return controller.NewPetController(m.config.Validator, useCases)
}

func (m *PetModule) registerRoutes(controller *controller.PetController) {
	routes.PetsRoutes(m.config.Router, controller)
}

func (m *PetModule) validateConfig() error {
	if m.config == nil {
		return fmt.Errorf("pet module configuration cannot be nil")
	}
	if m.config.Router == nil {
		return fmt.Errorf("router cannot be nil")
	}
	if m.config.Queries == nil {
		return fmt.Errorf("queries cannot be nil")
	}
	if m.config.Validator == nil {
		return fmt.Errorf("validator cannot be nil")
	}
	if m.config.OwnerRepo == nil {
		return fmt.Errorf("owner repository cannot be nil")
	}
	return nil
}

func (m *PetModule) GetComponents() (*PetModuleComponents, error) {
	if !m.isBuilt {
		if err := m.Bootstrap(); err != nil {
			return nil, err
		}
	}
	return m.components, nil
}

func (m *PetModule) GetRepository() (repository.PetRepository, error) {
	components, err := m.GetComponents()
	if err != nil {
		return nil, err
	}
	return components.Repository, nil
}

func (m *PetModule) GetUseCases() (usecase.PetUseCasesFacade, error) {
	components, err := m.GetComponents()
	if err != nil {
		return nil, err
	}
	return components.UseCases, nil
}

func (m *PetModule) GetController() (*controller.PetController, error) {
	components, err := m.GetComponents()
	if err != nil {
		return nil, err
	}
	return components.Controller, nil
}
