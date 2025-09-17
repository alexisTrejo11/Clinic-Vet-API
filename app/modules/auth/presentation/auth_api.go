package api

import (
	event "clinic-vet-api/app/core/domain/event/user_event"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/core/service"
	"clinic-vet-api/app/modules/auth/application/command"
	"clinic-vet-api/app/modules/auth/infrastructure/bus"
	"clinic-vet-api/app/modules/auth/infrastructure/jwt"
	repositoryimpl "clinic-vet-api/app/modules/auth/infrastructure/repository"
	"clinic-vet-api/app/modules/auth/presentation/controller"
	"clinic-vet-api/app/modules/auth/presentation/routes"
	userPersistence "clinic-vet-api/app/modules/users/infrastructure/repository"
	"clinic-vet-api/app/shared/password"
	"clinic-vet-api/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
)

// TODO: EVENT
func SetupAuthModule(
	r *gin.Engine,
	validator *validator.Validate,
	employeeRepo repository.EmployeeRepository,
	customerRepo repository.CustomerRepository,
	client *redis.Client,
	queries *sqlc.Queries,
	secretKet string,
) {
	userRepo := userPersistence.NewSQLCUserRepository(queries)
	jwtService := jwt.NewJWTService(secretKet)
	passwordEncoder := password.NewPasswordEncoder()
	emailService := service.NewEmailService()

	session := repositoryimpl.NewRedisSessionRepository(client)
	security_service := service.NewUserSecurityService(userRepo, employeeRepo, passwordEncoder)
	profileRepo := userPersistence.NewSQLCProfileRepository(queries)

	eventService := service.NewUserAccountService(userRepo, profileRepo, customerRepo, employeeRepo, emailService)
	userEventProducer := event.NewUserEventHandler(eventService)

	commandHandler := command.NewAuthCommandHandler(userRepo, *security_service, session, jwtService, passwordEncoder, userEventProducer)
	authCMDBus := bus.NewAuthBus(commandHandler)
	authController := controller.NewAuthController(validator, authCMDBus)
	routes.AuthRoutes(r, *authController)
}
