// Package controller contains the controllers for handling payment-related requests.
package controller

import (
	"clinic-vet-api/app/shared/error/infrastructure/http"
	ginutils "clinic-vet-api/app/shared/gin_utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AdminPaymentController struct {
	validator  *validator.Validate
	operations *PaymentControllerOperations
}

func NewAdminPaymentController(validator *validator.Validate, operations *PaymentControllerOperations) *AdminPaymentController {
	return &AdminPaymentController{
		validator:  validator,
		operations: operations,
	}
}

// SearchPayments delegates to query controller
func (ctrl *AdminPaymentController) SearchPayments(c *gin.Context) {
	ctrl.operations.SearchPayments(c)
}

// CreatePayment delegates to payment controller
func (ctrl *AdminPaymentController) CreatePayment(c *gin.Context) {
	ctrl.operations.CreatePayment(c)
}

// GetPayment godoc
// @Summary Get payment by ID
// @Description Retrieves a specific payment by its ID
// @Tags admin-payments
// @Accept json
// @Produce json
// @Param payment_id path int true "Payment ID"
// @Success 200 {object} response.SuccessResponse{data=dto.PaymentResponse} "Payment retrieved successfully"
// @Failure 400 {object} response.ErrorResponse "Bad request - Invalid payment ID"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 403 {object} response.ErrorResponse "Forbidden"
// @Failure 404 {object} response.ErrorResponse "Payment not found"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Security BearerAuth
// @Router /admin/payments/{id} [get]
func (ctrl *AdminPaymentController) GetPaymentByID(c *gin.Context) {
	ctrl.operations.GetPaymentByID(c, nil)
}

func (ctrl *AdminPaymentController) GetPaymentsByCustomer(c *gin.Context) {
	customerID, err := ginutils.ParseParamToUInt(c, "customer_id")
	if err != nil {
		http.RequestURLParamError(err, "customer_id", c.Param("customer_id"))
		return
	}

	ctrl.operations.GetPaymentsByCustomer(c, customerID)
}

// UpdatePayment delegates to payment controller
func (ctrl *AdminPaymentController) UpdatePayment(c *gin.Context) {
	ctrl.operations.UpdatePayment(c)
}

// DeletePayment delegates to payment controller
func (ctrl *AdminPaymentController) DeletePayment(c *gin.Context) {
	ctrl.operations.DeletePayment(c)
}

// ProcessPayment delegates to payment controller
func (ctrl *AdminPaymentController) ProcessPayment(c *gin.Context) {
	ctrl.operations.ProcessPayment(c)
}

// RefundPayment delegates to payment controller
func (ctrl *AdminPaymentController) RefundPayment(c *gin.Context) {
	ctrl.operations.RefundPayment(c)
}

// CancelPayment delegates to payment controller
func (ctrl *AdminPaymentController) CancelPayment(c *gin.Context) {
	ctrl.operations.CancelPayment(c)
}

// GetOverduePayments delegates to query controller
func (ctrl *AdminPaymentController) GetOverduePayments(c *gin.Context) {
	ctrl.operations.GetOverduePayments(c)
}

// GeneratePaymentReport delegates to query controller
/* func (c *AdminPaymentController) GeneratePaymentReport(c *gin.Context) {
	c.operations.GeneratePaymentReport(c)
} */

// GetPaymentsByDateRange delegates to query controller
func (ctrl *AdminPaymentController) GetPaymentsByDateRange(c *gin.Context) {
	ctrl.operations.GetPaymentsByDateRange(c)
}

// GetPaymentsByStatus delegates to query controller
func (ctrl *AdminPaymentController) GetPaymentsByStatus(c *gin.Context) {
	ctrl.operations.GetPaymentsByStatus(c)
}
