package api

import (
	"clinic-vet-api/app/middleware"
	localApp "clinic-vet-api/app/modules/account/auth/local/application"
	localHandler "clinic-vet-api/app/modules/account/auth/local/application/handler"
	localCtrl "clinic-vet-api/app/modules/account/auth/local/presentation/controller"
	localRoutes "clinic-vet-api/app/modules/account/auth/local/presentation/routes"
	sessionApp "clinic-vet-api/app/modules/account/auth/session/application"
	sessionHandler "clinic-vet-api/app/modules/account/auth/session/application/handler"
	sessionRepo "clinic-vet-api/app/modules/account/auth/session/infrastructure/repository"
	sessionCtrl "clinic-vet-api/app/modules/account/auth/session/presentation/controller"
	sessionRoutes "clinic-vet-api/app/modules/account/auth/session/presentation/routes"
	token "clinic-vet-api/app/modules/account/auth/token/service"
	twoFaApp "clinic-vet-api/app/modules/account/auth/two_factor/application"
	twoFaHandler "clinic-vet-api/app/modules/account/auth/two_factor/application/handler"
	twoFaCtrl "clinic-vet-api/app/modules/account/auth/two_factor/presentation/controller"
	twoFaRoutes "clinic-vet-api/app/modules/account/auth/two_factor/presentation/routes"
	userRepo "clinic-vet-api/app/modules/account/user/infrastructure/repository"
	event "clinic-vet-api/app/modules/core/domain/event/user_event"
	irepo "clinic-vet-api/app/modules/core/repository"
	iservice "clinic-vet-api/app/modules/core/service"
	"clinic-vet-api/app/shared/mapper"
	"clinic-vet-api/app/shared/password"
	"clinic-vet-api/sqlc"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
)

type AuthAPI struct {
	isBuilt bool
	Modules *AuthModules
	config  *AuthModuleConfig
}

func NewAuthAPI(config *AuthModuleConfig) *AuthAPI {
	return &AuthAPI{
		isBuilt: false,
		config:  config,
	}
}

type AuthModules struct {
	TwoFactor *TwoFactorModule
	OAuth     *OAuthModule
	LocalAuth *LocalAuthModule
	Session   *SessionModule
}

type AuthModuleConfig struct {
	RouterGroup         *gin.RouterGroup
	Validator           *validator.Validate
	Client              *redis.Client
	Queries             *sqlc.Queries
	SecretKey           string
	AuthMiddleware      *middleware.AuthMiddleware
	NotificationService iservice.NotificationService
	EmployeeRepo        irepo.EmployeeRepository
	CustomerRepo        irepo.CustomerRepository
}

type TwoFactorModule struct {
	CommandHandler *twoFaHandler.TwoFACommandHandler
	Bus            twoFaApp.TwoFaAuthFacadeService
	Controller     *twoFaCtrl.TwoFAAuthController
	Routes         *twoFaRoutes.TwoFARoutes
}

type OAuthModule struct {
	CommandHandler any
	Bus            any
	Controller     any
	Routes         any
}

type LocalAuthModule struct {
	CmdHandler    *localHandler.AuthCommandHandler
	FacadeService localApp.AuthApplicationFacade
	Controller    *localCtrl.AuthController
	Routes        *localRoutes.AuthRoutes
}

type SessionModule struct {
	SessionRepo        irepo.SessionRepository
	ApplicationService sessionApp.SessionFacadeService
	Controller         *sessionCtrl.SessionController
	Routes             *sessionRoutes.SessionRoutes
}

