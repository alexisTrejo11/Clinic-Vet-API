package service

import (
	"clinic-vet-api/app/modules/pets/application/cqrs/query"
	"clinic-vet-api/app/modules/pets/presentation/dto"
	httpError "clinic-vet-api/app/shared/error/infrastructure/http"
	ginutils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/page"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
)

func (s *PetControllerOperations) FindPetByID(c *gin.Context, customerID *uint) {
	petID, err := ginutils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	query := query.NewFindPetByIDQuery(petID, customerID)
	result, err := s.bus.FindPetByID(c.Request.Context(), *query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	petResponse := dto.ToResponse(result)
	response.Found(c, petResponse, "Pet")
}

func (s *PetControllerOperations) FindPetsByCustomerID(c *gin.Context, customerID uint) {
	var paginatioRequest page.PaginationRequest
	if err := ginutils.ShouldBindPageParams(&paginatioRequest, c, s.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	query := query.NewFindPetsByCustomerIDQuery(customerID, paginatioRequest)
	resultPage, err := s.bus.FindPetsByCustomerID(c.Request.Context(), *query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	petResults := dto.ToResponseList(resultPage.Items)
	response.SuccessWithPagination(c, petResults, "Pets found successfully", resultPage.Metadata)
}

func (s *PetControllerOperations) FindPetsBySpecification(c *gin.Context) {
	var requestURLParams dto.PetSearchRequest
	if err := ginutils.BindAndValidateQuery(c, &requestURLParams, s.validator); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	query := requestURLParams.ToQuery()
	resultPage, err := s.bus.FindPetBySpecification(c.Request.Context(), query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	petResults := dto.ToResponseList(resultPage.Items)
	response.SuccessWithPagination(c, petResults, "Pets found successfully", resultPage.Metadata)
}

func (s *PetControllerOperations) FindPetsBySpecies(c *gin.Context) {
	petSpecies := c.Query("species")
	if petSpecies == "" {
		response.BadRequest(c, httpError.RequestURLQueryError(nil, c.Request.URL.RawQuery))
		return
	}

	var pagination page.PaginationRequest
	if err := ginutils.ShouldBindPageParams(&pagination, c, s.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	query := query.NewFindPetsBySpeciesQuery(petSpecies, pagination)
	resultPage, err := s.bus.FindPetsBySpecies(c.Request.Context(), *query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	petResults := dto.ToResponseList(resultPage.Items)
	response.SuccessWithPagination(c, petResults, "Pets found successfully", resultPage.Metadata)
}
