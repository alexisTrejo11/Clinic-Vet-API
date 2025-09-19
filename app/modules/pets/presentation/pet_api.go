package api

import (
	"fmt"

	"clinic-vet-api/app/core/repository"
	petRepo "clinic-vet-api/app/modules/pets/infrastructure/repository"
	"clinic-vet-api/app/modules/pets/presentation/controller"
	"clinic-vet-api/app/modules/pets/presentation/routes"
	"clinic-vet-api/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PetModuleConfig struct {
	Router       *gin.Engine
	Queries      *sqlc.Queries
	Validator    *validator.Validate
	CustomerRepo repository.CustomerRepository
}

type PetModuleComponents struct {
	Repository repository.PetRepository
	UseCases   usecase.PetUseCases
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
	return petRepo.NewSqlcPetRepository(m.config.Queries)
}

func (m *PetModule) createUseCases(repository repository.PetRepository) usecase.PetUseCases {
	return usecase.NewPetUseCases(repository, m.config.CustomerRepo)
}

func (m *PetModule) createController(useCases usecase.PetUseCases) *controller.PetController {
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
	if m.config.CustomerRepo == nil {
		return fmt.Errorf("customer repository cannot be nil")
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

func (m *PetModule) GetUseCases() (usecase.PetUseCases, error) {
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
