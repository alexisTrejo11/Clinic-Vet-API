// Package routes contains the route definitions for payment-related endpoints.+
package routes

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/payments/presentation/controller"

	"github.com/gin-gonic/gin"
)

type PaymentRoutes struct {
	adminController    *controller.AdminPaymentController
	customerController *controller.ClientPaymentController
}

func NewPaymentRoutes(
	adminController *controller.AdminPaymentController,
	customerController *controller.ClientPaymentController,
) *PaymentRoutes {
	return &PaymentRoutes{
		adminController:    adminController,
		customerController: customerController,
	}
}

func (r *PaymentRoutes) RegisterAdminPaymentRoutes(router *gin.RouterGroup, authMiddleware *middleware.AuthMiddleware) {
	adminGroup := router.Group("/payments")
	adminGroup.Use(authMiddleware.Authenticate())
	adminGroup.Use(authMiddleware.RequireAnyRole("recepetionist", "admin"))
	{
		// Basic CRUD operations
		adminGroup.GET("/", r.adminController.SearchPayments)
		adminGroup.POST("/", r.adminController.CreatePayment)
		adminGroup.GET("/:id", r.adminController.GetPayment)
		adminGroup.PUT("/:id", r.adminController.UpdatePayment)
		adminGroup.DELETE("/:id", r.adminController.DeletePayment)

		// Payment operations
		adminGroup.POST("/:id/process", r.adminController.ProcessPayment)
		adminGroup.POST("/:id/refund", r.adminController.RefundPayment)
		adminGroup.POST("/:id/cancel", r.adminController.CancelPayment)

		// Batch operations
		// adminGroup.POST("/mark-overdue", r.adminController.MarkOverduePayments)

		// Query operations
		adminGroup.GET("/overdue", r.adminController.GetOverduePayments)
		adminGroup.GET("/status/:status", r.adminController.GetPaymentsByStatus)
		adminGroup.GET("/date-range", r.adminController.GetPaymentsByDateRange)
		// adminGroup.GET("/report", r.adminController.GeneratePaymentReport)
	}
}

func (r *PaymentRoutes) RegisterClientRoutes(router *gin.RouterGroup, authnMiddleware *middleware.AuthMiddleware) {
	clientGroup := router.Group("customers/payments")
	clientGroup.Use(authnMiddleware.Authenticate())
	clientGroup.Use(authnMiddleware.RequireAnyRole("customer"))
	{
		clientGroup.GET("/customers/:customer_id", r.customerController.GetMyPayments)
	}
}
