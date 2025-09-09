package appointmentAPI

import (
	"fmt"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/infrastructure/api/controller"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/infrastructure/api/routes"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/infrastructure/bus"
	repositoryimpl "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/infrastructure/persistence/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// AppointmentAPIConfig holds configuration dependencies
type AppointmentAPIConfig struct {
	Router    *gin.Engine
	Queries   *sqlc.Queries
	Validator *validator.Validate
	OwnerRepo repository.OwnerRepository
}

// AppointmentAPIComponents holds all created components
type AppointmentAPIComponents struct {
	Repository  repository.AppointmentRepository // Your appointment repository interface
	CommandBus  *cqrs.CommandBus
	QueryBus    *cqrs.QueryBus
	Controllers *AppointmentControllers
	Routes      *routes.AppointmentRoutes
}

// AppointmentControllers holds all appointment controllers
type AppointmentControllers struct {
	Command *controller.AppointmentCommandController
	Query   *controller.AppointmentQueryController
	Owner   *controller.OwnerAppointmentController
	Vet     *controller.VetAppointmentController
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
	repository := repositoryimpl.NewSQLCAppointmentRepository(f.config.Queries)

	// Create buses
	commandBus := bus.NewAppointmentCommandBus(repository)
	queryBus := bus.NewAppointmentQueryBus(repository, f.config.OwnerRepo)

	// Create controllers
	controllers := f.createControllers(&commandBus, &queryBus)

	// Create and register routes
	routes := f.createRoutes(controllers)

	// Store components
	f.components = &AppointmentAPIComponents{
		Repository:  repository,
		CommandBus:  &commandBus,
		QueryBus:    &queryBus,
		Controllers: controllers,
		Routes:      routes,
	}

	f.isBuilt = true
	return nil
}

// createControllers creates all appointment controllers
func (f *AppointmentAPIBuilder) createControllers(
	commandBus *cqrs.CommandBus,
	queryBus *cqrs.QueryBus,
) *AppointmentControllers {
	return &AppointmentControllers{
		Command: controller.NewAppointmentCommandController(*commandBus, f.config.Validator),
		Query:   controller.NewAppointmentQueryController(*queryBus, f.config.Validator),
		Owner:   controller.NewOwnerAppointmentController(*commandBus, *queryBus),
		Vet:     controller.NewVetAppointmentController(*commandBus, *queryBus, f.config.Validator),
	}
}

// createRoutes creates routes and registers them
func (f *AppointmentAPIBuilder) createRoutes(controllers *AppointmentControllers) *routes.AppointmentRoutes {
	routes := routes.NewAppointmentRoutes(
		controllers.Command,
		controllers.Query,
		controllers.Owner,
		controllers.Vet,
	)
	routes.RegisterAdminRoutes(f.config.Router)
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
	if f.config.OwnerRepo == nil {
		return fmt.Errorf("owner repository cannot be nil")
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
