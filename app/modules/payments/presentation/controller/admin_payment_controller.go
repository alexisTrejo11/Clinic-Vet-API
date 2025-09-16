package controller

import (
	"clinic-vet-api/app/shared/cqrs"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AdminPaymentController struct {
	validator         *validator.Validate
	commandBus        cqrs.CommandBus
	queryController   *PaymentQueryController
	paymentController *PaymentController
}

func NewAdminPaymentController(
	validator *validator.Validate,
	queryController *PaymentQueryController,
	paymentController *PaymentController,
) *AdminPaymentController {
	return &AdminPaymentController{
		validator:         validator,
		queryController:   queryController,
		paymentController: paymentController,
	}
}

// SearchPayments delegates to query controller
func (ctrl *AdminPaymentController) SearchPayments(c *gin.Context) {
	ctrl.queryController.SearchPayments(c)
}

// CreatePayment delegates to payment controller
func (ctrl *AdminPaymentController) CreatePayment(c *gin.Context) {
	ctrl.paymentController.CreatePayment(c)
}

// GetPayment delegates to payment controller
func (ctrl *AdminPaymentController) GetPayment(c *gin.Context) {
	ctrl.queryController.GetPayment(c)
}

// UpdatePayment delegates to payment controller
func (ctrl *AdminPaymentController) UpdatePayment(c *gin.Context) {
	ctrl.paymentController.UpdatePayment(c)
}

// DeletePayment delegates to payment controller
func (ctrl *AdminPaymentController) DeletePayment(c *gin.Context) {
	ctrl.paymentController.DeletePayment(c)
}

// ProcessPayment delegates to payment controller
func (ctrl *AdminPaymentController) ProcessPayment(c *gin.Context) {
	ctrl.paymentController.ProcessPayment(c)
}

// RefundPayment delegates to payment controller
func (ctrl *AdminPaymentController) RefundPayment(c *gin.Context) {
	ctrl.paymentController.RefundPayment(c)
}

// CancelPayment delegates to payment controller
func (ctrl *AdminPaymentController) CancelPayment(c *gin.Context) {
	ctrl.paymentController.CancelPayment(c)
}

// GetOverduePayments delegates to query controller
func (ctrl *AdminPaymentController) GetOverduePayments(c *gin.Context) {
	ctrl.queryController.GetOverduePayments(c)
}

// GeneratePaymentReport delegates to query controller
/* func (c *AdminPaymentController) GeneratePaymentReport(c *gin.Context) {
	c.queryController.GeneratePaymentReport(c)
} */

// GetPaymentsByDateRange delegates to query controller
func (ctrl *AdminPaymentController) GetPaymentsByDateRange(c *gin.Context) {
	ctrl.queryController.GetPaymentsByDateRange(c)
}

// GetPaymentsByStatus delegates to query controller
func (ctrl *AdminPaymentController) GetPaymentsByStatus(c *gin.Context) {
	ctrl.queryController.GetPaymentsByStatus(c)
}
