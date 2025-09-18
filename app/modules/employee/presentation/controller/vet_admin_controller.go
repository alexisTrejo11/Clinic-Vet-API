package controller

import (
	"clinic-vet-api/app/modules/employee/application/cqrs/command"
	"clinic-vet-api/app/modules/employee/application/cqrs/query"
	"clinic-vet-api/app/modules/employee/infrastructure/bus"
	"clinic-vet-api/app/modules/employee/presentation/dto"
	httpError "clinic-vet-api/app/shared/error/infrastructure/http"
	ginUtils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type EmployeeController struct {
	validator *validator.Validate
	bus       bus.EmployeeBus
}

func NewEmployeeController(
	validator *validator.Validate,
	bus bus.EmployeeBus,
) *EmployeeController {
	return &EmployeeController{
		validator: validator,
		bus:       bus,
	}
}

// This package contains wrapper functions that hold the Swagger annotations.
// They then delegate the actual work to the controller to keep it clean.

// @Summary List all employee
// @Description Retrieves a list of all employee with optional filtering and pagination.
// @Tags Employees
// @Produce json
// @Param page query int false "Page number for pagination"
// @Param limit query int false "Number of items per page"
// @Param specialty query string false "Filter by specialty"
// @Success 200 {array} dto.VetResponse "List of employee"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /employee [get]
func (ctrl *EmployeeController) SearchEmployees(c *gin.Context) {
	var requestParams dto.EmployeeSearchRequest

	if err := ginUtils.BindAndValidateQuery(c, &requestParams, ctrl.validator); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	requestParams.ProcessBooleanParams(c)

	query, err := requestParams.ToQuery()
	if err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	employeesPage, err := ctrl.bus.QueryBus.SearchEmployees(c.Request.Context(), *query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	responseMetadata := gin.H{"pagination": employeesPage.Metadata, "requestParams": c.Request.URL.Query()}
	response.SuccessWithMeta(c, employeesPage, "Employees retrieved successfully", responseMetadata)
}

// @Summary Get a employee by ID
// @Description Retrieves a single employee by their unique ID.
// @Tags Employees
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} dto.VetResponse "Employee details"
// @Failure 400 {object} response.APIResponse "Invalid ID supplied"
// @Failure 404 {object} response.APIResponse "Employee not found"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /employee/{id} [get]
func (ctrl *EmployeeController) GetEmployeeByID(c *gin.Context) {
	id, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	query := query.NewGetEmployeeByIDQuery(id)
	result, err := ctrl.bus.QueryBus.GetEmployeeByID(c.Request.Context(), query)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	employeeResponse := dto.ToEmployeeResponse(result)
	response.Found(c, employeeResponse, "Employee")
}

// @Summary Create a new employee
// @Description Creates a new employee with the provided details.
// @Tags Employees
// @Accept json
// @Produce json
// @Param vetCreate body dto.VetCreate true "Employee data"
// @Success 201 {object} dto.VetResponse "Employee created"
// @Failure 400 {object} response.APIResponse "Invalid request body"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /employee [post]
func (ctrl *EmployeeController) CreateEmployee(c *gin.Context) {
	var requestData dto.CreateEmployeeRequest
	if err := ginUtils.BindAndValidateBody(c, &requestData, ctrl.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	command, err := requestData.ToCommand()
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.CommandBus.CreateEmployee(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Created(c, result.ID(), "Employee")
}

// @Summary Update an existing employee
// @Description Updates the details of an existing employee by ID.
// @Tags Employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Param requestData body dto.VetUpdate true "Updated employee data"
// @Success 200 {object} dto.VetResponse "Employee updated"
// @Failure 400 {object} response.APIResponse "Invalid request data or ID"
// @Failure 404 {object} response.APIResponse "Employee not found"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /employee/{id} [patch]
func (ctrl *EmployeeController) UpdateEmployee(c *gin.Context) {
	id, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "Employee_id", c.Param("id")))
		return
	}

	var requestData *dto.UpdateEmployeeRequest
	if err := ginUtils.BindAndValidateBody(c, &requestData, ctrl.validator); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	command := requestData.ToCommand(id)
	result := ctrl.bus.CommandBus.UpdateEmployee(c.Request.Context(), *command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
}

// @Summary Delete a employee
// @Description Deletes a employee from the system by ID.
// @Tags Employees
// @Produce json
// @Param id path int true "Employee ID"
// @Success 204 "No Content"
// @Failure 400 {object} response.APIResponse "Invalid ID supplied"
// @Failure 404 {object} response.APIResponse "Employee not found"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /employee/{id} [delete]
func (ctrl *EmployeeController) DeleteEmployee(c *gin.Context) {
	id, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "Employee_id", c.Param("id")))
		return
	}

	command := command.NewDeleteEmployeeCommand(id)
	result := ctrl.bus.CommandBus.DeleteEmployee(c.Request.Context(), *command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.NoContent(c)
}
