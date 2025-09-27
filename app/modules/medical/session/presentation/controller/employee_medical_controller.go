package controller

import (
	"clinic-vet-api/app/middleware"
	autherror "clinic-vet-api/app/shared/error/auth"
	"clinic-vet-api/app/shared/response"

	"github.com/gin-gonic/gin"
)

type EmployeeMedicalSessionController struct {
	operations *MedSessionControllerOperations
}

func NewEmployeeMedicalSessionController(opeations *MedSessionControllerOperations) *EmployeeMedicalSessionController {
	return &EmployeeMedicalSessionController{
		operations: opeations,
	}
}

func (ctrl *EmployeeMedicalSessionController) GetMyMedicalSessions(c *gin.Context) {
	userCTX, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.BadRequest(c, autherror.UnauthorizedCTXError())
		return
	}

	ctrl.operations.GetMedSessionsByEmployeeID(c, userCTX.EmployeeID)
}

func (ctrl *EmployeeMedicalSessionController) RegisterMedicalSession(c *gin.Context) {
	userCTX, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.BadRequest(c, autherror.UnauthorizedCTXError())
		return
	}

	ctrl.operations.CreateMedicalSession(c, &userCTX.EmployeeID)
}

func (ctrl *EmployeeMedicalSessionController) GetMyMedicalSessionByID(c *gin.Context) {
	userCTX, exists := middleware.GetUserFromContext(c)
	if !exists {
		response.BadRequest(c, autherror.UnauthorizedCTXError())
		return
	}

	ctrl.operations.GetMedSessionsByID(c, &GetByIDExtraArgs{
		EmployeeID: &userCTX.EmployeeID,
	})
}
