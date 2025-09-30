// Package api contains the API builder for the users module.
package api

import (
	"clinic-vet-api/app/middleware"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/modules/core/service"
	employeeRepo "clinic-vet-api/app/modules/employee/infrastructure/repository"
	"clinic-vet-api/app/modules/users/application/usecase"
	"clinic-vet-api/app/modules/users/infrastructure/bus"
	"clinic-vet-api/app/modules/users/presentation/controller"
	"clinic-vet-api/app/modules/users/presentation/routes"
	"clinic-vet-api/app/shared/mapper"
	"clinic-vet-api/app/shared/password"
	"clinic-vet-api/sqlc"
	"fmt"

	persistence "clinic-vet-api/app/modules/users/infrastructure/repository"

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
	Queries        *sqlc.Queries
	Router         *gin.RouterGroup
	DataValidator  *validator.Validate
	AuthMiddleware *middleware.AuthMiddleware
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

	userRepo := persistence.NewSqlcUserRepository(u.config.Queries, mapper.NewSqlcFieldMapper())
	profileRepo := persistence.NewSQLCProfileRepository(u.config.Queries)
	employeeRepo := employeeRepo.NewSqlcEmployeeRepository(u.config.Queries, mapper.NewSqlcFieldMapper())
	passwordEncoder := password.NewPasswordEncoder()

	service := service.NewUserSecurityService(userRepo, employeeRepo, passwordEncoder)

	commandUserBus := bus.NewUserCommandBus(userRepo, service)
	queryUserBus := bus.NewUserQueryBus(userRepo)
	userServiceBus := bus.NewUserBus(queryUserBus, commandUserBus)

	userControllers := controller.NewUserAdminController(u.config.DataValidator, userServiceBus)
	profileUseCases := usecase.NewProfileUseCases(profileRepo)
	profileController := controller.NewProfileController(profileUseCases)

	routes.UserRoutes(u.config.Router, userControllers, u.config.AuthMiddleware)
	routes.ProfileRoutes(u.config.Router, profileController, u.config.AuthMiddleware)

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

	if u.config.AuthMiddleware == nil {
		return fmt.Errorf("auth middleware cannot be nil")
	}

	return nil
}

func (u *UserAPIModule) GetRepository() repository.UserRepository {
	if u.components.repository == nil {
		panic("repository not initialized, ensure Bootstrap() is called first")
	}
	return u.components.repository
}

func (u *UserAPIModule) SetAuthMiddleware(middleware *middleware.AuthMiddleware) {
	u.config.AuthMiddleware = middleware
}
