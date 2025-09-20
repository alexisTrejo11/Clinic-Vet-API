package api

import (
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/employee/application/cqrs/command"
	"clinic-vet-api/app/modules/employee/application/cqrs/query"
	"clinic-vet-api/app/modules/employee/infrastructure/bus"
	repositoryimpl "clinic-vet-api/app/modules/employee/infrastructure/repository"
	"clinic-vet-api/app/modules/employee/presentation/controller"
	"clinic-vet-api/app/modules/employee/presentation/routes"
	"clinic-vet-api/sqlc"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeAPIConfig struct {
	Queries        *sqlc.Queries
	DB             *pgxpool.Pool
	Router         *gin.RouterGroup
	DataValidator  *validator.Validate
	AuthMiddleware *middleware.AuthMiddleware
}

type EmployeeAPIComponents struct {
	bus        *bus.EmployeeBus
	controller *controller.EmployeeController
	repository repository.EmployeeRepository
}

type EmployeeModule struct {
	config     *EmployeeAPIConfig
	components *EmployeeAPIComponents
	isBuilt    bool
}

func NewEmployeeModule(config *EmployeeAPIConfig) *EmployeeModule {
	return &EmployeeModule{
		config:  config,
		isBuilt: false,
	}
}

func (f *EmployeeModule) Bootstrap() error {
	if f.isBuilt {
		return nil
	}

	if err := f.validateConfig(); err != nil {
		return err
	}

	f.components = &EmployeeAPIComponents{}
	vetRepo := repositoryimpl.NewSqlcEmployeeRepository(f.config.Queries, f.config.DB)

	employeeCommandBus := command.NewEmployeeCommandBus(vetRepo)
	employeeQueryBus := query.NewEmployeeQueryBus(vetRepo)
	vetServiceBus := bus.NewEmployeeBus(*employeeCommandBus, employeeQueryBus)
	vetControllers := controller.NewEmployeeController(f.config.DataValidator, vetServiceBus)

	routes.EmployeeRoutes(f.config.Router, vetControllers, f.config.AuthMiddleware)

	f.components.controller = vetControllers
	f.components.bus = &vetServiceBus
	f.components.repository = vetRepo
	f.isBuilt = true

	return nil
}

func (f *EmployeeModule) validateConfig() error {
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

	if f.config.AuthMiddleware == nil {
		return fmt.Errorf("auth middleware cannot be nil")
	}

	return nil
}

func (f *EmployeeModule) GetRepository() (repository.EmployeeRepository, error) {
	if !f.isBuilt {
		return nil, errors.New("module not bootstrapped")
	}
	return f.components.repository, nil

}
