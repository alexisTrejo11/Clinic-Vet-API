// Package controller handles appointment-related HTTP endpoints
package controller

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/appointment/infrastructure/bus"
	"clinic-vet-api/app/modules/appointment/presentation/dto"
	"clinic-vet-api/app/shared/response"
	"errors"

	authError "clinic-vet-api/app/shared/error/auth"
	httpError "clinic-vet-api/app/shared/error/infrastructure/http"
	ginUtils "clinic-vet-api/app/shared/gin_utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type customerQueryExtraArgs struct {
	PetID  *uint
	Status *string
}

// CustomerAppointmetController handles customer-specific appointment operations
// @title Veterinary Clinic API - customer Appt Management
// @version 1.0
// @description This controller manages appointment operations specific to pet customers including scheduling, rescheduling, and viewing appointments
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
type CustomerAppointmetController struct {
	bus             *bus.AppointmentBus
	validator       *validator.Validate
	queryController *AppointmentQueryController
}

func NewCustomerApptControleer(bus *bus.AppointmentBus, validator *validator.Validate, queryController *AppointmentQueryController) *CustomerAppointmetController {
	return &CustomerAppointmetController{
		bus:             bus,
		validator:       validator,
		queryController: queryController,
	}
}

// RequestAppt godoc
// @Summary Request a new appointment
// @Description customer creates a new appointment request for their pet
// @Tags customer-appointments
// @Accept json
// @Produce json
// @Param appointment body command.CreateApptCommand true "Appointment details"
// @Security BearerAuth
// @Router /customer/appointments [post]
func (ctrl *CustomerAppointmetController) RequestAppointment(c *gin.Context) {
	userCtx, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, authError.UnauthorizedCTXError())
		return
	}

	var requestData *dto.RequestApptmentData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := ctrl.validator.Struct(&requestData); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	commmand, err := requestData.ToCommand(c.Request.Context(), userCtx.CustomerID)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.CommandBus.RequestAppointmentByCustomer(*commmand)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}
	response.Created(c, result, "Appointment")
}

// GetMyAppts godoc
// @Summary Get customer's appointments
// @Description Retrieves a list of all appointments for the authenticated customer
// @Tags customer-appointments
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Security BearerAuth
// @Router /customer/appointments [get]
func (ctrl *CustomerAppointmetController) GetMyAppointments(c *gin.Context) {
	userCtx, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, authError.UnauthorizedCTXError())
		return
	}

	noArgs := customerQueryExtraArgs{}
	ctrl.queryController.FindAppointmentsByCustomer(c, userCtx.CustomerID, noArgs)
}

func (ctrl *CustomerAppointmetController) GetMyAppointmentByID(c *gin.Context) {
	userCtx, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, authError.UnauthorizedCTXError())
		return
	}

	args := GetByIDExtraArgs{employeeID: nil, customerID: &userCtx.CustomerID}
	ctrl.queryController.GetAppointmentDetailByID(c, args)
}

// GetApptsByPet godoc
// @Summary Get appointments for a specific pet
// @Description Retrieves all appointments for a specific pet owned by the authenticated customer
// @Tags customer-appointments
// @Accept json
// @Produce json
// @Param petID path int true "Pet ID"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Security BearerAuth
// @Router /customer/appointments/pet/{petID} [get]
func (ctrl *CustomerAppointmetController) GetAppointmentsByPet(c *gin.Context) {
	userCTX, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, errors.New("unauthorized"))
		return
	}

	petID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	extraArgs := &customerQueryExtraArgs{PetID: &petID}
	ctrl.queryController.FindAppointmentsByCustomer(c, userCTX.CustomerID, *extraArgs)
}

// GetUpcomingAppts godoc
// @Summary Get upcoming appointments
// @Description Retrieves upcoming appointments for the authenticated customer within a date range
// @Tags customer-appointments
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Security BearerAuth
// @Router /customer/appointments/upcoming [get]
func (ctrl *CustomerAppointmetController) GetUpcomingAppointments(c *gin.Context) {
	userCTX, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, errors.New("unauthorized"))
		return
	}

	upcomingStatus := "upcoming"
	extraArgs := &customerQueryExtraArgs{Status: &upcomingStatus}
	ctrl.queryController.FindAppointmentsByCustomer(c, userCTX.CustomerID, *extraArgs)
}
