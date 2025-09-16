package api

import (
	"os"

	service "clinic-vet-api/app/modules/notifications/application"
	repositoryimpl "clinic-vet-api/app/modules/notifications/infrastructure/repository"
	"clinic-vet-api/app/modules/notifications/infrastructure/sending/email"
	"clinic-vet-api/app/modules/notifications/infrastructure/sending/sms"
	"clinic-vet-api/app/modules/notifications/presentation/controller"
	"clinic-vet-api/app/modules/notifications/presentation/routes"
	"clinic-vet-api/config"

	"github.com/gin-gonic/gin"
	"github.com/twilio/twilio-go"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupNotificationModule(app *gin.Engine, mongoClient *mongo.Client, emailConfig config.EmailConfig, twilioClient *twilio.RestClient) {
	notificationRepo := repositoryimpl.NewMongoNotificationRepository(mongoClient)
	emailSender := email.NewEmailSender(emailConfig)
	smsSender := sms.NewTwilioPhoneSender(twilioClient, os.Getenv("TWILIO_PHONE_NUMBER"))

	notificationService := service.NewNotificationService(notificationRepo, map[string]service.Sender{
		"email": emailSender,
		"sms":   smsSender,
	})

	controller := controller.NewNotificationAdminController(notificationService)
	routes.SetupNotificationRoutes(app, controller)
}
