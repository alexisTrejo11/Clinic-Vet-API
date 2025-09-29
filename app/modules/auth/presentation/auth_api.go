package api

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/auth/application/handler"
	"clinic-vet-api/app/modules/auth/infrastructure/bus"
	"clinic-vet-api/app/modules/auth/infrastructure/jwt"
	repositoryimpl "clinic-vet-api/app/modules/auth/infrastructure/repository"
	"clinic-vet-api/app/modules/auth/presentation/controller"
	"clinic-vet-api/app/modules/auth/presentation/routes"
	event "clinic-vet-api/app/modules/core/domain/event/user_event"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/modules/core/service"
	userPersistence "clinic-vet-api/app/modules/users/infrastructure/repository"
	"clinic-vet-api/app/shared/mapper"
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
	userRepo := userPersistence.NewSqlcUserRepository(queries, mapper.NewSqlcFieldMapper())
	jwtService := jwt.NewJWTService(secretKet)
	passwordEncoder := password.NewPasswordEncoder()
	emailService := service.NewEmailService()

	session := repositoryimpl.NewRedisSessionRepository(client)
	security_service := service.NewUserSecurityService(userRepo, employeeRepo, passwordEncoder)
	profileRepo := userPersistence.NewSQLCProfileRepository(queries)

	eventService := service.NewUserAccountService(userRepo, profileRepo, customerRepo, employeeRepo, emailService)
	userEventProducer := event.NewUserEventHandler(eventService)

	authHandler := handler.NewAuthCommandHandler(
		userRepo,
		*security_service,
		session,
		jwtService,
		passwordEncoder,
		userEventProducer,
	)

	cmdBus := bus.NewAuthCommandBus(*authHandler)
	authController := controller.NewAuthController(validator, cmdBus)
	tokenController := controller.NewTokenController(validator, cmdBus)

	authRoutes := routes.NewAuthRoutes(authController, tokenController)
	authRoutes.RegisterRegistAndLoginRoutes(r, authMiddle)
	authRoutes.RegisterSessionRoutes(r, authMiddle)
}
