package main

import (
	"context"
	"log"
	"os"

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

	petRepo := sqlcPetRepository.NewSqlcPetRepository(queries)
	getPetUseCase := petUsecase.NewGetPetByIdUseCase(petRepo)
	listPetsUseCase := petUsecase.NewListPetsUseCase(petRepo)
	createPetsUseCase := petUsecase.NewCreatePetUseCase(petRepo, nil) // TO Be IMPL
	updatePetsUseCase := petUsecase.NewUpdatePetUseCase(petRepo, nil)
	deletePetsUseCase := petUsecase.NewDeletePetUseCase(petRepo)
	petController := petController.NewPetController(validator.New(), getPetUseCase, listPetsUseCase, createPetsUseCase, updatePetsUseCase, deletePetsUseCase)

	routes.PetsRoutes(router, petController)
	router.Run()
}
