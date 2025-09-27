package medSessionAPI

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/modules/medical/application/query"
	"clinic-vet-api/app/modules/medical/infrastructure/bus"
	repositoryimpl "clinic-vet-api/app/modules/medical/infrastructure/persistence/repository"
	controller "clinic-vet-api/app/modules/medical/presentation/controller/med_session"
	"clinic-vet-api/app/modules/medical/presentation/routes"
	"clinic-vet-api/sqlc"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MedicalSessionModuleConfig struct {
	Router         *gin.RouterGroup
	Queries        *sqlc.Queries
	Validator      *validator.Validate
	CustomerRepo   *repository.CustomerRepository
	EmployeeRepo   *repository.EmployeeRepository
	PetRepo        *repository.PetRepository
	AuthMiddleware *middleware.AuthMiddleware
}

type MedicalSessionControllers struct {
	AdminController    *controller.AdminMedicalSessionController
	CustomerController *controller.CustomerMedicalSessionController
	EmployeeController *controller.EmployeeMedicalSessionController
}

type MedicalSessionModuleComponents struct {
	Repository repository.MedicalSessionRepository
	Bus        *bus.MedicalSessionBus
	Controller *MedicalSessionControllers
}

type MedicalSessionModule struct {
	config     *MedicalSessionModuleConfig
	components *MedicalSessionModuleComponents
	isBuilt    bool
}

func NewMedicalSessionModule(config *MedicalSessionModuleConfig) *MedicalSessionModule {
	return &MedicalSessionModule{
		config:  config,
		isBuilt: false,
	}
}

func (m *MedicalSessionModule) Bootstrap() error {
	if m.isBuilt {
		return nil
	}

	if err := m.validateConfig(); err != nil {
		return err
	}

	repository := m.createRepository()
	medSessionbus := m.createBus(repository)
	controller := m.createController(medSessionbus)
	m.registerRoutes(controller)

	m.components = &MedicalSessionModuleComponents{
		Repository: repository,
		Bus:        medSessionbus,
		Controller: controller,
	}

	m.isBuilt = true
	return nil
}

func (m *MedicalSessionModule) createRepository() repository.MedicalSessionRepository {
	return repositoryimpl.NewSQLCMedSessionRepository(m.config.Queries)
}

func (m *MedicalSessionModule) createBus(repository repository.MedicalSessionRepository) *bus.MedicalSessionBus {
	return bus.NewMedicalSessionBus(
		medSessionCommand.NewMedicalSessionCommandHandlers(repository),
		query.NewMedicalSessionQueryHandlers(repository),
	)
}

func (m *MedicalSessionModule) createController(medSessionBus *bus.MedicalSessionBus) *MedicalSessionControllers {
	controllerOperations := controller.NewMedSessionControllerOperations(medSessionBus, m.config.Validator)
	adminController := controller.NewAdminMedicalSessionController(controllerOperations)
	customerController := controller.NewCustomerMedicalSessionController(controllerOperations)
	employeeController := controller.NewEmployeeMedicalSessionController(controllerOperations)

	return &MedicalSessionControllers{
		AdminController:    adminController,
		CustomerController: customerController,
		EmployeeController: employeeController,
	}
}

func (m *MedicalSessionModule) registerRoutes(ctrl *MedicalSessionControllers) {
	medSessionRoutes := routes.NewMedicalSessionRoutes(ctrl.AdminController, ctrl.CustomerController, ctrl.EmployeeController)
	medSessionRoutes.RegisterAdminRoutes(m.config.Router, m.config.AuthMiddleware)
	medSessionRoutes.RegisterCustomerRoutes(m.config.Router, m.config.AuthMiddleware)
	medSessionRoutes.RegisterEmployeeRoutes(m.config.Router, m.config.AuthMiddleware)
}

func (m *MedicalSessionModule) validateConfig() error {
	if m.config == nil {
		return fmt.Errorf("medical history module configuration cannot be nil")
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
	if m.config.EmployeeRepo == nil {
		return fmt.Errorf("employee repository cannot be nil")
	}
	if m.config.PetRepo == nil {
		return fmt.Errorf("pet repository cannot be nil")
	}

	if m.config.AuthMiddleware == nil {
		return fmt.Errorf("auth middleware cannot be nil")
	}
	return nil
}

func (m *MedicalSessionModule) GetComponents() (*MedicalSessionModuleComponents, error) {
	if !m.isBuilt {
		if err := m.Bootstrap(); err != nil {
			return nil, err
		}
	}
	return m.components, nil
}

func (m *MedicalSessionModule) GetRepository() (repository.MedicalSessionRepository, error) {
	components, err := m.GetComponents()
	if err != nil {
		return nil, err
	}
	return components.Repository, nil
}

func (m *MedicalSessionModule) GetBus() (*bus.MedicalSessionBus, error) {
	components, err := m.GetComponents()
	if err != nil {
		return nil, err
	}
	return components.Bus, nil
}

func (m *MedicalSessionModule) GetController() (*MedicalSessionControllers, error) {
	components, err := m.GetComponents()
	if err != nil {
		return nil, err
	}
	return components.Controller, nil
}
