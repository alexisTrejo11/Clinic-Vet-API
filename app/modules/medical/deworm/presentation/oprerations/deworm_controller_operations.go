package operations

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/medical/deworm/application"
	"clinic-vet-api/app/modules/medical/deworm/application/command"
	"clinic-vet-api/app/modules/medical/deworm/application/query"
	"clinic-vet-api/app/modules/medical/deworm/presentation/dto"
	"clinic-vet-api/app/shared/error/infrastructure/http"
	ginutils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/page"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type DewormControllerOperations struct {
	dewormService  application.DewormingFacadeService
	validate       *validator.Validate
	responseMapper dto.DewormResponseMapper
}

func NewDewormControllerOperations(
	dewormService application.DewormingFacadeService,
	validate *validator.Validate,
) *DewormControllerOperations {
	return &DewormControllerOperations{
		dewormService: dewormService,
		validate:      validate,
	}
}

type FindByDewormsExtraArgs struct {
	PetID       *uint
	AppliedByID *uint
	CustomerID  *uint
}

// Query Operations

func (d *DewormControllerOperations) FindDewormByID(c *gin.Context, extraArgs FindByDewormsExtraArgs) {
	dewormID, err := ginutils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, http.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	query := &query.FindDewormByIDQuery{
		ID:            valueobject.NewDewormID(dewormID),
		OptPetID:      valueobject.NewOptPetID(extraArgs.PetID),
		OptEmployeeID: valueobject.NewOptEmployeeID(extraArgs.AppliedByID),
	}

	result, err := d.dewormService.FindDewormByID(c.Request.Context(), *query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	dewormResponse := d.responseMapper.FromResult(result)
	response.Found(c, dewormResponse, "Pet Deworm")
}

func (d *DewormControllerOperations) FindDewormsByCustomerID(c *gin.Context, customerID uint, extraArgs FindByDewormsExtraArgs) {
	var pagination page.PaginationRequest
	if err := ginutils.ShouldBindPageParams(&pagination, c, d.validate); err != nil {
		response.BadRequest(c, err)
		return
	}

	query := &query.FindDewormsByCustomerQuery{
		CustomerID: valueobject.NewCustomerID(customerID),
		OptPetID:   valueobject.NewOptPetID(extraArgs.PetID),
		Pagination: pagination,
	}

	results, err := d.dewormService.FindDewormsByCustomer(c.Request.Context(), *query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	dewormResponses := d.responseMapper.FromResultsToResponses(results.Items)
	response.SuccessWithPagination(c, dewormResponses, "Pet Deworms Successfully Retrieved", results.Metadata)
}

func (d *DewormControllerOperations) FindDewormsByPetID(c *gin.Context, petID uint, extraArgs FindByDewormsExtraArgs) {
	var pagination page.PaginationRequest
	if err := ginutils.ShouldBindPageParams(&pagination, c, d.validate); err != nil {
		response.BadRequest(c, err)
		return
	}

	query := &query.FindDewormsByPetQuery{
		PetID:         valueobject.NewPetID(petID),
		OptEmployeeID: valueobject.NewOptEmployeeID(extraArgs.AppliedByID),
		Pagination:    pagination,
	}

	results, err := d.dewormService.FindDewormsByPet(c.Request.Context(), *query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	dewormResponses := d.responseMapper.FromResultsToResponses(results.Items)
	response.SuccessWithPagination(c, dewormResponses, "Pet Deworms Successfully Retrieved", results.Metadata)
}

func (d *DewormControllerOperations) FindDewormsByEmployeeID(c *gin.Context, employeeID uint) {
	var pagination page.PaginationRequest
	if err := ginutils.ShouldBindPageParams(&pagination, c, d.validate); err != nil {
		response.BadRequest(c, err)
		return
	}

	query := &query.FindDewormsByEmployeeQuery{
		EmployeeID: valueobject.NewEmployeeID(employeeID),
		Pagination: pagination,
	}

	results, err := d.dewormService.FindDewormsByEmployee(c.Request.Context(), *query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}
	dewormResponses := d.responseMapper.FromResultsToResponses(results.Items)
	response.SuccessWithPagination(c, dewormResponses, "Pet Deworms Successfully Retrieved", results.Metadata)
}

func (d *DewormControllerOperations) FindDewormsByDateRange(c *gin.Context) {
	var req dto.FindDewormsByDateRangeRequest
	if err := ginutils.BindAndValidateQuery(c, &req, d.validate); err != nil {
		response.BadRequest(c, err)
		return
	}

	query := req.ToQuery()
	results, err := d.dewormService.FindDewormsByDateRange(c.Request.Context(), *query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	dewormResponses := d.responseMapper.FromResultsToResponses(results.Items)
	response.SuccessWithPagination(c, dewormResponses, "Pet Deworms Successfully Retrieved", results.Metadata)
}

// Command Operations

func (d *DewormControllerOperations) RegisterDeworm(c *gin.Context) {
	var req dto.CreateDewormRequest
	if err := ginutils.BindAndValidateBody(c, &req, d.validate); err != nil {
		response.BadRequest(c, err)
		return
	}

	command := req.ToCommand()
	result := d.dewormService.RegisterDeworm(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Created(c, result.ID(), "Deworm")
}

func (d *DewormControllerOperations) UpdateDeworm(c *gin.Context) {
	dewormID, err := ginutils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, http.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	var req dto.UpdateDewormRequest
	if err := ginutils.BindAndValidateBody(c, &req, d.validate); err != nil {
		response.BadRequest(c, err)
		return
	}

	command := req.ToCommand(dewormID)
	result := d.dewormService.UpdateDeworm(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Updated(c, nil, "Deworm")
}

func (d *DewormControllerOperations) DeleteDeworm(c *gin.Context) {
	dewormID, err := ginutils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, http.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	command := &command.DewormDeleteCommand{ID: valueobject.NewDewormID(dewormID)}
	result := d.dewormService.DeleteDeworm(c.Request.Context(), *command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.NoContent(c)
}
