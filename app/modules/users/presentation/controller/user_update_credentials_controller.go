package controller

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/users/infrastructure/bus"
	"clinic-vet-api/app/modules/users/presentation/dto"
	autherror "clinic-vet-api/app/shared/error/auth"
	httpError "clinic-vet-api/app/shared/error/infrastructure/http"
	ginUtils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserUpdateCredentialsController struct {
	validator *validator.Validate
	bus       *bus.UserBus
}

func NewUserUpdateCredentialsController(validator *validator.Validate, bus *bus.UserBus) *UserUpdateCredentialsController {
	return &UserUpdateCredentialsController{
		validator: validator,
		bus:       bus,
	}
}

func (ctrl *UserUpdateCredentialsController) SendResetPasswordEmail(c *gin.Context) {
	var req dto.RequestEmailRequest
	if err := ginUtils.ShouldBindAndValidateQuery(c, &req, ctrl.validator); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	command, err := req.ToCommand()
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.RequestResetPassword(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
}

func (ctrl *UserUpdateCredentialsController) ResetPassword(c *gin.Context) {
	var req dto.ResetPasswordRequest
	if err := ginUtils.ShouldBindAndValidateQuery(c, &req, ctrl.validator); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	command, err := req.ToCommand()
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.ResetPassword(c.Request.Context(), command)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
}

func (ctrl *UserUpdateCredentialsController) UpdatePassword(c *gin.Context) {
	id, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		response.BadRequest(c, autherror.UnauthorizedCTXError())
		return
	}

	var req dto.UpdatePasswordRequest
	if err := ginUtils.ShouldBindAndValidateQuery(c, &req, ctrl.validator); err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	cmd, err := req.ToCommand(id.Value())
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	result := ctrl.bus.UpdatePassword(c.Request.Context(), cmd)
	if !result.IsSuccess() {
		response.ApplicationError(c, result.Error())
		return
	}

	response.Success(c, nil, result.Message())
}

func (ctrl *UserUpdateCredentialsController) UpdateEmail(c *gin.Context) {

}

func (ctrl *UserUpdateCredentialsController) UpdatePhoneNumber(c *gin.Context) {

}
