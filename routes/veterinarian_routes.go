package routes

import (
	"example.com/at/backend/api-vet/controller"
	"github.com/gofiber/fiber/v2"
)

func VeterinarianRoutes(app *fiber.App, vetController *controller.VeterinarianController) {
	vetV1 := app.Group("/api/v1/vet")
	vetV1.Post("/create", vetController.CreateVeterinarian())
	vetV1.Get("/:id", vetController.GetVeterinarianById())
	vetV1.Patch("/update", vetController.UpdateVeterinarian())
	vetV1.Delete("/remove/:id", vetController.DeleteVeterinarian())
}
