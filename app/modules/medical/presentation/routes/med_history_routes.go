// Package routes defines the API routes for the medical module.
package routes

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/medical/presentation/controller"

	"github.com/gin-gonic/gin"
)

func MedicalHistoryRoutes(appGroup *gin.RouterGroup, controller controller.AdminMedicalHistoryController, authMiddleware *middleware.AuthMiddleware) {
	medGroup := appGroup.Group("/medical-histories")
	medGroup.Use(authMiddleware.Authenticate())
	medGroup.Use(authMiddleware.RequireAnyRole("admin", "employee", "receptionist"))

	medGroup.GET("", controller.SearchMedicalHistories)
	medGroup.GET("/:id", controller.GetMedicalHistoryDetails)
	medGroup.POST("", controller.CreateMedicalHistory)
	medGroup.DELETE("/:id", controller.SoftDeleteMedicalHistory)
}
