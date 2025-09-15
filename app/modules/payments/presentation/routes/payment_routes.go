package routes

import (
	"clinic-vet-api/app/modules/payments/infrastructure/api/controller"
	"github.com/gin-gonic/gin"
)

func RegisterAdminPaymentRoutes(router *gin.Engine, controller *controller.AdminPaymentController) {
	adminGroup := router.Group("api/v2/admin/payments")
	{
		// Basic CRUD operations
		adminGroup.GET("/", controller.SearchPayments)
		adminGroup.POST("/", controller.CreatePayment)
		adminGroup.GET("/:id", controller.GetPayment)
		adminGroup.PUT("/:id", controller.UpdatePayment)
		adminGroup.DELETE("/:id", controller.DeletePayment)

		// Payment operations
		adminGroup.POST("/:id/process", controller.ProcessPayment)
		adminGroup.POST("/:id/refund", controller.RefundPayment)
		adminGroup.POST("/:id/cancel", controller.CancelPayment)

		// Batch operations
		adminGroup.POST("/mark-overdue", controller.MarkOverduePayments)

		// Query operations
		adminGroup.GET("/overdue", controller.GetOverduePayments)
		adminGroup.GET("/status/:status", controller.GetPaymentsByStatus)
		adminGroup.GET("/date-range", controller.GetPaymentsByDateRange)
		// adminGroup.GET("/report", controller.GeneratePaymentReport)
	}
}

func RegisterClientPaymentRoutes(router *gin.Engine, controller *controller.ClientPaymentController) {
	clientGroup := router.Group("api/v2/client/payments")
	{
		// Owner payment operations
		clientGroup.GET("/owners/:owner_id", controller.GetMyPayments)
		clientGroup.GET("/owners/:owner_id/history", controller.GetMyPaymentHistory)
		clientGroup.GET("/owners/:owner_id/overdue", controller.GetMyOverduePayments)
		clientGroup.GET("/owners/:owner_id/pending", controller.GetMyPendingPayments)
		clientGroup.GET("/:payment_id", controller.GetMyPayment)
	}
}

func RegisterPaymentQueryRoutes(router *gin.Engine, controller *controller.PaymentQueryController) {
	queryGroup := router.Group("api/v2/payments")
	{
		queryGroup.GET("/search", controller.SearchPayments)
		queryGroup.GET("/user/:user_id", controller.GetPaymentsByUser)
		queryGroup.GET("/status/:status", controller.GetPaymentsByStatus)
		queryGroup.GET("/overdue", controller.GetOverduePayments)
		// queryGroup.GET("/report", controller.GeneratePaymentReport)
		queryGroup.GET("/date-range", controller.GetPaymentsByDateRange)
	}
}

func RegisterPaymentCommandRoutes(router *gin.Engine, controller *controller.PaymentController) {
	commandGroup := router.Group("api/v2/payments")
	{
		commandGroup.POST("/", controller.CreatePayment)
		commandGroup.PUT("/:id", controller.UpdatePayment)
		commandGroup.DELETE("/:id", controller.DeletePayment)
		commandGroup.POST("/:id/process", controller.ProcessPayment)
		commandGroup.POST("/:id/refund", controller.RefundPayment)
		commandGroup.POST("/:id/cancel", controller.CancelPayment)
	}
}
