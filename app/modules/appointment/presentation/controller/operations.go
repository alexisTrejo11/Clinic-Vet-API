// Package controller handles all appointment-related HTTP endpoints
package controller

import (
	"context"

	"clinic-vet-api/app/modules/appointment/application/command"
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

	appointmentQuery := query.NewFindApptByIDQuery(
		c.Request.Context(),
		appointmentID,
		args.employeeID,
		args.customerID,
	)

	result, err := ctrl.bus.QueryBus.FindByID(*appointmentQuery)
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

	searchQuery, err := searchSpecification.ToQuery(c.Request.Context())
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	appointmentPage, err := ctrl.bus.QueryBus.FindBySpecification(*searchQuery)
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

	query, err := dateRangeQueryData.ToQuery(c.Request.Context())
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	appointmentPage, err := ctrl.bus.QueryBus.FindByDateRange(query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	ctrl.HandlePaginatedResult(c, appointmentPage, dateRangeQueryData.PageInput.ToMap())
}

func (ctrl *ApptControllerOperations) FindAppointmentsByCustomer(c *gin.Context, customerID uint, extraArgs customerQueryExtraArgs) {
	var pageParams page.PageInput
	if err := c.ShouldBindQuery(&pageParams); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	if err := ctrl.validate.Struct(&pageParams); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	query := query.NewFindApptsByCustomerIDQuery(c.Request.Context(), pageParams, customerID, extraArgs.PetID, extraArgs.Status)
	appointPage, err := ctrl.bus.QueryBus.FindByCustomerID(*query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	ctrl.HandlePaginatedResult(c, appointPage, pageParams.ToMap())
}

func (ctrl *ApptControllerOperations) GetAppointmentsByEmployee(c *gin.Context, employeeID uint) {
	var pageParams page.PageInput
	if err := c.ShouldBindQuery(&pageParams); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	query := query.NewFindApptsByEmployeeIDQuery(c.Request.Context(), employeeID, pageParams)
	appointmentPage, err := ctrl.bus.QueryBus.FindByEmployeeID(*query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	ctrl.HandlePaginatedResult(c, appointmentPage, pageParams.ToMap())
}

func (ctrl *ApptControllerOperations) GetAppointmentsByPet(c *gin.Context, petID uint) {
	var pageParams page.PageInput
	if err := c.ShouldBindQuery(&pageParams); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	query := query.NewFindApptsByPetQuery(c.Request.Context(), petID, pageParams)
	appointmentPage, err := ctrl.bus.QueryBus.FindByPetID(*query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	ctrl.HandlePaginatedResult(c, appointmentPage, pageParams.ToMap())
}

func (ctrl *ApptControllerOperations) GetAppointmentStats(c *gin.Context) {
}

func (ctrl *ApptControllerOperations) HandlePaginatedResult(c *gin.Context, pageResponse page.Page[query.ApptResult], queryParams map[string]any) {
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

	command, err := requestCreateData.ToCommand(c.Request.Context(), employeeID)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.CommandBus.CreateAppointment(*command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Created(c, result.ID(), "Appointment")
}

func (ctrl *ApptControllerOperations) UpdateAppointment(c *gin.Context) {
	var updateAppointData dto.UpdateApptRequest
	if err := ginUtils.BindAndValidateBody(c, &updateAppointData, ctrl.validate); err != nil {
		response.BadRequest(c, err)
		return
	}

	updateCommand, err := updateAppointData.ToCommand(context.TODO())
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.CommandBus.UpdateAppointment(updateCommand)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Updated(c, result.ToMap(), "Appointment")
}

func (ctrl *ApptControllerOperations) DeleteAppointment(c *gin.Context) {
	entityID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	command := command.NewDeleteApptCommand(entityID, c.Request.Context())
	result := ctrl.bus.CommandBus.DeleteAppointment(*command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, result.ToMap(), "Appointment deleted successfully")
}

func (ctrl *ApptControllerOperations) RescheduleAppointment(c *gin.Context, employeeID *uint) {
	var requestApptData dto.RescheduleApptRequest
	if err := c.ShouldBindJSON(&requestApptData); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	command, err := requestApptData.ToCommand(c.Request.Context(), employeeID)
	if err != nil {
		response.ApplicationError(c, err)
	}
	result := ctrl.bus.CommandBus.RescheduleAppointment(command)
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

	command := command.NewNotAttendApptCommand(c.Request.Context(), appointmentID, employeeID)
	result := ctrl.bus.CommandBus.MarkAppointmentAsNotAttend(*command)
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

	confirmApptCommand := command.NewConfirmAppointmentCommand(c.Request.Context(), appointmentID, employeeID)

	result := ctrl.bus.CommandBus.ConfirmAppointment(*confirmApptCommand)
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
	notes := c.Query("notes")

	completApptCommand := command.NewCompleteApptCommand(c.Request.Context(), appointmentID, employeeID, notes)
	result := ctrl.bus.CommandBus.CompleteAppointment(*completApptCommand)
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

	command := command.NewCancelApptCommand(c.Request.Context(), appointmentID, employeeID, reason)
	result := ctrl.bus.CommandBus.CancelAppointment(*command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, result.ToMap(), "Appointment cancelled successfully")
}
