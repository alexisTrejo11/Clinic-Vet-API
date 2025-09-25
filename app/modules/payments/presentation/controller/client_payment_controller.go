package controller

import (
	"clinic-vet-api/app/middleware"
	autherror "clinic-vet-api/app/shared/error/auth"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ClientPaymentController struct {
	validator  *validator.Validate
	operations *PaymentControllerOperations
}

func NewClientPaymentController(validator *validator.Validate, operations *PaymentControllerOperations) *ClientPaymentController {
	return &ClientPaymentController{
		validator:  validator,
		operations: operations,
	}
}

func (ctrl *ClientPaymentController) GetMyPayments(c *gin.Context) {
	userCTX, ok := middleware.GetUserFromContext(c)
	if !ok {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	ctrl.operations.GetPaymentsByCustomer(c, userCTX.CustomerID)
}

func (ctrl *ClientPaymentController) GetMyPaymentByID(c *gin.Context) {
	userCTX, ok := middleware.GetUserFromContext(c)
	if !ok {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	ctrl.operations.GetPaymentByID(c, &userCTX.CustomerID)
}
