package petAPI

import (
	"fmt"

	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	petUsecase "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/usecase"
	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
	petController "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/infrastructure/api/controller"
	petRoutes "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/infrastructure/api/routes"
	sqlcPetRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/infrastructure/persistence/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PetModuleConfig struct {
	Router    *gin.Engine
	Queries   *sqlc.Queries
	Validator *validator.Validate
	OwnerRepo ownerDomain.OwnerRepository
}

type PetModuleComponents struct {
	Repository petDomain.PetRepository
	UseCases   petUsecase.PetUseCasesFacade
	Controller *petController.PetController
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

func (m *PetModule) createRepository() petDomain.PetRepository {
	return sqlcPetRepository.NewSqlcPetRepository(m.config.Queries)
}

func (m *PetModule) createUseCases(repository petDomain.PetRepository) petUsecase.PetUseCasesFacade {
	return petUsecase.NewPetUseCasesFacade(repository, m.config.OwnerRepo)
}

func (m *PetModule) createController(useCases petUsecase.PetUseCasesFacade) *petController.PetController {
	return petController.NewPetController(m.config.Validator, useCases)
}

func (m *PetModule) registerRoutes(controller *petController.PetController) {
	petRoutes.PetsRoutes(m.config.Router, controller)
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

func (m *PetModule) GetRepository() (petDomain.PetRepository, error) {
	components, err := m.GetComponents()
	if err != nil {
		return nil, err
	}
	return components.Repository, nil
}

func (m *PetModule) GetUseCases() (petUsecase.PetUseCasesFacade, error) {
	components, err := m.GetComponents()
	if err != nil {
		return nil, err
	}
	return components.UseCases, nil
}

func (m *PetModule) GetController() (*petController.PetController, error) {
	components, err := m.GetComponents()
	if err != nil {
		return nil, err
	}
	return components.Controller, nil
}
