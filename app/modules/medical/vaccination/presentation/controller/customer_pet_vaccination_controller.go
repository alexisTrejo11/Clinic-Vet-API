package controller

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/medical/vaccination/application"
	"clinic-vet-api/app/modules/medical/vaccination/application/query"
	autherror "clinic-vet-api/app/shared/error/auth"
	ginutils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/page"
	"clinic-vet-api/app/shared/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CustomerPetVaccinationController struct {
	service   application.VaccinationFacadeService
	validator *validator.Validate
}

func NewCustomerPetVaccinationController(service application.VaccinationFacadeService, validator *validator.Validate) *CustomerPetVaccinationController {
	return &CustomerPetVaccinationController{
		service:   service,
		validator: validator,
	}
}

func (ctrl *CustomerPetVaccinationController) GetByMyPetsVaccinationHistory(c *gin.Context) {
	user, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	var pagination page.PaginationRequest
	if err := ginutils.ShouldBindPageParams(&pagination, c, ctrl.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	petID := c.Param("pet_id")
	var petIDUint *uint
	if petID != "" {
		petIDParsed, err := strconv.ParseUint(petID, 10, 32)
		if err != nil {
			response.BadRequest(c, err)
			return
		}
		petIDUint = new(uint)
		*petIDUint = uint(petIDParsed)
	}

	query, err := query.NewFindVaccinationsByCustomerQuery(user.CustomerID, petIDUint, pagination)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	results, err := ctrl.service.FindVaccinationsByCustomer(c.Request.Context(), query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Found(c, results, "PetVaccionations")
}

func (ctrl *CustomerPetVaccinationController) GetByMyPetVaccinationHistory(c *gin.Context) {
	user, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	petID, err := ginutils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	var pagination page.PaginationRequest
	if err := ginutils.ShouldBindPageParams(&pagination, c, ctrl.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	query, err := query.NewFindVaccinationsByCustomerQuery(user.CustomerID, &petID, pagination)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	results, err := ctrl.service.FindVaccinationsByCustomer(c.Request.Context(), query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Found(c, results, "PetVaccinations")

}

func (ctrl *CustomerPetVaccinationController) GetByMyPetVaccinationHistoryDetail(c *gin.Context) {
	user, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	petID, err := ginutils.ParseParamToUInt(c, "pet_id")
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	vaccinationID, err := ginutils.ParseParamToUInt(c, "vaccination_id")
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	query, err := query.NewFindVaccinationByIDQuery(vaccinationID, &petID, &user.CustomerID)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result, err := ctrl.service.FindVaccinationByID(c.Request.Context(), query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Found(c, result, "PetVaccination")
}
