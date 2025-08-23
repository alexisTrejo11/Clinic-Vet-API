// Package appointmentController handles all appointment-related HTTP endpoints
package appointmentController

import (
	"errors"

	appointmentQuery "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/application/queries"
	appControllerDTO "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/infrastructure/api/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	response "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// AppointmentQueryController handles appointment query operations
// @title Veterinary Clinic API - Appointment Queries
// @version 1.0
// @description This controller manages appointment queries including retrieving appointments by various criteria
type AppointmentQueryController struct {
	queryBus appointmentQuery.QueryBus
	validate *validator.Validate
}

func NewAppointmentQueryController(
	queryBus appointmentQuery.QueryBus,
	validate *validator.Validate,
) *AppointmentQueryController {
	return &AppointmentQueryController{
		queryBus: queryBus,
		validate: validate,
	}
}

// GetAppointmentById godoc
// @Summary Get appointment by ID
// @Description Retrieves detailed information about a specific appointment
// @Tags appointments-query
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID"
// @Success 200 {object} appointmentQuery.AppointmentResponse "Appointment details"
// @Failure 400 {object} response.APIResponse "Invalid appointment ID"
// @Failure 404 {object} response.APIResponse "Appointment not found"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments/{id} [get]
func (controller *AppointmentQueryController) GetAppointmentById(ctx *gin.Context) {
	appointmentId, err := shared.ParseParamToInt(ctx, "id")
	if err != nil {
		response.RequestURLQueryError(ctx, err)
		return
	}

	query := appointmentQuery.NewGetAppointmentByIdQuery(appointmentId)

	appointmentResponse, err := controller.queryBus.Execute(ctx, query)
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
// @Success 200 {object} response.APIResponse{data=[]appointmentQuery.AppointmentResponse,metadata=PaginationMetadata} "List of appointments"
// @Failure 400 {object} response.APIResponse "Invalid pagination parameters"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments [get]
func (controller *AppointmentQueryController) GetAllAppointments(ctx *gin.Context) {
	var pageParams appControllerDTO.PaginationRequest

	if err := ctx.ShouldBindQuery(&pageParams); err != nil {
		response.RequestURLQueryError(ctx, err)
		return
	}

	pageParams.SetDefaultsIfNotProvided()
	query := appointmentQuery.NewGetAllAppointmentsQuery(pageParams.PageNumber, pageParams.PageNumber)

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
// @Success 200 {object} response.APIResponse{data=[]appointmentQuery.AppointmentResponse,metadata=PaginationMetadata} "List of appointments"
// @Failure 400 {object} response.APIResponse "Invalid date range or pagination parameters"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments/date-range [get]
func (controller *AppointmentQueryController) GetAppointmentsByDateRange(ctx *gin.Context) {
	var queryParams appControllerDTO.GetAppointmentsByDateRangeRequest

	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		response.RequestURLQueryError(ctx, err)
		return
	}

	if err := controller.validate.Struct(&queryParams); err != nil {
		response.RequestBodyDataError(ctx, err)
		return
	}

	query := appointmentQuery.NewGetAppointmentsByDateRangeQuery(
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
// @Param ownerId path int true "Owner ID"
// @Param page_number query int false "Page number" default(1)
// @Param page_size query int false "Items per page" default(10)
// @Success 200 {object} response.APIResponse{data=[]appointmentQuery.AppointmentResponse,metadata=PaginationMetadata} "List of appointments"
// @Failure 400 {object} response.APIResponse "Invalid owner ID or pagination parameters"
// @Failure 404 {object} response.APIResponse "Owner not found"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments/owner/{ownerId} [get]
func (controller *AppointmentQueryController) GetAppointmentsByOwner(ctx *gin.Context) {
	ownerId, err := shared.ParseParamToInt(ctx, "ownerId")
	if err != nil {
		response.RequestURLParamError(ctx, err, "ownerId", ctx.Param("ownerId"))
		return
	}

	var pageParams appControllerDTO.PaginationRequest
	if err := ctx.ShouldBindQuery(&pageParams); err != nil {
		response.RequestURLQueryError(ctx, err)
		return
	}

	pageParams.SetDefaultsIfNotProvided()
	query := appointmentQuery.NewGetAppointmentsByOwnerQuery(ownerId, pageParams.PageNumber, pageParams.PageSize)

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
// @Param vetId path int true "Veterinarian ID"
// @Param page_number query int false "Page number" default(1)
// @Param page_size query int false "Items per page" default(10)
// @Success 200 {object} response.APIResponse{data=[]appointmentQuery.AppointmentResponse,metadata=PaginationMetadata} "List of appointments"
// @Failure 400 {object} response.APIResponse "Invalid veterinarian ID or pagination parameters"
// @Failure 404 {object} response.APIResponse "Veterinarian not found"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments/vet/{vetId} [get]
func (controller *AppointmentQueryController) GetAppointmentsByVet(ctx *gin.Context) {
	vetId, err := shared.ParseParamToInt(ctx, "vetId")
	if err != nil {
		response.RequestURLParamError(ctx, err, "vetId", ctx.Param("vetId"))
		return
	}

	var pageParams appControllerDTO.PaginationRequest
	if err := ctx.ShouldBindQuery(&pageParams); err != nil {
		response.RequestURLQueryError(ctx, err)
		return
	}
	pageParams.SetDefaultsIfNotProvided()

	query := appointmentQuery.NewGetAppointmentsByVetQuery(
		vetId,
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
// @Param petId path int true "Pet ID"
// @Param page_number query int false "Page number" default(1)
// @Param page_size query int false "Items per page" default(10)
// @Success 200 {object} response.APIResponse{data=[]appointmentQuery.AppointmentResponse,metadata=PaginationMetadata} "List of appointments"
// @Failure 400 {object} response.APIResponse "Invalid pet ID or pagination parameters"
// @Failure 404 {object} response.APIResponse "Pet not found"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments/pet/{petId} [get]
func (controller *AppointmentQueryController) GetAppointmentsByPet(ctx *gin.Context) {
	petId, err := shared.ParseParamToInt(ctx, "petId")
	if err != nil {
		response.RequestURLParamError(ctx, err, "petId", ctx.Param("petId"))
		return
	}

	var pageParams appControllerDTO.PaginationRequest
	if err := ctx.ShouldBindQuery(&pageParams); err != nil {
		response.RequestURLQueryError(ctx, err)
		return
	}
	pageParams.SetDefaultsIfNotProvided()

	query := appointmentQuery.NewGetAppointmentsByPetQuery(
		petId,
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
	query := appointmentQuery.GetAppointmentStatsQuery{}

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

func mapInterfaceToPageResponse(query interface{}) (appointmentQuery.AppointmentPageResponse, error) {
	if appointmentPage, ok := query.(appointmentQuery.AppointmentPageResponse); ok {
		return appointmentPage, nil
	}
	return appointmentQuery.AppointmentPageResponse{}, errors.New("invalid query type")
}
