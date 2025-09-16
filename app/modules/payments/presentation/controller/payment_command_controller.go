package controller

import (
	"clinic-vet-api/app/modules/payments/application/command"
	"clinic-vet-api/app/modules/payments/infrastructure/bus"
	dto "clinic-vet-api/app/modules/payments/presentation/dtos"
	httpError "clinic-vet-api/app/shared/error/infrastructure/http"
	ginUtils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PaymentController struct {
	validator *validator.Validate
	bus       *bus.PaymentBus
}

func NewPaymentController(
	validator *validator.Validate,
	bus *bus.PaymentBus,
) *PaymentController {
	return &PaymentController{
		validator: validator,
		bus:       bus,
	}
}

// CreatePayment creates a new payment
func (ctrl *PaymentController) CreatePayment(c *gin.Context) {
	var req dto.CreatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := ctrl.validator.Struct(&req); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	createCommand, err := req.ToCreatePaymentCommand()
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	commandResult := ctrl.bus.CommandBus.CreatePayment(createCommand)
	if !commandResult.IsSuccess() {
		response.ApplicationError(c, commandResult.Error())
		return
	}

	response.Created(c, gin.H{"id": commandResult.ID()}, commandResult.Message())
}

// UpdatePayment updates an existing payment
func (ctrl *PaymentController) UpdatePayment(c *gin.Context) {
	id, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "payment_id", c.Param("id")))
		return
	}

	var req dto.UpdatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := ctrl.validator.Struct(&req); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	updatePaymentCommand, err := req.ToUpdatePaymentCommand(c.Request.Context(), id)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	commandResult := ctrl.bus.CommandBus.UpdatePayment(updatePaymentCommand)
	if !commandResult.IsSuccess() {
		response.ApplicationError(c, commandResult.Error())
		return
	}

	response.Updated(c, nil, "Payment")
}

// DeletePayment soft deletes a payment
func (ctrl *PaymentController) DeletePayment(c *gin.Context) {
	id, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "payment_id", c.Param("id")))
		return
	}

	deleteCommand := command.NewDeletePaymentCommand(c.Request.Context(), id)

	commandResult := ctrl.bus.CommandBus.DeletePayment(*deleteCommand)
	if !commandResult.IsSuccess() {
		response.ApplicationError(c, commandResult.Error())
		return
	}

	response.NoContent(c)
}

// ProcessPayment processes a payment
func (ctrl *PaymentController) ProcessPayment(c *gin.Context) {
	id, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "payment_id", c.Param("id")))
		return
	}

	var req dto.ProcessPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := ctrl.validator.Struct(&req); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	proccessPaymentCommand, err := command.NewProcessPaymentCommand(id, req.TransactionID)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	commandResult := ctrl.bus.CommandBus.ProcessPayment(*proccessPaymentCommand)
	if !commandResult.IsSuccess() {
		response.ApplicationError(c, commandResult.Error())
		return
	}

	response.Success(c, nil, commandResult.Message())
}

// RefundPayment refunds a payment
func (ctrl *PaymentController) RefundPayment(c *gin.Context) {
	id, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "payment_id", c.Param("id")))
		return
	}

	var req dto.RefundPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := ctrl.validator.Struct(&req); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	refundPaymentCommand := command.NewRefundPaymentCommand(c.Request.Context(), id, req.Reason)

	commandResult := ctrl.bus.CommandBus.RefundPayment(*refundPaymentCommand)
	if !commandResult.IsSuccess() {
		response.ApplicationError(c, commandResult.Error())
		return
	}

	response.Success(c, nil, commandResult.Message())
}

// CancelPayment cancels a payment
func (ctrl *PaymentController) CancelPayment(c *gin.Context) {
	id, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "payment_id", c.Param("id")))
		return
	}

	var req dto.CancelPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := ctrl.validator.Struct(&req); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	cancelPaymentCommand := command.NewCancelPaymentCommand(id, req.Reason)

	commandResult := ctrl.bus.CommandBus.CancelPayment(*cancelPaymentCommand)
	if !commandResult.IsSuccess() {
		response.ApplicationError(c, commandResult.Error())
		return
	}

	response.Success(c, nil, "Payment Successfully Cancelled")
}
