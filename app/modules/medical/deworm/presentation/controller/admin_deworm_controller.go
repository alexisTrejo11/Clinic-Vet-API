package controller

import "github.com/gin-gonic/gin"

type AdminDewormController struct{}

func NewAdminDewormController() *AdminDewormController {
	return &AdminDewormController{}
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
