package routes

import (
	"example.com/at/backend/api-vet/controller"
	"github.com/gofiber/fiber/v2"
)

func OwnerRoutes(app *fiber.App, ownerController *controller.OwnerController) {
	ownerV1 := app.Group("/api/v1/owner")
	ownerV1.Post("/create", ownerController.CreateOwner())
	ownerV1.Get("/:id", ownerController.GetOwnerById())
	ownerV1.Patch("/update", ownerController.UpdateOwner())
	ownerV1.Delete("/remove/:id", ownerController.DeleteOwner())
}
