// Package controller handles all appointment-related HTTP endpoints
package controller

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/application/command"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/infrastructure/api/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	httpError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure/http"
	ginUtils "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/gin_utils"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// AppointmentCommandController handles appointment management operations
// @title Veterinary Clinic API - Appointment Management
// @version 1.0
// @description This controller manages veterinary appointments including creation, updates, rescheduling, and status changes
type AppointmentCommandController struct {
	commandBus cqrs.CommandBus
	validate   *validator.Validate
}

func NewAppointmentCommandController(
	commandBus cqrs.CommandBus,
	validate *validator.Validate,
) *AppointmentCommandController {
	return &AppointmentCommandController{
		commandBus: commandBus,
		validate:   validate,
	}
}

// CreateAppointment godoc
// @Summary Create a new veterinary appointment
// @Description Creates a new appointment for a pet with a veterinarian
// @Tags appointments
// @Accept json
// @Produce json
// @Param appointment body command.CreateAppointmentCommand true "Appointment details"
// @Failure 400 {object} response.APIResponse "Invalid input data"
// @Failure 422 {object} response.APIResponse "Business rule validation failed"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments [post]
func (controller *AppointmentCommandController) CreateAppointment(c *gin.Context) {
	var requestCreateData *dto.CreateApptRequest
	if err := c.ShouldBindJSON(&requestCreateData); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := controller.validate.Struct(&requestCreateData); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	createCommand, err := requestCreateData.ToCommand(c.Request.Context())
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	if err := c.ShouldBindJSON(&createCommand); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	result := controller.commandBus.Execute(createCommand)
	if !result.IsSuccess {
		response.ApplicationError(c, result.Error)
		return
	}

	response.Created(c, result)
}

// UpdateAppointment godoc
// @Summary Update an existing appointment
// @Description Updates the details of an existing veterinary appointment
// @Tags appointments
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID"
// @Param appointment body command.UpdateAppointmentCommand true "Updated appointment details"
// @Failure 400 {object} response.APIResponse "Invalid input data"
// @Failure 404 {object} response.APIResponse "Appointment not found"
// @Failure 422 {object} response.APIResponse "Business rule validation failed"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments/{id} [put]
func (controller *AppointmentCommandController) UpdateAppointment(c *gin.Context) {
	var updateAppointData *dto.UpdateApptRequest
	if err := c.ShouldBindJSON(&updateAppointData); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := controller.validate.Struct(&updateAppointData); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	updateCommand, err := updateAppointData.ToCommand(context.TODO())
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := controller.commandBus.Execute(updateCommand)
	if !result.IsSuccess {
		response.ApplicationError(c, result.Error)
		return
	}

	response.Success(c, result.ToMap())
}

// DeleteAppointment godoc
// @Summary Delete an appointment
// @Description Removes an appointment from the system
// @Tags appointments
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID"
// @Failure 400 {object} response.APIResponse "Invalid appointment ID"
// @Failure 404 {object} response.APIResponse "Appointment not found"
// @Failure 422 {object} response.APIResponse "Cannot delete appointment"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments/{id} [delete]
func (controller *AppointmentCommandController) DeleteAppointment(c *gin.Context) {
	entityID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	deleteCommand := command.NewDeleteApptCommand(entityID, c.Request.Context())
	result := controller.commandBus.Execute(deleteCommand)
	if !result.IsSuccess {
		response.ApplicationError(c, result.Error)
		return
	}

	response.Success(c, result.ToMap())
}

