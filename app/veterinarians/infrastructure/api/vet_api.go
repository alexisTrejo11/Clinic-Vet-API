package vetAPI

import (
	"errors"
	"fmt"

	vetRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/repositories"
	vetUsecase "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/usecase"
	vetController "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/infrastructure/api/controller"
	vetRoutes "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/infrastructure/api/routes"
	sqlcVetRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/infrastructure/persistence/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VeterinarianAPIConfig struct {
	Queries       *sqlc.Queries
	Router        *gin.Engine
	DataValidator *validator.Validate
}

type VeterinarianAPIComponents struct {
	useCase    *vetUsecase.VeterinarianUseCases
	controller *vetController.VeterinarianController
	repository vetRepo.VeterinarianRepository
}

type VeterinarianModule struct {
	config     *VeterinarianAPIConfig
	components *VeterinarianAPIComponents
	isBuilt    bool
}

func NewVeterinarianModule(config *VeterinarianAPIConfig) *VeterinarianModule {
	return &VeterinarianModule{
		config:  config,
		isBuilt: false,
	}
}

func (f *VeterinarianModule) Bootstrap() error {
	if f.isBuilt {
		return nil
	}

	if err := f.validateConfig(); err != nil {
		return err
	}

	vetRepo := sqlcVetRepo.NewSqlcVetRepository(f.config.Queries)

	getVetUseCase := vetUsecase.NewGetVetByIdUseCase(vetRepo)
	listVetUseCase := vetUsecase.NewListVetUseCase(vetRepo)
	createVetUseCase := vetUsecase.NewCreateVetUseCase(vetRepo)
	updateVetUseCase := vetUsecase.NewUpdateVetUseCase(vetRepo)
	deleteVetUseCase := vetUsecase.NewDeleteVetUseCase(vetRepo)

	vetUseCaseContainer := vetUsecase.NewVetUseCase(*listVetUseCase, *getVetUseCase, *createVetUseCase, *updateVetUseCase, *deleteVetUseCase)

	vetControllers := vetController.NewVeterinarianController(f.config.DataValidator, *vetUseCaseContainer)

	vetRoutes.VetRoutes(f.config.Router, vetControllers)

	f.components.controller = vetControllers
	f.components.useCase = vetUseCaseContainer
	f.components.repository = vetRepo
	f.isBuilt = true

	return nil
}

func (f *VeterinarianModule) validateConfig() error {
	if f.config == nil {
		return errors.New("invalid config: nil")
	}

	if f.config.Router == nil {
		return fmt.Errorf("router cannot be nil")
	}

	if f.config.Queries == nil {
		return fmt.Errorf("queries cannot be nil")
	}

	if f.config.DataValidator == nil {
		return fmt.Errorf("validator cannot be nil")
	}

	return nil
}

func (f *VeterinarianModule) GetRepository() (vetRepo.VeterinarianRepository, error) {
	if !f.isBuilt {
		return nil, errors.New("module not bootstrapped")
	}
	return f.components.repository, nil
}
