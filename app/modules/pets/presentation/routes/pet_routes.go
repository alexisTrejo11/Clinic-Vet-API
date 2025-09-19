package routes

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/pets/presentation/controller"

	"github.com/gin-gonic/gin"
)

func PetsRoutes(r *gin.Engine, controller *controller.PetController, authMiddleware *middleware.AuthMiddleware) {
	pet := r.Group("/pets")
	pet.Use(authMiddleware.Authenticate())
	pet.Use(authMiddleware.RequireAnyRole("admin"))

	pet.GET("/", controller.SearchPets)
	pet.GET("/:id", controller.FindPetByID)
	pet.POST("/", controller.CreatePet)
	pet.PATCH("/:id", controller.UpdatePet)
	pet.DELETE("/:id", controller.DeletePet)
}

func CustomerPetsRoutes(r *gin.Engine, controller *controller.CustomerPetController, authMiddleware *middleware.AuthMiddleware) {
	customer := r.Group("/customers/pets")
	customer.Use(authMiddleware.Authenticate())
	customer.Use(authMiddleware.RequireAnyRole("customer"))

	customer.GET("/", controller.GetMyPets)
	customer.GET("/:id", controller.GetMyPetDetails)
	customer.POST("/", controller.RegisterNewPet)
	customer.PATCH("/:id", controller.UpdateMyPet)
	customer.DELETE("/:id", controller.DeleteMyPet)
}
