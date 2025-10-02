package controller

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"

	command "clinic-vet-api/app/modules/medical/session/application/command"
	facade "clinic-vet-api/app/modules/medical/session/application/facade_service"
	query "clinic-vet-api/app/modules/medical/session/application/query"
	"clinic-vet-api/app/modules/medical/session/presentation/dto"
	httpError "clinic-vet-api/app/shared/error/infrastructure/http"
	ginUtils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/page"
	"clinic-vet-api/app/shared/response"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MedSessionControllerOperations struct {
	facade.MedicalApplicationService
	validator *validator.Validate
}

func NewMedSessionControllerOperations(bus facade.MedicalApplicationService, validator *validator.Validate) *MedSessionControllerOperations {
	return &MedSessionControllerOperations{
		MedicalApplicationService: bus,
		validator:                 validator,
	}
}

type GetByIDExtraArgs struct {
	CustomerID *uint
	EmployeeID *uint
	PetID      *uint
}

func (co *MedSessionControllerOperations) GetMedSessionsByID(c *gin.Context, getByIDExtraArgs *GetByIDExtraArgs) {
	idUint, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "medical-history", c.Param("id")))
		return
	}

	var getByIDQuery *query.FindMedSessionByIDQuery

	if getByIDExtraArgs == nil {
		getByIDQuery = query.NewFindMedSessionByIDQuery(idUint)
	} else {
		if getByIDExtraArgs.CustomerID != nil {
			getByIDQuery = query.FindMedSessionByIDQueryWithCustomerID(idUint, *getByIDExtraArgs.CustomerID)
		} else if getByIDExtraArgs.PetID != nil {
			getByIDQuery = query.FindMedSessionByIDQueryWithPetID(idUint, *getByIDExtraArgs.PetID)
		} else if getByIDExtraArgs.EmployeeID != nil {
			getByIDQuery = query.FindMedSessionByIDQueryWithEmployeeID(idUint, *getByIDExtraArgs.EmployeeID)
		}
	}

	result, err := co.QueryBus().FindMedSessionByID(c.Request.Context(), *getByIDQuery)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	medSessionResponse := dto.FromResult(result)
	response.Found(c, medSessionResponse, "Medical Session")
}

func (co *MedSessionControllerOperations) GetMedSessionsByPetID(c *gin.Context, petID uint, customerID *uint) {
	var pagination page.PaginationRequest
	if err := ginUtils.ShouldBindPageParams(&pagination, c, co.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	query := query.NewFindMedSessionByPetIDQuery(petID, customerID, pagination)
	resultPage, err := co.QueryBus().FindMedSessionByPetID(c.Request.Context(), *query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	medSessionResponse := dto.FromResultList(resultPage.Items)
	response.SuccessWithPagination(c, medSessionResponse, "Medical Sessions", resultPage.Metadata)
}

func (co *MedSessionControllerOperations) GetMedSessionsByEmployeeID(c *gin.Context, employeeID uint) {
	var pagination page.PaginationRequest
	if err := ginUtils.ShouldBindPageParams(&pagination, c, co.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	query := &query.FindMedSessionByEmployeeIDQuery{
		EmployeeID:        valueobject.NewEmployeeID(employeeID),
		PaginationRequest: pagination,
	}

	resultPage, err := co.QueryBus().FindMedSessionByEmployeeID(c.Request.Context(), *query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	medSessionResponse := dto.FromResultList(resultPage.Items)
	response.SuccessWithPagination(c, medSessionResponse, "Medical Sessions", resultPage.Metadata)
}

func (co *MedSessionControllerOperations) GetMedSessionsByDateRange(c *gin.Context, startDate, endDate time.Time) {
	var pagination page.PaginationRequest
	if err := ginUtils.ShouldBindPageParams(&pagination, c, co.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	query := &query.FindMedSessionByDateRangeQuery{StartDate: startDate, EndDate: endDate, PaginationRequest: pagination}
	resultPage, err := co.QueryBus().FindMedSessionByDateRange(c.Request.Context(), *query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	medSessionResponse := dto.FromResultList(resultPage.Items)
	response.SuccessWithPagination(c, medSessionResponse, "Medical Sessions", resultPage.Metadata)
}

func (co MedSessionControllerOperations) GetMedicalSessionByCustomerID(c *gin.Context, customerID uint, petID *uint) {
	var pagination page.PaginationRequest
	if err := ginUtils.ShouldBindPageParams(&pagination, c, co.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	query := &query.FindMedSessionByCustomerIDQuery{CustomerID: valueobject.NewCustomerID(customerID), PaginationRequest: pagination}
	resultPage, err := co.QueryBus().FindMedSessionByCustomerID(c.Request.Context(), *query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	medSessionResponse := dto.FromResultList(resultPage.Items)
	response.SuccessWithPagination(c, medSessionResponse, "Medical Sessions", resultPage.Metadata)
}

func (co *MedSessionControllerOperations) CreateMedicalSession(c *gin.Context, employeeID *uint) {
	var requestData dto.AdminCreateMedSessionRequest
	if err := ginUtils.ShouldBindAndValidateBody(c, &requestData, co.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	command := requestData.ToCommand()
	result := co.CommandBus().CreateMedicalSession(c.Request.Context(), *command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Created(c, result.ID, "Medical Session")
}

func (co *MedSessionControllerOperations) FindMedSessionsByDateRange(c *gin.Context, petID uint, startDate, endDate time.Time) {
	var pagination page.PaginationRequest
	if err := ginUtils.ShouldBindPageParams(&pagination, c, co.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	query := &query.FindMedSessionByDateRangeQuery{StartDate: startDate, EndDate: endDate, PaginationRequest: pagination}
	resultPage, err := co.QueryBus().FindMedSessionByDateRange(c.Request.Context(), *query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	medSessionResponse := dto.FromResultList(resultPage.Items)
	response.SuccessWithPagination(c, medSessionResponse, "Medical Sessions", resultPage.Metadata)
}

func (co *MedSessionControllerOperations) DeleteMedicalSession(c *gin.Context) {
	idUint, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "medical-history", c.Param("id")))
		return
	}

	command := &command.DeleteMedSessionCommand{ID: valueobject.NewMedSessionID(idUint)}
	result := co.CommandBus().DeleteMedSessionCommand(c.Request.Context(), *command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
}
