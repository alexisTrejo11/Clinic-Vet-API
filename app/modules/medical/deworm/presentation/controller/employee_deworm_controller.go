package controller

import "github.com/gin-gonic/gin"

type EmployeeDewormController struct{}

func NewEmployeeDewormController() *EmployeeDewormController {
	return &EmployeeDewormController{}
}

func (ctrl *EmployeeDewormController) RegisterNewDewormApplication(c *gin.Context) {
}

func (ctrl *EmployeeDewormController) GetMyDewormsApplied(c *gin.Context) {
}

func (ctrl *EmployeeDewormController) GetMyDewormAppliedByID(c *gin.Context) {
}

func (ctrl *EmployeeDewormController) UpdateMyDewormApplied(c *gin.Context) {
}
