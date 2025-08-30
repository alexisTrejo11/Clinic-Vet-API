// Package controller handles all appointment-related HTTP endpoints
package controller

import (
	"errors"
	"net/http"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/application/query"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	apiResponse "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/responses"
	response "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// AppointmentQueryController handles appointment query operations
// @title Veterinary Clinic API - Appointment Queries
// @version 1.0
// @description This controller manages appointment queries including retrieving appointments by various criteria
type AppointmentQueryController struct {
	queryBus query.QueryBus
	validate *validator.Validate
}

func NewAppointmentQueryController(
	queryBus query.QueryBus,
	validate *validator.Validate,
) *AppointmentQueryController {
	return &AppointmentQueryController{
		queryBus: queryBus,
		validate: validate,
	}
}

// GetAppointmentByID godoc
// @Summary Get appointment by ID
// @Description Retrieves detailed information about a specific appointment
// @Tags appointments-query
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID"
// @Success 200 {object} query.AppointmentResponse "Appointment details"
// @Failure 400 {object} response.APIResponse "Invalid appointment ID"
// @Failure 404 {object} response.APIResponse "Appointment not found"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments/{id} [get]
func (controller *AppointmentQueryController) GetAppointmentByID(ctx *gin.Context) {
	entityID, err := shared.ParseParamToEntityID(ctx, "id", "appointment")
	if err != nil {
		apiResponse.RequestURLParamError(ctx, err, "appointmentID", ctx.Param("id"))
		return
	}

	appointmentID, valid := entityID.(valueobject.AppointmentID)
	if !valid {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid id entity parse"})
		return
	}

	appointmentQuery := query.NewGetAppointmentByIDQuery(appointmentID)
	appointmentResponse, err := controller.queryBus.Execute(ctx, appointmentQuery)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	response.Success(ctx, appointmentResponse)
}

// GetAllAppointments godoc
// @Summary Get all appointments
// @Description Retrieves a paginated list of all appointments
// @Tags appointments-query
// @Accept json
// @Produce json
// @Param page_number query int false "Page number" default(1)
// @Param page_size query int false "Items per page" default(10)
// @Success 200 {object} response.APIResponse{data=[]query.AppointmentResponse,metadata=PaginationMetadata} "List of appointments"
// @Failure 400 {object} response.APIResponse "Invalid pagination parameters"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments [get]
func (controller *AppointmentQueryController) GetAllAppointments(ctx *gin.Context) {
	var pageParams PaginationRequest

	if err := ctx.ShouldBindQuery(&pageParams); err != nil {
		response.RequestURLQueryError(ctx, err)
		return
	}

	pageParams.SetDefaultsIfNotProvided()
	query := query.NewGetAllAppointmentsQuery(pageParams.PageNumber, pageParams.PageNumber)

	pageInterface, err := controller.queryBus.Execute(ctx, query)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	HandlePaginatedResponse(ctx, pageInterface, pageParams)
}

// GetAppointmentsByDateRange godoc
// @Summary Get appointments by date range
// @Description Retrieves appointments within a specified date range
// @Tags appointments-query
// @Accept json
// @Produce json
// @Param start_date query string true "Start date (YYYY-MM-DD)" format(date)
// @Param end_date query string true "End date (YYYY-MM-DD)" format(date)
// @Param page_number query int false "Page number" default(1)
// @Param page_size query int false "Items per page" default(10)
// @Success 200 {object} response.APIResponse{data=[]query.AppointmentResponse,metadata=PaginationMetadata} "List of appointments"
// @Failure 400 {object} response.APIResponse "Invalid date range or pagination parameters"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments/date-range [get]
func (controller *AppointmentQueryController) GetAppointmentsByDateRange(ctx *gin.Context) {
	var queryParams GetAppointmentsByDateRangeRequest

	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		response.RequestURLQueryError(ctx, err)
		return
	}

	if err := controller.validate.Struct(&queryParams); err != nil {
		response.RequestBodyDataError(ctx, err)
		return
	}

	query := query.NewGetAppointmentsByDateRangeQuery(
		queryParams.StartDate.Time,
		queryParams.EndDate.Time,
		queryParams.PageNumber,
		queryParams.PageSize,
	)

	pageInterface, err := controller.queryBus.Execute(ctx, query)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	HandlePaginatedResponse(ctx, pageInterface, query)
}

// GetAppointmentsByOwner godoc
// @Summary Get appointments by owner
// @Description Retrieves all appointments for a specific pet owner
// @Tags appointments-query
// @Accept json
// @Produce json
// @Param ownerID path int true "Owner ID"
// @Param page_number query int false "Page number" default(1)
// @Param page_size query int false "Items per page" default(10)
// @Success 200 {object} response.APIResponse{data=[]query.AppointmentResponse,metadata=PaginationMetadata} "List of appointments"
// @Failure 400 {object} response.APIResponse "Invalid owner ID or pagination parameters"
// @Failure 404 {object} response.APIResponse "Owner not found"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments/owner/{id} [get]
func (controller *AppointmentQueryController) GetAppointmentsByOwner(ctx *gin.Context) {
	entityID, err := shared.ParseParamToEntityID(ctx, "id", "owner")
	if err != nil {
		response.RequestURLParamError(ctx, err, "id", ctx.Param("ownerID"))
		return
	}

	ownerID, valid := entityID.(valueobject.OwnerID)
	if !valid {
		response.InvalidParseDataError(ctx, "id", entityID.String(), "can't parse owner ID")
	}

	var pageParams PaginationRequest
	if err := ctx.ShouldBindQuery(&pageParams); err != nil {
		response.RequestURLQueryError(ctx, err)
		return
	}

	pageParams.SetDefaultsIfNotProvided()
	query := query.NewGetAppointmentsByOwnerQuery(ownerID, pageParams.PageNumber, pageParams.PageSize)

	pageInterface, err := controller.queryBus.Execute(ctx, query)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	HandlePaginatedResponse(ctx, pageInterface, pageParams)
}

