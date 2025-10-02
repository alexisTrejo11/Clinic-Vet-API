package api

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/modules/medical/deworm/application"
	"clinic-vet-api/app/modules/medical/deworm/application/command"
	"clinic-vet-api/app/modules/medical/deworm/application/query"
	sqlcRepo "clinic-vet-api/app/modules/medical/deworm/infrastructure/repository"
	"clinic-vet-api/app/modules/medical/deworm/presentation/controller"
	"clinic-vet-api/app/modules/medical/deworm/presentation/routes"
	"clinic-vet-api/sqlc"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type DewormAPIModule struct {
	Config     *DewormAPIConfig
	isBuilt    bool
	Components DewormAPIComponents
}

type DewormAPIConfig struct {
	RouterGroup    *gin.RouterGroup
	Validator      *validator.Validate
	Queries        *sqlc.Queries
	AuthMiddleware *middleware.AuthMiddleware

	CustomerRepo repository.CustomerRepository
	PetRepo      repository.PetRepository
	EmployeeRepo repository.EmployeeRepository
}

type DewormAPIComponents struct {
	Repository    repository.DewormRepository
	CqrsHandler   DewormHandlers
	FacadeService application.DewormingFacadeService
	Controllers   DewormControllers
	Routes        routes.DewormRoutes
}

type DewormHandlers struct {
	CommandHandler *command.DewormCommandHandler
	QueryHandler   *query.DewormQueryHandler
}

type DewormControllers struct {
	AdminController    *controller.AdminDewormController
	CustomerController *controller.CustomerPetDewormController
	EmployeeController *controller.EmployeeDewormController
}

func NewDewormAPIModule(config *DewormAPIConfig) *DewormAPIModule {
	return &DewormAPIModule{
		Config:  config,
		isBuilt: false,
	}
}

func (b *DewormAPIModule) validateConfig() error {
	if b.Config.RouterGroup == nil {
		return errors.New("router group is nil")
	}
	if b.Config.Validator == nil {
		return errors.New("validator is nil")
	}
	if b.Config.Queries == nil {
		return errors.New("queries is nil")
	}

	if b.Config.CustomerRepo == nil {
		return errors.New("customer repository is nil")
	}

	if b.Config.PetRepo == nil {
		return errors.New("pet repository is nil")
	}

	if b.Config.EmployeeRepo == nil {
		return errors.New("employee repository is nil")
	}

	if b.Config.AuthMiddleware == nil {
		return errors.New("auth middleware is nil")
	}

	return nil
}

func (b *DewormAPIModule) Bootstrap() error {
	if err := b.validateConfig(); err != nil {
		return err
	}

	if b.isBuilt {
		return nil
	}

	repo := sqlcRepo.NewSqlcPetDeworming(b.Config.Queries)

	cmdHandler := command.NewDewormCommandHandler(repo, b.Config.EmployeeRepo, b.Config.PetRepo)
	qryHandler := query.NewDewormQueryHandler(repo, b.Config.EmployeeRepo, b.Config.PetRepo)

	facadeService := application.NewDewormingFacadeService(qryHandler, cmdHandler)

	controllers := DewormControllers{
		AdminController:    controller.NewAdminDewormController(facadeService, b.Config.Validator),
		CustomerController: controller.NewCustomerPetDewormController(facadeService, b.Config.Validator),
		EmployeeController: controller.NewEmployeeDewormController(facadeService, b.Config.Validator),
	}

	routes := routes.NewDewormRoutes(controllers.AdminController, controllers.CustomerController, controllers.EmployeeController)
	routes.RegisterAdminRoutes(b.Config.RouterGroup, b.Config.AuthMiddleware)
	routes.RegisterCustomerRoutes(b.Config.RouterGroup, b.Config.AuthMiddleware)
	routes.RegisterEmployees(b.Config.RouterGroup, b.Config.AuthMiddleware)

	b.Components = DewormAPIComponents{
		Repository:    repo,
		CqrsHandler:   DewormHandlers{CommandHandler: cmdHandler, QueryHandler: qryHandler},
		FacadeService: facadeService,
		Controllers:   controllers,
		Routes:        *routes,
	}
	b.isBuilt = true
	return nil
}
