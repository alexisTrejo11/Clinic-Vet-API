package paymentController

import (
	"context"

	paymentCmd "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/application/command"
	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	paymentDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/infrastructure/api/dtos"
	utils "github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	apiResponse "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PaymentController struct {
	validator      *validator.Validate
	commandBus     paymentCmd.CommandBus
	paymentService paymentDomain.PaymentService
}

func NewPaymentController(
	validator *validator.Validate,
	commandBus paymentCmd.CommandBus) *PaymentController {
	return &PaymentController{
		validator:  validator,
		commandBus: commandBus,
	}
}

// CreatePayment creates a new payment
func (c *PaymentController) CreatePayment(ctx *gin.Context) {
	var req paymentDTOs.CreatePaymentRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	if err := c.validator.Struct(&req); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	createCommand := req.ToCreatePaymentCommand()
	commandResult := c.commandBus.Execute(context.TODO(), createCommand)
	if !commandResult.IsSuccess {
		apiResponse.AppError(ctx, commandResult.Error)
		return
	}

	apiResponse.Created(ctx, commandResult)
}

// UpdatePayment updates an existing payment
func (c *PaymentController) UpdatePayment(ctx *gin.Context) {
	id, err := utils.ParseID(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "payment_id", ctx.Param("id"))
		return
	}

	var req paymentDTOs.UpdatePaymentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	if err := c.validator.Struct(&req); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	updatePaymentCommand := req.ToUpdatePaymentCommand(id)

	commandResult := c.commandBus.Execute(context.TODO(), updatePaymentCommand)
	if !commandResult.IsSuccess {
		apiResponse.AppError(ctx, commandResult.Error)
		return
	}

	apiResponse.Ok(ctx, commandResult)
}

// DeletePayment soft deletes a payment
func (c *PaymentController) DeletePayment(ctx *gin.Context) {
	id, err := utils.ParseID(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "payment_id", ctx.Param("id"))
		return
	}

	deleteCommand := paymentCmd.NewDeletePaymentCommand(id)

	commandResult := c.commandBus.Execute(context.TODO(), deleteCommand)
	if !commandResult.IsSuccess {
		apiResponse.AppError(ctx, commandResult.Error)
		return
	}

	apiResponse.NoContent(ctx)
}

// ProcessPayment processes a payment
func (c *PaymentController) ProcessPayment(ctx *gin.Context) {
	id, err := utils.ParseID(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "payment_id", ctx.Param("id"))
		return
	}

	var req paymentDTOs.ProcessPaymentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	if err := c.validator.Struct(&req); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	proccessPaymentCommand := paymentCmd.NewProcessPaymentCommand(id, req.TransactionId)

	commandResult := c.commandBus.Execute(context.TODO(), proccessPaymentCommand)
	if !commandResult.IsSuccess {
		apiResponse.AppError(ctx, commandResult.Error)
		return
	}

	apiResponse.Ok(ctx, commandResult)
}

// RefundPayment refunds a payment
func (c *PaymentController) RefundPayment(ctx *gin.Context) {
	id, err := utils.ParseID(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "payment_id", ctx.Param("id"))
		return
	}

	var req paymentDTOs.RefundPaymentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	if err := c.validator.Struct(&req); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	refundPaymentCommand := paymentCmd.NewRefundPaymentCommand(id, req.Reason)

	commandResult := c.commandBus.Execute(context.TODO(), refundPaymentCommand)
	if !commandResult.IsSuccess {
		apiResponse.AppError(ctx, commandResult.Error)
		return
	}

	apiResponse.Ok(ctx, commandResult)
}

// CancelPayment cancels a payment
func (c *PaymentController) CancelPayment(ctx *gin.Context) {
	id, err := utils.ParseID(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "payment_id", ctx.Param("id"))
		return
	}

	var req paymentDTOs.CancelPaymentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	if err := c.validator.Struct(&req); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	cancelPaymentCommand := paymentCmd.NewCancelPaymentCommand(id, req.Reason)

	commandResult := c.commandBus.Execute(context.TODO(), cancelPaymentCommand)
	if !commandResult.IsSuccess {
		apiResponse.AppError(ctx, commandResult.Error)
		return
	}

	apiResponse.Ok(ctx, commandResult)
}
