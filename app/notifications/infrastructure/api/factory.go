package notification_api

import (
	"os"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/notifications/application"
	notificationController "github.com/alexisTrejo11/Clinic-Vet-API/app/notifications/infrastructure/api/controller"
	notificationRoutes "github.com/alexisTrejo11/Clinic-Vet-API/app/notifications/infrastructure/api/routes"
	notificationRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/notifications/infrastructure/persistence"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/notifications/infrastructure/sending/email"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/notifications/infrastructure/sending/sms"
	"github.com/alexisTrejo11/Clinic-Vet-API/config"
	"github.com/gin-gonic/gin"
	"github.com/twilio/twilio-go"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupNotificationModule(app *gin.Engine, mongoClient *mongo.Client, emailConfig config.EmailConfig, twilioClient *twilio.RestClient) {
	notificationRepo := notificationRepo.NewMongoNotificationRepository(mongoClient)
	emailSender := email.NewEmailSender(emailConfig)
	smsSender := sms.NewTwilioPhoneSender(twilioClient, os.Getenv("TWILIO_PHONE_NUMBER"))

	notificationService := application.NewNotificationService(notificationRepo, map[string]application.Sender{
		"email": emailSender,
		"sms":   smsSender,
	})

	controller := notificationController.NewNotificationAdminController(notificationService)
	notificationRoutes.SetupNotificationRoutes(app, controller)
}
