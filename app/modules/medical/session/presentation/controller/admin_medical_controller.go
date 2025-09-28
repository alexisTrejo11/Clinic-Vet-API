// Package controller defines the controllers for handling HTTP requests related to medical histories.
package controller

import (
	"clinic-vet-api/app/shared/response"
	"time"

	httpError "clinic-vet-api/app/shared/error/infrastructure/http"
	ginUtils "clinic-vet-api/app/shared/gin_utils"

	"github.com/gin-gonic/gin"
)

type AdminMedicalSessionController struct {
	operations *MedSessionControllerOperations
}

func NewAdminMedicalSessionController(operations *MedSessionControllerOperations) *AdminMedicalSessionController {
	return &AdminMedicalSessionController{
		operations: operations,
	}
}

func (ctrl AdminMedicalSessionController) SearchMedSessions(c *gin.Context) {
}

func (ctrl AdminMedicalSessionController) GetMedicalSessionByID(c *gin.Context) {
	ctrl.operations.GetMedSessionsByID(c, nil)
}

func (ctrl AdminMedicalSessionController) GetMedSessionsByPetID(c *gin.Context) {
	petID, err := ginUtils.ParseParamToUInt(c, "pet_id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "pet", c.Param("pet_id")))
		return
	}

	ctrl.operations.GetMedSessionsByPetID(c, petID, nil)
}

func (ctrl AdminMedicalSessionController) GetMedicalSessionByEmployeeID(c *gin.Context) {
	employeeID, err := ginUtils.ParseParamToUInt(c, "employee_id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "employee", c.Param("employee_id")))
		return
	}

	ctrl.operations.GetMedSessionsByEmployeeID(c, employeeID)
}

func (ctrl AdminMedicalSessionController) GetMedicalSessionByCustomerID(c *gin.Context) {
	customerID, err := ginUtils.ParseParamToUInt(c, "customer_id")
	if err != nil {
		response.BadRequest(c, httpError.RequestURLParamError(err, "customer", c.Param("customer_id")))
		return
	}

	ctrl.operations.GetMedicalSessionByCustomerID(c, customerID, nil)
}

func (ctrl AdminMedicalSessionController) GetTodayMedSessions(c *gin.Context) {
	now := time.Now()
	startDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endDay := startDay.Add(24*time.Hour - time.Nanosecond)

	ctrl.operations.GetMedSessionsByDateRange(c, startDay, endDay)
}

// Command
func (ctrl AdminMedicalSessionController) CreateMedicalSession(c *gin.Context) {
	ctrl.operations.CreateMedicalSession(c, nil)
}

func (ctrl AdminMedicalSessionController) SoftDeleteMedicalSession(c *gin.Context) {
	ctrl.operations.DeleteMedicalSession(c)
}
