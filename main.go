package main

import (
	"context"
	"log"
	"os"

	appointmentAPI "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/infrastructure/api"
	mhDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/application/dtos"
	medHistUsecases "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/application/usecase"
	med_hist_controller "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/infrastructure/api/controller"
	medHistoryRoutes "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/infrastructure/api/routes"
	sqlcMedHistoryRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/infrastructure/persistence/repositories"
	ownerUsecase "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/usecase"
	ownerController "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/infrastructure/api/controller"
	ownerRoutes "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/infrastructure/api/routes"
	sqlcOwnerRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/infrastructure/persistence"
	paymentAPI "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/infrastructure/api"
	paymentRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/infrastructure/persistence"
	petUsecase "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/usecase"
	petController "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/infrastructure/api/controller"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/pets/infrastructure/api/routes"
	sqlcPetRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/infrastructure/persistence/repositories"
	userApplication "github.com/alexisTrejo11/Clinic-Vet-API/app/users/application"
	userController "github.com/alexisTrejo11/Clinic-Vet-API/app/users/infrastructure/api/controller"
	userRoutes "github.com/alexisTrejo11/Clinic-Vet-API/app/users/infrastructure/api/routes"
	sqlcUserRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/users/infrastructure/persistence/repository"
	vetUsecase "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/usecase"
	vetController "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/infrastructure/api/controller"
	vetRoutes "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/infrastructure/api/routes"
	sqlcVetRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/infrastructure/persistence/repositories"

	"github.com/alexisTrejo11/Clinic-Vet-API/config"
	"github.com/alexisTrejo11/Clinic-Vet-API/middleware"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	config.InitLogger()
	defer config.SyncLogger()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while loading .env: %v", err)
	}

	ctx := context.Background()
	dbConn := config.DbConn(os.Getenv("DATABASE_URL"))
	defer dbConn.Close(ctx)

	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.AuditLog())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	queries := sqlc.New(dbConn)

	// Repository
	petRepo := sqlcPetRepository.NewSqlcPetRepository(queries)
	ownerRepo := sqlcOwnerRepository.NewSlqcOwnerRepository(queries, petRepo)
	vetRepo := sqlcVetRepo.NewSqlcVetRepository(queries)
	sqlcMedHistRepo := sqlcMedHistoryRepo.NewSQLCMedHistRepository(queries)
	paymentRepo := paymentRepo.NewSQLCPaymentRepository(queries)

	// Medical History UseCase
	medHistUseCase := medHistUsecases.NewMedicalHistoryUseCase(sqlcMedHistRepo, ownerRepo, vetRepo)

	// Owner UseCase
	getOwnerUseCase := ownerUsecase.NewGetOwnerByIdUseCase(ownerRepo)
	listOwnerUseCase := ownerUsecase.NewListOwnersUseCase(ownerRepo)
	createOwnerUseCase := ownerUsecase.NewCreateOwnerUseCase(ownerRepo)
	updateOwnerUseCase := ownerUsecase.NewUpdateOwnerUseCase(ownerRepo)
	deleteOwnerUseCase := ownerUsecase.NewDeleteOwnerUseCase(ownerRepo)

	ownerUCContainer := ownerUsecase.NewOwnerUseCases(getOwnerUseCase, listOwnerUseCase, createOwnerUseCase, updateOwnerUseCase, deleteOwnerUseCase)

	// Pet UseCase
	getPetUseCase := petUsecase.NewGetPetByIdUseCase(petRepo)
	listPetsUseCase := petUsecase.NewListPetsUseCase(petRepo)
	createPetsUseCase := petUsecase.NewCreatePetUseCase(petRepo, ownerRepo)
	updatePetsUseCase := petUsecase.NewUpdatePetUseCase(petRepo, ownerRepo)
	deletePetsUseCase := petUsecase.NewDeletePetUseCase(petRepo)

	// Vet UseCase
	getVetUseCase := vetUsecase.NewGetVetByIdUseCase(vetRepo)
	listVetUseCase := vetUsecase.NewListVetUseCase(vetRepo)
	createVetUseCase := vetUsecase.NewCreateVetUseCase(vetRepo)
	updateVetUseCase := vetUsecase.NewUpdateVetUseCase(vetRepo)
	deleteVetUseCase := vetUsecase.NewDeleteVetUseCase(vetRepo)
	vetUseCaseContainer := vetUsecase.NewVetUseCase(*listVetUseCase, *getVetUseCase, *createVetUseCase, *updateVetUseCase, *deleteVetUseCase)

	// Pet Controller
	dataValidator := validator.New()
	petController := petController.NewPetController(dataValidator, getPetUseCase, listPetsUseCase, createPetsUseCase, updatePetsUseCase, deletePetsUseCase)
	ownerController := ownerController.NewOwnerController(dataValidator, ownerUCContainer)
	vetControllers := vetController.NewVeterinarianController(dataValidator, *vetUseCaseContainer)

	// Medical History Routes
	med_hist_controller := med_hist_controller.NewAdminMedicalHistoryController(medHistUseCase)

	// Custom Validations
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("validVisitReason", mhDTOs.IsValidVisitReason)
		v.RegisterValidation("validVisitType", mhDTOs.IsValidVisitType)
		v.RegisterValidation("validPetCondition", mhDTOs.IsValidPetCondition)
	}

	routes.PetsRoutes(router, petController)
	ownerRoutes.OwnerRoutes(router, ownerController)
	vetRoutes.VetRoutes(router, vetControllers)
	medHistoryRoutes.MedicalHistoryRoutes(router, *med_hist_controller)

	// Payment
	paymentAPI.SetupPaymentAPI(router, dataValidator, paymentRepo)

	// User
	userRepo := sqlcUserRepo.NewSQLCUserRepository(queries)
	userDispatcher := userApplication.NewCommandDispatcher()
	userDispatcher.RegisterCurrentCommands(userRepo)
	userControllers := userController.NewUserAdminController(dataValidator, userDispatcher)
	userRoutes.UserRoutes(router, userControllers)

	// User Profile
	profileUseCase := userApplication.NewProfileUseCases(userRepo)
	profileController := userController.NewProfileController(profileUseCase)
	userRoutes.ProfileRoutes(router, profileController)

	appointmentAPI.SetupAppoinmentAPI(router, queries, dataValidator)

	router.Run()
}
