package controller

import (
	"clinic-vet-api/app/modules/medical/deworm/application"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type EmployeeDewormController struct {
	dewormService application.DewormingFacadeService
	validator     *validator.Validate
}

func NewEmployeeDewormController(dewormService application.DewormingFacadeService, validator *validator.Validate) *EmployeeDewormController {
	return &EmployeeDewormController{
		dewormService: dewormService,
		validator:     validator,
	}
}

func (ctrl *EmployeeDewormController) RegisterNewDewormApplication(c *gin.Context) {
}

func (ctrl *EmployeeDewormController) GetMyDewormsApplied(c *gin.Context) {
}

func (ctrl *EmployeeDewormController) GetMyDewormAppliedByID(c *gin.Context) {
}

func (ctrl *EmployeeDewormController) UpdateMyDewormApplied(c *gin.Context) {
}
