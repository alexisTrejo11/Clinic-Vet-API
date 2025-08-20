package notificationController

import (
	notificationService "github.com/alexisTrejo11/Clinic-Vet-API/app/notifications/application"
	"github.com/gin-gonic/gin"
)

type NotificationAdminController struct {
	notificationService notificationService.NotificationService
}

func NewNotificationAdminController(notificationService notificationService.NotificationService) *NotificationAdminController {
	controller := &NotificationAdminController{
		notificationService: notificationService,
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
