// Package controller handles appointment-related HTTP endpoints
package controller

import (
	"clinic-vet-api/app/middleware"
	authError "clinic-vet-api/app/shared/error/auth"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// EmployeeAppointmentController handles veterinarian-specific appointment operations
// @title Veterinary Clinic API - Veterinarian Appointment Management
// @version 1.0
// @description This controller manages appointment operations specific to veterinarians including viewing, confirming, completing, and managing appointments
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
type EmployeeAppointmentController struct {
	validator  *validator.Validate
	operations *ApptControllerOperations
}

func NewEmployeeController(operations *ApptControllerOperations, validator *validator.Validate) *EmployeeAppointmentController {
	return &EmployeeAppointmentController{
		validator:  validator,
		operations: operations,
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
func (ctrl *EmployeeAppointmentController) GetMyAppointments(c *gin.Context) {
	userCTX, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, authError.UnauthorizedCTXError())
		return
	}

	ctrl.operations.GetAppointmentsByEmployee(c, userCTX.EmployeeID)
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
func (ctrl *EmployeeAppointmentController) CompleteAppointment(c *gin.Context) {
	userCTX, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, authError.UnauthorizedCTXError())
		return
	}

	ctrl.operations.CompleteAppointment(c, &userCTX.EmployeeID)
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
func (ctrl *EmployeeAppointmentController) CancelAppointment(c *gin.Context) {
	userCTX, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, authError.UnauthorizedCTXError())
		return
	}

	ctrl.operations.CancelAppointment(c, &userCTX.EmployeeID)
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
func (ctrl *EmployeeAppointmentController) ConfirmAppointment(c *gin.Context) {
	userCTX, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, authError.UnauthorizedCTXError())
		return
	}

	ctrl.operations.ConfirmAppointment(c, userCTX.EmployeeID)
}

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
func (ctrl *EmployeeAppointmentController) MarkAsNoShow(c *gin.Context) {
	userCTX, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, authError.UnauthorizedCTXError())
		return
	}

	ctrl.operations.NotAttend(c, &userCTX.EmployeeID)
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
func (ctrl *EmployeeAppointmentController) GetAppointmentStats(c *gin.Context) {
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
func (ctrl *EmployeeAppointmentController) RescheduleAppointment(c *gin.Context) {
	userCTX, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, authError.UnauthorizedCTXError())
		return
	}

	ctrl.operations.RescheduleAppointment(c, &userCTX.EmployeeID)
}
