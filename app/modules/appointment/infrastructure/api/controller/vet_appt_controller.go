// Package controller handles appointment-related HTTP endpoints
package controller

import (
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

// VetAppointmentController handles veterinarian-specific appointment operations
// @title Veterinary Clinic API - Veterinarian Appointment Management
// @version 1.0
// @description This controller manages appointment operations specific to veterinarians including viewing, confirming, completing, and managing appointments
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
type VetAppointmentController struct {
	commandBus cqrs.CommandBus
	queryBus   cqrs.QueryBus
	validator  *validator.Validate
}

func NewVetAppointmentController(commandBus cqrs.CommandBus, queryBus cqrs.QueryBus, validator *validator.Validate) *VetAppointmentController {
	return &VetAppointmentController{
		commandBus: commandBus,
		queryBus:   queryBus,
		validator:  validator,
	}
}

// GetMyAppointments godoc
// @Summary Get veterinarian's appointments
// @Description Retrieves all appointments assigned to the authenticated veterinarian
// @Tags vet-appointments
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Items per page" default(10)
// @Security BearerAuth
// @Success 200 {object} response.APIResponse "List of appointments"
// @Failure 400 {object} response.APIResponse "Invalid pagination parameters"
// @Failure 401 {object} response.APIResponse "Unauthorized - Veterinarian not authenticated"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /vet/appointments [get]
func (controller *VetAppointmentController) GetMyAppointments(c *gin.Context) {
	userCTX, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, authError.UnauthorizedCTXError())
		return
	}

	var pagination page.PageInput
	if err := c.ShouldBindQuery(&pagination); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	if err := controller.validator.Struct(&pagination); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	listApptByVetQuery := query.NewListApptByVetQuery(c.Request.Context(), userCTX.EmployeeID, pagination)
	result, err := controller.queryBus.Execute(listApptByVetQuery)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Success(c, result)
}

// CompleteAppointment godoc
// @Summary Complete an appointment
// @Description Marks an appointment as completed by the authenticated veterinarian
// @Tags vet-appointments
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID"
// @Security BearerAuth
// @Router /vet/appointments/{id}/complete [put]
func (controller *VetAppointmentController) CompleteAppointment(c *gin.Context) {
	userCTX, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, authError.UnauthorizedCTXError())
		return
	}

	appointmentID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "appointment-ID", c.Param("id")))
		return
	}
	notes := c.Query("notes")

	completeAppointmentCommand := command.NewCompleteApptCommand(c.Request.Context(), appointmentID, &userCTX.EmployeeID, notes)
	result := controller.commandBus.Execute(completeAppointmentCommand)
	if !result.IsSuccess {
		response.ApplicationError(c, result.Error)
		return
	}

	response.Success(c, result)
}

// CancelAppointment godoc
// @Summary Cancel an appointment
// @Description Cancels an appointment by the authenticated veterinarian
// @Tags vet-appointments
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID"
// @Security BearerAuth
// @Router /vet/appointments/{id} [delete]
func (controller *VetAppointmentController) CancelAppointment(c *gin.Context) {
	userCTX, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, authError.UnauthorizedCTXError())
		return
	}

	appointmentID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "appointment-ID", c.Param("id")))
		return
	}

	reason := c.Query("reason")

	cancelAppointmentCommand := command.NewCancelApptCommand(c.Request.Context(), appointmentID, &userCTX.EmployeeID, reason)
	result := controller.commandBus.Execute(cancelAppointmentCommand)
	if !result.IsSuccess {
		response.ApplicationError(c, result.Error)
		return
	}

	response.Success(c, result)
}

// ConfirmAppointment godoc
// @Summary Confirm an appointment
// @Description Confirms a pending appointment by the authenticated veterinarian
// @Tags vet-appointments
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID"
// @Security BearerAuth
// @Success 200 {object} response.APIResponse{message=string} "Appointment confirmed successfully"
// @Failure 400 {object} response.APIResponse "Invalid appointment ID"
// @Failure 401 {object} response.APIResponse "Unauthorized - Veterinarian not authenticated"
// @Failure 403 {object} response.APIResponse "Forbidden - Not assigned to this appointment"
// @Failure 404 {object} response.APIResponse "Appointment not found"
// @Failure 422 {object} response.APIResponse "Cannot confirm appointment"
// @Router /vet/appointments/{id}/confirm [put]
func (controller *VetAppointmentController) ConfirmAppointment(c *gin.Context) {
	userCTX, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, authError.UnauthorizedCTXError())
		return
	}

	appointmentID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "appointment-ID", c.Param("id")))
		return
	}

	vetConfirmApptCommand := command.NewConfirmAppointmentCommand(c.Request.Context(), appointmentID, userCTX.EmployeeID)
	result := controller.commandBus.Execute(vetConfirmApptCommand)
	if !result.IsSuccess {
		response.ApplicationError(c, result.Error)
		return
	}

	response.Success(c, result)
}

