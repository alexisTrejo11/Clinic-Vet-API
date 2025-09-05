package controller

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	dto "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/payments/infrastructure/api/dtos"
	utils "github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// TODO: Retrieve id from JWT or session instead URL
type ClientPaymentController struct {
	validator       *validator.Validate
	paymentRepo     repository.PaymentRepository
	queryController *PaymentQueryController
}

func NewClientPaymentController(
	validator *validator.Validate,
	paymentRepo repository.PaymentRepository,
	queryController *PaymentQueryController,
) *ClientPaymentController {
	return &ClientPaymentController{
		validator:       validator,
		paymentRepo:     paymentRepo,
		queryController: queryController,
	}
}

// GetMyPayments retrieves payments for the authenticated owner
func (c *ClientPaymentController) GetMyPayments(ctx *gin.Context) {
	c.queryController.GetPaymentsByUser(ctx)
}

// GetMyPayment retrieves a specific payment for the authenticated owner
func (c *ClientPaymentController) GetMyPayment(ctx *gin.Context) {
	paymentID, err := utils.ParseParamToInt(ctx, "payment_id")
	if err != nil {
		response.RequestURLParamError(ctx, err, "payment_id", ctx.Param("payment_id"))
		return
	}

	payment, err := c.paymentRepo.GetByID(context.TODO(), paymentID)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	paymentResponse := dto.ToPaymentResponse(payment)
	response.Success(ctx, paymentResponse)
}

// GetMyPaymentHistory retrieves payment history for the authenticated owner
func (c *ClientPaymentController) GetMyPaymentHistory(ctx *gin.Context) {
	c.queryController.GetPaymentsByUser(ctx)
}

// GetMyOverduePayments retrieves overdue payments for the authenticated owner
func (c *ClientPaymentController) GetMyOverduePayments(ctx *gin.Context) {
	searchReq := dto.PaymentSearchRequest{
		Status: func() *enum.PaymentStatus {
			status := enum.OVERDUE
			return &status
		}(),
		Page: c.queryController.parsePagination(ctx),
	}

	criteria := searchReq.ToSearchCriteria()
	payments, err := c.paymentRepo.Search(context.TODO(), searchReq.Page, criteria)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	response := dto.ToPaymentListResponse(payments)
	response.Success(ctx, response)
}

// GetMyPendingPayments retrieves pending payments for the authenticated owner
func (c *ClientPaymentController) GetMyPendingPayments(ctx *gin.Context) {
	searchReq := dto.PaymentSearchRequest{
		Status: func() *enum.PaymentStatus {
			status := enum.PENDING
			return &status
		}(),
		Page: c.queryController.parsePagination(ctx),
	}

	criteria := searchReq.ToSearchCriteria()
	payments, err := c.paymentRepo.Search(context.TODO(), searchReq.Page, criteria)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	response := dto.ToPaymentListResponse(payments)
	response.Success(ctx, response)
}
