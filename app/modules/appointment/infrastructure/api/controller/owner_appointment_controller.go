// Package controller handles appointment-related HTTP endpoints
package controller

import (
	"errors"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/application/command"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/application/query"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/infrastructure/api/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	authError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/auth"
	httpError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure/http"
	ginUtils "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/gin_utils"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/response"
	"github.com/alexisTrejo11/Clinic-Vet-API/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// OwnerAppointmentController handles owner-specific appointment operations
// @title Veterinary Clinic API - Owner Appointment Management
// @version 1.0
// @description This controller manages appointment operations specific to pet owners including scheduling, rescheduling, and viewing appointments
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
type OwnerAppointmentController struct {
	commandBus cqrs.CommandBus
	queryBus   cqrs.QueryBus
	validator  *validator.Validate
}

func NewOwnerAppointmentController(commandBus cqrs.CommandBus, queryBus cqrs.QueryBus) *OwnerAppointmentController {
	return &OwnerAppointmentController{
		commandBus: commandBus,
		queryBus:   queryBus,
		validator:  &validator.Validate{},
	}
}

// RequestAppointment godoc
// @Summary Request a new appointment
// @Description Owner creates a new appointment request for their pet
// @Tags owner-appointments
// @Accept json
// @Produce json
// @Param appointment body command.CreateAppointmentCommand true "Appointment details"
// @Security BearerAuth
// @Router /owner/appointments [post]
func (controller *OwnerAppointmentController) RequestAppointment(c *gin.Context) {
	userCtx, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, authError.UnauthorizedCTXError())
		return
	}

	var requestData *dto.RequestOwnerAppointmentData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	requestData.Clean()

	if err := controller.validator.Struct(&requestData); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	createAppointCommand, err := requestData.ToCommand(userCtx.CustomerID)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := controller.commandBus.Execute(createAppointCommand)
	if !result.IsSuccess {
		response.ApplicationError(c, err)
		return
	}

	response.Created(c, result)
}

// GetMyAppointments godoc
// @Summary Get owner's appointments
// @Description Retrieves a list of all appointments for the authenticated owner
// @Tags owner-appointments
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Security BearerAuth
// @Router /owner/appointments [get]
func (controller *OwnerAppointmentController) GetMyAppointments(c *gin.Context) {
	userCtx, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, authError.UnauthorizedCTXError())
		return
	}

	var pageInput *page.PageInput
	if err := c.ShouldBindQuery(&pageInput); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	getByOwnerQuery, err := query.NewListAppointmentsByOwnerQuery(userCtx.CustomerID, *pageInput)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	appointmentPage, err := controller.queryBus.Execute(getByOwnerQuery)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Success(c, appointmentPage)
}

// GetAppointmentByID godoc
// @Summary Get specific appointment details
// @Description Retrieves details of a specific appointment for the authenticated owner
// @Tags owner-appointments
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID"
// @Security BearerAuth
// @Router /owner/appointments/{id} [get]
func (controller *OwnerAppointmentController) GetMyAppointmentDetail(c *gin.Context) {
	userCtx, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, authError.UnauthorizedCTXError())
		return
	}

	// TODO: Add GetByIdAndUserId
	getByIDQuery, err := query.NewGetAppointmentByIDQuery(userCtx.CustomerID)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	iAmppointment, err := controller.queryBus.Execute(getByIDQuery)
	appointmentDetail := iAmppointment.(query.AppointmentResponse)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	if appointmentDetail.OwnerID != userCtx.CustomerID {
		response.Forbidden(c, errors.New("forbidden"))
		return
	}

	response.Success(c, appointmentDetail)
}

// RescheduleAppointment godoc
// @Summary Reschedule an appointment
// @Description Owner reschedules their existing appointment
// @Tags owner-appointments
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID"
// @Param reschedule body command.RescheduleAppointmentCommand true "New appointment time"
// @Security BearerAuth
// @Router /owner/appointments/{id}/reschedule [put]
func (controller *OwnerAppointmentController) RescheduleAppointment(c *gin.Context) {
	// TODO: Implement Command
	_, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, authError.UnauthorizedCTXError())
		return
	}

	var requestData dto.OwnerRescheduleAppointData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	requestData.Clean()

	if err := controller.validator.Struct(&requestData); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	rescheduleAppointmentCommand, err := requestData.ToCommand(0)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := controller.commandBus.Execute(rescheduleAppointmentCommand)
	if !result.IsSuccess {
		response.ApplicationError(c, result.Error)
		return
	}

	response.Success(c, result)
}

// CancelAppointment - Owner cancels their appointment
// CancelAppointment godoc
// @Summary Cancel an appointment
// @Description Owner cancels their existing appointment
// @Tags owner-appointments
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID"
// @Security BearerAuth
// @Router /owner/appointments/{id}/cancel [put]
func (controller *OwnerAppointmentController) CancelAppointment(c *gin.Context) {
	appointmentID, err := ginUtils.ParseParamToInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	command, err := command.NewCancelAppointmentCommand(c.Request.Context(), appointmentID, nil, "Cancelled by admin")
	if err != nil {
		response.ApplicationError(c, err)
	}

	result := controller.commandBus.Execute(command)
	if !result.IsSuccess {
		response.ApplicationError(c, result.Error)
		return
	}

	response.Success(c, result)
}

// GetAppointmentsByPet godoc
// @Summary Get appointments for a specific pet
// @Description Retrieves all appointments for a specific pet owned by the authenticated owner
// @Tags owner-appointments
// @Accept json
// @Produce json
// @Param petID path int true "Pet ID"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Security BearerAuth
// @Router /owner/appointments/pet/{petID} [get]
func (controller *OwnerAppointmentController) GetAppointmentsByPet(c *gin.Context) {
	petID, err := ginUtils.ParseParamToInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	var pagination *page.PageInput
	if err := c.ShouldBindQuery(&pagination); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := controller.validator.Struct(&pagination); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	getApptByPetQuery, err := query.NewListAppointmentsByPetQuery(petID, *pagination)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	petResponsePage, err := controller.queryBus.Execute(getApptByPetQuery)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Success(c, petResponsePage)
}

// GetUpcomingAppointments godoc
// @Summary Get upcoming appointments
// @Description Retrieves upcoming appointments for the authenticated owner within a date range
// @Tags owner-appointments
// @Accept json
// @Produce json
// @Param start_date query string false "Start date (YYYY-MM-DD)" format(date)
// @Param end_date query string false "End date (YYYY-MM-DD)" format(date)
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Security BearerAuth
// @Router /owner/appointments/upcoming [get]
func (controller *OwnerAppointmentController) GetUpcomingAppointments(c *gin.Context) {
	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 30)

	if startDateStr := c.Query("start_date"); startDateStr != "" {
		if parsed, err := time.Parse("2006-01-02", startDateStr); err == nil {
			startDate = parsed
		}
	}
	if endDateStr := c.Query("end_date"); endDateStr != "" {
		if parsed, err := time.Parse("2006-01-02", endDateStr); err == nil {
			endDate = parsed
		}
	}

	var pageInput page.PageInput
	if err := c.ShouldBindQuery(&pageInput); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := controller.validator.Struct(&pageInput); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	listApptByDataRangeQuery, err := query.NewListAppointmentsByDateRangeQuery(startDate, endDate, pageInput)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	appointmentPage, err := controller.queryBus.Execute(listApptByDataRangeQuery)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Success(c, appointmentPage)
}
