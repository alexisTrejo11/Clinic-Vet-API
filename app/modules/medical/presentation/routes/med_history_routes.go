// Package routes defines the API routes for the medical module.
package routes

import (
	"clinic-vet-api/app/modules/medical/presentation/controller"

	"github.com/gin-gonic/gin"
)

func MedicalHistoryRoutes(router *gin.Engine, controller controller.AdminMedicalHistoryController) {
	path := "/api/v2/admin/medical-history"
	router.GET(path, controller.SearchMedicalHistories)
	router.GET(path+"/:id", controller.GetMedicalHistoryDetails)
	router.POST(path, controller.CreateMedicalHistory)
	router.DELETE(path+"/:id", controller.DeleteMedicalHistory)
}
