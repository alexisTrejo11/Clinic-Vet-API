package controller

import (
	"clinic-vet-api/app/modules/payments/application/command"
	query "clinic-vet-api/app/modules/payments/application/queries"
	"clinic-vet-api/app/modules/payments/infrastructure/bus"
	dto "clinic-vet-api/app/modules/payments/presentation/dtos"
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

func NewPaymentControllerOperations(
	validator *validator.Validate,
	bus *bus.PaymentBus,
) *PaymentControllerOperations {
	return &PaymentControllerOperations{
		validator: validator,
		bus:       bus,
	}
}

// Query Operations

func (ctrl *PaymentControllerOperations) SearchPayments(c *gin.Context) {
	var searchRequestData *dto.PaymentSearchRequest
	if err := c.ShouldBindQuery(&searchRequestData); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	if err := ctrl.validator.Struct(searchRequestData); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	searchQuery, err := searchRequestData.ToQuery(c.Request.Context())
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	paymentsPage, err := ctrl.bus.QueryBus.FindBySpecification(*searchQuery)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	ctrl.HandlePaginationResult(c, paymentsPage)
}

func (ctrl *PaymentControllerOperations) GetPayment(c *gin.Context) {
	paymentID, err := ginUtils.ParseParamToUInt(c, c.Param("payment_id"))
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "payment_id", c.Param("payment_id")))
		return
	}

	query, err := query.NewFindPaymentByIDQuery(paymentID)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	payment, err := ctrl.bus.QueryBus.FindByID(*query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	paymentPage := dto.ToPaymentResponse(payment)
	response.Success(c, paymentPage, "Payment retrieved successfully")
}

func (ctrl *PaymentControllerOperations) GetPaymentsByCustomer(c *gin.Context, customerID uint) {
	var pagination *page.PageInput
	if err := c.ShouldBindQuery(&pagination); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	if err := ctrl.validator.Struct(pagination); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	pagination.SetDefaultsFieldsIfEmpty()

	query := query.NewFindPaymentsByCustomerQuery(c.Request.Context(), customerID, *pagination)
	paymentPage, err := ctrl.bus.QueryBus.FindByCustomer(query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	ctrl.HandlePaginationResult(c, paymentPage)
}

func (ctrl *PaymentControllerOperations) GetPaymentsByStatus(c *gin.Context) {
	statusStr := c.Param("status")
	var pagination *page.PageInput
	if err := c.ShouldBindQuery(&pagination); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	if err := ctrl.validator.Struct(pagination); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	pagination.SetDefaultsFieldsIfEmpty()

	query, err := query.NewFindPaymentsByStatusQuery(c.Request.Context(), statusStr, *pagination)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	paymentPage, err := ctrl.bus.QueryBus.FindByStatus(*query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	ctrl.HandlePaginationResult(c, paymentPage)
}

func (ctrl *PaymentControllerOperations) GetPaymentsByDateRange(c *gin.Context) {
	var requestData *dto.PaymentsByDateRangeRequest

	if err := c.ShouldBindQuery(&requestData); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	requestData.SetDefaultsFieldsIfEmpty()
	if err := ctrl.validator.Struct(requestData); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	query := requestData.ToQuery(c.Request.Context())

	paymentPage, err := ctrl.bus.QueryBus.FindByDateRange(*query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	ctrl.HandlePaginationResult(c, paymentPage)
}

func (ctrl *PaymentControllerOperations) GetOverduePayments(c *gin.Context) {
	var pagination *page.PageInput
	if err := c.ShouldBindQuery(&pagination); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	if err := ctrl.validator.Struct(pagination); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	paymentOverdueQuery := query.NewFindOverduePaymentsQuery(c.Request.Context(), *pagination)

	paymentPage, err := ctrl.bus.QueryBus.FindOverdues(paymentOverdueQuery)
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
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := ctrl.validator.Struct(&req); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	createCommand, err := req.ToCreatePaymentCommand(c.Request.Context())
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	commandResult := ctrl.bus.CommandBus.CreatePayment(createCommand)
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
func (ctrl *PaymentControllerOperations) DeletePayment(c *gin.Context) {
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
func (ctrl *PaymentControllerOperations) ProcessPayment(c *gin.Context) {
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

	refundPaymentCommand := command.NewRefundPaymentCommand(c.Request.Context(), id, req.Reason)

	commandResult := ctrl.bus.CommandBus.RefundPayment(*refundPaymentCommand)
	if !commandResult.IsSuccess() {
		response.ApplicationError(c, commandResult.Error())
		return
	}

	response.Success(c, nil, commandResult.Message())
}

// CancelPayment cancels a payment
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

	cancelPaymentCommand := command.NewCancelPaymentCommand(id, req.Reason)

	commandResult := ctrl.bus.CommandBus.CancelPayment(*cancelPaymentCommand)
	if !commandResult.IsSuccess() {
		response.ApplicationError(c, commandResult.Error())
		return
	}

	response.Success(c, nil, "Payment Successfully Cancelled")
}
