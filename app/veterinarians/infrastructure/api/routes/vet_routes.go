package vetRoutes

import (
	vetController "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/infrastructure/api/controller"
	"github.com/gin-gonic/gin"
)

func VetRoutes(router *gin.Engine, vetController *vetController.VeterinarianController) {
	v2Router := router.Group("/api/v2/veterinarians")
	v2Router.GET("/:id", vetController.GetVeterinarianById)
	v2Router.GET("/", vetController.ListVeterinarians)
	v2Router.POST("/", vetController.CreateVeterinarian)
	v2Router.PATCH("/:id", vetController.UpdateVeterinarian)
	v2Router.DELETE("/:id", vetController.DeleteVeterinarian)
}
