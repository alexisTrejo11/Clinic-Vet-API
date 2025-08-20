package userAPI

import (
	userApplication "github.com/alexisTrejo11/Clinic-Vet-API/app/users/application"
	userController "github.com/alexisTrejo11/Clinic-Vet-API/app/users/infrastructure/api/controller"
	userRoutes "github.com/alexisTrejo11/Clinic-Vet-API/app/users/infrastructure/api/routes"
	sqlcUserRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/users/infrastructure/persistence/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserAPI struct {
}

func NewUserAPI(queries *sqlc.Queries, dataValidator *validator.Validate, router *gin.Engine) *UserAPI {
	userRepo := sqlcUserRepo.NewSQLCUserRepository(queries)
	userDispatcher := userApplication.NewCommandDispatcher()
	userDispatcher.RegisterCurrentCommands(userRepo)
	userControllers := userController.NewUserAdminController(dataValidator, userDispatcher)
	userRoutes.UserRoutes(router, userControllers)

	return &UserAPI{}
}
