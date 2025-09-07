package api

import (
	"os"

	service "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/notifications/application"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/notifications/infrastructure/api/controller"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/notifications/infrastructure/api/routes"
	repositoryimpl "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/notifications/infrastructure/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/notifications/infrastructure/sending/email"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/notifications/infrastructure/sending/sms"
	"github.com/alexisTrejo11/Clinic-Vet-API/config"
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
