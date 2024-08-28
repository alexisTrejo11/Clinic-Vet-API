package routes

import (
	"example.com/at/backend/api-vet/controller"
	"github.com/gofiber/fiber/v2"
)

func AppointmentRoutes(app *fiber.App, appointmentController *controller.AppointmentController) {
	appointmentV1 := app.Group("v1/clinic-vet/appointment")
	appointmentV1.Post("/create", appointmentController.CreateAppointment())
	appointmentV1.Get("/:id", appointmentController.GetAppointmentById())
	appointmentV1.Put("/update", appointmentController.UpdateAnAppointment())
	appointmentV1.Delete("/remove/:id", appointmentController.DeleteAppointment())
}
