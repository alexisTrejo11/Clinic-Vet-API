// Package appointmentController handles all appointment-related HTTP endpoints
package controller

import (
	"context"
	"errors"
	"net/http"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/application/command"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	apiResponse "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// AppointmentCommandController handles appointment management operations
// @title Veterinary Clinic API - Appointment Management
// @version 1.0
// @description This controller manages veterinary appointments including creation, updates, rescheduling, and status changes
type AppointmentCommandController struct {
	commandBus command.CommandBus
	validate   *validator.Validate
}

func NewAppointmentCommandController(
	commandBus command.CommandBus,
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
// @Failure 400 {object} apiResponse.APIResponse "Invalid input data"
// @Failure 422 {object} apiResponse.APIResponse "Business rule validation failed"
// @Failure 500 {object} apiResponse.APIResponse "Internal server error"
// @Router /appointments [post]
func (controller *AppointmentCommandController) CreateAppointment(ctx *gin.Context) {
	var command command.CreateAppointmentCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	result := controller.commandBus.Execute(ctx, command)
	if !result.IsSuccess {
		apiResponse.ApplicationError(ctx, result.Error)
		return
	}

	apiResponse.Created(ctx, gin.H{"message": result.Message, "appointment_id": result.ID})
}

// UpdateAppointment godoc
// @Summary Update an existing appointment
// @Description Updates the details of an existing veterinary appointment
// @Tags appointments
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID"
// @Param appointment body command.UpdateAppointmentCommand true "Updated appointment details"
// @Failure 400 {object} apiResponse.APIResponse "Invalid input data"
// @Failure 404 {object} apiResponse.APIResponse "Appointment not found"
// @Failure 422 {object} apiResponse.APIResponse "Business rule validation failed"
// @Failure 500 {object} apiResponse.APIResponse "Internal server error"
// @Router /appointments/{id} [put]
func (controller *AppointmentCommandController) UpdateAppointment(ctx *gin.Context) {
	idInt, err := shared.ParseParamToInt(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
		return
	}

	var command command.UpdateAppointmentCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.validate.Struct(command); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	command.AppointmentID = idInt
	result := controller.commandBus.Execute(ctx, command)
	if !result.IsSuccess {
		apiResponse.ApplicationError(ctx, result.Error)
		return
	}

	apiResponse.Success(ctx, result.Message)
}

// DeleteAppointment godoc
// @Summary Delete an appointment
// @Description Removes an appointment from the system
// @Tags appointments
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID"
// @Failure 400 {object} apiResponse.APIResponse "Invalid appointment ID"
// @Failure 404 {object} apiResponse.APIResponse "Appointment not found"
// @Failure 422 {object} apiResponse.APIResponse "Cannot delete appointment"
// @Failure 500 {object} apiResponse.APIResponse "Internal server error"
// @Router /appointments/{id} [delete]
func (controller *AppointmentCommandController) DeleteAppointment(ctx *gin.Context) {
	idInt, err := shared.ParseParamToInt(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
		return
	}

	command := command.NewDeleteAppointmentCommand(idInt)

	result := controller.commandBus.Execute(ctx, command)
	if !result.IsSuccess {
		apiResponse.ApplicationError(ctx, result.Error)
		return
	}

	apiResponse.Success(ctx, result.Message)
}

// RescheduleAppointment godoc
// @Summary Reschedule an appointment
// @Description Changes the date and time of an existing appointment
// @Tags appointments
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID"
// @Param reschedule body command.RescheduleAppointmentCommand true "New appointment time details"
// @Failure 400 {object} apiResponse.APIResponse "Invalid input data"
// @Failure 404 {object} apiResponse.APIResponse "Appointment not found"
// @Failure 422 {object} apiResponse.APIResponse "Invalid time slot or scheduling conflict"
// @Failure 500 {object} apiResponse.APIResponse "Internal server error"
// @Router /appointments/{id}/reschedule [put]
func (controller *AppointmentCommandController) RescheduleAppointment(ctx *gin.Context) {
	appointmentID, err := shared.ParseParamToInt(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
		return
	}

	var command command.RescheduleAppointmentCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	command.AppointmentID = appointmentID

	result := controller.commandBus.Execute(context.Background(), command)
	if !result.IsSuccess {
		apiResponse.ApplicationError(ctx, result.Error)
		return
	}

	apiResponse.Success(ctx, result.Message)
}

// MarkAsNoShow godoc
// @Summary Mark appointment as no-show
// @Description Marks an appointment as no-show when the client doesn't attend
// @Tags appointments
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID"
// @Failure 400 {object} apiResponse.APIResponse "Invalid appointment ID"
// @Failure 404 {object} apiResponse.APIResponse "Appointment not found"
// @Failure 422 {object} apiResponse.APIResponse "Cannot mark as no-show"
// @Failure 500 {object} apiResponse.APIResponse "Internal server error"
// @Router /appointments/{id}/no-show [put]
func (controller *AppointmentCommandController) MarkAsNoShow(ctx *gin.Context) {
	appointmentID, err := shared.ParseParamToInt(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
		return
	}

	command := command.MarkAsNotPresentedCommand{
		ID: appointmentID,
	}

	result := controller.commandBus.Execute(context.Background(), command)
	if !result.IsSuccess {
		apiResponse.ApplicationError(ctx, result.Error)
		return
	}

	apiResponse.Success(ctx, result)
}

// ConfirmAppointment godoc
// @Summary Confirm an appointment
// @Description Confirms an appointment by a veterinarian
// @Tags appointments
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID"
// @Security BearerAuth
// @Failure 400 {object} apiResponse.APIResponse "Invalid appointment ID"
// @Failure 401 {object} apiResponse.APIResponse "Unauthorized - Veterinarian not authenticated"
// @Failure 403 {object} apiResponse.APIResponse "Forbidden - Not assigned to this appointment"
// @Failure 404 {object} apiResponse.APIResponse "Appointment not found"
// @Failure 422 {object} apiResponse.APIResponse "Cannot confirm appointment"
// @Failure 500 {object} apiResponse.APIResponse "Internal server error"
// @Router /appointments/{id}/confirm [put]
func (controller *AppointmentCommandController) ConfirmAppointment(ctx *gin.Context) {
	appointmentID, err := shared.ParseParamToInt(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
		return
	}

	// Get vet id from context
	vetIDInterface, exists := ctx.Get("vet_id")
	if !exists {
		apiResponse.Unauthorized(ctx, errors.New("vet id not found in context"))
		return
	}

	vetIDInt, ok := vetIDInterface.(int)
	if !ok {
		apiResponse.BadRequest(ctx, errors.New("invalid vet id format"))
		return
	}

	vetID, err := valueobject.NewVetID(vetIDInt)
	if err != nil {
		apiResponse.BadRequest(ctx, err)
		return
	}

	command := command.ConfirmAppointmentCommand{
		ID:    appointmentID,
		VetID: &vetID,
	}

	result := controller.commandBus.Execute(context.Background(), command)
	if !result.IsSuccess {
		apiResponse.ApplicationError(ctx, result.Error)
		return
	}

	apiResponse.Success(ctx, result)
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
// @Failure 400 {object} apiResponse.APIResponse "Invalid input"
// @Failure 401 {object} apiResponse.APIResponse "Unauthorized - Veterinarian not authenticated"
// @Failure 403 {object} apiResponse.APIResponse "Forbidden - Not assigned to this appointment"
// @Failure 404 {object} apiResponse.APIResponse "Appointment not found"
// @Failure 422 {object} apiResponse.APIResponse "Cannot complete appointment"
// @Failure 500 {object} apiResponse.APIResponse "Internal server error"
// @Router /appointments/{id}/complete [put]
func (controller *AppointmentCommandController) CompleteAppointment(ctx *gin.Context) {
	appointmentID, err := shared.ParseParamToInt(ctx, "id")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "id", ctx.Param("id"))
		return
	}

	var requestBody *CompleteAppointmentRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		apiResponse.RequestBodyDataError(ctx, err)
		return
	}

	command := command.CompleteAppointmentCommand{
		ID:    appointmentID,
		Notes: requestBody.Notes,
	}

	result := controller.commandBus.Execute(context.Background(), command)
	if !result.IsSuccess {
		apiResponse.ApplicationError(ctx, result.Error)
		return
	}

	apiResponse.Success(ctx, result)
}

type CompleteAppointmentRequest struct {
	Notes *string `json:"notes,omitempty"`
}
