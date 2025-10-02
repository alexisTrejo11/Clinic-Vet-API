package controller

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/medical/vaccination/application"
	"clinic-vet-api/app/modules/medical/vaccination/application/query"
	"clinic-vet-api/app/modules/medical/vaccination/presentation/dto"
	autherror "clinic-vet-api/app/shared/error/auth"
	ginutils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/page"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type EmployeePetVaccinationController struct {
	vaccinationService application.VaccinationFacadeService
	validator          *validator.Validate
}

func NewEmployeePetVaccinationController(
	vaccinationService application.VaccinationFacadeService,
	validator *validator.Validate,
) *EmployeePetVaccinationController {
	return &EmployeePetVaccinationController{
		vaccinationService: vaccinationService,
		validator:          validator,
	}
}

func (ctrl *EmployeePetVaccinationController) RegisterNewVaccination(c *gin.Context) {
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

	var req dto.RegisterVaccineRequest
	if err := ginutils.ShouldBindAndValidateBody(c, &req, ctrl.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	cmd, err := req.ToCommand(petID, user.EmployeeID)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.vaccinationService.RegisterVaccine(c.Request.Context(), cmd)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Created(c, result.ID(), "PetVaccination")
}

func (ctrl *EmployeePetVaccinationController) UpdateMyVaccinationApplied(c *gin.Context) {
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

	var req dto.UpdateVaccineRequest
	if err := ginutils.ShouldBindAndValidateBody(c, &req, ctrl.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	cmd, err := req.ToCommand(petID, user.EmployeeID)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.vaccinationService.UpdateVaccine(c.Request.Context(), cmd)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Updated(c, nil, "PetVaccination")

}

func (ctrl *EmployeePetVaccinationController) GetMyVaccinationAppliedDetail(c *gin.Context) {
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

	vaccinationID, err := ginutils.ParseParamToUInt(c, "vaccination_id")
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	query, err := query.NewFindVaccinationByIDQuery(vaccinationID, &petID, &user.EmployeeID)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	result, err := ctrl.vaccinationService.FindVaccinationByID(c.Request.Context(), query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Found(c, result, "PetVaccination")
}

func (ctrl *EmployeePetVaccinationController) GetMyVaccinationsHistoryByPet(c *gin.Context) {
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

	query, err := query.NewFindVaccinationsByPetQuery(petID, &user.EmployeeID, pagination)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	results, err := ctrl.vaccinationService.FindVaccinationsByPet(c.Request.Context(), query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Found(c, results, "PetVaccinations")
}

func (ctrl *EmployeePetVaccinationController) GetMyVaccinationHistory(c *gin.Context) {
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

	query, err := query.NewFindVaccinationsByEmployeeQuery(user.EmployeeID, pagination)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	results, err := ctrl.vaccinationService.FindVaccinationsByEmployee(c.Request.Context(), query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Found(c, results, "PetVaccinations")
}
