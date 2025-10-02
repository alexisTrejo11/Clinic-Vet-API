package controller

import (
	"clinic-vet-api/app/modules/employee/application/command"
	"clinic-vet-api/app/modules/employee/application/query"
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
	bus       bus.EmployeeCqrsBus
}

func NewEmployeeController(
	validator *validator.Validate,
	bus bus.EmployeeCqrsBus,
) *EmployeeController {
	return &EmployeeController{
		validator: validator,
		bus:       bus,
	}
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

	query, err := query.NewFindEmployeeByIDQuery(id)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result, err := ctrl.bus.FindEmployeeByID(c.Request.Context(), query)
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
	if err := ginUtils.ShouldBindAndValidateBody(c, &requestData, ctrl.validator); err != nil {
		response.BadRequest(c, err)
		return
	}

	command, err := requestData.ToCommand()
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.CreateEmployee(c.Request.Context(), command)
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

	var requestData dto.UpdateEmployeeRequest
	if err := ginUtils.ShouldBindAndValidateBody(c, &requestData, ctrl.validator); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	command, err := requestData.ToCommand(id)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.UpdateEmployee(c.Request.Context(), command)
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

	command, err := command.NewDeleteEmployeeCommand(id)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.DeleteEmployee(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.NoContent(c)
}
