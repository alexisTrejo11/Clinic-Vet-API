// Package controller defines the HTTP controllers for the customer module.
package controller

import (
	"clinic-vet-api/app/modules/customer/application/command"
	"clinic-vet-api/app/modules/customer/application/query"
	"clinic-vet-api/app/modules/customer/infrastructure/bus"
	"clinic-vet-api/app/modules/customer/presentation/dto"
	httpError "clinic-vet-api/app/shared/error/infrastructure/http"
	ginUtils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CustomerController struct {
	validator *validator.Validate
	bus       bus.CustomerBus
}

func NewCustomerController(validator *validator.Validate, bus bus.CustomerBus) *CustomerController {
	return &CustomerController{
		validator: validator,
		bus:       bus,
	}
}

func (ctrl *CustomerController) SearchCustomers(c *gin.Context) {
	var searchParams dto.CustomerSearchQuery
	if err := c.ShouldBindQuery(&searchParams); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	if err := ctrl.validator.Struct(&searchParams); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	query, err := searchParams.ToQuery()
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	customer, err := ctrl.bus.FindCustomerByCriteria(c.Request.Context(), query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}
	response.Found(c, customer, "Customer")
}

func (ctrl *CustomerController) GetCustomerByID(c *gin.Context) {
	id, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "customer_id", c.Param("id")))
		return
	}

	query, err := query.NewFindCustomerByIDQuery(id)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result, err := ctrl.bus.FindCustomerByID(c.Request.Context(), query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	customerResponse := dto.FromResult(result)
	response.Found(c, customerResponse, "Customer")
}

func (ctrl *CustomerController) CreateCustomer(c *gin.Context) {
	var customerCreate dto.CreateCustomerRequest
	if err := ginUtils.ShouldBindAndValidateBody(c, &customerCreate, ctrl.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	command, err := customerCreate.ToCommand()
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.CreateCustomer(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Created(c, result.ID(), "Customer")
}

func (ctrl *CustomerController) UpdateCustomer(c *gin.Context) {
	id, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "customer_id", c.Param("id")))
		return
	}

	var requestData dto.UpdateCustomerRequest
	if err := ginUtils.ShouldBindAndValidateBody(c, &requestData, ctrl.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	command, err := requestData.ToCommand(id)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.UpdateCustomer(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
}

func (ctrl *CustomerController) DeactivateCustomer(c *gin.Context) {
	id, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "customer_id", c.Param("id")))
		return
	}

	command, err := command.NewDeactivateCustomerCommand(id)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.DeactivateCustomer(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
}
