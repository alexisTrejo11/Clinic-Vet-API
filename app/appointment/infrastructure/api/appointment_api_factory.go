package appointmentAPI

import (
	"fmt"

	appointmentCmd "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/application/command"
	appointmentQuery "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/application/queries"
	appointmentController "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/infrastructure/api/controller"
	appointmentRoutes "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/infrastructure/api/routes"
	appointmentRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/infrastructure/persistence/repositories"
	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// AppointmentAPIConfig holds configuration dependencies
type AppointmentAPIConfig struct {
	Router    *gin.Engine
	Queries   *sqlc.Queries
	Validator *validator.Validate
	OwnerRepo ownerDomain.OwnerRepository
}

// AppointmentAPIComponents holds all created components
type AppointmentAPIComponents struct {
	Repository  interface{} // Your appointment repository interface
	CommandBus  *appointmentCmd.CommandBus
	QueryBus    *appointmentQuery.QueryBus
	Controllers *AppointmentControllers
	Routes      *appointmentRoutes.AppointmentRoutes
}

// AppointmentControllers holds all appointment controllers
type AppointmentControllers struct {
	Command *appointmentController.AppointmentCommandController
	Query   *appointmentController.AppointmentQueryController
	Owner   *appointmentController.OwnerAppointmentController
	Vet     *appointmentController.VetAppointmentController
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
	repository := appointmentRepo.NewSQLCAppointmentRepository(f.config.Queries)

	// Create buses
	commandBus := appointmentCmd.NewAppointmentCommandBus(repository)
	queryBus := appointmentQuery.NewAppointmentQueryBus(repository, f.config.OwnerRepo)

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
	commandBus *appointmentCmd.CommandBus,
	queryBus *appointmentQuery.QueryBus,
) *AppointmentControllers {
	return &AppointmentControllers{
		Command: appointmentController.NewAppointmentCommandController(*commandBus, f.config.Validator),
		Query:   appointmentController.NewAppointmentQueryController(*queryBus, f.config.Validator),
		Owner:   appointmentController.NewOwnerAppointmentController(*commandBus, *queryBus),
		Vet:     appointmentController.NewVetAppointmentController(*commandBus, *queryBus),
	}
}

// createRoutes creates routes and registers them
func (f *AppointmentAPIBuilder) createRoutes(controllers *AppointmentControllers) *appointmentRoutes.AppointmentRoutes {
	routes := appointmentRoutes.NewAppointmentRoutes(
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
