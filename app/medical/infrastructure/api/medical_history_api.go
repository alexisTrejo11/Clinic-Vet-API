package medHistoryAPI

import (
	"fmt"

	medHistUsecases "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/application/usecase"
	medHistRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/domain/repositories"
	medHistController "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/infrastructure/api/controller"
	medHistoryRoutes "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/infrastructure/api/routes"
	sqlcMedHistoryRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/infrastructure/persistence/repositories"
	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
	vetRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MedicalHistoryModuleConfig struct {
	Router    *gin.Engine
	Queries   *sqlc.Queries
	Validator *validator.Validate
	OwnerRepo ownerDomain.OwnerRepository
	VetRepo   vetRepo.VeterinarianRepository
	PetRepo   petDomain.PetRepository
}

type MedicalHistoryModuleComponents struct {
	Repository medHistRepo.MedicalHistoryRepository
	UseCase    *medHistUsecases.MedicalHistoryUseCase
	Controller *medHistController.AdminMedicalHistoryController
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

func (m *MedicalHistoryModule) createRepository() medHistRepo.MedicalHistoryRepository {
	return sqlcMedHistoryRepo.NewSQLCMedHistRepository(m.config.Queries)
}

func (m *MedicalHistoryModule) createUseCase(repository medHistRepo.MedicalHistoryRepository) *medHistUsecases.MedicalHistoryUseCase {
	return medHistUsecases.NewMedicalHistoryUseCase(
		repository,
		m.config.OwnerRepo,
		m.config.VetRepo,
		m.config.PetRepo,
	)
}

func (m *MedicalHistoryModule) createController(useCase *medHistUsecases.MedicalHistoryUseCase) *medHistController.AdminMedicalHistoryController {
	return medHistController.NewAdminMedicalHistoryController(useCase)
}

func (m *MedicalHistoryModule) registerRoutes(controller *medHistController.AdminMedicalHistoryController) {
	medHistoryRoutes.MedicalHistoryRoutes(m.config.Router, *controller)
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

func (m *MedicalHistoryModule) GetRepository() (medHistRepo.MedicalHistoryRepository, error) {
	components, err := m.GetComponents()
	if err != nil {
		return nil, err
	}
	return components.Repository, nil
}

func (m *MedicalHistoryModule) GetUseCase() (*medHistUsecases.MedicalHistoryUseCase, error) {
	components, err := m.GetComponents()
	if err != nil {
		return nil, err
	}
	return components.UseCase, nil
}

func (m *MedicalHistoryModule) GetController() (*medHistController.AdminMedicalHistoryController, error) {
	components, err := m.GetComponents()
	if err != nil {
		return nil, err
	}
	return components.Controller, nil
}
