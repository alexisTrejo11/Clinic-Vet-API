package paymentController

import (
	"context"

	paymentCmd "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/application/command"
	paymentDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/infrastructure/api/dtos"
	apiResponse "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AdminPaymentController struct {
	validator         *validator.Validate
	commandFacade     *paymentCmd.PaymentCommandFacade
	queryController   *PaymentQueryController
	paymentController *PaymentController
}

func NewAdminPaymentController(
	validator *validator.Validate,
	commandFacade *paymentCmd.PaymentCommandFacade,
	queryController *PaymentQueryController,
	paymentController *PaymentController,
) *AdminPaymentController {
	return &AdminPaymentController{
		validator:         validator,
		commandFacade:     commandFacade,
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
	c.paymentController.GetPayment(ctx)
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
	updatedCount, err := c.commandFacade.MarkOverduePayments(context.TODO())
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	response := paymentDTOs.MarkOverdueResponse{
		UpdatedCount: updatedCount,
		Message:      "Successfully marked overdue payments",
	}

	apiResponse.Ok(ctx, response)
}

// GetOverduePayments delegates to query controller
func (c *AdminPaymentController) GetOverduePayments(ctx *gin.Context) {
	c.queryController.GetOverduePayments(ctx)
}

// GeneratePaymentReport delegates to query controller
func (c *AdminPaymentController) GeneratePaymentReport(ctx *gin.Context) {
	c.queryController.GeneratePaymentReport(ctx)
}

// GetPaymentsByDateRange delegates to query controller
func (c *AdminPaymentController) GetPaymentsByDateRange(ctx *gin.Context) {
	c.queryController.GetPaymentsByDateRange(ctx)
}

// GetPaymentsByStatus delegates to query controller
func (c *AdminPaymentController) GetPaymentsByStatus(ctx *gin.Context) {
	c.queryController.GetPaymentsByStatus(ctx)
}
