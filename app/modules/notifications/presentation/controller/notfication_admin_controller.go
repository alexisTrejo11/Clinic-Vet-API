package controller

import (
	service "clinic-vet-api/app/modules/notifications/application"
	"github.com/gin-gonic/gin"
)

type NotificationAdminController struct {
	notificationService service.NotificationService
}

func NewNotificationAdminController(notificationService service.NotificationService) *NotificationAdminController {
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
