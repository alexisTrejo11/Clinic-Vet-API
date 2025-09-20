package api

import (
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/appointment/application/command"
	"clinic-vet-api/app/modules/appointment/application/query"
	"clinic-vet-api/app/modules/appointment/infrastructure/bus"
	"clinic-vet-api/app/modules/appointment/presentation/controller"
	"clinic-vet-api/app/modules/appointment/presentation/routes"
	"clinic-vet-api/sqlc"
	"fmt"

	apptRepo "clinic-vet-api/app/modules/appointment/infrastructure/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// AppointmentAPIConfig holds configuration dependencies
type AppointmentAPIConfig struct {
	Router         *gin.RouterGroup
	Queries        *sqlc.Queries
	Validator      *validator.Validate
	CustomerRepo   repository.CustomerRepository
	AuthMiddleware *middleware.AuthMiddleware
}

// AppointmentAPIComponents holds all created components
type AppointmentAPIComponents struct {
	Repository  repository.AppointmentRepository
	Bus         *bus.AppointmentBus
	Controllers *AppointmentControllers
	Routes      *routes.AppointmentRoutes
}

// AppointmentControllers holds all appointment controllers
type AppointmentControllers struct {
	Customer   *controller.CustomerAppointmetController
	Employee   *controller.EmployeeAppointmentController
	Operations *controller.ApptControllerOperations
}

// AppointmentAPIBuilder creates and manages appointment API components
type AppointmentAPIBuilder struct {
	config     *AppointmentAPIConfig
	components *AppointmentAPIComponents
	isBuilt    bool
}

// NewAppointmentAPIBuilder creates a new Builder with configuration
func NewAppointmentAPIBuilder(config *AppointmentAPIConfig) *AppointmentAPIBuilder {
	return &AppointmentAPIBuilder{
		config:  config,
		isBuilt: false,
	}
}

// Build creates all components and registers routes
func (f *AppointmentAPIBuilder) Build() error {
	if f.isBuilt {
		return nil
	}

	if err := f.validateConfig(); err != nil {
		return err
	}

	// Create repository (single instance)
	repository := apptRepo.NewSqlcAppointmentRepository(f.config.Queries)

	// Create buses
	commandHandler := command.NewAppointmentCommandHandler(repository)
	queryHandler := query.NewAppointmentQueryHandler(repository)

	commandBus := bus.NewAppointmentCommandBus(commandHandler)
	queryBus := bus.NewAppointmentQueryBus(queryHandler)
	apptBus := bus.NewAppointmentBus(*commandBus, *queryBus)

	// Create controllers
	controllers := f.createControllers(*apptBus)

	// Create and register routes
	routes := f.createRoutes(controllers)

	// Store components
	f.components = &AppointmentAPIComponents{
		Repository:  repository,
		Bus:         apptBus,
		Controllers: controllers,
		Routes:      routes,
	}

	f.isBuilt = true
	return nil
}

// createControllers creates all appointment controllers
func (f *AppointmentAPIBuilder) createControllers(apptBus bus.AppointmentBus) *AppointmentControllers {
	ctrlOperations := controller.NewApptControllerOperations(apptBus, f.config.Validator)

	return &AppointmentControllers{
		Customer: controller.NewCustomerApptControleer(&apptBus, f.config.Validator, ctrlOperations),
		Employee: controller.NewEmployeeController(ctrlOperations, f.config.Validator),
	}
}

// createRoutes creates routes and registers them
func (f *AppointmentAPIBuilder) createRoutes(controllers *AppointmentControllers) *routes.AppointmentRoutes {
	routes := routes.NewAppointmentRoutes(controllers.Customer, controllers.Employee)
	routes.RegisterAdminRoutes(f.config.Router)
	routes.RegisterCustomerRoutes(f.config.Router, f.config.AuthMiddleware)
	routes.RegisterEmployeeRoutes(f.config.Router, f.config.AuthMiddleware)
	return routes
}

// validateConfig validates the Builder configuration
func (f *AppointmentAPIBuilder) validateConfig() error {
	if f.config == nil {
		return fmt.Errorf("configuration cannot be nil")
	}
	if f.config.Router == nil {
		return fmt.Errorf("router cannot be nil")
	}
	if f.config.Queries == nil {
		return fmt.Errorf("queries cannot be nil")
	}
	if f.config.Validator == nil {
		return fmt.Errorf("validator cannot be nil")
	}
	if f.config.CustomerRepo == nil {
		return fmt.Errorf("customer repository cannot be nil")
	}

	if f.config.AuthMiddleware == nil {
		return fmt.Errorf("auth middleware cannot be nil")
	}

	return nil
}

// GetComponents returns the built components (builds if necessary)
func (f *AppointmentAPIBuilder) GetComponents() (*AppointmentAPIComponents, error) {
	if !f.isBuilt {
		if err := f.Build(); err != nil {
			return nil, err
		}
	}
	return f.components, nil
}

// GetControllers returns all controllers
func (f *AppointmentAPIBuilder) GetControllers() (*AppointmentControllers, error) {
	components, err := f.GetComponents()
	if err != nil {
		return nil, err
	}
	return components.Controllers, nil
}
