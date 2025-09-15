package api

import (
	"fmt"

	repository "clinic-vet-api/app/core/repositories"
	"clinic-vet-api/app/modules/users/application/usecase"
	"clinic-vet-api/app/modules/users/infrastructure/api/controller"
	"clinic-vet-api/app/modules/users/infrastructure/api/routes"
	"clinic-vet-api/app/modules/users/infrastructure/bus"
	persistence "clinic-vet-api/app/modules/users/infrastructure/persistence/repository"
	"clinic-vet-api/app/shared/cqrs"

	"clinic-vet-api/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserAPIComponents struct {
	repository        repository.UserRepository
	adminController   controller.UserAdminController
	profileController controller.ProfileController
	commandBus        cqrs.CommandBus
	quertBus          cqrs.QueryBus
}

type UserAPIConfig struct {
	Queries       *sqlc.Queries
	Router        *gin.Engine
	DataValidator *validator.Validate
}

func NewUserAPIConfig(queries *sqlc.Queries, router *gin.Engine, dataValidator *validator.Validate) *UserAPIConfig {
	return &UserAPIConfig{
		Queries:       queries,
		Router:        router,
		DataValidator: dataValidator,
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

	commandUserBus := bus.NewUserCommandBus(userRepo)
	queryUserBus := bus.NewUserQueryBus(userRepo)

	userControllers := controller.NewUserAdminController(u.config.DataValidator, commandUserBus, queryUserBus)
	profileUseCases := usecase.NewProfileUseCases(profileRepo)
	profileController := controller.NewProfileController(profileUseCases)

	routes.UserRoutes(u.config.Router, userControllers)

	u.components = &UserAPIComponents{
		repository:        userRepo,
		adminController:   *userControllers,
		profileController: *profileController,
		commandBus:        commandUserBus,
		quertBus:          queryUserBus,
	}
	u.isBuilt = true

	return nil
}

func (u *UserAPIModule) GetComponents() UserAPIComponents {
	return UserAPIComponents{
		repository:        u.components.repository,
		adminController:   u.components.adminController,
		profileController: u.components.profileController,
		commandBus:        u.components.commandBus,
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
