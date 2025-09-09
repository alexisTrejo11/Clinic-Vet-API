package controller

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/users/application/usecase/query"
	httpError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure/http"
	ginUtils "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/gin_utils"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/response"
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
func (controller *UserAdminController) GetUserByID(c *gin.Context) {
	userID, err := ginUtils.ParseParamToInt(c, "id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "id", c.Param("id")))
		return
	}

	getUserByIDQuery, err := query.NewGetUserByIDQuery(c.Request.Context(), userID, false)
	if err != nil {
		response.BadRequest(c, httpError.InvalidDataError(err))
		return
	}

	userPage, err := controller.queryBus.Execute(getUserByIDQuery)
	if err != nil {
		response.ApplicationError(c, err)
		return
	}

	response.Success(c, userPage)
}

func (controller *UserAdminController) SearchUsers(c *gin.Context) {
}

func (c *UserAdminController) GetUserByEmail(ctx *gin.Context) {
}

func (c *UserAdminController) GetUserByPhone(ctx *gin.Context) {
}
