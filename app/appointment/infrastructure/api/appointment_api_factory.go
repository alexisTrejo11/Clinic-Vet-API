package appointmentAPI

import (
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

type AppointmentApiFactory struct {
	queries   *sqlc.Queries
	validator *validator.Validate
	ownerRepo ownerDomain.OwnerRepository
}

func NewAppointmentApiFactory(queries *sqlc.Queries, validator *validator.Validate, ownerRepo ownerDomain.OwnerRepository) *AppointmentApiFactory {
	return &AppointmentApiFactory{
		queries:   queries,
		validator: validator,
		ownerRepo: ownerRepo,
	}
}

func (factory *AppointmentApiFactory) CreateCommandController() *appointmentController.AppointmentCommandController {
	appointmentRepository := appointmentRepo.NewSQLCAppointmentRepository(factory.queries)
	commandBus := appointmentCmd.NewAppointmentCommandBus(appointmentRepository)

	return appointmentController.NewAppointmentCommandController(commandBus, factory.validator)
}

func (factory *AppointmentApiFactory) CreateQueryController() *appointmentController.AppointmentQueryController {
	appointmentRepository := appointmentRepo.NewSQLCAppointmentRepository(factory.queries)
	queryBus := appointmentQuery.NewAppointmentQueryBus(appointmentRepository, factory.ownerRepo)

	return appointmentController.NewAppointmentQueryController(queryBus, factory.validator)
}

func (factory *AppointmentApiFactory) CreateOwnerController() *appointmentController.OwnerAppointmentController {
	appointmentRepository := appointmentRepo.NewSQLCAppointmentRepository(factory.queries)
	commandBus := appointmentCmd.NewAppointmentCommandBus(appointmentRepository)
	queryBus := appointmentQuery.NewAppointmentQueryBus(appointmentRepository, factory.ownerRepo)

	return appointmentController.NewOwnerAppointmentController(commandBus, queryBus)
}

func (factory *AppointmentApiFactory) CreateVetController() *appointmentController.VetAppointmentController {
	appointmentRepository := appointmentRepo.NewSQLCAppointmentRepository(factory.queries)

	commandBus := appointmentCmd.NewAppointmentCommandBus(appointmentRepository)
	queryBus := appointmentQuery.NewAppointmentQueryBus(appointmentRepository, factory.ownerRepo)

	return appointmentController.NewVetAppointmentController(commandBus, queryBus)
}

func (factory *AppointmentApiFactory) CreateRoutes(router *gin.Engine) *appointmentRoutes.AppointmentRoutes {
	appointmentController := factory.CreateCommandController()
	ownerAppointmentController := factory.CreateOwnerController()
	vetAppointmentController := factory.CreateVetController()
	appointmentQueryController := factory.CreateQueryController()

	routes := appointmentRoutes.NewAppointmentRoutes(appointmentController, appointmentQueryController, ownerAppointmentController, vetAppointmentController)
	routes.RegisterAdminRoutes(router)
	return routes
}

func SetupAppointmentAPI(router *gin.Engine, queries *sqlc.Queries, validator *validator.Validate, ownerRepo ownerDomain.OwnerRepository) {
	factory := NewAppointmentApiFactory(queries, validator, ownerRepo)
	factory.CreateRoutes(router)
}
