package authApi

import (
	authCmd "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/application/command"
	authController "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/infrastructure/api/controller"
	authRoutes "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/infrastructure/api/routes"
	sessionRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/infrastructure/persistence"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
)

func SetupAuthModule(r *gin.Engine, validator *validator.Validate, client *redis.Client) {
	session := sessionRepository.NewRedisSessionRepository(client)
	authCMDBus := authCmd.NewAuthCommandBus(session)
	authController := authController.NewAuthController(validator, authCMDBus)
	authRoutes.AuthRoutes(r, *authController)
}
