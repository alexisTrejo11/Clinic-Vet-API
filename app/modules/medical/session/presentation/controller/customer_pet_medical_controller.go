package controller

import (
	"clinic-vet-api/app/middleware"
	autherror "clinic-vet-api/app/shared/error/auth"
	"clinic-vet-api/app/shared/error/infrastructure/http"
	ginUtils "clinic-vet-api/app/shared/gin_utils"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
)

type CustomerMedicalSessionController struct {
	operation *MedSessionControllerOperations
}

func NewCustomerMedicalSessionController(operation *MedSessionControllerOperations) *CustomerMedicalSessionController {
	return &CustomerMedicalSessionController{
		operation: operation,
	}
}

func (ctrl *CustomerMedicalSessionController) GetMyPetSessions(c *gin.Context) {
	userCTX, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	ctrl.operation.GetMedicalSessionByCustomerID(c, userCTX.CustomerID, nil)
}

func (ctrl *CustomerMedicalSessionController) GetSessionByID(c *gin.Context) {
	userCTX, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	ctrl.operation.GetMedSessionsByID(c, &GetByIDExtraArgs{
		CustomerID: &userCTX.CustomerID,
	})
}

func (ctrl *CustomerMedicalSessionController) GetMyPetSessionsByPetID(c *gin.Context) {
	userCTX, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	petID, err := ginUtils.ParseParamToUInt(c, "pet_id")
	if err != nil {
		response.BadRequest(c, http.RequestURLParamError(err, "pet", c.Param("pet_id")))
		return
	}

	ctrl.operation.GetMedicalSessionByCustomerID(c, userCTX.CustomerID, &petID)
}

func (ctrl *CustomerMedicalSessionController) GetMyPetDewormingHistory(c *gin.Context) {
	userCTX, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	petID, err := ginUtils.ParseParamToUInt(c, "pet_id")
	if err != nil {
		response.BadRequest(c, http.RequestURLParamError(err, "pet", c.Param("pet_id")))
		return
	}

	ctrl.operation.GetDewormingHistoryByCustomerID(c, userCTX.CustomerID, petID)
}

func (ctrl *CustomerMedicalSessionController) GetMyPetVaccinationHistory(c *gin.Context) {
	userCTX, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.Unauthorized(c, autherror.UnauthorizedCTXError())
		return
	}

	petID, err := ginUtils.ParseParamToUInt(c, "pet_id")
	if err != nil {
		response.BadRequest(c, http.RequestURLParamError(err, "pet", c.Param("pet_id")))
		return
	}

	ctrl.operation.GetVaccinationHistoryByCustomerID(c, userCTX.CustomerID, petID)
}
