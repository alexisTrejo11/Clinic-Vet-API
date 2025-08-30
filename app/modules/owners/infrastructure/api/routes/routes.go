package routes

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/owners/infrastructure/api/controller"
	"github.com/gin-gonic/gin"
)

func OwnerRoutes(app *gin.Engine, ownerController *controller.OwnerController) {
	ownerV2 := app.Group("/api/v2/owners")
	ownerV2.GET("/:id", ownerController.GetOwnerById)
	ownerV2.GET("/", ownerController.ListOwners)
	ownerV2.POST("/", ownerController.CreateOwner)
	ownerV2.PATCH("/:id", ownerController.UpdateOwner)
	ownerV2.DELETE("/:id", ownerController.DeleteOwner)
}
