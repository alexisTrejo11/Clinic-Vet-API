package main

import (
	"context"
	"log"
	"net/http"
	"os"

	appointmentAPI "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/infrastructure/api"
	authApi "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/infrastructure/api"
	mhDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/application/dtos"
	medHistoryAPI "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/infrastructure/api"
	notificationApi "github.com/alexisTrejo11/Clinic-Vet-API/app/notifications/infrastructure/api"
	ownerAPI "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/infrastructure/api"
	paymentAPI "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/infrastructure/api"
	petAPI "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/infrastructure/api"
	userAPI "github.com/alexisTrejo11/Clinic-Vet-API/app/users/infrastructure/api"

	vetAPI "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/infrastructure/api"

	"github.com/alexisTrejo11/Clinic-Vet-API/config"
	"github.com/alexisTrejo11/Clinic-Vet-API/middleware"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"

	_ "github.com/alexisTrejo11/Clinic-Vet-API/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API Clínica Veterinaria
// @version 1.0
// @description Esta es la documentación para la API de la clínica veterinaria.
// @termsOfService http://swagger.io/terms/
// @contact.name Equipo de Soporte API
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v2
func main() {
	config.InitLogger()
	defer config.SyncLogger()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while loading .env: %v", err)
	}

	config.InitTwilio()
	mongoClient := config.InitMongoDB()
	emailConfig := config.InitEmailConfig()
	config.InitRedis(os.Getenv("REDIS_URL"))
	dbConn := config.PostgresConn(os.Getenv("DATABASE_URL"))
	dataValidator := validator.New()

	ctx := context.Background()
	defer dbConn.Close(ctx)

	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.AuditLog())

	queries := sqlc.New(dbConn)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("validVisitReason", mhDTOs.IsValidVisitReason)
		v.RegisterValidation("validVisitType", mhDTOs.IsValidVisitType)
		v.RegisterValidation("validPetCondition", mhDTOs.IsValidPetCondition)
	}

	vetAPI := vetAPI.NewVeterinarianAPI(queries, router, dataValidator)
	petAPI := petAPI.NewPetAPI(router, queries, dataValidator)
	ownerAPI := ownerAPI.NewOwnerAPI(router, dataValidator, queries, petAPI.GetRepository())
	medHistoryAPI.NewMedicalHistoryAPI(queries, router, dataValidator, ownerAPI.GetRepository(), vetAPI.GetRepository(), petAPI.GetRepository())
	paymentAPI.SetupPaymentAPI(router, dataValidator, queries)
	userAPI.NewUserAPI(queries, dataValidator, router)
	appointmentAPI.SetupAppointmentAPI(router, queries, dataValidator, ownerAPI.GetRepository())
	notificationApi.SetupNotificationModule(router, mongoClient, emailConfig, config.GetTwilioClient())
	authApi.SetupAuthModule(router, dataValidator, config.RedisClient)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/ping", Ping)

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		router.Run(":8080")
	} else {
		router.Run(":" + serverPort)
	}
}

// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Ping
// @Router /ping [get]
func Ping(g *gin.Context) {
	g.JSON(http.StatusOK, "Ping")
}
