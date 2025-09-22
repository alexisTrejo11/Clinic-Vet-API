package api

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/auth/application/command"
	"clinic-vet-api/app/modules/auth/application/command/authentication"
	"clinic-vet-api/app/modules/auth/application/command/register"
	sesionCmd "clinic-vet-api/app/modules/auth/application/command/session"
	"clinic-vet-api/app/modules/auth/infrastructure/jwt"
	repositoryimpl "clinic-vet-api/app/modules/auth/infrastructure/repository"
	"clinic-vet-api/app/modules/auth/presentation/controller"
	"clinic-vet-api/app/modules/auth/presentation/routes"
	event "clinic-vet-api/app/modules/core/domain/event/user_event"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/modules/core/service"
	userPersistence "clinic-vet-api/app/modules/users/infrastructure/repository"
	"clinic-vet-api/app/shared/password"
	"clinic-vet-api/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
)

// TODO: EVENT
func SetupAuthModule(
	r *gin.RouterGroup,
	validator *validator.Validate,
	employeeRepo repository.EmployeeRepository,
	customerRepo repository.CustomerRepository,
	client *redis.Client,
	queries *sqlc.Queries,
	secretKet string,
	authMiddle *middleware.AuthMiddleware,
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

	registerHandler := register.NewRegisterCommandHandler(userRepo, passwordEncoder, userEventProducer, *security_service)
	loginHandler := authentication.NewLoginCommandHandler(userRepo, *security_service, session, jwtService)
	sessionHandler := sesionCmd.NewSessionCommandHandler(userRepo, session, jwtService)

	cmdBus := command.NewAuthCommandBus(loginHandler, registerHandler, sessionHandler)
	authController := controller.NewAuthController(validator, cmdBus)
	tokenController := controller.NewTokenController(validator, cmdBus)

	authRoutes := routes.NewAuthRoutes(authController, tokenController)
	authRoutes.RegisterRegistAndLoginRoutes(r, authMiddle)
	authRoutes.RegisterSessionRoutes(r, authMiddle)
}
