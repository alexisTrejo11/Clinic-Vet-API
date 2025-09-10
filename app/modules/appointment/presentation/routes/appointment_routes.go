package routes

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/presentation/controller"
	"github.com/gin-gonic/gin"
)

type AppointmentRoutes struct {
	// Admin
	appointmentCommandController *controller.ApptAdminCommandController
	appointmentQueryController   *controller.ApptAdminQueryController

	// Customer
	ownerAppointmentController *controller.CustomerAppointmentController
	employeeApptController     *controller.VetAppointmentController
}

func NewAppointmentRoutes(
	appointmentCommandController *controller.AppointmentCommandController,
	appointmentQueryController *controller.AppointmentQueryController,
	ownerApptController *controller.OwnerAppointmentController,
	employeeApptController *controller.VetAppointmentController,
) *AppointmentRoutes {
	return &AppointmentRoutes{
		appointmentCommandController: appointmentCommandController,
		appointmentQueryController:   appointmentQueryController,
		ownerApptController:          ownerApptController,
		employeeApptController:       employeeApptController,
	}
}

// RegisterRoutes sets up all appointment-related routes
func (r *AppointmentRoutes) RegisterAdminRoutes(router *gin.Engine) {
	// Admin/General appointment routes
	appointmentGroup := router.Group("/api/v2/appointments")
	{
		// Command operations (admin access)
		appointmentGroup.POST("", r.appointmentCommandController.CreateAppointment)
		appointmentGroup.PUT("/:id", r.appointmentCommandController.UpdateAppointment)
		appointmentGroup.DELETE("/:id", r.appointmentCommandController.DeleteAppointment)

		// Query operations (admin access)
		appointmentGroup.GET("", r.appointmentQueryController.SearchAppointments)
		appointmentGroup.GET("/:id", r.appointmentQueryController.GetAppointmentByID)
		appointmentGroup.GET("/date-range", r.appointmentQueryController.GetAppointmentsByDateRange)
		appointmentGroup.GET("/owner/:ownerID", r.appointmentQueryController.GetAppointmentsByOwner)
		appointmentGroup.GET("/vet/:vetID", r.appointmentQueryController.GetAppointmentsByVet)
		appointmentGroup.GET("/pet/:petID", r.appointmentQueryController.GetAppointmentsByPet)
		appointmentGroup.GET("/stats", r.appointmentQueryController.GetAppointmentStats)
	}

	// Owner-specific appointment routes
	ownerGroup := router.Group("/api/v2/owner")
	{
		// Owner appointment management
		ownerGroup.POST("/appointments", r.ownerApptController.RequestAppointment)
		ownerGroup.GET("/appointments", r.ownerApptController.GetMyAppointments)
		ownerGroup.GET("/appointments/:id", r.ownerApptController.GetMyAppointmentDetail)
		ownerGroup.PUT("/appointments/:id/reschedule", r.ownerApptController.RescheduleAppointment)
		ownerGroup.DELETE("/appointments/:id", r.ownerApptController.CancelAppointment)

		// Pet-specific appointments for owners
		ownerGroup.GET("/pets/:petID/appointments", r.ownerApptController.GetAppointmentsByPet)
	}

	// Veterinarian-specific appointment routes
	vetGroup := router.Group("/api/v2/vet")
	{
		// Vet appointment management
		vetGroup.GET("/appointments", r.employeeApptController.GetMyAppointments)
		// vetGroup.GET("/appointments/today", r.employeeApptController.GetTodayAppointments)
		vetGroup.GET("/appointments/stats", r.employeeApptController.GetAppointmentStats)

		// Appointment state management
		vetGroup.PUT("/appointments/:id/confirm", r.employeeApptController.ConfirmAppointment)
		vetGroup.PUT("/appointments/:id/complete", r.employeeApptController.CompleteAppointment)
		vetGroup.PUT("/appointments/:id/reschedule", r.employeeApptController.RescheduleAppointment)
		vetGroup.PUT("/appointments/:id/no-show", r.employeeApptController.MarkAsNoShow)
		vetGroup.DELETE("/appointments/:id", r.employeeApptController.CancelAppointment)
	}
}

// RegisterOwnerRoutes registers only owner-specific appointment routes
func (r *AppointmentRoutes) RegisterOwnerRoutes(router *gin.Engine) {
	ownerGroup := router.Group("/api/v2/owner")
	{
		ownerGroup.POST("/appointments", r.ownerApptController.RequestAppointment)
		ownerGroup.GET("/appointments", r.ownerApptController.GetMyAppointments)
		ownerGroup.GET("/appointments/:id", r.ownerApptController.GetMyAppointmentDetail)
		ownerGroup.PUT("/appointments/:id/reschedule", r.ownerApptController.RescheduleAppointment)
		ownerGroup.DELETE("/appointments/:id", r.ownerApptController.CancelAppointment)
		ownerGroup.GET("/pets/:petID/appointments", r.ownerApptController.GetAppointmentsByPet)
	}
}

// RegisterVetRoutes registers only veterinarian-specific appointment routes
func (r *AppointmentRoutes) RegisterVetRoutes(router *gin.Engine) {
	vetGroup := router.Group("/api/v2/vet")
	{
		vetGroup.GET("/appointments", r.employeeApptController.GetMyAppointments)
		// vetGroup.GET("/appointments/today", r.employeeApptController.GetTodayAppointments)
		vetGroup.GET("/appointments/stats", r.employeeApptController.GetAppointmentStats)
		vetGroup.PUT("/appointments/:id/confirm", r.employeeApptController.ConfirmAppointment)
		vetGroup.PUT("/appointments/:id/complete", r.employeeApptController.CompleteAppointment)
		vetGroup.PUT("/appointments/:id/reschedule", r.employeeApptController.RescheduleAppointment)
		vetGroup.PUT("/appointments/:id/no-show", r.employeeApptController.MarkAsNoShow)
		vetGroup.DELETE("/appointments/:id", r.employeeApptController.CancelAppointment)
	}
}
