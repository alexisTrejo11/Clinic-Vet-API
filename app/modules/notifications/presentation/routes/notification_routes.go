package routes

import (
	"clinic-vet-api/app/modules/notifications/presentation/controller"

	"github.com/gin-gonic/gin"
)

func SetupNotificationRoutes(app *gin.RouterGroup, adminController *controller.NotificationAdminController) {
	notificationGroup := app.Group("/notifications")
	{
		notificationGroup.POST("/send", adminController.SendNotification)
		notificationGroup.GET("/:id", adminController.GetNotificationById)
		notificationGroup.GET("/users/:id", adminController.GetNotificationByUserId)
		notificationGroup.GET("/types/:type", adminController.GetNotificationBySpecies)
		notificationGroup.GET("/status/:status", adminController.GetNotificationByStatus)
	}
}