func (api *AuthAPI) Boostrap() error {
	if api.isBuilt {
		return nil
	}

	if err := api.ValidateConfig(); err != nil {
		return err
	}

	// Repositories
	userRepo := userRepo.NewSqlcUserRepository(api.config.Queries, mapper.NewSqlcFieldMapper())
	sessionRepo := sessionRepo.NewRedisSessionRepository(api.config.Client)

	passwordEncoder := password.NewPasswordEncoder()

	// Services
	jwtService := token.NewJWTService(api.config.SecretKey)
	tokeManager := token.NewTokenManager([]byte(api.config.SecretKey), api.config.Client)
	securityService := iservice.NewUserSecurityService(userRepo, api.config.EmployeeRepo, passwordEncoder)
	authService := iservice.NewUserAccountService(userRepo, api.config.CustomerRepo, api.config.EmployeeRepo, tokeManager, api.config.NotificationService)
	tokenManager := token.NewTokenManager([]byte(api.config.SecretKey), api.config.Client)
	userEventProducer := event.NewUserEventHandler(*authService)

	// Local Auth Module
	localCmdHandler := localHandler.NewAuthCommandHandler(userRepo, *securityService, sessionRepo, jwtService, passwordEncoder, userEventProducer)
	localFacade := localApp.NewAuthCommandBus(localCmdHandler)
	localController := localCtrl.NewAuthController(api.config.Validator, localFacade)
	localRoutes := localRoutes.NewAuthRoutes(localController, api.config.RouterGroup, api.config.AuthMiddleware)

	localRoutes.RegisterRoutes()

	// Two-Factor Auth Module
	twoFAHandler := twoFaHandler.NewTwoFACommandHandler(userRepo, sessionRepo, *securityService, tokenManager, api.config.NotificationService, jwtService)
	twoFAFacade := twoFaApp.NewTwoFaAuthFacadeService(twoFAHandler)
	twoFAController := twoFaCtrl.NewTwoFAAuthController(twoFAFacade)
	twoFARoutes := twoFaRoutes.NewTwoFARoutes(twoFAController, api.config.RouterGroup, api.config.AuthMiddleware)

	twoFARoutes.RegisterRoutes()

	// Session Module
	sessionCmdHandler := sessionHandler.NewSessionCommandHandler(userRepo, sessionRepo, jwtService)
	sessionApplicationService := sessionApp.NewSessionFacadeService(sessionCmdHandler)
	sessionController := sessionCtrl.NewSessionController(api.config.Validator, sessionApplicationService)
	sessionRoutes := sessionRoutes.NewSessionRoutes(api.config.AuthMiddleware, sessionController, api.config.RouterGroup)

	sessionRoutes.RegisterRoutes()

	localAuthModule := &LocalAuthModule{
		CmdHandler:    localCmdHandler,
		FacadeService: localFacade,
		Controller:    localController,
		Routes:        localRoutes,
	}

	twoFactorModule := &TwoFactorModule{
		CommandHandler: twoFAHandler,
		Bus:            twoFAFacade,
		Controller:     twoFAController,
		Routes:         twoFARoutes,
	}

	sessionModule := &SessionModule{
		SessionRepo:        sessionRepo,
		ApplicationService: sessionApplicationService,
		Controller:         sessionController,
		Routes:             sessionRoutes,
	}

	api.Modules = &AuthModules{
		LocalAuth: localAuthModule,
		TwoFactor: twoFactorModule,
		Session:   sessionModule,
	}

	api.isBuilt = true

	return nil
}

type SessionAuthModule struct {
	Repository     irepo.SessionRepository
	CommandHandler any
	Bus            any
	Controller     any
	Routes         any
}

func (api *AuthAPI) ValidateConfig() error {
	errs := make([]error, 0)
	if api.config.RouterGroup == nil {
		errs = append(errs, errors.New("router group is nil"))
	}

	if api.config.Validator == nil {
		errs = append(errs, errors.New("validator is nil"))
	}

	if api.config.EmployeeRepo == nil {
		errs = append(errs, errors.New("employee repository is nil"))
	}

	if api.config.CustomerRepo == nil {
		errs = append(errs, errors.New("customer repository is nil"))
	}

	if api.config.Client == nil {
		errs = append(errs, errors.New("redis client is nil"))
	}

	if api.config.Queries == nil {
		errs = append(errs, errors.New("queries is nil"))
	}

	if api.config.SecretKey == "" {
		errs = append(errs, errors.New("secret key is empty"))
	}

	if api.config.AuthMiddleware == nil {
		errs = append(errs, errors.New("auth middleware is nil"))
	}

	if api.config.NotificationService == nil {
		errs = append(errs, errors.New("notification service is nil"))
	}

	if len(errs) > 0 {
		return errs[0]
	}
	return nil
}
