package ownerAPI

import (
	ownerUsecase "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/usecase"
	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	ownerController "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/infrastructure/api/controller"
	ownerRoutes "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/infrastructure/api/routes"
	sqlcOwnerRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/infrastructure/persistence"
	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type OwnerAPI struct {
	repository ownerDomain.OwnerRepository
	useCase    ownerUsecase.OwnerServiceFacade
	controller *ownerController.OwnerController
}

func NewOwnerAPI(r *gin.Engine, validator *validator.Validate, queries *sqlc.Queries, petRepo petDomain.PetRepository) *OwnerAPI {
	ownerRepo := sqlcOwnerRepository.NewSqlcOwnerRepository(queries, petRepo)
	getOwnerUseCase := ownerUsecase.NewGetOwnerByIdUseCase(ownerRepo)
	listOwnerUseCase := ownerUsecase.NewListOwnersUseCase(ownerRepo)
	createOwnerUseCase := ownerUsecase.NewCreateOwnerUseCase(ownerRepo)
	updateOwnerUseCase := ownerUsecase.NewUpdateOwnerUseCase(ownerRepo)
	deleteOwnerUseCase := ownerUsecase.NewSoftDeleteOwnerUseCase(ownerRepo)
	ownerUCContainer := ownerUsecase.NewOwnerUseCases(getOwnerUseCase, listOwnerUseCase, createOwnerUseCase, updateOwnerUseCase, deleteOwnerUseCase)
	ownerController := ownerController.NewOwnerController(validator, ownerUCContainer)
	ownerRoutes.OwnerRoutes(r, ownerController)

	return &OwnerAPI{
		repository: ownerRepo,
		useCase:    ownerUCContainer,
		controller: ownerController,
	}
}

func (api *OwnerAPI) GetRepository() ownerDomain.OwnerRepository {
	return api.repository
}

func (api *OwnerAPI) GetUseCase() ownerUsecase.OwnerServiceFacade {
	return api.useCase
}

func (api *OwnerAPI) GetController() *ownerController.OwnerController {
	return api.controller
}
