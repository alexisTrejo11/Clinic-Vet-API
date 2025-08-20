package petAPI

import (
	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	ownerRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/infrastructure/persistence"
	petUsecase "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/usecase"
	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
	petController "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/infrastructure/api/controller"
	petRoutes "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/infrastructure/api/routes"
	sqlcPetRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/infrastructure/persistence/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PetAPI struct {
	repository petDomain.PetRepository
	useCases   petUsecase.PetUseCasesFacade
	controller *petController.PetController
}

func NewPetAPI(
	router *gin.Engine,
	queries *sqlc.Queries,
	dataValidator *validator.Validate,
) *PetAPI {
	petRepository := sqlcPetRepository.NewSqlcPetRepository(queries)
	ownerRepo := ownerRepository.NewSqlcOwnerRepository(queries, petRepository)

	repository := setUpRepository(queries)
	useCases := setupApplication(repository, ownerRepo)
	controller := setupController(dataValidator, useCases)
	setupRoutes(router, controller)
	return &PetAPI{
		repository: repository,
		useCases:   useCases,
		controller: controller,
	}
}

func setupRoutes(app *gin.Engine, petController *petController.PetController) {
	petRoutes.PetsRoutes(app, petController)
}

func setUpRepository(queries *sqlc.Queries) petDomain.PetRepository {
	return sqlcPetRepository.NewSqlcPetRepository(queries)
}

func setupApplication(repository petDomain.PetRepository, ownerRepo ownerDomain.OwnerRepository) petUsecase.PetUseCasesFacade {
	return petUsecase.NewPetUseCasesFacade(repository, ownerRepo)
}

func setupController(validator *validator.Validate, facade petUsecase.PetUseCasesFacade) *petController.PetController {
	return petController.NewPetController(validator, facade)
}

func RegisterRoutes(app *gin.Engine, petController *petController.PetController) {
	petRoutes.PetsRoutes(app, petController)
}

func (api *PetAPI) GetUseCases() petUsecase.PetUseCasesFacade {
	return api.useCases
}

func (api *PetAPI) GetController() *petController.PetController {
	return api.controller
}

func (api *PetAPI) GetRepository() petDomain.PetRepository {
	return api.repository
}
