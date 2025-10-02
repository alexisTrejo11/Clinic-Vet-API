package controller

import (
	"clinic-vet-api/app/modules/medical/deworm/application"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AdminDewormController struct {
	service   application.DewormingFacadeService
	validator *validator.Validate
}

func NewAdminDewormController(service application.DewormingFacadeService, validator *validator.Validate) *AdminDewormController {
	return &AdminDewormController{
		service:   service,
		validator: validator,
	}
}

func (ctrl *AdminDewormController) GetDewormByID(c *gin.Context) {
}

func (ctrl *AdminDewormController) SearhDeworms(c *gin.Context) {
}

func (ctrl *AdminDewormController) CreateDeworm(c *gin.Context) {
}

func (ctrl *AdminDewormController) UpdateDeworm(c *gin.Context) {
}

func (ctrl *AdminDewormController) DeleteDeworm(c *gin.Context) {
}
