package medHistoryRoutes

import (
	med_hist_controller "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/infrastructure/api/controller"
	"github.com/gin-gonic/gin"
)

func MedicalHistoryRoutes(router *gin.Engine, controller med_hist_controller.AdminMedicalHistoryController) {
	path := "/api/v2/admin/medical-history"
	router.GET(path, controller.SearchMedicalHistories)
	router.GET(path+"/:id", controller.GetMedicalHistory)
	router.GET(path+"/:id/details", controller.GetMedicalHistoryDetails)

	router.POST(path, controller.CreateMedicalHistories)
	router.PUT(path+"/:id", controller.UpdateMedicalHistories)
	router.DELETE(path+"/:id", controller.DeleteMedicalHistories)
}
