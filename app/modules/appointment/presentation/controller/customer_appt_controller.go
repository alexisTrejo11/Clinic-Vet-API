// Package controller handles appointment-related HTTP endpoints
package controller

import (
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/infrastructure/bus"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/presentation/dto"
	authError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/auth"
	httpError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure/http"
	ginUtils "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/gin_utils"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/response"
	"github.com/alexisTrejo11/Clinic-Vet-API/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type OwnerQueryExtraArgs struct {
	PetID  *uint
	Status *string
}

// CustomerApptControleer handles owner-specific appointment operations
// @title Veterinary Clinic API - Owner Appt Management
// @version 1.0
// @description This controller manages appointment operations specific to pet owners including scheduling, rescheduling, and viewing appointments
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
type CustomerApptControleer struct {
	bus             *bus.AppointmentBus
	validator       *validator.Validate
	queryController *AppointmentQueryController
}

func NewCustomerApptControleer(bus *bus.AppointmentBus, validator *validator.Validate, queryController *AppointmentQueryController) *CustomerApptControleer {
	return &CustomerApptControleer{
		bus:             bus,
		validator:       validator,
		queryController: queryController,
	}
}

// RequestAppt godoc
// @Summary Request a new appointment
// @Description Owner creates a new appointment request for their pet
// @Tags owner-appointments
// @Accept json
// @Produce json
// @Param appointment body command.CreateApptCommand true "Appointment details"
// @Security BearerAuth
// @Router /owner/appointments [post]
func (ctrl *CustomerApptControleer) RequestAppointment(c *gin.Context) {
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
// @Summary Get owner's appointments
// @Description Retrieves a list of all appointments for the authenticated owner
// @Tags owner-appointments
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Security BearerAuth
// @Router /owner/appointments [get]
func (ctrl *CustomerApptControleer) GetMyAppointments(c *gin.Context) {
	userCtx, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, authError.UnauthorizedCTXError())
		return
	}

	noArgs := OwnerQueryExtraArgs{}
	ctrl.queryController.FindAppointmentsByCustomer(c, userCtx.CustomerID, noArgs)
}

// GetApptsByPet godoc
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
func (ctrl *CustomerApptControleer) GetAppointmentsByPet(c *gin.Context) {
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

	extraArgs := &OwnerQueryExtraArgs{PetID: &petID}
	ctrl.queryController.FindAppointmentsByCustomer(c, userCTX.CustomerID, *extraArgs)
}

// GetUpcomingAppts godoc
// @Summary Get upcoming appointments
// @Description Retrieves upcoming appointments for the authenticated owner within a date range
// @Tags owner-appointments
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Security BearerAuth
// @Router /owner/appointments/upcoming [get]
func (controller *CustomerApptControleer) GetUpcomingAppointments(c *gin.Context) {
	userCTX, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, errors.New("unauthorized"))
		return
	}

	upcomingStatus := "upcoming"
	extraArgs := &OwnerQueryExtraArgs{Status: &upcomingStatus}
	controller.queryController.FindAppointmentsByCustomer(c, userCTX.CustomerID, *extraArgs)
}
