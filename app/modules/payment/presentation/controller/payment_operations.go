package controller

import (
	"clinic-vet-api/app/modules/payment/application/command"
	"clinic-vet-api/app/modules/payment/application/query"
	"clinic-vet-api/app/modules/payment/infrastructure/bus"
	dto "clinic-vet-api/app/modules/payment/presentation/dtos"
	httpError "clinic-vet-api/app/shared/error/infrastructure/http"
	ginUtils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/page"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PaymentControllerOperations struct {
	validator *validator.Validate
	bus       *bus.PaymentBus
}

func NewPaymentControllerOperations(validator *validator.Validate, bus *bus.PaymentBus) *PaymentControllerOperations {
	return &PaymentControllerOperations{
		validator: validator,
		bus:       bus,
	}
}

// Query Operations

func (ctrl *PaymentControllerOperations) SearchPayments(c *gin.Context) {
	var searchRequestData *dto.PaymentSearchRequest
	if err := ginUtils.ShouldBindAndValidateQuery(c, &searchRequestData, ctrl.validator); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	query := searchRequestData.ToQuery()
	paymentsPage, err := ctrl.bus.QueryBus.FindBySpecification(c.Request.Context(), *query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	ctrl.HandlePaginationResult(c, paymentsPage)
}

func (ctrl *PaymentControllerOperations) GetPaymentByID(c *gin.Context, customerID *uint) {
	paymentID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	query := query.NewFindPaymentByIDQuery(paymentID, customerID)
	payment, err := ctrl.bus.QueryBus.FindByID(c.Request.Context(), *query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	paymentPage := dto.ToPaymentResponse(payment)
	response.Success(c, paymentPage, "Payment retrieved successfully")
}

func (ctrl *PaymentControllerOperations) GetPaymentsByCustomer(c *gin.Context, customerID uint) {
	var pagination page.PaginationRequest
	if err := ginUtils.ShouldBindPageParams(&pagination, c, ctrl.validator); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	query := query.NewFindPaymentsByCustomerQuery(customerID, pagination)
	paymentPage, err := ctrl.bus.QueryBus.FindByCustomer(c.Request.Context(), query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	ctrl.HandlePaginationResult(c, paymentPage)
}

func (ctrl *PaymentControllerOperations) GetPaymentsByStatus(c *gin.Context) {
	statusStr := c.Param("status")
	var pagination page.PaginationRequest
	if err := ginUtils.ShouldBindPageParams(&pagination, c, ctrl.validator); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	query := query.NewFindPaymentsByStatusQuery(statusStr, pagination)
	paymentPage, err := ctrl.bus.QueryBus.FindByStatus(c.Request.Context(), *query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	ctrl.HandlePaginationResult(c, paymentPage)
}

func (ctrl *PaymentControllerOperations) GetPaymentsByDateRange(c *gin.Context) {
	var requestData dto.PaymentsByDateRangeRequest
	if err := ginUtils.ShouldBindAndValidateQuery(c, &requestData, ctrl.validator); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	query := requestData.ToQuery()
	paymentPage, err := ctrl.bus.QueryBus.FindByDateRange(c.Request.Context(), *query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	ctrl.HandlePaginationResult(c, paymentPage)
}

func (ctrl *PaymentControllerOperations) GetOverduePayments(c *gin.Context) {
	var pagination page.PaginationRequest
	if err := ginUtils.ShouldBindPageParams(&pagination, c, ctrl.validator); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	query := query.NewFindOverduePaymentsQuery(pagination)
	paymentPage, err := ctrl.bus.QueryBus.FindOverdues(c.Request.Context(), query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	ctrl.HandlePaginationResult(c, paymentPage)
}

/* func (c *PaymentControllerOperations) GeneratePaymentReport(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	if startDateStr == "" || endDateStr == "" {
		response.RequestURLQueryError(c, paymentDomain.NewPaymentError("MISSING_DATES", "start_date and end_date are required", 0, ""))
		return
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		response.RequestURLQueryError(c, err)
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		response.RequestURLQueryError(c, err)
		return
	}

	report, err := c.paymentService.GeneratePaymentReport(startDate, endDate)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response := dto.ToPaymentReportResponse(report)
	response.Success(c, response)
} */

func (ctrl *PaymentControllerOperations) HandlePaginationResult(c *gin.Context, paymentsPage page.Page[query.PaymentResult]) {
	paymentsResponses := dto.ToPaymentListResponse(paymentsPage.Items)
	metadata := gin.H{"metadata": gin.H{"pagination": paymentsPage.Metadata, "requestURLParam": c.Request.URL.RawQuery}}

	response.SuccessWithMeta(c, paymentsResponses, "Payments retrieved successfully", metadata)
}

// Command Operations

// CreatePayment creates a new payment
func (ctrl *PaymentControllerOperations) CreatePayment(c *gin.Context) {
	var req dto.CreatePaymentRequest
	if err := ginUtils.ShouldBindAndValidateBody(c, &req, ctrl.validator); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	command := req.ToCommand()
	commandResult := ctrl.bus.CommandBus.CreatePayment(c.Request.Context(), command)
	if !commandResult.IsSuccess() {
		response.ApplicationError(c, commandResult.Error())
		return
	}

	response.Created(c, commandResult.ID(), "payment")
}

// UpdatePayment updates an existing payment
func (ctrl *PaymentControllerOperations) UpdatePayment(c *gin.Context) {
	id, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "payment_id", c.Param("id")))
		return
	}

	var req dto.UpdatePaymentRequest
	if err := ginUtils.ShouldBindAndValidateBody(c, &req, ctrl.validator); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	command := req.ToCommand(id)
	commandResult := ctrl.bus.CommandBus.UpdatePayment(c.Request.Context(), command)
	if !commandResult.IsSuccess() {
		response.ApplicationError(c, commandResult.Error())
		return
	}

	response.Updated(c, nil, "Payment")
}

func (ctrl *PaymentControllerOperations) DeletePayment(c *gin.Context) {
	id, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "payment_id", c.Param("id")))
		return
	}

	command := command.NewDeletePaymentCommand(id)
	commandResult := ctrl.bus.CommandBus.DeletePayment(c.Request.Context(), *command)
	if !commandResult.IsSuccess() {
		response.ApplicationError(c, commandResult.Error())
		return
	}

	response.NoContent(c)
}