// GetAppointmentsByVet godoc
// @Summary Get appointments by veterinarian
// @Description Retrieves all appointments assigned to a specific veterinarian
// @Tags appointments-query
// @Accept json
// @Produce json
// @Param vetID path int true "Veterinarian ID"
// @Param page_number query int false "Page number" default(1)
// @Param page_size query int false "Items per page" default(10)
// @Success 200 {object} response.APIResponse{data=[]query.AppointmentResponse,metadata=PaginationMetadata} "List of appointments"
// @Failure 400 {object} response.APIResponse "Invalid veterinarian ID or pagination parameters"
// @Failure 404 {object} response.APIResponse "Veterinarian not found"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments/vet/{id} [get]
func (controller *AppointmentQueryController) GetAppointmentsByVet(ctx *gin.Context) {
	entityID, err := shared.ParseParamToEntityID(ctx, "id", "veterinarian")
	if err != nil {
		response.RequestURLParamError(ctx, err, "vetID", ctx.Param("vetID"))
		return
	}

	vetID := entityID.(valueobject.VetID)

	var pageParams PaginationRequest
	if err := ctx.ShouldBindQuery(&pageParams); err != nil {
		response.RequestURLQueryError(ctx, err)
		return
	}
	pageParams.SetDefaultsIfNotProvided()

	query := query.NewGetAppointmentsByVetQuery(
		vetID,
		pageParams.PageNumber,
		pageParams.PageSize,
	)

	pageInterface, err := controller.queryBus.Execute(ctx, query)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	HandlePaginatedResponse(ctx, pageInterface, query)
}

// GetAppointmentsByPet godoc
// @Summary Get appointments by pet
// @Description Retrieves all appointments for a specific pet
// @Tags appointments-query
// @Accept json
// @Produce json
// @Param petID path int true "Pet ID"
// @Param page_number query int false "Page number" default(1)
// @Param page_size query int false "Items per page" default(10)
// @Success 200 {object} response.APIResponse{data=[]query.AppointmentResponse,metadata=PaginationMetadata} "List of appointments"
// @Failure 400 {object} response.APIResponse "Invalid pet ID or pagination parameters"
// @Failure 404 {object} response.APIResponse "Pet not found"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments/pet/{id} [get]
func (controller *AppointmentQueryController) GetAppointmentsByPet(ctx *gin.Context) {
	entityID, err := shared.ParseParamToEntityID(ctx, "id", "id")
	if err != nil {
		response.RequestURLParamError(ctx, err, "petID", ctx.Param("id"))
		return
	}

	petID := entityID.(valueobject.PetID)

	var pageParams PaginationRequest
	if err := ctx.ShouldBindQuery(&pageParams); err != nil {
		response.RequestURLQueryError(ctx, err)
		return
	}
	pageParams.SetDefaultsIfNotProvided()

	query := query.NewGetAppointmentsByPetQuery(
		petID,
		pageParams.PageNumber,
		pageParams.PageSize,
	)

	pageInterface, err := controller.queryBus.Execute(ctx, query)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	HandlePaginatedResponse(ctx, pageInterface, query)
}

// GetAppointmentStats godoc
// @Summary Get appointment statistics
// @Description Retrieves statistical information about appointments
// @Tags appointments-query
// @Accept json
// @Produce json
// @Success 200 {object} response.APIResponse "Appointment statistics"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments/stats [get]
func (controller *AppointmentQueryController) GetAppointmentStats(ctx *gin.Context) {
	query := query.GetAppointmentStatsQuery{}

	pageInterface, err := controller.queryBus.Execute(ctx, query)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	HandlePaginatedResponse(ctx, pageInterface, query)
}

func HandlePaginatedResponse(ctx *gin.Context, pageResponseInterface interface{}, queryParams interface{}) {
	appointmentPage, err := mapInterfaceToPageResponse(pageResponseInterface)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	metadata := gin.H{"page_meta": appointmentPage.Metadata, "request_params": queryParams}
	response.SuccessWithMeta(ctx, &appointmentPage.Data, metadata)
}

func mapInterfaceToPageResponse(qry interface{}) (query.AppointmentPageResponse, error) {
	if appointmentPage, ok := qry.(query.AppointmentPageResponse); ok {
		return appointmentPage, nil
	}
	return query.AppointmentPageResponse{}, errors.New("invalid query type")
}
