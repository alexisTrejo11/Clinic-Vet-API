package customerAPI

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/modules/customer/application/handler"
	"clinic-vet-api/app/modules/customer/infrastructure/bus"
	customerRepo "clinic-vet-api/app/modules/customer/infrastructure/repository"
	"clinic-vet-api/app/modules/customer/presentation/controller"
	"clinic-vet-api/app/modules/customer/presentation/routes"
	"clinic-vet-api/sqlc"

	petRepo "clinic-vet-api/app/modules/pet/infrastructure/repository"
	appError "clinic-vet-api/app/shared/error/application"
	"clinic-vet-api/app/shared/mapper"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CustomerAPIConfig struct {
	Router         *gin.RouterGroup
	Validator      *validator.Validate
	Queries        *sqlc.Queries
	AuthMiddleware *middleware.AuthMiddleware
}

type CustomerAPIComponents struct {
	Repository repository.CustomerRepository
	Bus        bus.CustomerBus
	Controller *controller.CustomerController
}

type CustomerAPIModule struct {
	config     *CustomerAPIConfig
	components *CustomerAPIComponents
	isBuilt    bool
}

func NewCustomerAPIModule(config *CustomerAPIConfig) *CustomerAPIModule {
	return &CustomerAPIModule{
		config:  config,
		isBuilt: false,
	}
}

func (f *CustomerAPIModule) Bootstrap() error {
	if f.isBuilt {
		return nil
	}

	if err := f.validateConfig(); err != nil {
		return err
	}

	// Create repository
	petRepo := petRepo.NewSqlcPetRepository(f.config.Queries, mapper.NewSqlcFieldMapper())
	customerRepo := customerRepo.NewSqlcCustomerRepository(f.config.Queries, petRepo)

	// Create use cases
	customerBus := f.createBus(customerRepo)

	// Create controller
	controller := controller.NewCustomerController(f.config.Validator, customerBus)

	// Register routes
	routes.CustomerRoutes(f.config.Router, controller, f.config.AuthMiddleware)

	// Store components
	f.components = &CustomerAPIComponents{
		Repository: customerRepo,
		Bus:        customerBus,
		Controller: controller,
	}

	f.isBuilt = true
	return nil
}

// createUseCases creates and wires all use cases
func (f *CustomerAPIModule) createBus(repo repository.CustomerRepository) bus.CustomerBus {
	customerCommandHandler := handler.NewCustomerCommandHandler(repo)
	customerQueryHandler := handler.NewCustomerQueryHandler(repo)

	return bus.NewCustomerBus(*customerCommandHandler, *customerQueryHandler)
}

// validateConfig validates the Module configuration
func (f *CustomerAPIModule) validateConfig() error {
	if f.config == nil {
		return appError.ConflictError("INVALID_CONFIG", "configuration cannot be nil")
	}
	if f.config.Router == nil {
		return appError.ConflictError("INVALID_CONFIG", "router cannot be nil")
	}
	if f.config.Validator == nil {
		return appError.ConflictError("INVALID_CONFIG", "validator cannot be nil")
	}
	if f.config.Queries == nil {
		return appError.ConflictError("INVALID_CONFIG", "queries cannot be nil")
	}
	if f.config.AuthMiddleware == nil {
		return appError.ConflictError("INVALID_CONFIG", "auth middleware cannot be nil")
	}

	return nil
}

func (f *CustomerAPIModule) GetComponents() (*CustomerAPIComponents, error) {
	if !f.isBuilt {
		if err := f.Bootstrap(); err != nil {
			return nil, err
		}
	}
	return f.components, nil
}

func (f *CustomerAPIModule) GetRepository() (repository.CustomerRepository, error) {
	components, err := f.GetComponents()
	if err != nil {
		return nil, err
	}
	return components.Repository, nil
}

func (f *CustomerAPIModule) GetBus() (*bus.CustomerBus, error) {
	components, err := f.GetComponents()
	if err != nil {
		return nil, err
	}
	return &components.Bus, nil
}

func (f *CustomerAPIModule) GetController() (*controller.CustomerController, error) {
	components, err := f.GetComponents()
	if err != nil {
		return nil, err
	}
	return components.Controller, nil
}
