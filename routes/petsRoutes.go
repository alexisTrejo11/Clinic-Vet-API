package routes

import (
	"example.com/at/backend/api-vet/controller"
	"github.com/gofiber/fiber/v2"
)

func PetsRoutes(app *fiber.App, petController *controller.PetController) {
	petV1 := app.Group("/api/v1/pet")
	petV1.Post("/create", petController.CreatePet())
	petV1.Get("/:petId", petController.GetPetById())
	petV1.Put("/update", petController.UpdatePet())
	petV1.Get("/:petId", petController.DeletePet())
}
