package routes

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/veterinarians/infrastructure/api/controller"
	"github.com/gin-gonic/gin"
)

func VetRoutes(router *gin.Engine, vetController *controller.VeterinarianController) {
	v2Router := router.Group("/api/v2/veterinarians")
	v2Router.GET("/:id", vetController.GetVeterinarianByID)
	v2Router.GET("/", vetController.SearchVeterinarians)
	v2Router.POST("/", vetController.CreateVeterinarian)
	v2Router.PATCH("/:id", vetController.UpdateVeterinarian)
	v2Router.DELETE("/:id", vetController.DeleteVeterinarian)
}
