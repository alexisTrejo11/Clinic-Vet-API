package routes

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/pet/presentation/controller"

	"github.com/gin-gonic/gin"
)

func PetsRoutes(appGroup *gin.RouterGroup, controller *controller.PetController, authMiddleware *middleware.AuthMiddleware) {
	petGroup := appGroup.Group("/pets")
	petGroup.Use(authMiddleware.Authenticate())
	petGroup.Use(authMiddleware.RequireAnyRole("admin"))

	petGroup.GET("/", controller.SearchPets)
	petGroup.GET("/:id", controller.FindPetByID)
	petGroup.POST("/", controller.CreatePet)
	petGroup.PATCH("/:id", controller.UpdatePet)
	petGroup.DELETE("/:id", controller.DeletePet)
}

func CustomerPetsRoutes(r *gin.RouterGroup, controller *controller.CustomerPetController, authMiddleware *middleware.AuthMiddleware) {
	customer := r.Group("/customers/pets")
	customer.Use(authMiddleware.Authenticate())
	customer.Use(authMiddleware.RequireAnyRole("customer"))

	customer.GET("/", controller.GetMyPets)
	customer.GET("/:id", controller.GetMyPetDetails)
	customer.POST("/", controller.RegisterNewPet)
	customer.PATCH("/:id", controller.UpdateMyPet)
	customer.DELETE("/:id", controller.DeleteMyPet)
}
