package controller

import (
	"clinic-vet-api/app/modules/medical/deworm/application"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CustomerPetDewormController struct {
	dewormService application.DewormingFacadeService
	validator     *validator.Validate
}

func NewCustomerPetDewormController(dewormService application.DewormingFacadeService, validator *validator.Validate) *CustomerPetDewormController {
	return &CustomerPetDewormController{
		dewormService: dewormService,
		validator:     validator,
	}
}

func (ctrl *CustomerPetDewormController) GetMyPetDewormHistory(c *gin.Context) {
}

func (ctrl *CustomerPetDewormController) GetMyPetDewormHistoryByPetID(c *gin.Context) {
}

func (ctrl *CustomerPetDewormController) GetMyPetDewormDetailByID(c *gin.Context) {
}
