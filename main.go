package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"clinic-vet-api/app/config"
	"clinic-vet-api/app/middleware"
	notiAPI "clinic-vet-api/app/modules/notifications/presentation"
	"clinic-vet-api/app/shared/log"
	"clinic-vet-api/sqlc"

	_ "clinic-vet-api/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Application holds all application dependencies
type Application struct {
	Settings  *config.AppSettings
	Router    *gin.Engine
	Server    *http.Server
	Validator *validator.Validate
	Queries   *sqlc.Queries
}

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
	// Initialize application
	app, err := initializeApplication()
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize application: %v", err))
	}

	// Setup graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start server in a goroutine
	go func() {
		if err := app.startServer(); err != nil && err != http.ErrServerClosed {
			log.App.Fatal(fmt.Sprintf("Failed to start server: %v", err))
		}
	}()

	log.App.Info(fmt.Sprintf("Server started on %s", app.Settings.GetServerAddr()))
	log.App.Info(fmt.Sprintf("Environment: %s", app.Settings.Server.Environment))
	log.App.Info(fmt.Sprintf("Debug mode: %t", app.Settings.App.Debug))

	// Wait for shutdown signal
	app.waitForShutdown(ctx)
}

// initializeApplication initializes all application components
func initializeApplication() (*Application, error) {
	// Load configuration
	settings, err := config.LoadSettings()
	if err != nil {
		return nil, fmt.Errorf("failed to load settings: %w", err)
	}

	// Validate configuration
	if err := settings.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	// Initialize logger
	config.InitLogger()

	// Set gin mode based on environment
	if settings.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	} else if settings.App.Debug {
		gin.SetMode(gin.DebugMode)
	}

	// Initialize external services
	if err := initializeExternalServices(settings); err != nil {
		return nil, fmt.Errorf("failed to initialize external services: %w", err)
	}

	// Initialize database
	pxpool := config.CreatePgxPool(settings.Database.URL)
	queries := sqlc.New(pxpool)

	// Initialize validator
	dataValidator := validator.New()

	// Setup router
	router := setupRouter(settings)

	// Setup modules
	if err := setupModules(router, settings, queries, pxpool, dataValidator); err != nil {
		return nil, fmt.Errorf("failed to setup modules: %w", err)
	}

	// Create HTTP server
	server := createHTTPServer(router, settings)

	app := &Application{
		Settings:  settings,
		Router:    router,
		Server:    server,
		Validator: dataValidator,
		Queries:   queries,
	}

	return app, nil
}

// initializeExternalServices initializes all external service connections
func initializeExternalServices(settings *config.AppSettings) error {
	// Initialize Twilio
	config.InitTwilio(settings.Services.Twilio)

	// Initialize Redis
	config.InitRedis(settings.Redis)

	// MongoDB is initialized in setupModules since it's passed to specific modules
	config.InitMongoDB(settings.Services.Mongo)

	return nil
}

// setupRouter configures the Gin router with middleware
func setupRouter(settings *config.AppSettings) *gin.Engine {
	router := gin.New()

	// Recovery middleware
	router.Use(gin.Recovery())

	// CORS middleware
	router.Use(setupCORS(settings))

	// Audit logging middleware
	router.Use(middleware.AuditLog())

	// Rate limiting middleware (if enabled)
	if settings.RateLimit.Enabled {
		rateLimitConfig := middleware.RateLimiterConfig{
			MaxRequests:    settings.RateLimit.RequestsPerSecond,
			WindowDuration: settings.RateLimit.WindowDuration,
			//BurstSize:         settings.RateLimit.BurstSize,
		}
		router.Use(middleware.RateLimiter(rateLimitConfig))
	}

	// Health check endpoint
	router.GET("/health", healthCheck)

	// Swagger documentation (only in development)
	if settings.App.EnableSwagger && !settings.IsProduction() {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		log.App.Info("Swagger UI available at http://" + settings.GetServerAddr() + "/swagger/index.html")
	}

	return router
}

// setupCORS configures CORS middleware based on settings
func setupCORS(settings *config.AppSettings) gin.HandlerFunc {
	corsConfig := cors.Config{
		AllowOrigins:     settings.CORS.AllowOrigins,
		AllowMethods:     settings.CORS.AllowMethods,
		AllowHeaders:     settings.CORS.AllowHeaders,
		ExposeHeaders:    []string{"Content-Length", "X-Total-Count", "X-Rate-Limit-Remaining"},
		AllowCredentials: settings.CORS.AllowCredentials,
		MaxAge:           settings.CORS.MaxAge,
	}

	// In development, be more permissive
	if settings.IsDevelopment() {
		corsConfig.AllowOrigins = []string{"*"}
		corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "*")
	}

	return cors.New(corsConfig)
}

// setupModules initializes and registers all application modules
func setupModules(router *gin.Engine, settings *config.AppSettings, queries *sqlc.Queries, pxpool *pgxpool.Pool, validator *validator.Validate) error {
	// Initialize MongoDB for notification module
	mongoClient := config.InitMongoDB(settings.Services.Mongo)

	// Setup notification module
	notiAPI.SetupNotificationModule(router, mongoClient, settings.Services.Email, config.GetTwilioClient())

	// Bootstrap other API modules
	if err := config.BootstrapAPIModules(router, queries, pxpool, validator, config.RedisClient, settings.Auth.JWTSecret); err != nil {
		return fmt.Errorf("failed to bootstrap API modules: %w", err)
	}

	return nil
}

// createHTTPServer creates and configures the HTTP server
func createHTTPServer(handler http.Handler, settings *config.AppSettings) *http.Server {
	return &http.Server{
		Addr:           settings.GetServerAddr(),
		Handler:        handler,
		ReadTimeout:    settings.Server.ReadTimeout,
		WriteTimeout:   settings.Server.WriteTimeout,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}
}

// startServer starts the HTTP server
func (app *Application) startServer() error {
	log.App.Info(fmt.Sprintf("Starting server on %s", app.Settings.GetServerAddr()))
	return app.Server.ListenAndServe()
}

// waitForShutdown waits for shutdown signals and gracefully shuts down the server
func (app *Application) waitForShutdown(ctx context.Context) {
	// Create a channel to receive OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Wait for signal
	<-quit
	log.App.Info("Shutting down server...")

	// Create a timeout context for shutdown
	shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Shutdown server
	if err := app.Server.Shutdown(shutdownCtx); err != nil {
		log.App.Error(fmt.Sprintf("Server forced to shutdown: %v", err))
	}

	// Cleanup resources
	app.cleanup()

	log.App.Info("Server exited")
}

// cleanup handles cleanup of resources
func (app *Application) cleanup() {
	// Sync logger before exit
	config.SyncLogger()

	// Close database connections
	// Note: Add any other cleanup tasks here
	log.App.Info("Cleanup completed")
}

// Health check endpoint
// @Summary Health Check
// @Description Check if the API is running
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"timestamp": time.Now().Unix(),
		"service":   "clinic-vet-api",
		"version":   "2.0.0",
	})
}
