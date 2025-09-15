package medHistoryAPI

import (
	"fmt"

	"clinic-vet-api/app/core/repository"
	medHistCommand "clinic-vet-api/app/modules/medical/application/command"
	"clinic-vet-api/app/modules/medical/application/query"
	"clinic-vet-api/app/modules/medical/infrastructure/bus"
	repositoryimpl "clinic-vet-api/app/modules/medical/infrastructure/persistence/repository"
	"clinic-vet-api/app/modules/medical/infrastructure/presentation/controller"
	"clinic-vet-api/app/modules/medical/infrastructure/presentation/routes"
	"clinic-vet-api/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MedicalHistoryModuleConfig struct {
	Router    *gin.Engine
	Queries   *sqlc.Queries
	Validator *validator.Validate
	OwnerRepo *repository.CustomerRepository
	VetRepo   *repository.EmployeeRepository
	PetRepo   *repository.PetRepository
}

type MedicalHistoryModuleComponents struct {
	Repository repository.MedicalHistoryRepository
	Bus        *bus.MedicalHistoryBus
	Controller *controller.AdminMedicalHistoryController
}

type MedicalHistoryModule struct {
	config     *MedicalHistoryModuleConfig
	components *MedicalHistoryModuleComponents
	isBuilt    bool
}

func NewMedicalHistoryModule(config *MedicalHistoryModuleConfig) *MedicalHistoryModule {
	return &MedicalHistoryModule{
		config:  config,
		isBuilt: false,
	}
}

func (m *MedicalHistoryModule) Bootstrap() error {
	if m.isBuilt {
		return nil
	}

	if err := m.validateConfig(); err != nil {
		return err
	}

	repository := m.createRepository()
	medHistbus := m.createBus(repository)
	controller := m.createController(medHistbus)
	m.registerRoutes(controller)

	m.components = &MedicalHistoryModuleComponents{
		Repository: repository,
		Bus:        medHistbus,
		Controller: controller,
	}

	m.isBuilt = true
	return nil
}

func (m *MedicalHistoryModule) createRepository() repository.MedicalHistoryRepository {
	return repositoryimpl.NewSQLCMedHistRepository(m.config.Queries)
}

func (m *MedicalHistoryModule) createBus(repository repository.MedicalHistoryRepository) *bus.MedicalHistoryBus {
	return bus.NewMedicalHistoryBus(
		medHistCommand.NewMedicalHistoryCommandHandlers(repository),
		query.NewMedicalHistoryQueryHandlers(repository),
	)
}

func (m *MedicalHistoryModule) createController(medHistBus *bus.MedicalHistoryBus) *controller.AdminMedicalHistoryController {
	return controller.NewAdminMedicalHistoryController(medHistBus)
}

func (m *MedicalHistoryModule) registerRoutes(controller *controller.AdminMedicalHistoryController) {
	routes.MedicalHistoryRoutes(m.config.Router, *controller)
}

func (m *MedicalHistoryModule) validateConfig() error {
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
	if m.config.OwnerRepo == nil {
		return fmt.Errorf("owner repository cannot be nil")
	}
	if m.config.VetRepo == nil {
		return fmt.Errorf("veterinarian repository cannot be nil")
	}
	if m.config.PetRepo == nil {
		return fmt.Errorf("pet repository cannot be nil")
	}
	return nil
}

func (m *MedicalHistoryModule) GetComponents() (*MedicalHistoryModuleComponents, error) {
	if !m.isBuilt {
		if err := m.Bootstrap(); err != nil {
			return nil, err
		}
	}
	return m.components, nil
}

func (m *MedicalHistoryModule) GetRepository() (repository.MedicalHistoryRepository, error) {
	components, err := m.GetComponents()
	if err != nil {
		return nil, err
	}
	return components.Repository, nil
}

func (m *MedicalHistoryModule) GetBus() (*usecase.MedicalHistoryUseCase, error) {
	components, err := m.GetComponents()
	if err != nil {
		return nil, err
	}
	return components.UseCase, nil
}

func (m *MedicalHistoryModule) GetController() (*controller.AdminMedicalHistoryController, error) {
	components, err := m.GetComponents()
	if err != nil {
		return nil, err
	}
	return components.Controller, nil
}
