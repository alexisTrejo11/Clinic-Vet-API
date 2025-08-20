package userAPI

import (
	"fmt"

	userApplication "github.com/alexisTrejo11/Clinic-Vet-API/app/users/application"
	userRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
	userController "github.com/alexisTrejo11/Clinic-Vet-API/app/users/infrastructure/api/controller"
	userRoutes "github.com/alexisTrejo11/Clinic-Vet-API/app/users/infrastructure/api/routes"
	sqlcUserRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/users/infrastructure/persistence/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserAPIComponents struct {
	repository        userRepository.UserRepository
	adminController   userController.UserAdminController
	profileController userController.ProfileController
	dispatcher        userApplication.CommandDispatcher
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

	userRepo := sqlcUserRepo.NewSQLCUserRepository(u.config.Queries)
	userDispatcher := userApplication.NewCommandDispatcher()
	userDispatcher.RegisterCurrentCommands(userRepo)
	userControllers := userController.NewUserAdminController(u.config.DataValidator, userDispatcher)
	profileUseCases := userApplication.NewProfileUseCases(userRepo)
	profileController := userController.NewProfileController(profileUseCases)

	userRoutes.UserRoutes(u.config.Router, userControllers)

	u.components = &UserAPIComponents{
		repository:        userRepo,
		adminController:   *userControllers,
		profileController: *profileController,
		dispatcher:        *userDispatcher,
	}
	u.isBuilt = true

	return nil
}

func (u *UserAPIModule) GetComponents() UserAPIComponents {
	return UserAPIComponents{
		repository:        u.components.repository,
		adminController:   u.components.adminController,
		profileController: u.components.profileController,
		dispatcher:        u.components.dispatcher,
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
