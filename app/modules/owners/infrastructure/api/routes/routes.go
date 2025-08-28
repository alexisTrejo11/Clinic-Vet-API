package ownerRoutes

import (
	ownerController "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/infrastructure/api/controller"
	"github.com/gin-gonic/gin"
)

func OwnerRoutes(app *gin.Engine, ownerController *ownerController.OwnerController) {
	ownerV2 := app.Group("/api/v2/owners")
	ownerV2.GET("/:id", ownerController.GetOwnerById)
	ownerV2.GET("/", ownerController.ListOwners)
	ownerV2.POST("/", ownerController.CreateOwner)
	ownerV2.PATCH("/:id", ownerController.UpdateOwner)
	ownerV2.DELETE("/:id", ownerController.DeleteOwner)
}
