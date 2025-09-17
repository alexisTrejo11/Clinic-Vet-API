package api

import (
	"fmt"

	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/core/service"
	"clinic-vet-api/app/modules/users/application/usecase"
	"clinic-vet-api/app/modules/users/infrastructure/bus"
	persistence "clinic-vet-api/app/modules/users/infrastructure/repository"
	"clinic-vet-api/app/modules/users/presentation/controller"
	"clinic-vet-api/app/modules/users/presentation/routes"
	"clinic-vet-api/app/shared/password"

	"clinic-vet-api/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserAPIComponents struct {
	repository        repository.UserRepository
	adminController   controller.UserAdminController
	profileController controller.ProfileController
	userServiceBus    bus.UserBus
}

type UserAPIConfig struct {
	Queries       *sqlc.Queries
	Router        *gin.Engine
	DataValidator *validator.Validate
	employeeRepo  repository.EmployeeRepository
}

func NewUserAPIConfig(queries *sqlc.Queries, router *gin.Engine, dataValidator *validator.Validate, employeeRepo repository.EmployeeRepository) *UserAPIConfig {
	return &UserAPIConfig{
		Queries:       queries,
		Router:        router,
		DataValidator: dataValidator,
		employeeRepo:  employeeRepo,
	}
}

type UserAPIModule struct {
	config     *UserAPIConfig
	components *UserAPIComponents
	isBuilt    bool
}

func NewUserAPIModule(config UserAPIConfig) *UserAPIModule {
	return &UserAPIModule{
		config:     &config,
		components: &UserAPIComponents{},
		isBuilt:    false,
	}
}

func (u *UserAPIModule) Bootstrap() error {
	if u.isBuilt {
		return nil
	}

	if err := u.validateConfig(); err != nil {
		return err
	}

	userRepo := persistence.NewSQLCUserRepository(u.config.Queries)
	profileRepo := persistence.NewSQLCProfileRepository(u.config.Queries)
	employeeRepo := u.config.employeeRepo
	passwordEncoder := password.NewPasswordEncoder()

	service := service.NewUserSecurityService(userRepo, employeeRepo, passwordEncoder)

	commandUserBus := bus.NewUserCommandBus(userRepo, service)
	queryUserBus := bus.NewUserQueryBus(userRepo)
	userServiceBus := bus.NewUserBus(queryUserBus, commandUserBus)

	userControllers := controller.NewUserAdminController(u.config.DataValidator, userServiceBus)
	profileUseCases := usecase.NewProfileUseCases(profileRepo)
	profileController := controller.NewProfileController(profileUseCases)

	routes.UserRoutes(u.config.Router, userControllers)

	u.components = &UserAPIComponents{
		repository:        userRepo,
		adminController:   *userControllers,
		profileController: *profileController,
		userServiceBus:    *userServiceBus,
	}
	u.isBuilt = true

	return nil
}

func (u *UserAPIModule) GetComponents() UserAPIComponents {
	return UserAPIComponents{
		repository:        u.components.repository,
		adminController:   u.components.adminController,
		profileController: u.components.profileController,
		userServiceBus:    u.components.userServiceBus,
	}
}

func (u *UserAPIModule) validateConfig() error {
	if u.config == nil {
		return fmt.Errorf("configuration cannot be nil")
	}
	if u.config.Router == nil {
		return fmt.Errorf("router cannot be nil")
	}
	if u.config.Queries == nil {
		return fmt.Errorf("queries cannot be nil")
	}
	if u.config.DataValidator == nil {
		return fmt.Errorf("data validator cannot be nil")
	}

	return nil
}
