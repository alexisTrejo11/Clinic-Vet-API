package api

import (
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/pets/application/cqrs"
	"clinic-vet-api/app/modules/pets/presentation/controller"
	"clinic-vet-api/app/modules/pets/presentation/routes"
	"clinic-vet-api/app/modules/pets/presentation/service"
	"clinic-vet-api/sqlc"
	"fmt"

	petRepo "clinic-vet-api/app/modules/pets/infrastructure/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PetModuleConfig struct {
	Router         *gin.RouterGroup
	Queries        *sqlc.Queries
	Validator      *validator.Validate
	CustomerRepo   repository.CustomerRepository
	AuthMiddleware *middleware.AuthMiddleware
}

type PetModuleComponents struct {
	Repository            repository.PetRepository
	PetServiceBus         cqrs.PetServiceBus
	PetCtrlOperations     service.PetControllerOperations
	PetController         controller.PetController
	CustomerPetController controller.CustomerPetController
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

	bus := m.createServiceBus(repository, m.config.CustomerRepo)

	controllerOperations := m.createControllerOperations(bus, m.config.Validator)
	petController := m.createPetController(controllerOperations)
	customerPetController := m.createCustomerPetController(controllerOperations)

	m.registerRoutes(petController, customerPetController, m.config.AuthMiddleware)

	m.components = &PetModuleComponents{
		Repository:            repository,
		PetServiceBus:         bus,
		PetCtrlOperations:     *controllerOperations,
		PetController:         *petController,
		CustomerPetController: *customerPetController,
	}

	m.isBuilt = true
	return nil
}

func (m *PetModule) createRepository() repository.PetRepository {
	return petRepo.NewSqlcPetRepository(m.config.Queries)
}

func (m *PetModule) createServiceBus(petRepo repository.PetRepository, customerRepo repository.CustomerRepository) cqrs.PetServiceBus {
	bus := cqrs.NewPetServiceBus(petRepo, customerRepo)
	return bus
}

func (m *PetModule) createControllerOperations(serviceBus cqrs.PetServiceBus, validator *validator.Validate) *service.PetControllerOperations {
	return service.NewPetControllerOperations(serviceBus, validator)
}

func (m *PetModule) createPetController(controllerOperations *service.PetControllerOperations) *controller.PetController {
	return controller.NewPetController(m.config.Validator, controllerOperations)
}

func (m *PetModule) createCustomerPetController(controllerOperations *service.PetControllerOperations) *controller.CustomerPetController {
	return controller.NewCustomerPetController(m.config.Validator, controllerOperations)
}

func (m *PetModule) registerRoutes(
	adminCtrl *controller.PetController,
	customerCtrl *controller.CustomerPetController,
	authMiddle *middleware.AuthMiddleware,
) {
	routes.PetsRoutes(m.config.Router, adminCtrl, authMiddle)
	routes.CustomerPetsRoutes(m.config.Router, customerCtrl, authMiddle)
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
	if m.config.AuthMiddleware == nil {
		return fmt.Errorf("auth middleware cannot be nil")
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

func (m *PetModule) GetServiceBus() (cqrs.PetServiceBus, error) {
	components, err := m.GetComponents()
	if err != nil {
		return nil, err
	}
	return components.PetServiceBus, nil
}

func (m *PetModule) GetPetControllerOperations() (*service.PetControllerOperations, error) {
	components, err := m.GetComponents()
	if err != nil {
		return nil, err
	}
	return &components.PetCtrlOperations, nil
}

func (m *PetModule) GetPetController() (*controller.PetController, error) {
	components, err := m.GetComponents()
	if err != nil {
		return nil, err
	}
	return &components.PetController, nil
}
