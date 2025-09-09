package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	config "command-line-arguments/Users/alexis/Documents/Clinic-Vet-API/config/postgres_config.go"

	authApi "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/auth/infrastructure/api"
	mhDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/medical/application/dto"
	notification_api "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/notifications/infrastructure/api"

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
	config.InitRedis(os.Getenv("REDIS_URL"))
	mongoClient := config.InitMongoDB()
	emailConfig := config.InitEmailConfig()
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

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is not set")
	}

	authApi.SetupAuthModule(router, dataValidator, config.RedisClient, queries, jwtSecret)
	notification_api.SetupNotificationModule(router, mongoClient, emailConfig, config.GetTwilioClient())

	pxpool := config.CreatePgxPool(os.Getenv("DATABASE_URL"))

	if err := config.BootstrapAPIModules(router, queries, pxpool, dataValidator); err != nil {
		panic(fmt.Sprintf("Failed to bootstrap modules: %v", err))
	}

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
