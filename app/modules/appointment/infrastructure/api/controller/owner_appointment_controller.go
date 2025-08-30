// Package controller handles appointment-related HTTP endpoints
package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/application/command"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/application/query"
	"github.com/gin-gonic/gin"
)

// OwnerAppointmentController handles owner-specific appointment operations
// @title Veterinary Clinic API - Owner Appointment Management
// @version 1.0
// @description This controller manages appointment operations specific to pet owners including scheduling, rescheduling, and viewing appointments
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
type OwnerAppointmentController struct {
	commandBus command.CommandBus
	queryBus   query.QueryBus
}

func NewOwnerAppointmentController(commandBus command.CommandBus, queryBus query.QueryBus) *OwnerAppointmentController {
	return &OwnerAppointmentController{
		commandBus: commandBus,
		queryBus:   queryBus,
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
func (controller *OwnerAppointmentController) RequestAppointment(ctx *gin.Context) {
	var command command.CreateAppointmentCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get owner ID from context (should be set by auth middleware)
	ownerIDStr, exists := ctx.Get("owner_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Owner ID not found in context"})
		return
	}

	ownerID, ok := ownerIDStr.(int)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid owner ID format"})
		return
	}

	command.OwnerID = ownerID

	result := controller.commandBus.Execute(ctx, command)
	if !result.IsSuccess {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Message})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": result.Message, "appointment_id": result.ID})
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
func (controller *OwnerAppointmentController) GetMyAppointments(ctx *gin.Context) {
	// Get owner ID from context (should be set by auth middleware)
	ownerIDStr, exists := ctx.Get("owner_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Owner ID not found in context"})
		return
	}

	idint, ok := ownerIDStr.(int)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid owner ID format"})
		return
	}

	ownerID, _ := valueobject.NewOwnerID(idint)

	pageParam := ctx.DefaultQuery("page", "1")
	limitParam := ctx.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	query := query.NewGetAppointmentsByOwnerQuery(ownerID, page, limit)

	response, err := controller.queryBus.Execute(ctx, query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
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
func (controller *OwnerAppointmentController) GetAppointmentByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}

	// Get owner ID from context (should be set by auth middleware)

	appointmentID, err := valueobject.NewAppointmentID(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	query := query.NewGetAppointmentByIDQuery(appointmentID)

	response, err := controller.queryBus.Execute(ctx, query)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Verify appointment belongs to the owner
	// This would require checking the appointment's owner_id
	// For now, we'll return the appointment

	ctx.JSON(http.StatusOK, response)
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
func (controller *OwnerAppointmentController) RescheduleAppointment(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}

	appointmentID, err := valueobject.NewAppointmentID(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var command command.RescheduleAppointmentCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	command.AppointmentID = appointmentID

	result := controller.commandBus.Execute(ctx, command)
	if !result.IsSuccess {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": result.Message})
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
func (controller *OwnerAppointmentController) CancelAppointment(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}

	appointmentID, err := valueobject.NewAppointmentID(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	command := command.NewCancelAppointmentCommand(appointmentID, "")

	result := controller.commandBus.Execute(ctx, command)
	if !result.IsSuccess {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": result.Message})
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
func (controller *OwnerAppointmentController) GetAppointmentsByPet(ctx *gin.Context) {
	petIDParam := ctx.Param("petID")
	idInt, err := strconv.Atoi(petIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pet ID"})
		return
	}

	pageParam := ctx.DefaultQuery("page", "1")
	limitParam := ctx.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	petID, _ := valueobject.NewPetID(idInt)
	query := query.NewGetAppointmentsByPetQuery(petID, page, limit)
	response, err := controller.queryBus.Execute(ctx, query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
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
func (controller *OwnerAppointmentController) GetUpcomingAppointments(ctx *gin.Context) {
	// Get owner ID from context (should be set by auth middleware)

	// Default to next 30 days if no date range specified
	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 30)

	// Allow custom date range via query parameters
	if startDateStr := ctx.Query("start_date"); startDateStr != "" {
		if parsed, err := time.Parse("2006-01-02", startDateStr); err == nil {
			startDate = parsed
		}
	}
	if endDateStr := ctx.Query("end_date"); endDateStr != "" {
		if parsed, err := time.Parse("2006-01-02", endDateStr); err == nil {
			endDate = parsed
		}
	}

	pageParam := ctx.DefaultQuery("page", "1")
	limitParam := ctx.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	query := query.NewGetAppointmentsByDateRangeQuery(startDate, endDate, page, limit)

	response, err := controller.queryBus.Execute(ctx, query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
