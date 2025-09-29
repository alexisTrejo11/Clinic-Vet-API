// Package controller handles all appointment-related HTTP endpoints
package controller

import (
	"clinic-vet-api/app/modules/appointment/application/command"
	"clinic-vet-api/app/modules/appointment/application/handler"
	"clinic-vet-api/app/modules/appointment/application/query"
	"clinic-vet-api/app/modules/appointment/infrastructure/bus"
	"clinic-vet-api/app/modules/appointment/presentation/dto"
	httpError "clinic-vet-api/app/shared/error/infrastructure/http"
	ginUtils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/page"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ApptControllerOperations struct {
	bus      bus.AppointmentBus
	validate *validator.Validate
	mapper   dto.ResponseMapper
}

func NewApptControllerOperations(bus bus.AppointmentBus, validate *validator.Validate) *ApptControllerOperations {
	return &ApptControllerOperations{
		bus:      bus,
		validate: validate,
		mapper:   dto.ResponseMapper{},
	}
}

type GetByIDExtraArgs struct {
	employeeID *uint
	customerID *uint
}

func (ctrl *ApptControllerOperations) GetAppointmentDetailByID(c *gin.Context, args GetByIDExtraArgs) {
	appointmentID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	appointmentQuery := query.NewFindApptByIDQuery(appointmentID, args.employeeID, args.customerID)
	result, err := ctrl.bus.QueryBus.FindByID(c.Request.Context(), *appointmentQuery)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	appointmentResponse := ctrl.mapper.FromResult(result)
	response.Found(c, appointmentResponse, "Appointment")
}

func (ctrl *ApptControllerOperations) FindAppointmentsBySpecification(c *gin.Context) {
	searchSpecification, err := dto.NewApptSearchRequestFromContext(c)
	if err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	if err := ctrl.validate.Struct(searchSpecification); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	searchQuery, err := searchSpecification.ToQuery()
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	appointmentPage, err := ctrl.bus.QueryBus.FindBySpecification(c.Request.Context(), *searchQuery)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	ctrl.HandlePaginatedResult(c, appointmentPage, searchSpecification.Pagination())
}

