package appointmentController

import (
	"net/http"
	"strconv"

	"errors"

	appointmentQuery "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/application/queries"
	appControllerDTO "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/infrastructure/api/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
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
	appointmentId, err := shared.ParseID(ctx, "id")
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
	pageParam := ctx.DefaultQuery("pageNumber", "1")
	limitParam := ctx.DefaultQuery("pageSize", "10")

	pageNumber, err := strconv.Atoi(pageParam)
	if err != nil {
		response.RequestURLQueryError(ctx, err)
		return
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		response.RequestURLQueryError(ctx, err)
		return
	}

	query := appointmentQuery.NewGetAllAppointmentsQuery(pageNumber, limit)

	appointment, err := controller.queryBus.Execute(ctx, query)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	response.Success(ctx, appointment.(page.Page[[]appointmentQuery.AppointmentResponse]))
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

	query := appointmentQuery.NewGetAppointmentsByOwnerQuery(ownerId, page, limit)

	pageInterface, err := controller.queryBus.Execute(ctx, query)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	HandlePaginatedResponse(ctx, pageInterface, query)
}

// Get appointments by vet
func (controller *AppointmentQueryController) GetAppointmentsByVet(ctx *gin.Context) {
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

	query := appointmentQuery.NewGetAppointmentsByVetQuery(vetId, page, limit)

	pageInterface, err := controller.queryBus.Execute(ctx, query)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	HandlePaginatedResponse(ctx, pageInterface, query)
}

// Get appointments by pet
func (controller *AppointmentQueryController) GetAppointmentsByPet(ctx *gin.Context) {
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

	query := appointmentQuery.NewGetAppointmentsByPetQuery(petId, page, limit)

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
