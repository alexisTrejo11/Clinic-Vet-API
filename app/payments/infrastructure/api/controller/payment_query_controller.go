package paymentController

import (
	"context"
	"strconv"
	"time"

	paymentQuery "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/application/queries"
	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	paymentDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/infrastructure/api/dtos"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	apiResponse "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PaymentQueryController struct {
	validator *validator.Validate
	queryBus  paymentQuery.QueryBus
}

func NewPaymentQueryController(
	validator *validator.Validate,
	paymentRepo paymentDomain.PaymentRepository,
	queryBus paymentQuery.QueryBus,

) *PaymentQueryController {
	return &PaymentQueryController{
		validator: validator,
		queryBus:  queryBus,
	}
}

func (c *PaymentQueryController) SearchPayments(ctx *gin.Context) {
	searchReq := c.parseSearchRequest(ctx)
	searchPaymentCommand := ToSearchComand(searchReq)

	payments, err := c.queryBus.Execute(context.TODO(), searchPaymentCommand)
	if err != nil {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	response := paymentDTOs.ToPaymentListResponse(payments)
	apiResponse.Success(ctx, response)
}

func (c *PaymentQueryController) GetPayment(ctx *gin.Context) {
	paymentId, err := shared.ParseID(ctx, ctx.Param("payment_id"))
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "payment_id", ctx.Param("payment_id"))
		return
	}

	getPaymentQuery := paymentQuery.NewGetPaymentByIdQuery(paymentId)
	payment, err := c.queryBus.Execute(context.TODO(), getPaymentQuery)
	if err != nil {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	response := paymentDTOs.ToPaymentResponse(payment)
	apiResponse.Success(ctx, response)
}

func (c *PaymentQueryController) GetPaymentsByUser(ctx *gin.Context) {
	userId, err := shared.ParseID(ctx, ctx.Param("user_id"))
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "user_id", ctx.Param("user_id"))
		return
	}

	listByUserQuery := paymentQuery.NewListPaymentsByUserQuery(userId, c.parsePagination(ctx))
	paymentPage, err := c.queryBus.Execute(context.TODO(), listByUserQuery)
	if err != nil {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	response := paymentDTOs.ToPaymentListResponse(paymentPage)
	apiResponse.Success(ctx, response)
}

func (c *PaymentQueryController) GetPaymentsByStatus(ctx *gin.Context) {
	statusStr := ctx.Param("status")
	status := paymentDomain.PaymentStatus(statusStr)
	if !status.IsValid() {
		apiResponse.RequestURLParamError(ctx, paymentDomain.ErrInvalidPaymentStatus, "status", statusStr)
		return
	}

	payments, err := c.queryBus.Execute(
		context.TODO(),
		paymentQuery.NewListPaymentsByStatusQuery(status, c.parsePagination(ctx)),
	)

	if err != nil {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	response := paymentDTOs.ToPaymentListResponse(payments)
	apiResponse.Success(ctx, response)
}

func (c *PaymentQueryController) GetOverduePayments(ctx *gin.Context) {
	pagination := c.parsePagination(ctx)
	paymentOverdueQuery := paymentQuery.NewListOverduePaymentsQuery(pagination)

	payments, err := c.queryBus.Execute(context.TODO(), paymentOverdueQuery)
	if err != nil {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	response := paymentDTOs.ToPaymentListResponse(payments)
	apiResponse.Success(ctx, response)
}

/* func (c *PaymentQueryController) GeneratePaymentReport(ctx *gin.Context) {
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
		apiResponse.ApplicationError(ctx, err)
		return
	}

	response := paymentDTOs.ToPaymentReportResponse(report)
	apiResponse.Success(ctx, response)
} */

func (c *PaymentQueryController) GetPaymentsByDateRange(ctx *gin.Context) {
	startDate, endDate, err := parseStartEndDates(ctx)
	if err != nil {
		apiResponse.RequestURLQueryError(ctx, err)
		return
	}

	paymentsByDateRangeQuery := paymentQuery.NewListPaymentsByDateRangeQuery(startDate, endDate, c.parsePagination(ctx))

	payments, err := c.queryBus.Execute(context.TODO(), paymentsByDateRangeQuery)
	if err != nil {
		apiResponse.ApplicationError(ctx, err)
		return
	}

	response := paymentDTOs.ToPaymentListResponse(payments)
	apiResponse.Success(ctx, response)
}

func (c *PaymentQueryController) parseSearchRequest(ctx *gin.Context) paymentDTOs.PaymentSearchRequest {
	req := paymentDTOs.PaymentSearchRequest{
		Page: c.parsePagination(ctx),
	}

	if UserIdStr := ctx.Query("owner_id"); UserIdStr != "" {
		if UserId, err := strconv.Atoi(UserIdStr); err == nil {
			req.UserId = &UserId
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

func ToSearchComand(req paymentDTOs.PaymentSearchRequest) paymentQuery.SearchPaymentsQuery {
	return paymentQuery.SearchPaymentsQuery{
		UserId:        req.UserId,
		AppointmentId: req.AppointmentId,
		Status:        req.Status,
		PaymentMethod: req.PaymentMethod,
		MinAmount:     req.MinAmount,
		MaxAmount:     req.MaxAmount,
		Currency:      req.Currency,
		StartDate:     req.StartDate,
		EndDate:       req.EndDate,
		Pagination:    req.Page,
	}
}

func parseStartEndDates(ctx *gin.Context) (time.Time, time.Time, error) {
	startDateStr := ctx.Query("start_date")
	endDateStr := ctx.Query("end_date")

	if startDateStr == "" || endDateStr == "" {
		return time.Time{}, time.Time{}, paymentDomain.NewPaymentError("MISSING_DATES", "start_date and end_date are required", 0, "")
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	return startDate, endDate, nil
}
