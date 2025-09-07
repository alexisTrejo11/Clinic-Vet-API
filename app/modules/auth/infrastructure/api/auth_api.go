package api

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/auth/infrastructure/api/controller"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/auth/infrastructure/api/routes"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/auth/infrastructure/bus"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/auth/infrastructure/jwt"
	repositoryimpl "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/auth/infrastructure/repository"
	userPersistence "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/users/infrastructure/persistence/repository"
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
	userRepo := userPersistence.NewSQLCUserRepository(queries)
	jwtService := jwt.NewJWTService(secretKet)

	session := repositoryimpl.NewRedisSessionRepository(client)
	authCMDBus := bus.NewAuthCommandBus(session, userRepo, jwtService)
	authController := controller.NewAuthController(validator, authCMDBus)
	routes.AuthRoutes(r, *authController)
}
