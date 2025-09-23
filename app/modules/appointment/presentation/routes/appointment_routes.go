// Package routes contains all appointment-related route definitions
package routes

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/appointment/presentation/controller"

	"github.com/gin-gonic/gin"
)

type AppointmentRoutes struct {
	customerController *controller.CustomerAppointmetController
	employeeController *controller.EmployeeAppointmentController
}

func NewAppointmentRoutes(
	customerController *controller.CustomerAppointmetController,
	employeeController *controller.EmployeeAppointmentController,
) *AppointmentRoutes {
	return &AppointmentRoutes{
		customerController: customerController,
		employeeController: employeeController,
	}
}

func (r *AppointmentRoutes) RegisterAdminRoutes(Router *gin.RouterGroup) {
	// Admin/General appointment routes
	// appointmentGroup := router.Group("/api/v2/appointments")
	{
		// Command operations (admin access)

		// Query operations (admin access)
	}
}

func (r *AppointmentRoutes) RegisterCustomerRoutes(router *gin.RouterGroup, authMiddleware *middleware.AuthMiddleware) {
	customerGroup := router.Group("/customers/appointments")
	customerGroup.Use(authMiddleware.Authenticate())
	customerGroup.Use(authMiddleware.RequireAnyRole("customer"))

	customerGroup.GET("/", r.customerController.GetMyAppointments)
	customerGroup.GET("/upcoming", r.customerController.GetMyUpcomingAppointments)
	customerGroup.GET("/:id", r.customerController.GetMyAppointmentByID)
	customerGroup.GET("/pets/:petID/", r.customerController.GetAppointmentsByPet)
	customerGroup.POST("/", r.customerController.RequestAppointment)
	//customerGroup.PUT("//:id/reschedule", controller.RescheduleAppointment)
	//customerGroup.DELETE("/:id", controller.CancelAppointment)

}

func (r *AppointmentRoutes) RegisterEmployeeRoutes(router *gin.RouterGroup, authMiddleware *middleware.AuthMiddleware) {
	employeeRoutes := router.Group("employees/appointments")
	employeeRoutes.Use(authMiddleware.Authenticate())
	employeeRoutes.Use(authMiddleware.RequireAnyRole("veterinarian", "receptionist"))

	employeeRoutes.GET("/", r.employeeController.GetMyAppointments)
	employeeRoutes.GET("/stats", r.employeeController.GetAppointmentStats)
	employeeRoutes.PUT("/:id/confirm", r.employeeController.ConfirmAppointment)
	employeeRoutes.PUT("/:id/complete", r.employeeController.CompleteAppointment)
	employeeRoutes.PUT("/:id/reschedule", r.employeeController.RescheduleAppointment)
	employeeRoutes.PUT("/:id/no-show", r.employeeController.MarkAsNoShow)
	employeeRoutes.DELETE("/:id", r.employeeController.CancelAppointment)
}
