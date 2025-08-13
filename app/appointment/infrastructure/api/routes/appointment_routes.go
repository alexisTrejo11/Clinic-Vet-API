package appointmentRoutes

import (
	appointmentController "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/infrastructure/api/controller"
	"github.com/gin-gonic/gin"
)

type AppointmentRoutes struct {
	appointmentCommandController *appointmentController.AppointmentCommandController
	appointmentQueryController   *appointmentController.AppointmentQueryController
	ownerAppointmentController   *appointmentController.OwnerAppointmentController
	vetAppointmentController     *appointmentController.VetAppointmentController
}

func NewAppointmentRoutes(
	appointmentCommandController *appointmentController.AppointmentCommandController,
	appointmentQueryController *appointmentController.AppointmentQueryController,
	ownerAppointmentController *appointmentController.OwnerAppointmentController,
	vetAppointmentController *appointmentController.VetAppointmentController,
) *AppointmentRoutes {
	return &AppointmentRoutes{
		appointmentCommandController: appointmentCommandController,
		appointmentQueryController:   appointmentQueryController,
		ownerAppointmentController:   ownerAppointmentController,
		vetAppointmentController:     vetAppointmentController,
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
		appointmentGroup.GET("", r.appointmentQueryController.GetAllAppointments)
		appointmentGroup.GET("/:id", r.appointmentQueryController.GetAppointmentById)
		appointmentGroup.GET("/date-range", r.appointmentQueryController.GetAppointmentsByDateRange)
		appointmentGroup.GET("/owner/:ownerId", r.appointmentQueryController.GetAppointmentsByOwner)
		appointmentGroup.GET("/vet/:vetId", r.appointmentQueryController.GetAppointmentsByVet)
		appointmentGroup.GET("/pet/:petId", r.appointmentQueryController.GetAppointmentsByPet)
		appointmentGroup.GET("/stats", r.appointmentQueryController.GetAppointmentStats)
	}

	// Owner-specific appointment routes
	ownerGroup := router.Group("/api/v2/owner")
	{
		// Owner appointment management
		ownerGroup.POST("/appointments", r.ownerAppointmentController.RequestAppointment)
		ownerGroup.GET("/appointments", r.ownerAppointmentController.GetMyAppointments)
		ownerGroup.GET("/appointments/:id", r.ownerAppointmentController.GetAppointmentById)
		ownerGroup.PUT("/appointments/:id/reschedule", r.ownerAppointmentController.RescheduleAppointment)
		ownerGroup.DELETE("/appointments/:id", r.ownerAppointmentController.CancelAppointment)

		// Pet-specific appointments for owners
		ownerGroup.GET("/pets/:petId/appointments", r.ownerAppointmentController.GetAppointmentsByPet)
	}

	// Veterinarian-specific appointment routes
	vetGroup := router.Group("/api/v2/vet")
	{
		// Vet appointment management
		vetGroup.GET("/appointments", r.vetAppointmentController.GetMyAppointments)
		vetGroup.GET("/appointments/today", r.vetAppointmentController.GetTodayAppointments)
		vetGroup.GET("/appointments/stats", r.vetAppointmentController.GetAppointmentStats)

		// Appointment state management
		vetGroup.PUT("/appointments/:id/confirm", r.vetAppointmentController.ConfirmAppointment)
		vetGroup.PUT("/appointments/:id/complete", r.vetAppointmentController.CompleteAppointment)
		vetGroup.PUT("/appointments/:id/reschedule", r.vetAppointmentController.RescheduleAppointment)
		vetGroup.PUT("/appointments/:id/no-show", r.vetAppointmentController.MarkAsNoShow)
		vetGroup.DELETE("/appointments/:id", r.vetAppointmentController.CancelAppointment)
	}
}

// RegisterOwnerRoutes registers only owner-specific appointment routes
func (r *AppointmentRoutes) RegisterOwnerRoutes(router *gin.Engine) {
	ownerGroup := router.Group("/api/v2/owner")
	{
		ownerGroup.POST("/appointments", r.ownerAppointmentController.RequestAppointment)
		ownerGroup.GET("/appointments", r.ownerAppointmentController.GetMyAppointments)
		ownerGroup.GET("/appointments/:id", r.ownerAppointmentController.GetAppointmentById)
		ownerGroup.PUT("/appointments/:id/reschedule", r.ownerAppointmentController.RescheduleAppointment)
		ownerGroup.DELETE("/appointments/:id", r.ownerAppointmentController.CancelAppointment)
		ownerGroup.GET("/pets/:petId/appointments", r.ownerAppointmentController.GetAppointmentsByPet)
	}
}

// RegisterVetRoutes registers only veterinarian-specific appointment routes
func (r *AppointmentRoutes) RegisterVetRoutes(router *gin.Engine) {
	vetGroup := router.Group("/api/v2/vet")
	{
		vetGroup.GET("/appointments", r.vetAppointmentController.GetMyAppointments)
		vetGroup.GET("/appointments/today", r.vetAppointmentController.GetTodayAppointments)
		vetGroup.GET("/appointments/stats", r.vetAppointmentController.GetAppointmentStats)
		vetGroup.PUT("/appointments/:id/confirm", r.vetAppointmentController.ConfirmAppointment)
		vetGroup.PUT("/appointments/:id/complete", r.vetAppointmentController.CompleteAppointment)
		vetGroup.PUT("/appointments/:id/reschedule", r.vetAppointmentController.RescheduleAppointment)
		vetGroup.PUT("/appointments/:id/no-show", r.vetAppointmentController.MarkAsNoShow)
		vetGroup.DELETE("/appointments/:id", r.vetAppointmentController.CancelAppointment)
	}
}
