package authApi

import (
	authCmd "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/application/command"
	jwtService "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/application/jwt"
	authController "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/infrastructure/api/controller"
	authRoutes "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/infrastructure/api/routes"
	sessionRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/auth/infrastructure/persistence"
	sqlcUserRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/users/infrastructure/persistence/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
)

func SetupAuthModule(
	r *gin.Engine,
	validator *validator.Validate,
	client *redis.Client,
	queries *sqlc.Queries,
	secretKet string,
) {
	userRepo := sqlcUserRepo.NewSQLCUserRepository(queries)
	jwtService := jwtService.NewJWTService(secretKet)

	session := sessionRepository.NewRedisSessionRepository(client)
	authCMDBus := authCmd.NewAuthCommandBus(session, userRepo, jwtService)
	authController := authController.NewAuthController(validator, authCMDBus)
	authRoutes.AuthRoutes(r, *authController)
}
