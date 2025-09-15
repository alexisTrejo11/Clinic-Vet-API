package routes

import (
	"clinic-vet-api/app/modules/notifications/infrastructure/api/controller"
	"github.com/gin-gonic/gin"
)

func SetupNotificationRoutes(app *gin.Engine, adminController *controller.NotificationAdminController) {
	notificationGroup := app.Group("api/v2/admin/notifications")
	{
		notificationGroup.POST("/send", adminController.SendNotification)
		notificationGroup.GET("/:id", adminController.GetNotificationById)
		notificationGroup.GET("/users/:id", adminController.GetNotificationByUserId)
		notificationGroup.GET("/types/:type", adminController.GetNotificationByType)
		notificationGroup.GET("/status/:status", adminController.GetNotificationByStatus)
	}
}
