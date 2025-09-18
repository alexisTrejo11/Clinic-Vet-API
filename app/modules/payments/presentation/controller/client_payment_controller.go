package controller

import (
	"clinic-vet-api/app/middleware"
	autherror "clinic-vet-api/app/shared/error/auth"
	"clinic-vet-api/app/shared/response"

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

// GetMyPayments retrieves payments for the authenticated customer
func (ctrl *ClientPaymentController) GetMyPayments(c *gin.Context) {
	userCTX, ok := middleware.GetUserFromContext(c)
	if !ok {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	ctrl.queryController.GetPaymentsByCustomer(c, userCTX.CustomerID)
}
