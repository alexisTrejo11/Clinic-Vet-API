package appointmentAPI

import (
	appointmentCmd "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/application/command"
	appointmentQuery "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/application/queries"
	appointmentController "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/infrastructure/api/controller"
	appointmentRoutes "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/infrastructure/api/routes"
	appointmentRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/infrastructure/persistence/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/gin-gonic/gin"
)

type AppointmentApiFactory struct {
	queries *sqlc.Queries
}

func NewAppointmentApiFactory(queries *sqlc.Queries) *AppointmentApiFactory {
	return &AppointmentApiFactory{
		queries: queries,
	}
}

func (factory *AppointmentApiFactory) CreateController() *appointmentController.AppointmentController {
	appointmentRepository := appointmentRepo.NewPostgresAppointmentRepository(factory.queries)
	commandBus := appointmentCmd.NewAppointmentCommandBus(appointmentRepository)
	queryBus := appointmentQuery.NewAppointmentQueryBus(appointmentRepository)

	return appointmentController.NewAppointmentController(commandBus, queryBus)
}

func (factory *AppointmentApiFactory) CreateOwnerController() *appointmentController.OwnerAppointmentController {
	appointmentRepository := appointmentRepo.NewPostgresAppointmentRepository(factory.queries)
	commandBus := appointmentCmd.NewAppointmentCommandBus(appointmentRepository)
	queryBus := appointmentQuery.NewAppointmentQueryBus(appointmentRepository)

	return appointmentController.NewOwnerAppointmentController(commandBus, queryBus)
}

func (factory *AppointmentApiFactory) CreateVetController() *appointmentController.VetAppointmentController {
	appointmentRepository := appointmentRepo.NewPostgresAppointmentRepository(factory.queries)
	commandBus := appointmentCmd.NewAppointmentCommandBus(appointmentRepository)
	queryBus := appointmentQuery.NewAppointmentQueryBus(appointmentRepository)

	return appointmentController.NewVetAppointmentController(commandBus, queryBus)
}

func SetupAppoinmentAPI(router *gin.Engine, queries *sqlc.Queries) {
	appointmentRepository := appointmentRepo.NewPostgresAppointmentRepository(queries)
	commandBus := appointmentCmd.NewAppointmentCommandBus(appointmentRepository)
	queryBus := appointmentQuery.NewAppointmentQueryBus(appointmentRepository)
	controllers := appointmentController.NewAppointmentController(commandBus, queryBus)
	vetControllers := appointmentController.NewVetAppointmentController(commandBus, queryBus)
	ownerControllers := appointmentController.NewOwnerAppointmentController(commandBus, queryBus)

	routes := appointmentRoutes.NewAppointmentRoutes(controllers, ownerControllers, vetControllers)
	routes.RegisterRoutes(router)
}
