package routes

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/infrastructure/api/controller"
	"github.com/gin-gonic/gin"
)

func PetsRoutes(app *gin.Engine, petController *controller.PetController) {
	path := app.Group("api/v2/pets")
	path.GET("/", petController.SearchPets)
	path.GET("/:id", petController.GetPetByID)
	path.POST("/", petController.CreatePet)
	path.PATCH("/:id", petController.UpdatePet)
	path.DELETE("/:id", petController.SoftDeletePet)
}
