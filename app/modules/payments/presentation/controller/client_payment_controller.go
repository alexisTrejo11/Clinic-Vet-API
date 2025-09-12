package controller

import (
	autherror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/auth"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/response"
	"github.com/alexisTrejo11/Clinic-Vet-API/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ClientPaymentController struct {
	validator       *validator.Validate
	queryController *PaymentQueryController
}

func NewClientPaymentController(
	validator *validator.Validate,
	queryController *PaymentQueryController,
) *ClientPaymentController {
	return &ClientPaymentController{
		validator:       validator,
		queryController: queryController,
	}
}

// GetMyPayments retrieves payments for the authenticated owner
func (ctrl *ClientPaymentController) GetMyPayments(c *gin.Context) {
	userCTX, ok := middleware.GetUserFromContext(c)
	if !ok {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	ctrl.queryController.GetPaymentsByCustomer(c, userCTX.CustomerID)
}

// GetMyPayment retrieves a specific payment for the authenticated owner
func (ctrl *ClientPaymentController) GetMyPayment(c *gin.Context) {
}
