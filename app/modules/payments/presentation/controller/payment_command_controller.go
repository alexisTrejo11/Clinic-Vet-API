package controller

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/payments/application/command"
	dto "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/payments/infrastructure/api/dtos"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	httpError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure/http"
	ginUtils "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/gin_utils"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PaymentController struct {
	validator  *validator.Validate
	commandBus cqrs.CommandBus
}

func NewPaymentController(
	validator *validator.Validate,
	commandBus cqrs.CommandBus,
) *PaymentController {
	return &PaymentController{
		validator:  validator,
		commandBus: commandBus,
	}
}

// CreatePayment creates a new payment
func (c *PaymentController) CreatePayment(ctx *gin.Context) {
	var req dto.CreatePaymentRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, httpError.RequestBodyDataError(err))
		return
	}

	if err := c.validator.Struct(&req); err != nil {
		response.BadRequest(ctx, httpError.InvalidDataError(err))
		return
	}

	createCommand, err := req.ToCreatePaymentCommand()
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	commandResult := c.commandBus.Execute(createCommand)
	if !commandResult.IsSuccess {
		response.ApplicationError(ctx, commandResult.Error)
		return
	}

	response.Created(ctx, commandResult)
}

// UpdatePayment updates an existing payment
func (c *PaymentController) UpdatePayment(ctx *gin.Context) {
	id, err := ginUtils.ParseParamToInt(ctx, "id")
	if err != nil {
		response.BadRequest(ctx, httpError.RequestURLParamError(err, "payment_id", ctx.Param("id")))
		return
	}

	var req dto.UpdatePaymentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, httpError.RequestBodyDataError(err))
		return
	}

	if err := c.validator.Struct(&req); err != nil {
		response.BadRequest(ctx, httpError.InvalidDataError(err))
		return
	}

	updatePaymentCommand, err := req.ToUpdatePaymentCommand(ctx.Request.Context(), id)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	commandResult := c.commandBus.Execute(updatePaymentCommand)
	if !commandResult.IsSuccess {
		response.ApplicationError(ctx, commandResult.Error)
		return
	}

	response.Success(ctx, commandResult)
}

// DeletePayment soft deletes a payment
func (c *PaymentController) DeletePayment(ctx *gin.Context) {
	id, err := ginUtils.ParseParamToInt(ctx, "id")
	if err != nil {
		response.BadRequest(ctx, httpError.RequestURLParamError(err, "payment_id", ctx.Param("id")))
		return
	}

	deleteCommand, err := command.NewDeletePaymentCommand(id)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	commandResult := c.commandBus.Execute(deleteCommand)
	if !commandResult.IsSuccess {
		response.ApplicationError(ctx, commandResult.Error)
		return
	}

	response.NoContent(ctx)
}

// ProcessPayment processes a payment
func (c *PaymentController) ProcessPayment(ctx *gin.Context) {
	id, err := ginUtils.ParseParamToInt(ctx, "id")
	if err != nil {
		response.BadRequest(ctx, httpError.RequestURLParamError(err, "payment_id", ctx.Param("id")))
		return
	}

	var req dto.ProcessPaymentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, httpError.RequestBodyDataError(err))
		return
	}

	if err := c.validator.Struct(&req); err != nil {
		response.BadRequest(ctx, httpError.InvalidDataError(err))
		return
	}

	proccessPaymentCommand, err := command.NewProcessPaymentCommand(id, req.TransactionID)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	commandResult := c.commandBus.Execute(proccessPaymentCommand)
	if !commandResult.IsSuccess {
		response.ApplicationError(ctx, commandResult.Error)
		return
	}

	response.Success(ctx, commandResult)
}

// RefundPayment refunds a payment
func (c *PaymentController) RefundPayment(ctx *gin.Context) {
	id, err := ginUtils.ParseParamToInt(ctx, "id")
	if err != nil {
		response.BadRequest(ctx, httpError.RequestURLParamError(err, "payment_id", ctx.Param("id")))
		return
	}

	var req dto.RefundPaymentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, httpError.RequestBodyDataError(err))
		return
	}

	if err := c.validator.Struct(&req); err != nil {
		response.BadRequest(ctx, httpError.InvalidDataError(err))
		return
	}

	refundPaymentCommand, err := command.NewRefundPaymentCommand(id, req.Reason)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}
	commandResult := c.commandBus.Execute(refundPaymentCommand)
	if !commandResult.IsSuccess {
		response.ApplicationError(ctx, commandResult.Error)
		return
	}

	response.Success(ctx, commandResult)
}

// CancelPayment cancels a payment
func (c *PaymentController) CancelPayment(ctx *gin.Context) {
	id, err := ginUtils.ParseParamToInt(ctx, "id")
	if err != nil {
		response.BadRequest(ctx, httpError.RequestURLParamError(err, "payment_id", ctx.Param("id")))
		return
	}

	var req dto.CancelPaymentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, httpError.RequestBodyDataError(err))
		return
	}

	if err := c.validator.Struct(&req); err != nil {
		response.BadRequest(ctx, httpError.InvalidDataError(err))
		return
	}

	cancelPaymentCommand, err := command.NewCancelPaymentCommand(id, req.Reason)
	if err != nil {
		response.BadRequest(ctx, err)
		return
	}

	commandResult := c.commandBus.Execute(cancelPaymentCommand)
	if !commandResult.IsSuccess {
		response.ApplicationError(ctx, commandResult.Error)
		return
	}

	response.Success(ctx, commandResult)
}
