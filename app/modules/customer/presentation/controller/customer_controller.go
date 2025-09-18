// Package controller defines the HTTP controllers for the customer module.
package controller

import (
	"clinic-vet-api/app/core/domain/valueobject"
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
	bus       *bus.CustomerBus
}

func NewCustomerController(validator *validator.Validate, bus *bus.CustomerBus) *CustomerController {
	return &CustomerController{
		validator: validator,
		bus:       bus,
	}
}

func (controller *CustomerController) SearchCustomers(c *gin.Context) {
	var searchParams *dto.CustomerSearchQuery
	if err := c.ShouldBindQuery(&searchParams); err != nil {
		response.BadRequest(c, httpError.RequestURLQueryError(err, c.Request.URL.RawQuery))
		return
	}

	if err := controller.validator.Struct(searchParams); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	query := searchParams.ToQuery(c.Request.Context())
	customer, err := controller.bus.QueryBus.FindCustomerByCriteria(*query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}
	response.Found(c, customer, "Customer")
}

func (controller *CustomerController) GetCustomerByID(c *gin.Context) {
	id, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "customer_id", c.Param("id")))
		return
	}

	query := query.NewFindCustomerByIDQuery(c.Request.Context(), id)

	result, err := controller.bus.QueryBus.GetCustomerByID(*query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	customerResponse := dto.FromResult(result)
	response.Found(c, customerResponse, "Customer")
}

func (controller *CustomerController) CreateCustomer(c *gin.Context) {
	var customerCreate dto.CreateCustomerRequest
	if err := ginUtils.BindAndValidateBody(c, &customerCreate, controller.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	createCommand, err := customerCreate.ToCommand(c.Request.Context())
	if err != nil {
		response.ApplicationError(c, err)
		return
	}
	result := controller.bus.CommandBus.CreateCustomer(*createCommand)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}
	response.Created(c, result.ID(), "Customer")
}

func (controller *CustomerController) UpdateCustomer(c *gin.Context) {
	id, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "customer_id", c.Param("id")))
		return
	}

	var requestData dto.UpdateCustomerRequest
	if err := c.ShouldBindBodyWithJSON(&requestData); err != nil {
		response.BadRequest(c, httpError.RequestBodyDataError(err))
		return
	}

	if err := controller.validator.Struct(&requestData); err != nil {
		response.ApplicationError(c, err)
		return
	}

	updateCommand, err := requestData.ToCommand(c.Request.Context(), id)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := controller.bus.CommandBus.UpdateCustomer(*updateCommand)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
}

func (controller *CustomerController) DeactivateCustomer(c *gin.Context) {
	id, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "customer_id", c.Param("id")))
		return
	}

	deactivateCommand := command.DeactivateCustomerCommand{
		CTX: c.Request.Context(),
		ID:  valueobject.NewCustomerID(id),
	}

	result := controller.bus.CommandBus.DeactivateCustomer(deactivateCommand)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
}
