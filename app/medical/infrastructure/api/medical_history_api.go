package medHistoryAPI

import (
	medHistUsecases "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/application/usecase"
	medHistRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/domain/repositories"
	med_hist_controller "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/infrastructure/api/controller"
	medHistoryRoutes "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/infrastructure/api/routes"
	sqlcMedHistoryRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/infrastructure/persistence/repositories"
	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
	vetRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MedicalHistoryAPI struct {
	sqlcMedHistRepo     medHistRepo.MedicalHistoryRepository
	medHistUseCase      *medHistUsecases.MedicalHistoryUseCase
	med_hist_controller *med_hist_controller.AdminMedicalHistoryController
}

func NewMedicalHistoryAPI(
	queries *sqlc.Queries,
	router *gin.Engine,
	dataValidator *validator.Validate,
	ownerRepo ownerDomain.OwnerRepository,
	vetRepo vetRepo.VeterinarianRepository,
	petRepo petDomain.PetRepository,
) *MedicalHistoryAPI {
	sqlcMedHistRepo := sqlcMedHistoryRepo.NewSQLCMedHistRepository(queries)
	medHistUseCase := medHistUsecases.NewMedicalHistoryUseCase(sqlcMedHistRepo, ownerRepo, vetRepo, petRepo)
	med_hist_controller := med_hist_controller.NewAdminMedicalHistoryController(medHistUseCase)

	medHistoryRoutes.MedicalHistoryRoutes(router, *med_hist_controller)

	return &MedicalHistoryAPI{
		sqlcMedHistRepo:     sqlcMedHistRepo,
		medHistUseCase:      medHistUseCase,
		med_hist_controller: med_hist_controller,
	}
}

func (m *MedicalHistoryAPI) GetUseCase() *medHistUsecases.MedicalHistoryUseCase {
	return m.medHistUseCase
}

func (m *MedicalHistoryAPI) GetController() *med_hist_controller.AdminMedicalHistoryController {
	return m.med_hist_controller
}

func (m *MedicalHistoryAPI) GetRepository() medHistRepo.MedicalHistoryRepository {
	return m.sqlcMedHistRepo
}