// MarkAsNoShow godoc
// @Summary Mark appointment as no-show
// @Description Marks an appointment as no-show when the client doesn't attend
// @Tags vet-appointments
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID"
// @Security BearerAuth
// @Success 200 {object} response.APIResponse{message=string} "Appointment marked as no-show"
// @Failure 400 {object} response.APIResponse "Invalid appointment ID"
// @Failure 401 {object} response.APIResponse "Unauthorized - Veterinarian not authenticated"
// @Failure 403 {object} response.APIResponse "Forbidden - Not assigned to this appointment"
// @Failure 404 {object} response.APIResponse "Appointment not found"
// @Failure 422 {object} response.APIResponse "Cannot mark as no-show"
// @Router /vet/appointments/{id}/no-show [put]
func (controller *VetAppointmentController) MarkAsNoShow(c *gin.Context) {
	appointmentID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "appointment-ID", c.Param("id")))
		return
	}
	userCTX, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, authError.UnauthorizedCTXError())
		return
	}

	vetApptCommand := command.NewNotAttendApptCommand(c, appointmentID, &userCTX.EmployeeID)
	result := controller.commandBus.Execute(vetApptCommand)
	if !result.IsSuccess {
		response.ApplicationError(c, result.Error)
		return
	}

	response.Success(c, result)
}

// GetAppointmentStats godoc
// @Summary Get appointment statistics
// @Description Retrieves statistical information about appointments for the authenticated veterinarian
// @Tags vet-appointments
// @Accept json
// @Produce json
// @Param start_date query string false "Start date (YYYY-MM-DD)" format(date)
// @Param end_date query string false "End date (YYYY-MM-DD)" format(date)
// @Security BearerAuth
// @Success 200 {object} response.APIResponse"Appointment statistics"
// @Failure 400 {object} response.APIResponse "Invalid date parameters"
// @Failure 401 {object} response.APIResponse "Unauthorized - Veterinarian not authenticated"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /vet/appointments/stats [get]
func (controller *VetAppointmentController) GetAppointmentStats(c *gin.Context) {
	/*
		// Get vet id from JWT context
		vetIDInterface, exists := c.Get("vet_id")
		if !exists {
			response.Unauthorized(c, errors.New("vet id not found in context"))
			return
		}

		vetID, ok := vetIDInterface.(int)
		if !ok {
			response.BadRequest(c, errors.New("invalid vet id format"))
			return
		}

		// Parse date range from query parameters (optional)
		var startDate, endDate *time.Time
		if startDateStr := c.Query("start_date"); startDateStr != "" {
			if parsed, err := time.Parse("2006-01-02", startDateStr); err == nil {
				startDate = &parsed
			}
		}
		if endDateStr := c.Query("end_date"); endDateStr != "" {
			if parsed, err := time.Parse("2006-01-02", endDateStr); err == nil {
				endDate = &parsed
			}
		}

		query := query.NewGetAppointmentStatsQuery(&vetID, nil, startDate, endDate)
		result, err := c.queryBus.Execute(context.Background(), query)
		if err != nil {
			response.ApplicationError(c, err)
			return
		}

		response.Success(c, result)
	*/
}

// RescheduleAppointment godoc
// @Summary Reschedule an appointment
// @Description Allows a veterinarian to reschedule their assigned appointment
// @Tags vet-appointments
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID"
// @Param reschedule body command.RescheduleAppointmentCommand true "New appointment time details"
// @Security BearerAuth
// @Success 200 {object} response.APIResponse{message=string} "Appointment rescheduled successfully"
// @Failure 400 {object} response.APIResponse "Invalid input data"
// @Failure 401 {object} response.APIResponse "Unauthorized - Veterinarian not authenticated"
// @Failure 403 {object} response.APIResponse "Forbidden - Not assigned to this appointment"
// @Failure 404 {object} response.APIResponse "Appointment not found"
// @Failure 422 {object} response.APIResponse "Invalid time slot or scheduling conflict"
// @Router /vet/appointments/{id}/reschedule [put]
func (controller *VetAppointmentController) RescheduleAppointment(c *gin.Context) {
	userCTX, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, authError.UnauthorizedCTXError())
		return
	}

	appointmentID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "appointment-ID", c.Param("id")))
		return
	}

	var requestData *dto.RescheduleAppointData
	if err := c.ShouldBindQuery(&requestData); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	if err := controller.validator.Struct(&requestData); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	vetRescheduleApptCommand := requestData.ToCommandWithVetID(c.Request.Context(), userCTX.EmployeeID, appointmentID)
	result := controller.commandBus.Execute(vetRescheduleApptCommand)
	if !result.IsSuccess {
		response.ApplicationError(c, result.Error)
		return
	}

	response.Success(c, result)
}
