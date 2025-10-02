package controller

import (
	"errors"

	"clinic-vet-api/app/modules/account/user/application/query"
	"clinic-vet-api/app/modules/account/user/presentation/dto"
	httpError "clinic-vet-api/app/shared/error/infrastructure/http"
	ginUtils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/page"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
)

// GetUserByID retrieves a user by their ID.
// @Summary Get a user by ID
// @Description Retrieves a single user record by their unique ID.
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.APIResponse{data=object} "User found"
// @Failure 400 {object} response.APIResponse "Invalid URL parameter"
// @Router /v1/users/{id} [get]
func (ctrl *UserAdminController) GetUserByID(c *gin.Context) {
	userID, err := ginUtils.ParseParamToUInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	getUserByIDQuery := query.NewFindUserByIDQuery(userID, false)
	userResult, err := ctrl.bus.FindUserByID(c.Request.Context(), *getUserByIDQuery)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	userResponse := dto.UserResultToResponse(&userResult)
	response.Found(c, userResponse, "User")
}

func (ctrl *UserAdminController) SearchUsers(c *gin.Context) {
}

func (ctrl *UserAdminController) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		response.BadRequest(c, httpError.RequestURLParamError(errors.New("email is required"), "email", ""))
		return
	}

	getUserByEmailQuery := query.NewFindUserByEmailQuery(email, false)
	userResult, err := ctrl.bus.FindUserByEmail(c.Request.Context(), *getUserByEmailQuery)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	userResponse := dto.UserResultToResponse(&userResult)
	response.Found(c, userResponse, "User")
}

func (c *UserAdminController) GetUserByPhone(ctx *gin.Context) {
	phone := ctx.Query("phone")
	if phone == "" {
		response.BadRequest(ctx, httpError.RequestURLParamError(errors.New("phone is required"), "phone", ""))
		return
	}

	getUserByPhoneQuery, err := query.NewFindUserByPhoneQuery(phone, false)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	userResult, err := c.bus.FindUserByPhone(ctx.Request.Context(), *getUserByPhoneQuery)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	userResponse := dto.UserResultToResponse(&userResult)
	response.Found(ctx, userResponse, "User")
}

func (ctrl *UserAdminController) FindByRole(ctx *gin.Context) {
	role := ctx.Param("role")
	if role == "" {
		response.BadRequest(ctx, httpError.RequestURLParamError(errors.New("role is required"), "role", ""))
		return
	}

	var pagination page.PaginationRequest
	if err := ginUtils.ShouldBindPageParams(&pagination, ctx, ctrl.validator); err != nil {
		response.BadRequest(ctx, httpError.RequestURLQueryError(err, ctx.Request.URL.RawQuery))
		return
	}

	findUsersByRoleQuery := query.NewFindUsersByRoleQuery(role, pagination)
	userResults, err := ctrl.bus.FindUsersByRole(ctx.Request.Context(), *findUsersByRoleQuery)
	if err != nil {
		response.ApplicationError(ctx, err)
		return
	}

	userResponses := dto.UserResultsToResponses(userResults.Items)
	response.SuccessWithPagination(ctx, userResponses, "Users retrieved successfully", userResults.Metadata)
}
