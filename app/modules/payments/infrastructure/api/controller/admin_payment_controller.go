package controller

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/payments/application/command"
	apiResponse "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AdminPaymentController struct {
	validator         *validator.Validate
	commandBus        command.CommandBus
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
func (c *AdminPaymentController) SearchPayments(ctx *gin.Context) {
	c.queryController.SearchPayments(ctx)
}

// CreatePayment delegates to payment controller
func (c *AdminPaymentController) CreatePayment(ctx *gin.Context) {
	c.paymentController.CreatePayment(ctx)
}

// GetPayment delegates to payment controller
func (c *AdminPaymentController) GetPayment(ctx *gin.Context) {
	c.queryController.GetPayment(ctx)
}

// UpdatePayment delegates to payment controller
func (c *AdminPaymentController) UpdatePayment(ctx *gin.Context) {
	c.paymentController.UpdatePayment(ctx)
}

// DeletePayment delegates to payment controller
func (c *AdminPaymentController) DeletePayment(ctx *gin.Context) {
	c.paymentController.DeletePayment(ctx)
}

// ProcessPayment delegates to payment controller
func (c *AdminPaymentController) ProcessPayment(ctx *gin.Context) {
	c.paymentController.ProcessPayment(ctx)
}

// RefundPayment delegates to payment controller
func (c *AdminPaymentController) RefundPayment(ctx *gin.Context) {
	c.paymentController.RefundPayment(ctx)
}

// CancelPayment delegates to payment controller
func (c *AdminPaymentController) CancelPayment(ctx *gin.Context) {
	c.paymentController.CancelPayment(ctx)
}

// MarkOverduePayments marks all overdue payments
func (c *AdminPaymentController) MarkOverduePayments(ctx *gin.Context) {
	commandResult := c.commandBus.Execute(context.TODO(), command.MarkOverduePaymentsCommand{})
	if !commandResult.IsSuccess {
		apiResponse.ApplicationError(ctx, commandResult.Error)
		return
	}

	apiResponse.Success(ctx, commandResult)
}

// GetOverduePayments delegates to query controller
func (c *AdminPaymentController) GetOverduePayments(ctx *gin.Context) {
	c.queryController.GetOverduePayments(ctx)
}

// GeneratePaymentReport delegates to query controller
/* func (c *AdminPaymentController) GeneratePaymentReport(ctx *gin.Context) {
	c.queryController.GeneratePaymentReport(ctx)
} */

// GetPaymentsByDateRange delegates to query controller
func (c *AdminPaymentController) GetPaymentsByDateRange(ctx *gin.Context) {
	c.queryController.GetPaymentsByDateRange(ctx)
}

// GetPaymentsByStatus delegates to query controller
func (c *AdminPaymentController) GetPaymentsByStatus(ctx *gin.Context) {
	c.queryController.GetPaymentsByStatus(ctx)
}
