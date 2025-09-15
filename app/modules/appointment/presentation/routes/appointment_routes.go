// Package routes contains all appointment-related route definitions
package routes

import (
	"clinic-vet-api/app/modules/appointment/presentation/controller"

	"github.com/gin-gonic/gin"
)

type AppointmentRoutes struct {
	// CQRS
	appointmentCommandController *controller.AppointmentCommandController
	appointmentQueryController   *controller.AppointmentQueryController

	// Users
	customerAppointmentController *controller.CustomerAppointmetController
	employeeAppointmentController *controller.EmployeeAppointmentController
}

func NewAppointmentRoutes(
	appointmentCommandController *controller.AppointmentCommandController,
	appointmentQueryController *controller.AppointmentQueryController,
	employeeAppointmentController *controller.EmployeeAppointmentController,
	customerAppointmetController *controller.CustomerAppointmetController,
) *AppointmentRoutes {
	return &AppointmentRoutes{
		appointmentCommandController:  appointmentCommandController,
		appointmentQueryController:    appointmentQueryController,
		customerAppointmentController: customerAppointmetController,
		employeeAppointmentController: employeeAppointmentController,
	}
}

func (r *AppointmentRoutes) RegisterAdminRoutes(router *gin.Engine) {
	// Admin/General appointment routes
	// appointmentGroup := router.Group("/api/v2/appointments")
	{
		// Command operations (admin access)

		// Query operations (admin access)
	}
}

func (r *AppointmentRoutes) RegisterCustomerRoutes(router *gin.Engine) {
	customerGroup := router.Group("/api/v2/customers")
	{
		customerGroup.POST("/appointments", r.customerAppointmentController.RequestAppointment)
		customerGroup.GET("/appointments", r.customerAppointmentController.GetMyAppointments)
		customerGroup.GET("/appointments/:id", r.customerAppointmentController.GetMyAppointmentByID)
		//		customerGroup.PUT("/appointments/:id/reschedule", r.customerAppointmentController.RescheduleAppointment)
		//		customerGroup.DELETE("/appointments/:id", r.customerAppointmentController.CancelAppointment)
		customerGroup.GET("/pets/:petID/appointments", r.customerAppointmentController.GetAppointmentsByPet)
	}
}

func (r *AppointmentRoutes) RegisterEmployeeRoutes(router *gin.Engine) {
	employeeRoutes := router.Group("/api/v2/employee")
	{
		employeeRoutes.GET("/appointments", r.employeeAppointmentController.GetMyAppointments)
		employeeRoutes.GET("/appointments/stats", r.employeeAppointmentController.GetAppointmentStats)
		employeeRoutes.PUT("/appointments/:id/confirm", r.employeeAppointmentController.ConfirmAppointment)
		employeeRoutes.PUT("/appointments/:id/complete", r.employeeAppointmentController.CompleteAppointment)
		employeeRoutes.PUT("/appointments/:id/reschedule", r.employeeAppointmentController.RescheduleAppointment)
		employeeRoutes.PUT("/appointments/:id/no-show", r.employeeAppointmentController.MarkAsNoShow)
		employeeRoutes.DELETE("/appointments/:id", r.employeeAppointmentController.CancelAppointment)
	}
}
