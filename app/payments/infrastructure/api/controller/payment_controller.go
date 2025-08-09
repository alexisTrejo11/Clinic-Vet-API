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
	commandFacade  *paymentCmd.PaymentCommandFacade
	paymentRepo    paymentDomain.PaymentRepository
	paymentService paymentDomain.PaymentService
}

func NewPaymentController(
	validator *validator.Validate,
	commandFacade *paymentCmd.PaymentCommandFacade,
	paymentRepo paymentDomain.PaymentRepository,
	paymentService paymentDomain.PaymentService,
) *PaymentController {
	return &PaymentController{
		validator:      validator,
		commandFacade:  commandFacade,
		paymentRepo:    paymentRepo,
		paymentService: paymentService,
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

	cmd := req.ToCreatePaymentCommand()
	payment, err := c.commandFacade.CreatePayment(context.TODO(), cmd)
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	response := paymentDTOs.ToPaymentResponse(payment)
	apiResponse.Created(ctx, response)
}

// GetPayment retrieves a payment by ID
func (c *PaymentController) GetPayment(ctx *gin.Context) {
	id, err := utils.ParseID(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "payment_id", ctx.Param("id"))
		return
	}

	payment, err := c.paymentRepo.GetById(context.TODO(), id)
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	response := paymentDTOs.ToPaymentResponse(payment)
	apiResponse.Ok(ctx, response)
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

	cmd := req.ToUpdatePaymentCommand(id)
	payment, err := c.commandFacade.UpdatePayment(context.TODO(), cmd)
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	response := paymentDTOs.ToPaymentResponse(payment)
	apiResponse.Ok(ctx, response)
}

// DeletePayment soft deletes a payment
func (c *PaymentController) DeletePayment(ctx *gin.Context) {
	id, err := utils.ParseID(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "payment_id", ctx.Param("id"))
		return
	}

	err = c.commandFacade.DeletePayment(context.TODO(), id)
	if err != nil {
		apiResponse.AppError(ctx, err)
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

	err = c.commandFacade.ProcessPayment(context.TODO(), id, req.TransactionId)
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	payment, err := c.paymentRepo.GetById(context.TODO(), id)
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	response := paymentDTOs.ToPaymentResponse(payment)
	apiResponse.Ok(ctx, response)
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

	err = c.commandFacade.RefundPayment(context.TODO(), id, req.Reason)
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	payment, err := c.paymentRepo.GetById(context.TODO(), id)
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	response := paymentDTOs.ToPaymentResponse(payment)
	apiResponse.Ok(ctx, response)
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

	err = c.commandFacade.CancelPayment(context.TODO(), id, req.Reason)
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	payment, err := c.paymentRepo.GetById(context.TODO(), id)
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	response := paymentDTOs.ToPaymentResponse(payment)
	apiResponse.Ok(ctx, response)
}
