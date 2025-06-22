package routes

import (
	petController "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/infrastructure/api/controller"
	"github.com/gin-gonic/gin"
)

func PetsRoutes(app *gin.Engine, petController *petController.PetController) {
	path := app.Group("api/v2/pets")
	path.GET("/", petController.ListPets)
	path.GET("/:id", petController.GetPetById)
	path.POST("/", petController.CreatePet)
	path.PATCH("/:id", petController.UpdatePet)
	path.DELETE("/:id", petController.DeletePet)
}
