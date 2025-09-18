package main

import (
	"fmt"
	"net/http"
	"os"

	notiAPI "clinic-vet-api/app/modules/notifications/presentation"
	"clinic-vet-api/app/shared/log"
	"clinic-vet-api/config"
	"clinic-vet-api/middleware"
	"clinic-vet-api/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"

	_ "clinic-vet-api/docs"

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
		log.App.DPanic(fmt.Sprintf("Error while loading .env: %v", err))
	}

	config.InitTwilio()
	config.InitRedis(os.Getenv("REDIS_URL"))
	mongoClient := config.InitMongoDB()
	emailConfig := config.InitEmailConfig()
	dataValidator := validator.New()

	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.AuditLog())
	router.Use(middleware.RateLimiter(middleware.DefaultConfig()))

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.App.Fatal("JWT_SECRET environment variable is not set")
	}

	pxpool := config.CreatePgxPool(os.Getenv("DATABASE_URL"))
	queries := sqlc.New(pxpool)

	notiAPI.SetupNotificationModule(router, mongoClient, emailConfig, config.GetTwilioClient())
	if err := config.BootstrapAPIModules(router, queries, pxpool, dataValidator, config.RedisClient, jwtSecret); err != nil {
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
	log.App.Info("pong")
	g.JSON(http.StatusOK, "pong")
}
