// Package controller handles all appointment-related HTTP endpoints
package controller

import (
	"clinic-vet-api/app/modules/appointment/application/query"
	"clinic-vet-api/app/modules/appointment/infrastructure/bus"
	"clinic-vet-api/app/modules/appointment/presentation/dto"
	httpError "clinic-vet-api/app/shared/error/infrastructure/http"
	ginutils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/page"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AppointmentQueryController struct {
	bus            bus.AppointmentBus
	validate       *validator.Validate
	responseMapper *dto.ResponseMapper
}

func NewApptQueryController(
	bus bus.AppointmentBus,
	validate *validator.Validate,
	responseMapper *dto.ResponseMapper,
) *AppointmentQueryController {
	return &AppointmentQueryController{
		bus:            bus,
		validate:       validate,
		responseMapper: responseMapper,
	}
}

type GetByIDExtraArgs struct {
	employeeID *uint
	customerID *uint
}

func (ctrl *AppointmentQueryController) GetAppointmentDetailByID(c *gin.Context, args GetByIDExtraArgs) {
	appointmentID, err := ginutils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	appointmentQuery := query.NewFindApptByIDQuery(
		c.Request.Context(),
		appointmentID,
		args.employeeID,
		args.customerID,
	)

	result, err := ctrl.bus.QueryBus.FindByID(*appointmentQuery)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	appointmentResponse := ctrl.responseMapper.FromResult(result)
	response.Found(c, appointmentResponse, "Appointment")
}

func (ctrl *AppointmentQueryController) FindAppointmentsBySpecification(c *gin.Context) {
	searchSpecification, err := dto.NewApptSearchRequestFromContext(c)
	if err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	if err := ctrl.validate.Struct(searchSpecification); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	searchQuery, err := searchSpecification.ToQuery(c.Request.Context())
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	appointmentPage, err := ctrl.bus.QueryBus.FindBySpecification(*searchQuery)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	ctrl.HandlePaginatedResult(c, appointmentPage, searchSpecification.Pagination())
}

func (ctrl *AppointmentQueryController) FindAppointmentsByDateRange(c *gin.Context) {
	var dateRangeQueryData dto.ListApptByDateRangeRequest
	if err := c.ShouldBindQuery(&dateRangeQueryData); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := ctrl.validate.Struct(&dateRangeQueryData); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	query, err := dateRangeQueryData.ToQuery(c.Request.Context())
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	appointmentPage, err := ctrl.bus.QueryBus.FindByDateRange(query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	ctrl.HandlePaginatedResult(c, appointmentPage, dateRangeQueryData.PageInput.ToMap())
}

func (ctrl *AppointmentQueryController) FindAppointmentsByCustomer(c *gin.Context, customerID uint, extraArgs OwnerQueryExtraArgs) {
	var pageParams page.PageInput
	if err := c.ShouldBindQuery(&pageParams); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	if err := ctrl.validate.Struct(&pageParams); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	query := query.NewFindApptsByCustomerIDQuery(c.Request.Context(), pageParams, customerID, extraArgs.PetID, extraArgs.Status)
	appointPage, err := ctrl.bus.QueryBus.FindByCustomerID(*query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	ctrl.HandlePaginatedResult(c, appointPage, pageParams.ToMap())
}

func (ctrl *AppointmentQueryController) GetAppointmentsByEmployee(c *gin.Context, employeeID uint) {
	var pageParams page.PageInput
	if err := c.ShouldBindQuery(&pageParams); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	query := query.NewFindApptsByEmployeeIDQuery(c.Request.Context(), employeeID, pageParams)
	appointmentPage, err := ctrl.bus.QueryBus.FindByEmployeeID(*query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	ctrl.HandlePaginatedResult(c, appointmentPage, pageParams.ToMap())
}

func (ctrl *AppointmentQueryController) GetAppointmentsByPet(c *gin.Context, petID uint) {
	var pageParams page.PageInput
	if err := c.ShouldBindQuery(&pageParams); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	query := query.NewFindApptsByPetQuery(c.Request.Context(), petID, pageParams)
	appointmentPage, err := ctrl.bus.QueryBus.FindByPetID(*query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	ctrl.HandlePaginatedResult(c, appointmentPage, pageParams.ToMap())
}

func (ctrl *AppointmentQueryController) GetAppointmentStats(c *gin.Context) {
}

func (ctrl *AppointmentQueryController) HandlePaginatedResult(c *gin.Context, pageResponse page.Page[query.ApptResult], queryParams map[string]any) {
	appointmentsResponse := ctrl.responseMapper.FromResults(pageResponse.Items)
	metadata := gin.H{"page_meta": pageResponse.Metadata, "request_params": queryParams}
	response.SuccessWithMeta(c, appointmentsResponse, "Appointment ", metadata)
}
