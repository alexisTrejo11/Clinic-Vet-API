package vetAPI

import (
	vetRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/repositories"
	vetUsecase "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/usecase"
	vetController "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/infrastructure/api/controller"
	vetRoutes "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/infrastructure/api/routes"
	sqlcVetRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/infrastructure/persistence/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VeterinarianAPI struct {
	useCase    *vetUsecase.VeterinarianUseCases
	controller *vetController.VeterinarianController
	repository vetRepo.VeterinarianRepository
}

func NewVeterinarianAPI(
	queries *sqlc.Queries,
	router *gin.Engine,
	dataValidator *validator.Validate,
) *VeterinarianAPI {
	vetRepo := sqlcVetRepo.NewSqlcVetRepository(queries)

	getVetUseCase := vetUsecase.NewGetVetByIdUseCase(vetRepo)
	listVetUseCase := vetUsecase.NewListVetUseCase(vetRepo)
	createVetUseCase := vetUsecase.NewCreateVetUseCase(vetRepo)
	updateVetUseCase := vetUsecase.NewUpdateVetUseCase(vetRepo)
	deleteVetUseCase := vetUsecase.NewDeleteVetUseCase(vetRepo)

	vetUseCaseContainer := vetUsecase.NewVetUseCase(*listVetUseCase, *getVetUseCase, *createVetUseCase, *updateVetUseCase, *deleteVetUseCase)

	vetControllers := vetController.NewVeterinarianController(dataValidator, *vetUseCaseContainer)

	vetRoutes.VetRoutes(router, vetControllers)
	return &VeterinarianAPI{
		useCase:    vetUseCaseContainer,
		controller: vetControllers,
		repository: vetRepo,
	}
}

func (s *VeterinarianAPI) GetUseCase() *vetUsecase.VeterinarianUseCases {
	return s.useCase
}

func (s *VeterinarianAPI) GetController() *vetController.VeterinarianController {
	return s.controller
}

func (s *VeterinarianAPI) GetRepository() vetRepo.VeterinarianRepository {
	return s.repository
}