// RescheduleAppointment godoc
// @Summary Reschedule an appointment
// @Description Changes the date and time of an existing appointment
// @Tags appointments
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID"
// @Param reschedule body command.RescheduleAppointmentCommand true "New appointment time details"
// @Failure 400 {object} response.APIResponse "Invalid input data"
// @Failure 404 {object} response.APIResponse "Appointment not found"
// @Failure 422 {object} response.APIResponse "Invalid time slot or scheduling conflict"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments/{id}/reschedule [put]
func (controller *AppointmentCommandController) RescheduleAppointment(c *gin.Context) {
	var requestApptData dto.RescheduleApptRequest
	if err := c.ShouldBindJSON(&requestApptData); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	rescheduleCommand, err := requestApptData.ToCommand(c.Request.Context())
	if err != nil {
		response.ApplicationError(c, err)
	}

	result := controller.commandBus.Execute(rescheduleCommand)
	if !result.IsSuccess {
		response.ApplicationError(c, result.Error)
		return
	}

	response.Success(c, result.ToMap())
}

// MarkAsNoShow godoc
// @Summary Mark appointment as no-show
// @Description Marks an appointment as no-show when the client doesn't attend
// @Tags appointments
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID"
// @Failure 400 {object} response.APIResponse "Invalid appointment ID"
// @Failure 404 {object} response.APIResponse "Appointment not found"
// @Failure 422 {object} response.APIResponse "Cannot mark as no-show"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments/{id}/no-show [put]
func (controller *AppointmentCommandController) NotAttend(c *gin.Context) {
	appointmentID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	notAttendCommand := command.NewNotAttendApptCommand(c.Request.Context(), appointmentID, nil)
	result := controller.commandBus.Execute(notAttendCommand)
	if !result.IsSuccess {
		response.ApplicationError(c, result.Error)
		return
	}

	response.Success(c, result)
}

// ConfirmAppointment godoc
// @Summary Confirm an appointment
// @Description Confirms an appointment by a veterinarian
// @Tags appointments
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID"
// @Security BearerAuth
// @Failure 400 {object} response.APIResponse "Invalid appointment ID"
// @Failure 401 {object} response.APIResponse "Unauthorized - Veterinarian not authenticated"
// @Failure 403 {object} response.APIResponse "Forbidden - Not assigned to this appointment"
// @Failure 404 {object} response.APIResponse "Appointment not found"
// @Failure 422 {object} response.APIResponse "Cannot confirm appointment"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments/{id}/confirm [put]
func (controller *AppointmentCommandController) ConfirmAppointment(c *gin.Context) {
	appointmentID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	vetIDStr := c.Query("vet_id")
	if vetIDStr == "" {
		response.BadRequest(c, httpError.RequestURLParamError(errors.New("vetID is required"), "vet_id", vetIDStr))
		return
	}

	vetID, err := ginUtils.ParseParamToUInt(c, "vet_id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "vet_id", vetIDStr))
		return
	}

	confirmApptCommand := command.NewConfirmAppointmentCommand(c.Request.Context(), appointmentID, vetID)
	result := controller.commandBus.Execute(confirmApptCommand)
	if !result.IsSuccess {
		response.ApplicationError(c, result.Error)
		return
	}

	response.Success(c, result.ToMap())
}

// CompleteAppointment godoc
// @Summary Complete an appointment
// @Description Marks an appointment as completed and adds optional notes
// @Tags appointments
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID"
// @Param notes body CompleteAppointmentRequest true "Completion notes"
// @Security BearerAuth
// @Failure 400 {object} response.APIResponse "Invalid input"
// @Failure 401 {object} response.APIResponse "Unauthorized - Veterinarian not authenticated"
// @Failure 403 {object} response.APIResponse "Forbidden - Not assigned to this appointment"
// @Failure 404 {object} response.APIResponse "Appointment not found"
// @Failure 422 {object} response.APIResponse "Cannot complete appointment"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments/{id}/complete [put]
func (controller *AppointmentCommandController) CompleteAppointment(c *gin.Context) {
	appointmentID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}
	notes := c.Query("notes")

	completApptCommand := command.NewCompleteApptCommand(c.Request.Context(), appointmentID, nil, notes)
	result := controller.commandBus.Execute(completApptCommand)
	if !result.IsSuccess {
		response.ApplicationError(c, result.Error)
		return
	}

	response.Success(c, result)
}
