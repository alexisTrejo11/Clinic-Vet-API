package appointmentController

import (
	"context"
	"errors"
	"strconv"
	"time"

	appointmentCmd "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/application/command"
	appointmentQuery "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/application/queries"
	responses "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/responses"
	"github.com/gin-gonic/gin"
)

type VetAppointmentController struct {
	commandBus appointmentCmd.CommandBus
	queryBus   appointmentQuery.QueryBus
}

func NewVetAppointmentController(commandBus appointmentCmd.CommandBus, queryBus appointmentQuery.QueryBus) *VetAppointmentController {
	return &VetAppointmentController{
		commandBus: commandBus,
		queryBus:   queryBus,
	}
}

// GetMyAppointments retrieves all appointments assigned to the current veterinarian
// GET /vet/appointments
func (c *VetAppointmentController) GetMyAppointments(ctx *gin.Context) {
	// Parse pagination parameters
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))

	// Get vet id from JWT context (assuming it's set by auth middleware)
	vetIdInterface, exists := ctx.Get("vet_id")
	if !exists {
		responses.Unauthorized(ctx, errors.New("vet id not found in context"))
		return
	}

	vetId, ok := vetIdInterface.(int)
	if !ok {
		responses.BadRequest(ctx, errors.New("invalid vet id format"))
		return
	}

	query := appointmentQuery.NewGetAppointmentsByVetQuery(vetId, page, pageSize)
	result, err := c.queryBus.Execute(context.Background(), query)
	if err != nil {
		responses.ApplicationError(ctx, err)
		return
	}

	responses.Success(ctx, result)
}

// GetTodayAppointments retrieves today's appointments for the current veterinarian
// GET /vet/appointments/today
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
		responses.ApplicationError(ctx, err)
		return
	}

	responses.Success(ctx, result)
}

// ConfirmAppointment allows veterinarians to confirm pending appointments
// PUT /vet/appointments/:id/confirm
func (c *VetAppointmentController) ConfirmAppointment(ctx *gin.Context) {
	appointmentId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.RequestURLQueryError(ctx, err)
		return
	}

	// Get vet id from JWT context
	vetIdInterface, exists := ctx.Get("vet_id")
	if !exists {
		responses.Unauthorized(ctx, errors.New("vet id not found in context"))
		return
	}

	vetId, ok := vetIdInterface.(int)
	if !ok {
		responses.BadRequest(ctx, errors.New("invalid vet id format"))
		return
	}

	command := appointmentCmd.ConfirmAppointmentCommand{
		Id:    appointmentId,
		VetId: &vetId,
	}

	result := c.commandBus.Execute(context.Background(), command)
	if !result.IsSuccess {
		responses.ApplicationError(ctx, result.Error)
		return
	}

	responses.Success(ctx, result)
}

// CompleteAppointment allows veterinarians to mark appointments as completed
// PUT /vet/appointments/:id/complete
func (c *VetAppointmentController) CompleteAppointment(ctx *gin.Context) {
	appointmentId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.RequestURLQueryError(ctx, err)
		return
	}

	var requestBody struct {
		Notes *string `json:"notes,omitempty"`
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		responses.RequestBodyDataError(ctx, err)
		return
	}

	// Get vet id from JWT context to verify the appointment belongs to this vet
	vetIdInterface, exists := ctx.Get("vet_id")
	if !exists {
		responses.Unauthorized(ctx, errors.New("vet id not found in context"))
		return
	}

	vetId, ok := vetIdInterface.(int)
	if !ok {
		responses.BadRequest(ctx, errors.New("invalid vet id format"))
		return
	}

	// Verify the appointment belongs to this vet
	getQuery := appointmentQuery.NewGetAppointmentByIdQuery(appointmentId)
	getResult, err := c.queryBus.Execute(context.Background(), getQuery)
	if err != nil {
		responses.ApplicationError(ctx, err)
		return
	}

	appointment := getResult.(*appointmentQuery.AppointmentResponse)

	if appointment.VetId == nil || *appointment.VetId != vetId {
		responses.Forbidden(ctx, errors.New("access denied: you can only complete your own appointments"))
		return
	}

	command := appointmentCmd.CompleteAppointmentCommand{
		Id:    appointmentId,
		Notes: requestBody.Notes,
	}

	result := c.commandBus.Execute(context.Background(), command)
	if !result.IsSuccess {
		responses.ApplicationError(ctx, result.Error)
		return
	}

	responses.Success(ctx, result)
}

