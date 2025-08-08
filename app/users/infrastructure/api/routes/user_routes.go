package routes

import (
	userController "github.com/alexisTrejo11/Clinic-Vet-API/app/users/infrastructure/api/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, userQueryController *userController.UserQueryController) {
	// User Query Routes
	path := "api/v2/admin/users"

	router.GET(path+"/:id", userQueryController.GetUserById)
	router.GET(path+"/search", userQueryController.SearchUsers)
	router.GET(path+"email/:email", userQueryController.GetUserByEmail)
	router.GET(path+"/phone/:phone", userQueryController.GetUserByPhone)
}
