package paymentController

import (
	"context"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	paymentDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/infrastructure/api/dtos"
	utils "github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	apiResponse "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// TODO: Retrieve id from JWT or session instead URL
type ClientPaymentController struct {
	validator       *validator.Validate
	paymentRepo     paymentDomain.PaymentRepository
	paymentService  paymentDomain.PaymentService
	queryController *PaymentQueryController
}

func NewClientPaymentController(
	validator *validator.Validate,
	paymentRepo paymentDomain.PaymentRepository,
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
	paymentId, err := utils.ParseID(ctx, "payment_id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "payment_id", ctx.Param("payment_id"))
		return
	}

	payment, err := c.paymentRepo.GetById(context.TODO(), paymentId)
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	response := paymentDTOs.ToPaymentResponse(payment)
	apiResponse.Ok(ctx, response)
}

// GetMyPaymentHistory retrieves payment history for the authenticated owner
func (c *ClientPaymentController) GetMyPaymentHistory(ctx *gin.Context) {
	c.queryController.GetPaymentsByUser(ctx)
}

// GetMyOverduePayments retrieves overdue payments for the authenticated owner
func (c *ClientPaymentController) GetMyOverduePayments(ctx *gin.Context) {
	searchReq := paymentDTOs.PaymentSearchRequest{
		Status: func() *paymentDomain.PaymentStatus {
			status := paymentDomain.OVERDUE
			return &status
		}(),
		Page: c.queryController.parsePagination(ctx),
	}

	criteria := searchReq.ToSearchCriteria()
	payments, err := c.paymentRepo.Search(context.TODO(), searchReq.Page, criteria)
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	response := paymentDTOs.ToPaymentListResponse(payments)
	apiResponse.Ok(ctx, response)
}

// GetMyPendingPayments retrieves pending payments for the authenticated owner
func (c *ClientPaymentController) GetMyPendingPayments(ctx *gin.Context) {
	searchReq := paymentDTOs.PaymentSearchRequest{
		Status: func() *paymentDomain.PaymentStatus {
			status := paymentDomain.PENDING
			return &status
		}(),
		Page: c.queryController.parsePagination(ctx),
	}

	criteria := searchReq.ToSearchCriteria()
	payments, err := c.paymentRepo.Search(context.TODO(), searchReq.Page, criteria)
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	response := paymentDTOs.ToPaymentListResponse(payments)
	apiResponse.Ok(ctx, response)
}
