package medHistoryAPI

import (
	"fmt"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/medical/application/usecase"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/medical/infrastructure/api/controller"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/medical/infrastructure/api/routes"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/medical/infrastructure/persistence"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MedicalHistoryModuleConfig struct {
	Router    *gin.Engine
	Queries   *sqlc.Queries
	Validator *validator.Validate
	OwnerRepo repository.OwnerRepository
	VetRepo   *repository.VetRepository
	PetRepo   repository.PetRepository
}

type MedicalHistoryModuleComponents struct {
	Repository repository.MedicalHistoryRepository
	UseCase    *usecase.MedicalHistoryUseCase
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

	useCase := m.createUseCase(repository)

	controller := m.createController(useCase)

	m.registerRoutes(controller)

	m.components = &MedicalHistoryModuleComponents{
		Repository: repository,
		UseCase:    useCase,
		Controller: controller,
	}

	m.isBuilt = true
	return nil
}

func (m *MedicalHistoryModule) createRepository() repository.MedicalHistoryRepository {
	return persistence.NewSQLCMedHistRepository(m.config.Queries)
}

func (m *MedicalHistoryModule) createUseCase(repository repository.MedicalHistoryRepository) *usecase.MedicalHistoryUseCase {
	return usecase.NewMedicalHistoryUseCase(
		repository,
		m.config.OwnerRepo,
		*m.config.VetRepo,
		m.config.PetRepo,
	)
}

func (m *MedicalHistoryModule) createController(useCase *usecase.MedicalHistoryUseCase) *controller.AdminMedicalHistoryController {
	return controller.NewAdminMedicalHistoryController(useCase)
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

func (m *MedicalHistoryModule) GetUseCase() (*usecase.MedicalHistoryUseCase, error) {
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
