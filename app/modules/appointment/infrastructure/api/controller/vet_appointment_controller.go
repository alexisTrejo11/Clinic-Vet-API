// Package controller handles appointment-related HTTP endpoints
package controller

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/application/command"
	appointmentQuery "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/application/queries"
	response "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/responses"
	"github.com/gin-gonic/gin"
)

// VetAppointmentController handles veterinarian-specific appointment operations
// @title Veterinary Clinic API - Veterinarian Appointment Management
// @version 1.0
// @description This controller manages appointment operations specific to veterinarians including viewing, confirming, completing, and managing appointments
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
type VetAppointmentController struct {
	commandBus command.CommandBus
	queryBus   appointmentQuery.QueryBus
}

func NewVetAppointmentController(commandBus command.CommandBus, queryBus appointmentQuery.QueryBus) *VetAppointmentController {
	return &VetAppointmentController{
		commandBus: commandBus,
		queryBus:   queryBus,
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
func (c *VetAppointmentController) GetMyAppointments(ctx *gin.Context) {
	// Parse pagination parameters
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))

	// Get vet id from JWT context (assuming it's set by auth middleware)
	vetIdInterface, exists := ctx.Get("vet_id")
	if !exists {
		response.Unauthorized(ctx, errors.New("vet id not found in context"))
		return
	}

	vetId, ok := vetIdInterface.(int)
	if !ok {
		response.BadRequest(ctx, errors.New("invalid vet id format"))
		return
	}

	query := appointmentQuery.NewGetAppointmentsByVetQuery(vetId, page, pageSize)
	result, err := c.queryBus.Execute(context.Background(), query)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	response.Success(ctx, result)
}

// GetTodayAppointments godoc
// @Summary Get today's appointments
// @Description Retrieves all appointments scheduled for today for the authenticated veterinarian
// @Tags vet-appointments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.APIResponse{data=[]appointmentQuery.AppointmentResponse} "Today's appointments"
// @Failure 401 {object} response.APIResponse "Unauthorized - Veterinarian not authenticated"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /vet/appointments/today [get]
func (c *VetAppointmentController) GetTodayAppointments(ctx *gin.Context) {
	// Get today's date range
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24 * time.Hour).Add(-time.Nanosecond)

	query := appointmentQuery.NewGetAppointmentsByDateRangeQuery(
		startOfDay, endOfDay, 1, 5000,
	)
	result, err := c.queryBus.Execute(context.Background(), query)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	response.Success(ctx, result)
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
func (c *VetAppointmentController) ConfirmAppointment(ctx *gin.Context) {
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
func (c *VetAppointmentController) CompleteAppointment(ctx *gin.Context) {
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
func (c *VetAppointmentController) CancelAppointment(ctx *gin.Context) {
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
func (c *VetAppointmentController) MarkAsNoShow(ctx *gin.Context) {
	appointmentId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.RequestURLQueryError(ctx, err)
		return
	}

	// Get vet id from JWT context to verify the appointment belongs to this vet
	vetIdInterface, exists := ctx.Get("vet_id")
	if !exists {
		response.Unauthorized(ctx, errors.New("vet id not found in context"))
		return
	}

	vetId, ok := vetIdInterface.(int)
	if !ok {
		response.BadRequest(ctx, errors.New("invalid vet id format"))
		return
	}

	// Verify the appointment belongs to this vet
	getQuery := appointmentQuery.NewGetAppointmentByIdQuery(appointmentId)
	getResult, err := c.queryBus.Execute(context.Background(), getQuery)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	appointment := getResult.(*appointmentQuery.AppointmentResponse)

	if appointment.VetId == nil || *appointment.VetId != vetId {
		response.Forbidden(ctx, errors.New("access denied: you can only mark your own appointments as no-show"))
		return
	}

	command := command.MarkAsNotPresentedCommand{
		Id: appointmentId,
	}

	result := c.commandBus.Execute(context.Background(), command)
	if !result.IsSuccess {
		response.ApplicationError(ctx, result.Error)
		return
	}

	response.Success(ctx, result)
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
func (c *VetAppointmentController) GetAppointmentStats(ctx *gin.Context) {
	// Get vet id from JWT context
	vetIdInterface, exists := ctx.Get("vet_id")
	if !exists {
		response.Unauthorized(ctx, errors.New("vet id not found in context"))
		return
	}

	vetId, ok := vetIdInterface.(int)
	if !ok {
		response.BadRequest(ctx, errors.New("invalid vet id format"))
		return
	}

	// Parse date range from query parameters (optional)
	var startDate, endDate *time.Time
	if startDateStr := ctx.Query("start_date"); startDateStr != "" {
		if parsed, err := time.Parse("2006-01-02", startDateStr); err == nil {
			startDate = &parsed
		}
	}
	if endDateStr := ctx.Query("end_date"); endDateStr != "" {
		if parsed, err := time.Parse("2006-01-02", endDateStr); err == nil {
			endDate = &parsed
		}
	}

	query := appointmentQuery.NewGetAppointmentStatsQuery(&vetId, nil, startDate, endDate)
	result, err := c.queryBus.Execute(context.Background(), query)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	response.Success(ctx, result)
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
func (c *VetAppointmentController) RescheduleAppointment(ctx *gin.Context) {
	appointmentId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.RequestURLQueryError(ctx, err)
		return
	}

	var command command.RescheduleAppointmentCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		response.RequestBodyDataError(ctx, err)
		return
	}

	command.AppointmentId = appointmentId

	// Get vet id from JWT context to verify the appointment belongs to this vet
	vetIdInterface, exists := ctx.Get("vet_id")
	if !exists {
		response.Unauthorized(ctx, errors.New("vet id not found in context"))
		return
	}

	vetId, ok := vetIdInterface.(int)
	if !ok {
		response.BadRequest(ctx, errors.New("invalid vet id format"))
		return
	}

	// Verify the appointment belongs to this vet
	getQuery := appointmentQuery.NewGetAppointmentByIdQuery(appointmentId)
	getResult, err := c.queryBus.Execute(context.Background(), getQuery)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	appointment := getResult.(*appointmentQuery.AppointmentResponse)

	if appointment.VetId == nil || *appointment.VetId != vetId {
		response.Forbidden(ctx, errors.New("acccess denied: you can only reschedule your own appointments"))
		return
	}

	result := c.commandBus.Execute(context.Background(), command)
	if !result.IsSuccess {
		response.ApplicationError(ctx, result.Error)
		return
	}

	response.Success(ctx, result)
}
