// Package controller handles all appointment-related HTTP endpoints
package controller

import (
	"context"

	"clinic-vet-api/app/modules/appointment/application/command"
	"clinic-vet-api/app/modules/appointment/infrastructure/bus"
	"clinic-vet-api/app/modules/appointment/presentation/dto"
	httpError "clinic-vet-api/app/shared/error/infrastructure/http"
	ginUtils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AppointmentCommandController struct {
	bus      bus.AppointmentBus
	validate *validator.Validate
}

func NewApptCommandController(bus bus.AppointmentBus, validate *validator.Validate) *AppointmentCommandController {
	return &AppointmentCommandController{
		bus:      bus,
		validate: validate,
	}
}

func (ctrl *AppointmentCommandController) CreateAppointment(c *gin.Context, employeeID *uint) {
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

func (ctrl *AppointmentCommandController) UpdateAppointment(c *gin.Context) {
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

func (ctrl *AppointmentCommandController) DeleteAppointment(c *gin.Context) {
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

func (ctrl *AppointmentCommandController) RescheduleAppointment(c *gin.Context, employeeID *uint) {
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

func (ctrl *AppointmentCommandController) NotAttend(c *gin.Context, employeeID *uint) {
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

func (ctrl *AppointmentCommandController) ConfirmAppointment(c *gin.Context, employeeID uint) {
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

func (ctrl *AppointmentCommandController) CompleteAppointment(c *gin.Context, employeeID *uint) {
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

func (ctrl *AppointmentCommandController) CancelAppointment(c *gin.Context, employeeID *uint) {
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
