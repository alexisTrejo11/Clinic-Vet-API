package paymentController

import (
	"context"
	"strconv"
	"time"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	paymentDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/infrastructure/api/dtos"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	apiResponse "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PaymentQueryController struct {
	validator   *validator.Validate
	paymentRepo paymentDomain.PaymentRepository
}

func NewPaymentQueryController(
	validator *validator.Validate,
	paymentRepo paymentDomain.PaymentRepository,

) *PaymentQueryController {
	return &PaymentQueryController{
		validator:   validator,
		paymentRepo: paymentRepo,
	}
}

func (c *PaymentQueryController) SearchPayments(ctx *gin.Context) {
	searchReq := c.parseSearchRequest(ctx)

	pagination := searchReq.Page
	criteria := searchReq.ToSearchCriteria()

	payments, err := c.paymentRepo.Search(context.TODO(), pagination, criteria)
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	response := paymentDTOs.ToPaymentListResponse(payments)
	apiResponse.Ok(ctx, response)
}

func (c *PaymentQueryController) GetPaymentsByOwner(ctx *gin.Context) {
	ownerIdStr := ctx.Param("owner_id")
	ownerId, err := strconv.Atoi(ownerIdStr)
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "owner_id", ownerIdStr)
		return
	}

	pagination := c.parsePagination(ctx)

	payments, err := c.paymentRepo.ListByUserId(context.TODO(), ownerId, pagination)
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	response := paymentDTOs.ToPaymentListResponse(payments)
	apiResponse.Ok(ctx, response)
}

func (c *PaymentQueryController) GetPaymentsByStatus(ctx *gin.Context) {
	statusStr := ctx.Param("status")
	status := paymentDomain.PaymentStatus(statusStr)

	if !status.IsValid() {
		apiResponse.RequestURLParamError(ctx, paymentDomain.ErrInvalidPaymentStatus, "status", statusStr)
		return
	}

	pagination := c.parsePagination(ctx)

	payments, err := c.paymentRepo.ListByStatus(context.TODO(), status, pagination)
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	response := paymentDTOs.ToPaymentListResponse(payments)
	apiResponse.Ok(ctx, response)
}

func (c *PaymentQueryController) GetOverduePayments(ctx *gin.Context) {
	pagination := c.parsePagination(ctx)

	payments, err := c.paymentRepo.ListOverduePayments(context.TODO(), pagination)
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	response := paymentDTOs.ToPaymentListResponse(payments)
	apiResponse.Ok(ctx, response)
}

func (c *PaymentQueryController) GetPaymentHistory(ctx *gin.Context) {
	ownerIdStr := ctx.Param("owner_id")
	ownerId, err := strconv.Atoi(ownerIdStr)
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "owner_id", ownerIdStr)
		return
	}

	payments, err := c.paymentService.GetPaymentHistory(ownerId)
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	response := paymentDTOs.ToPaymentListResponse(payments)
	apiResponse.Ok(ctx, response)
}

func (c *PaymentQueryController) GeneratePaymentReport(ctx *gin.Context) {
	startDateStr := ctx.Query("start_date")
	endDateStr := ctx.Query("end_date")

	if startDateStr == "" || endDateStr == "" {
		apiResponse.RequestURLQueryError(ctx, paymentDomain.NewPaymentError("MISSING_DATES", "start_date and end_date are required", 0, ""))
		return
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		apiResponse.RequestURLQueryError(ctx, err)
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		apiResponse.RequestURLQueryError(ctx, err)
		return
	}

	report, err := c.paymentService.GeneratePaymentReport(startDate, endDate)
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	response := paymentDTOs.ToPaymentReportResponse(report)
	apiResponse.Ok(ctx, response)
}

func (c *PaymentQueryController) GetPaymentsByDateRange(ctx *gin.Context) {
	startDateStr := ctx.Query("start_date")
	endDateStr := ctx.Query("end_date")

	if startDateStr == "" || endDateStr == "" {
		apiResponse.RequestURLQueryError(ctx, paymentDomain.NewPaymentError("MISSING_DATES", "start_date and end_date are required", 0, ""))
		return
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		apiResponse.RequestURLQueryError(ctx, err)
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		apiResponse.RequestURLQueryError(ctx, err)
		return
	}

	payments, err := c.paymentRepo.ListPaymentsByDateRange(startDate, endDate)
	if err != nil {
		apiResponse.AppError(ctx, err)
		return
	}

	response := paymentDTOs.ToPaymentListResponse(payments)
	apiResponse.Ok(ctx, response)
}

func (c *PaymentQueryController) parseSearchRequest(ctx *gin.Context) paymentDTOs.PaymentSearchRequest {
	req := paymentDTOs.PaymentSearchRequest{
		Page: c.parsePagination(ctx),
	}

	if ownerIdStr := ctx.Query("owner_id"); ownerIdStr != "" {
		if ownerId, err := strconv.Atoi(ownerIdStr); err == nil {
			req.OwnerId = &ownerId
		}
	}

	if appointmentIdStr := ctx.Query("appointment_id"); appointmentIdStr != "" {
		if appointmentId, err := strconv.Atoi(appointmentIdStr); err == nil {
			req.AppointmentId = &appointmentId
		}
	}

	if statusStr := ctx.Query("status"); statusStr != "" {
		status := paymentDomain.PaymentStatus(statusStr)
		if status.IsValid() {
			req.Status = &status
		}
	}

	if methodStr := ctx.Query("payment_method"); methodStr != "" {
		method := paymentDomain.PaymentMethod(methodStr)
		if method.IsValid() {
			req.PaymentMethod = &method
		}
	}

	if minAmountStr := ctx.Query("min_amount"); minAmountStr != "" {
		if minAmount, err := strconv.ParseFloat(minAmountStr, 64); err == nil {
			req.MinAmount = &minAmount
		}
	}

	if maxAmountStr := ctx.Query("max_amount"); maxAmountStr != "" {
		if maxAmount, err := strconv.ParseFloat(maxAmountStr, 64); err == nil {
			req.MaxAmount = &maxAmount
		}
	}

	if currency := ctx.Query("currency"); currency != "" {
		req.Currency = &currency
	}

	if startDateStr := ctx.Query("start_date"); startDateStr != "" {
		if startDate, err := time.Parse("2006-01-02", startDateStr); err == nil {
			req.StartDate = &startDate
		}
	}

	if endDateStr := ctx.Query("end_date"); endDateStr != "" {
		if endDate, err := time.Parse("2006-01-02", endDateStr); err == nil {
			req.EndDate = &endDate
		}
	}

	return req
}

func (c *PaymentQueryController) parsePagination(ctx *gin.Context) page.PageData {
	pageSize := 20            // default
	pageNumber := 1           // default
	sortDirection := page.ASC // default

	if pageSizeStr := ctx.Query("page_size"); pageSizeStr != "" {
		if ps, err := strconv.Atoi(pageSizeStr); err == nil && ps > 0 && ps <= 100 {
			pageSize = ps
		}
	}

	if pageNumberStr := ctx.Query("page_number"); pageNumberStr != "" {
		if pn, err := strconv.Atoi(pageNumberStr); err == nil && pn > 0 {
			pageNumber = pn
		}
	}

	if sortStr := ctx.Query("sort_direction"); sortStr != "" {
		if sortStr == "DESC" {
			sortDirection = page.DESC
		}
	}

	return page.PageData{
		PageSize:      pageSize,
		PageNumber:    pageNumber,
		SortDirection: sortDirection,
	}
}
