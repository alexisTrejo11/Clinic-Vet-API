package appointmentController

import (
	"net/http"
	"strconv"
	"time"

	appointmentCmd "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/application/command"
	appointmentQuery "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/application/queries"
	"github.com/gin-gonic/gin"
)

type AppointmentController struct {
	commandBus appointmentCmd.CommandBus
	queryBus   appointmentQuery.QueryBus
}

func NewAppointmentController(commandBus appointmentCmd.CommandBus, queryBus appointmentQuery.QueryBus) *AppointmentController {
	return &AppointmentController{
		commandBus: commandBus,
		queryBus:   queryBus,
	}
}

// Create appointment
func (controller *AppointmentController) CreateAppointment(ctx *gin.Context) {
	var command appointmentCmd.CreateAppointmentCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := controller.commandBus.Execute(ctx, command)
	if !result.IsSuccess {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Message})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": result.Message, "appointment_id": result.Id})
}

// Get appointment by ID
func (controller *AppointmentController) GetAppointmentById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}

	query := appointmentQuery.GetAppointmentByIdQuery{AppointmentId: id}

	response, err := controller.queryBus.Execute(ctx, query)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Get all appointments with pagination
func (controller *AppointmentController) GetAllAppointments(ctx *gin.Context) {
	pageParam := ctx.DefaultQuery("page", "1")
	limitParam := ctx.DefaultQuery("pageSize", "10")

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

	query := appointmentQuery.GetAllAppointmentsQuery{
		Page:     page,
		PageSize: limit,
	}

	response, err := controller.queryBus.Execute(ctx, query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Update appointment
func (controller *AppointmentController) UpdateAppointment(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}

	var command appointmentCmd.UpdateAppointmentCommand
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

// Delete appointment
func (controller *AppointmentController) DeleteAppointment(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}

	command := appointmentCmd.DeleteAppointmentCommand{
		AppointmentId: id,
	}

	result := controller.commandBus.Execute(ctx, command)
	if !result.IsSuccess {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": result.Message})
}

// Get appointments by date range
func (controller *AppointmentController) GetAppointmentsByDateRange(ctx *gin.Context) {
	startDateStr := ctx.Query("start_date")
	endDateStr := ctx.Query("end_date")

	if startDateStr == "" || endDateStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Start date and end date are required"})
		return
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format, expected format: YYYY-MM-DD"})
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format, expected format: YYYY-MM-DD"})
		return
	}

	pageParam := ctx.DefaultQuery("page", "1")
	limitParam := ctx.DefaultQuery("pageSize", "10")

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

	query := appointmentQuery.GetAppointmentsByDateRangeQuery{
		StartDate: startDate,
		EndDate:   endDate,
		Page:      page,
		PageSize:  limit,
	}

	response, err := controller.queryBus.Execute(ctx, query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Get appointments by owner
func (controller *AppointmentController) GetAppointmentsByOwner(ctx *gin.Context) {
	ownerIdParam := ctx.Param("ownerId")
	ownerId, err := strconv.Atoi(ownerIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid owner ID"})
		return
	}

	pageParam := ctx.DefaultQuery("page", "1")
	limitParam := ctx.DefaultQuery("pageSize", "10")

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

	query := appointmentQuery.GetAppointmentsByOwnerQuery{
		OwnerId:  ownerId,
		Page:     page,
		PageSize: limit,
	}

	response, err := controller.queryBus.Execute(ctx, query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Get appointments by vet
func (controller *AppointmentController) GetAppointmentsByVet(ctx *gin.Context) {
	vetIdParam := ctx.Param("vetId")
	vetId, err := strconv.Atoi(vetIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vet ID"})
		return
	}

	pageParam := ctx.DefaultQuery("page", "1")
	limitParam := ctx.DefaultQuery("pageSize", "10")

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

	query := appointmentQuery.GetAppointmentsByVetQuery{
		VetId:    vetId,
		Page:     page,
		PageSize: limit,
	}

	response, err := controller.queryBus.Execute(ctx, query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Get appointments by pet
func (controller *AppointmentController) GetAppointmentsByPet(ctx *gin.Context) {
	petIdParam := ctx.Param("petId")
	petId, err := strconv.Atoi(petIdParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pet ID"})
		return
	}

	pageParam := ctx.DefaultQuery("page", "1")
	limitParam := ctx.DefaultQuery("pageSize", "10")

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

	query := appointmentQuery.GetAppointmentsByPetQuery{
		PetId:    petId,
		Page:     page,
		PageSize: limit,
	}

	response, err := controller.queryBus.Execute(ctx, query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Get appointment statistics
func (controller *AppointmentController) GetAppointmentStats(ctx *gin.Context) {
	query := appointmentQuery.GetAppointmentStatsQuery{}

	response, err := controller.queryBus.Execute(ctx, query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
