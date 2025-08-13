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

// Get appointment by ID
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

// Get all appointments with pagination
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

// Get appointments by date range
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

// Get appointments by owner
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

// Get appointments by vet
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

// Get appointments by pet
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

// Get appointment statistics
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