// CancelAppointment allows veterinarians to cancel appointments
// DELETE /vet/appointments/:id
func (c *VetAppointmentController) CancelAppointment(ctx *gin.Context) {
	appointmentId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.RequestURLQueryError(ctx, err)
		return
	}

	var requestBody struct {
		Reason string `json:"reason" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		responses.RequestBodyDataError(ctx, err)
		return
	}

	// Get vet id from JWT context to verify the appointment belongs to this vet
	vetIdInterface, exists := ctx.Get("vet_id")
	if !exists {
		responses.Unauthorized(ctx, errors.New("vet id not found in context"))
		return
	}

	vetId, ok := vetIdInterface.(int)
	if !ok {
		responses.BadRequest(ctx, errors.New("invalid vet id format"))
		return
	}

	// Verify the appointment belongs to this vet
	getQuery := appointmentQuery.NewGetAppointmentByIdQuery(appointmentId)
	getResult, err := c.queryBus.Execute(context.Background(), getQuery)
	if err != nil {
		responses.ApplicationError(ctx, err)
		return
	}

	appointment := getResult.(*appointmentQuery.AppointmentResponse)

	if appointment.VetId == nil || *appointment.VetId != vetId {
		responses.Forbidden(ctx, errors.New("access denied: you can only cancel your own appointments"))
		return
	}

	command := appointmentCmd.CancelAppointmentCommand{
		AppointmentId: appointmentId,
		Reason:        requestBody.Reason,
	}

	result := c.commandBus.Execute(context.Background(), command)
	if !result.IsSuccess {
		responses.ApplicationError(ctx, result.Error)
		return
	}

	responses.Success(ctx, result)
}

// MarkAsNoShow allows veterinarians to mark appointments as no-show
// PUT /vet/appointments/:id/no-show
func (c *VetAppointmentController) MarkAsNoShow(ctx *gin.Context) {
	appointmentId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.RequestURLQueryError(ctx, err)
		return
	}

	// Get vet id from JWT context to verify the appointment belongs to this vet
	vetIdInterface, exists := ctx.Get("vet_id")
	if !exists {
		responses.Unauthorized(ctx, errors.New("vet id not found in context"))
		return
	}

	vetId, ok := vetIdInterface.(int)
	if !ok {
		responses.BadRequest(ctx, errors.New("invalid vet id format"))
		return
	}

	// Verify the appointment belongs to this vet
	getQuery := appointmentQuery.NewGetAppointmentByIdQuery(appointmentId)
	getResult, err := c.queryBus.Execute(context.Background(), getQuery)
	if err != nil {
		responses.ApplicationError(ctx, err)
		return
	}

	appointment := getResult.(*appointmentQuery.AppointmentResponse)

	if appointment.VetId == nil || *appointment.VetId != vetId {
		responses.Forbidden(ctx, errors.New("access denied: you can only mark your own appointments as no-show"))
		return
	}

	command := appointmentCmd.MarkAsNotPresentedCommand{
		Id: appointmentId,
	}

	result := c.commandBus.Execute(context.Background(), command)
	if !result.IsSuccess {
		responses.ApplicationError(ctx, result.Error)
		return
	}

	responses.Success(ctx, result)
}

// GetAppointmentStats retrieves appointment statistics for the veterinarian
// GET /vet/appointments/stats
func (c *VetAppointmentController) GetAppointmentStats(ctx *gin.Context) {
	// Get vet id from JWT context
	vetIdInterface, exists := ctx.Get("vet_id")
	if !exists {
		responses.Unauthorized(ctx, errors.New("vet id not found in context"))
		return
	}

	vetId, ok := vetIdInterface.(int)
	if !ok {
		responses.BadRequest(ctx, errors.New("invalid vet id format"))
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
		responses.ApplicationError(ctx, err)
		return
	}

	responses.Success(ctx, result)
}

// RescheduleAppointment allows veterinarians to reschedule appointments
// PUT /vet/appointments/:id/reschedule
func (c *VetAppointmentController) RescheduleAppointment(ctx *gin.Context) {
	appointmentId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.RequestURLQueryError(ctx, err)
		return
	}

	var command appointmentCmd.RescheduleAppointmentCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		responses.RequestBodyDataError(ctx, err)
		return
	}

	command.AppointmentId = appointmentId

	// Get vet id from JWT context to verify the appointment belongs to this vet
	vetIdInterface, exists := ctx.Get("vet_id")
	if !exists {
		responses.Unauthorized(ctx, errors.New("vet id not found in context"))
		return
	}

	vetId, ok := vetIdInterface.(int)
	if !ok {
		responses.BadRequest(ctx, errors.New("invalid vet id format"))
		return
	}

	// Verify the appointment belongs to this vet
	getQuery := appointmentQuery.NewGetAppointmentByIdQuery(appointmentId)
	getResult, err := c.queryBus.Execute(context.Background(), getQuery)
	if err != nil {
		responses.ApplicationError(ctx, err)
		return
	}

	appointment := getResult.(*appointmentQuery.AppointmentResponse)

	if appointment.VetId == nil || *appointment.VetId != vetId {
		responses.Forbidden(ctx, errors.New("acccess denied: you can only reschedule your own appointments"))
		return
	}

	result := c.commandBus.Execute(context.Background(), command)
	if !result.IsSuccess {
		responses.ApplicationError(ctx, result.Error)
		return
	}

	responses.Success(ctx, result)
}
