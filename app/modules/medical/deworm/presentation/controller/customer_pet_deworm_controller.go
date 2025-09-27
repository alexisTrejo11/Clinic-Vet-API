package controller

import (
	"clinic-vet-api/app/modules/medical/deworm/application"

	"github.com/gin-gonic/gin"
)

type CustomerPetDewormController struct {
	dewormService application.DewormingFacadeService
}

func NewCustomerPetDewormController(dewormService application.DewormingFacadeService) *CustomerPetDewormController {
	return &CustomerPetDewormController{
		dewormService: dewormService,
	}
}

func (ctrl *CustomerPetDewormController) GetMyPetDewormHistory(c *gin.Context) {
}

func (ctrl *CustomerPetDewormController) GetMyPetDewormHistoryByPetID(c *gin.Context) {
}

func (ctrl *CustomerPetDewormController) GetMyPetDewormDetailByID(c *gin.Context) {
}
