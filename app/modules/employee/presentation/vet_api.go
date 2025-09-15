package api

import (
	"errors"
	"fmt"

	repository "clinic-vet-api/app/core/repositories"
	"clinic-vet-api/app/modules/veterinarians/application/usecase"
	"clinic-vet-api/app/modules/veterinarians/infrastructure/api/controller"
	"clinic-vet-api/app/modules/veterinarians/infrastructure/api/routes"
	"clinic-vet-api/app/modules/veterinarians/infrastructure/persistence"
	"clinic-vet-api/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgconn"
)

type VeterinarianAPIConfig struct {
	Queries       *sqlc.Queries
	DB            *pgconn.Conn
	Router        *gin.Engine
	DataValidator *validator.Validate
}

type VeterinarianAPIComponents struct {
	useCase    *usecase.VeterinarianUseCases
	controller *controller.VeterinarianController
	repository repository.VetRepository
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

	f.components = &VeterinarianAPIComponents{}
	vetRepo := persistence.NewSqlcVetRepository(f.config.Queries, f.config.DB)

	getVetUseCase := usecase.NewGetVetByIDUseCase(vetRepo)
	listVetUseCase := usecase.NewListVetUseCase(vetRepo)
	createVetUseCase := usecase.NewCreateVetUseCase(vetRepo)
	updateVetUseCase := usecase.NewUpdateVetUseCase(vetRepo)
	deleteVetUseCase := usecase.NewDeleteVetUseCase(vetRepo)

	vetUseCaseContainer := usecase.NewVetUseCase(*listVetUseCase, *getVetUseCase, *createVetUseCase, *updateVetUseCase, *deleteVetUseCase)

	vetControllers := controller.NewVeterinarianController(f.config.DataValidator, *vetUseCaseContainer)

	routes.VetRoutes(f.config.Router, vetControllers)

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

func (f *VeterinarianModule) GetRepository() (repository.VetRepository, error) {
	if !f.isBuilt {
		return nil, errors.New("module not bootstrapped")
	}
	return f.components.repository, nil

