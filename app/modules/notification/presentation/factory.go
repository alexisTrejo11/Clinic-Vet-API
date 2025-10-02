package api

import (
	"os"

	"clinic-vet-api/app/config"
	iservice "clinic-vet-api/app/modules/core/service"
	service "clinic-vet-api/app/modules/notification/application"
	repositoryimpl "clinic-vet-api/app/modules/notification/infrastructure/repository"
	"clinic-vet-api/app/modules/notification/infrastructure/sending/email"
	"clinic-vet-api/app/modules/notification/infrastructure/sending/sms"
	"clinic-vet-api/app/modules/notification/presentation/controller"
	"clinic-vet-api/app/modules/notification/presentation/routes"

	"github.com/gin-gonic/gin"
	"github.com/twilio/twilio-go"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupNotificationModule(app *gin.RouterGroup, mongoClient *mongo.Client, emailConfig config.EmailConfig, twilioClient *twilio.RestClient) iservice.NotificationService {
	notificationRepo := repositoryimpl.NewMongoNotificationRepository(mongoClient)
	emailSender := email.NewEmailSender(emailConfig)
	smsSender := sms.NewTwilioPhoneSender(twilioClient, os.Getenv("TWILIO_PHONE_NUMBER"))

	notificationService := service.NewNotificationService(emailSender, smsSender, notificationRepo)
	service.NewNotificationQueryService(notificationRepo)

	controller := controller.NewNotificationAdminController(notificationService)
	routes.SetupNotificationRoutes(app, controller)

	return notificationService
}
