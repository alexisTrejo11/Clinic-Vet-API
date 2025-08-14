package notificationController

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/notifications/application"
	"github.com/gin-gonic/gin"
)

type NotificationAdminController struct {
	notificationFacade application.NotificationFacade
}

func NewNotificationAdminController(notificationFacade application.NotificationFacade) *NotificationAdminController {
	controller := &NotificationAdminController{
		notificationFacade: notificationFacade,
	}
	return controller
}

func (c *NotificationAdminController) SendNotification(ctx *gin.Context) {

}

func (c *NotificationAdminController) GetNotificationByUserId(ctx *gin.Context) {

}

func (c *NotificationAdminController) GetNotificationById(ctx *gin.Context) {

}

func (c *NotificationAdminController) GetNotificationByType(ctx *gin.Context) {
}

func (c *NotificationAdminController) GetNotificationByStatus(ctx *gin.Context) {

}
