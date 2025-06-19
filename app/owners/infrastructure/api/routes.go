package routes

import (
	"example.com/at/backend/api-vet/controller"
	"github.com/gofiber/fiber/v2"
)

func OwnerRoutes(app *fiber.App, ownerController *controller.OwnerController) {
	ownerV1 := app.Group("v1/clinic-vet/owner")
	ownerV1.Post("/create", ownerController.CreateOwner())
	ownerV1.Get("/:id", ownerController.GetOwnerById())
	ownerV1.Put("/update", ownerController.UpdateOwner())
	ownerV1.Delete("/remove/:id", ownerController.DeleteOwner())
}

func OwnerPetRoutes(app *fiber.App, ownerPetController *controller.OwnerPetController) {
	ownerV1 := app.Group("v1/clinic-vet/owner-pet")
	ownerV1.Post("/create", ownerPetController.AddPet())
	ownerV1.Get("/my-pets", ownerPetController.GetMyPets())
	ownerV1.Put("/update", ownerPetController.UpdatePet()) // To be Tested
	ownerV1.Delete("/remove/:id", ownerPetController.DeletePet())
}

func OwnerAppointmentRoutes(app *fiber.App, ownerAppController *controller.OwnerAppointmentController) {
	ownerV1 := app.Group("/v1/clinic-vet/owner-appointment")
	ownerV1.Post("/request", ownerAppController.RequestAnAppointment())
	ownerV1.Get("/my-appointments", ownerAppController.GetMyAppointments())
	ownerV1.Delete("/cancel/:id", ownerAppController.CancelAnAppointment())

}

func OwnerMedicalHistoryRoutes(app *fiber.App, ownerAppController *controller.ClientMedicalHistory) {
	ownerV1 := app.Group("v1/clinic-vet/owner-medical-history")
	ownerV1.Get("/my-pets", ownerAppController.GetMyPetsMedicalHistories())
	ownerV1.Get("/pet/:id", ownerAppController.GetMyPetsMedicalHistoryByPetID())

}
