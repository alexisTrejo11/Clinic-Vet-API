// Package controller handles all appointment-related HTTP endpoints
package controller

import (
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/application/query"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/infrastructure/api/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	httpError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure/http"
	ginUtils "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/gin_utils"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ApptQueryController handles appointment query operations
// @title Veterinary Clinic API - Appt Queries
// @version 1.0
// @description This controller manages appointment queries including retrieving appointments by various criteria
type ApptQueryController struct {
	queryBus cqrs.QueryBus
	validate *validator.Validate
}

func NewApptQueryController(
	queryBus cqrs.QueryBus,
	validate *validator.Validate,
) *ApptQueryController {
	return &ApptQueryController{
		queryBus: queryBus,
		validate: validate,
	}
}

// GetApptByID godoc
// @Summary Get appointment by ID
// @Description Retrieves detailed information about a specific appointment
// @Tags appointments-query
// @Accept json
// @Produce json
// @Param id path int true "Appt ID"
// @Success 200 {object} query.ApptResponse "Appointment details"
// @Failure 400 {object} response.APIResponse "Invalid appointment ID"
// @Failure 404 {object} response.APIResponse "Appt not found"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments/{id} [get]
func (controller *ApptQueryController) GetAppointmentByID(c *gin.Context) {
	appointmentID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	appointmentQuery := query.NewGetApptByIDQuery(c.Request.Context(), appointmentID)
	appointmentResponse, err := controller.queryBus.Execute(appointmentQuery)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Success(c, appointmentResponse)
}

// GetAllAppts godoc
// @Summary Get all appointments
// @Description Retrieves a paginated list of all appointments
// @Tags appointments-query
// @Accept json
// @Produce json
// @Param page_number query int false "Page number" default(1)
// @Param page_size query int false "Items per page" default(10)
// @Success 200 {object} response.APIResponse{data=[]query.ApptResponse,metadata=PaginationMetadata} "List of appointments"
// @Failure 400 {object} response.APIResponse "Invalid pagination parameters"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments [get]
func (controller *ApptQueryController) SearchAppointments(c *gin.Context) {
	searchSpecification, err := dto.NewApptSearchRequestFromContext(c)
	if err := controller.validate.Struct(searchSpecification); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	appointmentPage, err := controller.queryBus.Execute(searchSpecification)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	HandlePaginatedResponse(c, appointmentPage, appointmentPage)
}

// GetApptsByDateRange godoc
// @Summary Get appointments by date range
// @Description Retrieves appointments within a specified date range
// @Tags appointments-query
// @Accept json
// @Produce json
// @Param start_date query string true "Start date (YYYY-MM-DD)" format(date)
// @Param end_date query string true "End date (YYYY-MM-DD)" format(date)
// @Param page_number query int false "Page number" default(1)
// @Param page_size query int false "Items per page" default(10)
// @Success 200 {object} response.APIResponse{data=[]query.ApptResponse,metadata=PaginationMetadata} "List of appointments"
// @Failure 400 {object} response.APIResponse "Invalid date range or pagination parameters"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments/date-range [get]
func (controller *ApptQueryController) GetAppointmentsByDateRange(c *gin.Context) {
	var dateRangeQueryData dto.ListApptByDateRangeRequest

	if err := c.ShouldBindQuery(&dateRangeQueryData); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := controller.validate.Struct(&dateRangeQueryData); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	listByDateRangeQuery, err := dateRangeQueryData.ToQuery()
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	appointmentPage, err := controller.queryBus.Execute(listByDateRangeQuery)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	HandlePaginatedResponse(c, appointmentPage, dateRangeQueryData.PageInput)
}

// GetApptsByOwner godoc
// @Summary Get appointments by owner
// @Description Retrieves all appointments for a specific pet owner
// @Tags appointments-query
// @Accept json
// @Produce json
// @Param ownerID path int true "Owner ID"
// @Param page_number query int false "Page number" default(1)
// @Param page_size query int false "Items per page" default(10)
// @Success 200 {object} response.APIResponse{data=[]query.ApptResponse,metadata=PaginationMetadata} "List of appointments"
// @Failure 400 {object} response.APIResponse "Invalid owner ID or pagination parameters"
// @Failure 404 {object} response.APIResponse "Owner not found"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments/owner/{id} [get]
func (controller *ApptQueryController) GetAppointmentsByOwner(c *gin.Context) {
	ownerID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	var pageParams page.PageInput
	if err := c.ShouldBindQuery(&pageParams); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	if err := controller.validate.Struct(&pageParams); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	listAppointByOwnerQuery := query.NewListApptByOwnerQuery(c.Request.Context(), ownerID, pageParams)
	appointPage, err := controller.queryBus.Execute(listAppointByOwnerQuery)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	HandlePaginatedResponse(c, appointPage, pageParams)
}

// GetApptsByVet godoc
// @Summary Get appointments by veterinarian
// @Description Retrieves all appointments assigned to a specific veterinarian
// @Tags appointments-query
// @Accept json
// @Produce json
// @Param vetID path int true "Veterinarian ID"
// @Param page_number query int false "Page number" default(1)
// @Param page_size query int false "Items per page" default(10)
// @Success 200 {object} response.APIResponse{data=[]query.ApptResponse,metadata=PaginationMetadata} "List of appointments"
// @Failure 400 {object} response.APIResponse "Invalid veterinarian ID or pagination parameters"
// @Failure 404 {object} response.APIResponse "Veterinarian not found"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments/vet/{id} [get]
func (controller *ApptQueryController) GetAppointmentsByVet(c *gin.Context) {
	ownerID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	var pageParams page.PageInput
	if err := c.ShouldBindQuery(&pageParams); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	query := query.NewListApptByVetQuery(c.Request.Context(), ownerID, pageParams)
	appointmentPage, err := controller.queryBus.Execute(query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	HandlePaginatedResponse(c, appointmentPage, pageParams)
}

// GetApptsByPet godoc
// @Summary Get appointments by pet
// @Description Retrieves all appointments for a specific pet
// @Tags appointments-query
// @Accept json
// @Produce json
// @Param petID path int true "Pet ID"
// @Param page_number query int false "Page number" default(1)
// @Param page_size query int false "Items per page" default(10)
// @Success 200 {object} response.APIResponse{data=[]query.ApptResponse,metadata=PaginationMetadata} "List of appointments"
// @Failure 400 {object} response.APIResponse "Invalid pet ID or pagination parameters"
// @Failure 404 {object} response.APIResponse "Pet not found"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments/pet/{id} [get]
func (controller *ApptQueryController) GetAppointmentsByPet(c *gin.Context) {
	petID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	var pageParams page.PageInput
	if err := c.ShouldBindQuery(&pageParams); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	listByPetQuery := query.NewListApptByPetQuery(c.Request.Context(), petID, pageParams)
	appointmentPage, err := controller.queryBus.Execute(listByPetQuery)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	HandlePaginatedResponse(c, appointmentPage, pageParams)
}

// GetApptStats godoc
// @Summary Get appointment statistics
// @Description Retrieves statistical information about appointments
// @Tags appointments-query
// @Accept json
// @Produce json
// @Success 200 {object} response.APIResponse "Appt statistics"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /appointments/stats [get]
func (controller *ApptQueryController) GetAppointmentStats(c *gin.Context) {
	query := query.GetAppointmentStatsHandler{}

	appointmentStats, err := controller.queryBus.Execute(query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Success(c, appointmentStats)
}

func HandlePaginatedResponse(c *gin.Context, pageResponseInterface any, queryParams any) {
	appointmentPage, err := mapInterfaceToPageResponse(pageResponseInterface)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	metadata := gin.H{"page_meta": appointmentPage.Metadata, "request_params": queryParams}
	response.SuccessWithMeta(c, appointmentPage.Data, metadata)
}

func mapInterfaceToPageResponse(qry any) (page.Page[query.ApptResponse], error) {
	if appointmentPage, ok := qry.(page.Page[query.ApptResponse]); ok {
		return appointmentPage, nil
	}
	return page.Page[query.ApptResponse]{}, errors.New("invalid query type")
}
