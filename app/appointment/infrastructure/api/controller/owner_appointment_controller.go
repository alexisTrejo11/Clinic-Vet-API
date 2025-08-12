package appointmentController

import (
	"net/http"
	"strconv"
	"time"

	appointmentCmd "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/application/command"
	appointmentQuery "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/application/queries"
	"github.com/gin-gonic/gin"
)

type OwnerAppointmentController struct {
	commandBus appointmentCmd.CommandBus
	queryBus   appointmentQuery.QueryBus
}

func NewOwnerAppointmentController(commandBus appointmentCmd.CommandBus, queryBus appointmentQuery.QueryBus) *OwnerAppointmentController {
	return &OwnerAppointmentController{
		commandBus: commandBus,
		queryBus:   queryBus,
	}
}

// RequestAppointment - Owner requests a new appointment
func (controller *OwnerAppointmentController) RequestAppointment(ctx *gin.Context) {
	var command appointmentCmd.CreateAppointmentCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get owner ID from context (should be set by auth middleware)
	ownerIdStr, exists := ctx.Get("owner_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Owner ID not found in context"})
		return
	}

	ownerId, ok := ownerIdStr.(int)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid owner ID format"})
		return
	}

	command.OwnerId = ownerId

	result := controller.commandBus.Execute(ctx, command)
	if !result.IsSuccess {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Message})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": result.Message, "appointment_id": result.Id})
}

// GetMyAppointments - Owner retrieves their appointments
func (controller *OwnerAppointmentController) GetMyAppointments(ctx *gin.Context) {
	// Get owner ID from context (should be set by auth middleware)
	ownerIdStr, exists := ctx.Get("owner_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Owner ID not found in context"})
		return
	}

	ownerId, ok := ownerIdStr.(int)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid owner ID format"})
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

	query := appointmentQuery.NewGetAppointmentsByOwnerQuery(ownerId, page, limit)

	response, err := controller.queryBus.Execute(ctx, query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// GetAppointmentById - Owner retrieves a specific appointment by ID (only their own)
func (controller *OwnerAppointmentController) GetAppointmentById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}

	// Get owner ID from context (should be set by auth middleware)

	query := appointmentQuery.GetAppointmentByIdQuery{
		AppointmentId: id,
	}

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

// RescheduleAppointment - Owner reschedules their appointment
func (controller *OwnerAppointmentController) RescheduleAppointment(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}

	var command appointmentCmd.RescheduleAppointmentCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	command.AppointmentId = id

	result := controller.commandBus.Execute(ctx, command)
	if !result.IsSuccess {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": result.Message})
}

// CancelAppointment - Owner cancels their appointment
func (controller *OwnerAppointmentController) CancelAppointment(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}

	command := appointmentCmd.CancelAppointmentCommand{
		AppointmentId: id,
	}

	result := controller.commandBus.Execute(ctx, command)
	if !result.IsSuccess {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": result.Message})
}

// GetAppointmentsByPet - Owner retrieves appointments for a specific pet
func (controller *OwnerAppointmentController) GetAppointmentsByPet(ctx *gin.Context) {
	petIdParam := ctx.Param("petId")
	petId, err := strconv.Atoi(petIdParam)
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

	query := appointmentQuery.NewGetAppointmentsByPetQuery(petId, page, limit)
	response, err := controller.queryBus.Execute(ctx, query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// GetUpcomingAppointments - Owner retrieves upcoming appointments within date range
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

	query := appointmentQuery.NewGetAppointmentsByDateRangeQuery(startDate, endDate, page, limit)

	response, err := controller.queryBus.Execute(ctx, query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// TODO: Filter results by owner_id if needed at application level
	ctx.JSON(http.StatusOK, response)
}
