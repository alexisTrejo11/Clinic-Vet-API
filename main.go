package main

import (
	"context"
	"log"
	"os"

	ownerUsecase "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/usecase"
	ownerController "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/infrastructure/api/controller"
	ownerRoutes "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/infrastructure/api/routes"
	sqlcOwnerRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/infrastructure/persistence"
	petUsecase "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/usecase"
	petController "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/infrastructure/api/controller"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/pets/infrastructure/api/routes"
	sqlcPetRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/infrastructure/persistence/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/config"
	"github.com/alexisTrejo11/Clinic-Vet-API/middleware"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/gin-gonic/gin"
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

	// Pet Controller
	petController := petController.NewPetController(validator.New(), getPetUseCase, listPetsUseCase, createPetsUseCase, updatePetsUseCase, deletePetsUseCase)
	ownerController := ownerController.NewOwnerController(validator.New(), ownerUCContainer)

	routes.PetsRoutes(router, petController)
	ownerRoutes.OwnerRoutes(router, ownerController)

	router.Run()
}
