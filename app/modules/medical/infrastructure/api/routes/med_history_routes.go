package routes

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/medical/infrastructure/api/controller"
	"github.com/gin-gonic/gin"
)

func MedicalHistoryRoutes(router *gin.Engine, controller controller.AdminMedicalHistoryController) {
	path := "/api/v2/admin/medical-history"
	router.GET(path, controller.SearchMedicalHistories)
	router.GET(path+"/:id", controller.GetMedicalHistory)
	router.GET(path+"/:id/details", controller.GetMedicalHistoryDetails)

	router.POST(path, controller.CreateMedicalHistories)
	router.DELETE(path+"/:id", controller.DeleteMedicalHistories)
}
