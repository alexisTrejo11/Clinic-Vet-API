package controller

import (
	"strconv"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	query "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/payments/application/queries"
	dto "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/payments/infrastructure/api/dtos"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	httpError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure/http"
	ginUtils "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/gin_utils"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PaymentQueryController struct {
	validator *validator.Validate
	queryBus  cqrs.QueryBus
}

func NewPaymentQueryController(
	validator *validator.Validate,
	queryBus cqrs.QueryBus,
) *PaymentQueryController {
	return &PaymentQueryController{
		validator: validator,
		queryBus:  queryBus,
	}
}

func (c *PaymentQueryController) SearchPayments(ctx *gin.Context) {
	searchReq := c.parseSearchRequest(ctx)
	searchPaymentCommand := ToSearchComand(searchReq)

	payments, err := c.queryBus.Execute(searchPaymentCommand)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	paymentPage := dto.ToPaymentListResponse(payments)
	response.Success(ctx, paymentPage)
}

func (c *PaymentQueryController) GetPayment(ctx *gin.Context) {
	paymentID, err := ginUtils.ParseParamToInt(ctx, ctx.Param("payment_id"))
	if err != nil {
		response.BadRequest(ctx, httpError.RequestURLParamError(err, "payment_id", ctx.Param("payment_id")))
		return
	}

	getPaymentQuery, err := query.NewGetPaymentByIDQuery(paymentID)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	payment, err := c.queryBus.Execute(getPaymentQuery)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	paymentPage := dto.ToPaymentResponse(payment)
	response.Success(ctx, paymentPage)
}

func (c *PaymentQueryController) GetPaymentsByUser(ctx *gin.Context) {
	userID, err := ginUtils.ParseParamToInt(ctx, ctx.Param("user_id"))
	if err != nil {
		response.BadRequest(ctx, httpError.RequestURLParamError(err, "user_id", ctx.Param("user_id")))
		return
	}

	var pagination *page.PageInput
	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		response.BadRequest(ctx, httpError.RequestURLQueryError(err, ctx.Request.URL.RawQuery))
		return
	}

	if err := c.validator.Struct(pagination); err != nil {
		response.BadRequest(ctx, httpError.InvalidDataError(err))
		return
	}

	listByUserQuery := query.NewListPaymentsByUserQuery(userID, *pagination)
	paymentPage, err := c.queryBus.Execute(listByUserQuery)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	response.Success(ctx, paymentPage)
}

func (c *PaymentQueryController) GetPaymentsByStatus(ctx *gin.Context) {
	statusStr := ctx.Param("status")
	status, err := enum.ParsePaymentStatus(statusStr)
	if err != nil {
		response.BadRequest(ctx, httpError.RequestURLParamError(err, "status", statusStr))
		return
	}

	var pagination *page.PageInput
	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		response.BadRequest(ctx, httpError.RequestURLQueryError(err, ctx.Request.URL.RawQuery))
		return
	}

	if err := c.validator.Struct(pagination); err != nil {
		response.BadRequest(ctx, httpError.InvalidDataError(err))
		return
	}

	paymentPage, err := c.queryBus.Execute(query.NewListPaymentsByStatusQuery(status, *pagination))
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	response.Success(ctx, paymentPage)
}

func (c *PaymentQueryController) GetOverduePayments(ctx *gin.Context) {
	var pagination *page.PageInput
	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		response.BadRequest(ctx, httpError.RequestURLQueryError(err, ctx.Request.URL.RawQuery))
		return
	}

	if err := c.validator.Struct(pagination); err != nil {
		response.BadRequest(ctx, httpError.InvalidDataError(err))
		return
	}

	paymentOverdueQuery := query.NewListOverduePaymentsQuery(*pagination)

	paymentPage, err := c.queryBus.Execute(paymentOverdueQuery)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	response.Success(ctx, paymentPage)
}

/* func (c *PaymentQueryController) GeneratePaymentReport(ctx *gin.Context) {
	startDateStr := ctx.Query("start_date")
	endDateStr := ctx.Query("end_date")

	if startDateStr == "" || endDateStr == "" {
		response.RequestURLQueryError(ctx, paymentDomain.NewPaymentError("MISSING_DATES", "start_date and end_date are required", 0, ""))
		return
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		response.RequestURLQueryError(ctx, err)
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		response.RequestURLQueryError(ctx, err)
		return
	}

	report, err := c.paymentService.GeneratePaymentReport(startDate, endDate)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	response := dto.ToPaymentReportResponse(report)
	response.Success(ctx, response)
} */

func (c *PaymentQueryController) GetPaymentsByDateRange(ctx *gin.Context) {
	var pagination *page.PageInput
	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		response.BadRequest(ctx, httpError.RequestURLQueryError(err, ctx.Request.URL.RawQuery))
		return
	}

	if err := c.validator.Struct(pagination); err != nil {
		response.BadRequest(ctx, httpError.InvalidDataError(err))
		return
	}
	// TODO: parse actual date range from query params
	paymentsByDateRangeQuery := query.NewListPaymentsByDateRangeQuery(time.Now(), time.Now(), *pagination)

	paymentPage, err := c.queryBus.Execute(paymentsByDateRangeQuery)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	response.Success(ctx, paymentPage)
}

func (c *PaymentQueryController) parseSearchRequest(ctx *gin.Context) dto.PaymentSearchRequest {
	req := dto.PaymentSearchRequest{}

	if UserIDStr := ctx.Query("owner_id"); UserIDStr != "" {
		if UserID, err := strconv.Atoi(UserIDStr); err == nil {
			req.UserID = &UserID
		}
	}

	if appointmentIDStr := ctx.Query("appointment_id"); appointmentIDStr != "" {
		if appointmentID, err := strconv.Atoi(appointmentIDStr); err == nil {
			req.AppointmentID = &appointmentID
		}
	}

	if statusStr := ctx.Query("status"); statusStr != "" {
		status := enum.PaymentStatus(statusStr)
		if status.IsValid() {
			req.Status = &status
		}
	}

	if methodStr := ctx.Query("payment_method"); methodStr != "" {
		method := enum.PaymentMethod(methodStr)
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

func ToSearchComand(req dto.PaymentSearchRequest) query.SearchPaymentsQuery {
	return query.SearchPaymentsQuery{
		UserID:        req.UserID,
		AppointmentID: req.AppointmentID,
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