func (ctrl *PaymentControllerOperations) ProcessPayment(c *gin.Context) {
	id, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "payment_id", c.Param("id")))
		return
	}

	var req dto.ProcessPaymentRequest
	if err := ginUtils.ShouldBindAndValidateBody(c, &req, ctrl.validator); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	command := command.NewProcessPaymentCommand(id, req.TransactionID)
	commandResult := ctrl.bus.CommandBus.ProcessPayment(c.Request.Context(), *command)
	if !commandResult.IsSuccess() {
		response.ApplicationError(c, commandResult.Error())
		return
	}

	response.Success(c, nil, commandResult.Message())
}

func (ctrl *PaymentControllerOperations) RefundPayment(c *gin.Context) {
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

	command := command.NewRefundPaymentCommand(id, req.Reason)
	commandResult := ctrl.bus.CommandBus.RefundPayment(c.Request.Context(), *command)
	if !commandResult.IsSuccess() {
		response.ApplicationError(c, commandResult.Error())
		return
	}

	response.Success(c, nil, commandResult.Message())
}

func (ctrl *PaymentControllerOperations) CancelPayment(c *gin.Context) {
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

	command := command.NewCancelPaymentCommand(id, req.Reason)
	commandResult := ctrl.bus.CommandBus.CancelPayment(c.Request.Context(), *command)
	if !commandResult.IsSuccess() {
		response.ApplicationError(c, commandResult.Error())
		return
	}

	response.Success(c, nil, commandResult.Message())
}
