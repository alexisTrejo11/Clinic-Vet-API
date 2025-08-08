package paymentRoutes

import (
	paymentController "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/infrastructure/api/controller"
	"github.com/gin-gonic/gin"
)

func RegisterAdminPaymentRoutes(router *gin.Engine, controller *paymentController.AdminPaymentController) {
	adminGroup := router.Group("api/v2/admin/payments")
	adminGroup.GET("/", controller.SearchPayments)
	adminGroup.POST("/", controller.CreatePayment)
	adminGroup.GET("/:id", controller.GetPayment)
	adminGroup.PUT("/:id", controller.UpdatePayment)
	adminGroup.DELETE("/:id", controller.DeletePayment)
}
