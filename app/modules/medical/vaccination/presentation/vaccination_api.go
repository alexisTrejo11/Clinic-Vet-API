package api

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/modules/core/service"
	"clinic-vet-api/app/modules/medical/vaccination/application"
	"clinic-vet-api/app/modules/medical/vaccination/application/handler"
	sqlcRepo "clinic-vet-api/app/modules/medical/vaccination/infrastructure/repository"
	"clinic-vet-api/app/modules/medical/vaccination/presentation/controller"
	"clinic-vet-api/app/modules/medical/vaccination/presentation/routes"
	"clinic-vet-api/app/shared/mapper"
	"clinic-vet-api/sqlc"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VaccinationConfig struct {
	Router         *gin.RouterGroup
	Validator      *validator.Validate
	AuthMiddleware *middleware.AuthMiddleware
	PetRepo        repository.PetRepository
	EmployeeRepo   repository.EmployeeRepository
	CustomerRepo   repository.CustomerRepository
	Queries        *sqlc.Queries
}

type VaccinationComponents struct {
	Repository repository.VaccinationRepository
	Service    *application.VaccinationFacadeService
	Controller *VaccinationController
	Routes     *routes.VaccinationRoutes
}

type VaccinationController struct {
	Client   *controller.CustomerPetVaccinationController
	Employee *controller.EmployeePetVaccinationController
}

type VaccinationAPIModule struct {
	config     *VaccinationConfig
	isBuilt    bool
	Components VaccinationComponents
}

func NewVaccinationAPIModule(config *VaccinationConfig) *VaccinationAPIModule {
	return &VaccinationAPIModule{
		config:  config,
		isBuilt: false,
	}
}

func (b *VaccinationAPIModule) Bootstrap() error {
	if b.isBuilt {
		return nil
	}

	if err := b.validateConfig(); err != nil {
		return err
	}

	repo := sqlcRepo.NewSqlcPetVaccinationRepository(b.config.Queries, mapper.SqlcFieldMapper{})
	vaccinationScheduleService := service.NewVaccinationScheduleService(nil)

	petVaccineCmdHandler := handler.NewPetVaccineCmdHandler(b.config.PetRepo, repo, vaccinationScheduleService)
	petVaccineQryHandler := handler.NewVaccinationQueryHandler(repo, b.config.EmployeeRepo, b.config.PetRepo)
	service := application.NewVaccinationFacadeService(petVaccineQryHandler, petVaccineCmdHandler)

	controllers := &VaccinationController{
		Client:   controller.NewCustomerPetVaccinationController(service, b.config.Validator),
		Employee: controller.NewEmployeePetVaccinationController(service, b.config.Validator),
	}

	dRoutes := routes.NewVaccinationRoutes(controllers.Client, controllers.Employee)
	dRoutes.RegisterCustomerRoutes(b.config.Router, b.config.AuthMiddleware, controllers.Client)
	dRoutes.RegisterEmployeeRoutes(b.config.Router, b.config.AuthMiddleware, controllers.Employee)

	b.Components = VaccinationComponents{
		Repository: repo,
		Service:    &service,
		Controller: controllers,
		Routes:     dRoutes,
	}
	b.isBuilt = true

	return nil
}

func (b *VaccinationAPIModule) validateConfig() error {
	if b.config == nil {
		return errors.New("vaccination api config is nil")
	}
	if b.config.Router == nil {
		return errors.New("router is nil")
	}

	if b.config.Validator == nil {
		return errors.New("validator is nil")
	}

	if b.config.AuthMiddleware == nil {
		return errors.New("auth middleware is nil")
	}

	if b.config.Queries == nil {
		return errors.New("queries is nil")
	}

	if b.config.PetRepo == nil {
		return errors.New("pet repository is nil")
	}

	if b.config.EmployeeRepo == nil {
		return errors.New("employee repository is nil")
	}

	if b.config.CustomerRepo == nil {
		return errors.New("customer repository is nil")
	}

	return nil
}