func (ctrl *ApptControllerOperations) FindAppointmentsByDateRange(c *gin.Context) {
	var dateRangeQueryData dto.ListApptByDateRangeRequest
	if err := c.ShouldBindQuery(&dateRangeQueryData); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := ctrl.validate.Struct(&dateRangeQueryData); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	query, err := dateRangeQueryData.ToQuery()
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	appointmentPage, err := ctrl.bus.QueryBus.FindByDateRange(c.Request.Context(), query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	ctrl.HandlePaginatedResult(c, appointmentPage, dateRangeQueryData.PaginationRequest.ToMap())
}

func (ctrl *ApptControllerOperations) FindAppointmentsByCustomer(c *gin.Context, customerID uint, extraArgs customerQueryExtraArgs) {
	var pageParams page.PaginationRequest
	if err := ginUtils.ShouldBindPageParams(&pageParams, c, ctrl.validate); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	query := query.NewFindApptsByCustomerIDQuery(pageParams, customerID, extraArgs.PetID, extraArgs.Status)
	appointPage, err := ctrl.bus.QueryBus.FindByCustomerID(c.Request.Context(), *query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	ctrl.HandlePaginatedResult(c, appointPage, pageParams.ToMap())
}

func (ctrl *ApptControllerOperations) GetAppointmentsByEmployee(c *gin.Context, employeeID uint) {
	var pageParams page.PaginationRequest
	if err := c.ShouldBindQuery(&pageParams); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	query := query.NewFindApptsByEmployeeIDQuery(employeeID, pageParams)
	appointmentPage, err := ctrl.bus.QueryBus.FindByEmployeeID(c.Request.Context(), *query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	ctrl.HandlePaginatedResult(c, appointmentPage, pageParams.ToMap())
}

func (ctrl *ApptControllerOperations) GetAppointmentsByPet(c *gin.Context, petID uint) {
	var pageParams page.PaginationRequest
	if err := c.ShouldBindQuery(&pageParams); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	query := query.NewFindApptsByPetQuery(petID, pageParams)
	appointmentPage, err := ctrl.bus.QueryBus.FindByPetID(c.Request.Context(), *query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	ctrl.HandlePaginatedResult(c, appointmentPage, pageParams.ToMap())
}

func (ctrl *ApptControllerOperations) GetAppointmentStats(c *gin.Context) {
	// TODO: Implement appointment stats logic or remove this function if not needed
}

func (ctrl *ApptControllerOperations) HandlePaginatedResult(c *gin.Context, pageResponse page.Page[handler.ApptResult], queryParams map[string]any) {
	appointmentsResponse := ctrl.mapper.FromResults(pageResponse.Items)
	metadata := gin.H{"page_meta": pageResponse.Metadata, "request_params": queryParams}
	response.SuccessWithMeta(c, appointmentsResponse, "Appointment ", metadata)
}

// Command Operations

func (ctrl *ApptControllerOperations) CreateAppointment(c *gin.Context, employeeID *uint) {
	var requestCreateData dto.CreateApptRequest
	if err := ginUtils.BindAndValidateBody(c, &requestCreateData, ctrl.validate); err != nil {
		response.BadRequest(c, err)
		return
	}

	command, err := requestCreateData.ToCommand(employeeID)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.CommandBus.CreateAppointment(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Created(c, result.ID(), "Appointment")
}

func (ctrl *ApptControllerOperations) UpdateAppointment(c *gin.Context, apptID uint) {
	var updateAppointData dto.UpdateApptRequest
	if err := ginUtils.BindAndValidateBody(c, &updateAppointData, ctrl.validate); err != nil {
		response.BadRequest(c, err)
		return
	}

	updateCommand, err := updateAppointData.ToCommand(apptID)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.CommandBus.UpdateAppointment(c.Request.Context(), updateCommand)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Updated(c, nil, "Appointment")
}

func (ctrl *ApptControllerOperations) DeleteAppointment(c *gin.Context) {
	entityID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	command, err := command.NewDeleteApptCommand(entityID)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.CommandBus.DeleteAppointment(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, "Appointment deleted successfully")
}

func (ctrl *ApptControllerOperations) RescheduleAppointment(c *gin.Context, employeeID *uint) {
	appointmentID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	var requestApptData dto.RescheduleAppointData
	if err := c.ShouldBindJSON(&requestApptData); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	command, err := requestApptData.ToCommand(appointmentID)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.CommandBus.RescheduleAppointment(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, result.ToMap(), "Appointment rescheduled successfully")
}

func (ctrl *ApptControllerOperations) NotAttend(c *gin.Context, employeeID *uint) {
	appointmentID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	command, err := command.NewNotAttendApptCommand(appointmentID, employeeID)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.CommandBus.MarkAppointmentAsNotAttend(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
}

func (ctrl *ApptControllerOperations) ConfirmAppointment(c *gin.Context, employeeID uint) {
	appointmentID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	command, err := command.NewConfirmAppointmentCommand(appointmentID, employeeID)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.CommandBus.ConfirmAppointment(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
}

func (ctrl *ApptControllerOperations) CompleteAppointment(c *gin.Context, employeeID *uint) {
	appointmentID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	command, err := command.NewCompleteApptCommand(appointmentID, employeeID)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.CommandBus.CompleteAppointment(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, result.ToMap(), "Appointment completed successfully")
}

func (ctrl *ApptControllerOperations) CancelAppointment(c *gin.Context, employeeID *uint) {
	appointmentID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}
	reason := c.Query("reason")

	command, err := command.NewCancelApptCommand(appointmentID, employeeID, reason)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.CommandBus.CancelAppointment(c.Request.Context(), *command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, result.ToMap(), "Appointment cancelled successfully")
}
