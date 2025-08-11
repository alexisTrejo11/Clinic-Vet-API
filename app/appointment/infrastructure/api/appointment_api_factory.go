package api

import (
	"database/sql"

	appointmentCmd "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/application/command"
	appointmentQuery "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/application/queries"
	appointmentController "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/infrastructure/api/controller"
	appointmentRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/infrastructure/persistence/repositories"
)

type AppointmentApiFactory struct {
	db *sql.DB
}

func NewAppointmentApiFactory(db *sql.DB) *AppointmentApiFactory {
	return &AppointmentApiFactory{
		db: db,
	}
}

func (factory *AppointmentApiFactory) CreateController() *appointmentController.AppointmentController {
	appointmentRepository := appointmentRepo.NewPostgresAppointmentRepository(factory.db)
	commandBus := appointmentCmd.NewAppointmentCommandBus(appointmentRepository)
	queryBus := appointmentQuery.NewAppointmentQueryBus(appointmentRepository)

	return appointmentController.NewAppointmentController(commandBus, queryBus)
}

func (factory *AppointmentApiFactory) CreateOwnerController() *appointmentController.OwnerAppointmentController {
	appointmentRepository := appointmentRepo.NewPostgresAppointmentRepository(factory.db)
	commandBus := appointmentCmd.NewAppointmentCommandBus(appointmentRepository)
	queryBus := appointmentQuery.NewAppointmentQueryBus(appointmentRepository)

	return appointmentController.NewOwnerAppointmentController(commandBus, queryBus)
}

func (factory *AppointmentApiFactory) CreateVetController() *appointmentController.VetAppointmentController {
	appointmentRepository := appointmentRepo.NewPostgresAppointmentRepository(factory.db)
	commandBus := appointmentCmd.NewAppointmentCommandBus(appointmentRepository)
	queryBus := appointmentQuery.NewAppointmentQueryBus(appointmentRepository)

	return appointmentController.NewVetAppointmentController(commandBus, queryBus)
}
