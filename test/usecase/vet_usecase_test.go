package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	vetDtos "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/dtos"
	vetUsecase "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/usecase"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/test/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type VetUseCaseTestSuite struct {
	suite.Suite
	ctrl           *gomock.Controller
	mockRepository *mock.MockVeterinarianRepository
	useCases       *vetUsecase.VeterinarianUseCases
}

func TestVetUseCaseSuite(t *testing.T) {
	suite.Run(t, new(VetUseCaseTestSuite))
}

func (s *VetUseCaseTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.mockRepository = mock.NewMockVeterinarianRepository(s.ctrl)

	// Inicializar los use cases individuales
	listVetUC := vetUsecase.NewListVetUseCase(s.mockRepository)
	getVetUC := vetUsecase.NewGetVetByIdUseCase(s.mockRepository)
	createVetUC := vetUsecase.NewCreateVetUseCase(s.mockRepository)
	updateVetUC := vetUsecase.NewUpdateVetUseCase(s.mockRepository)
	deleteVetUC := vetUsecase.NewDeleteVetUseCase(s.mockRepository)

	// Crear el agregado de use cases
	s.useCases = vetUsecase.NewVetUseCase(
		*listVetUC,
		*getVetUC,
		*createVetUC,
		*updateVetUC,
		*deleteVetUC,
	)
}

func (s *VetUseCaseTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *VetUseCaseTestSuite) TestGetVetById_Success() {
	// Arrange
	ctx := context.Background()
	now := time.Now()
	vetID := int(1)

	name, _ := shared.NewPersonName("John", "Doe")
	mockVet := vetDomain.Veterinarian{
		ID:              vetID,
		Name:            name,
		LicenseNumber:   "VET123",
		Specialty:       vetDomain.GeneralPracticeSpecialty,
		YearsExperience: 5,
		IsActive:        true,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	expectedResponse := vetDtos.VetResponse{
		Id:              vetID,
		FirstName:       "John",
		LastName:        "Doe",
		LicenseNumber:   "VET123",
		Specialty:       "general_practice",
		YearsExperience: 5,
	}

	s.mockRepository.EXPECT().
		GetByID(ctx, vetID).
		Return(mockVet, nil)

	// Assert
	result, err := s.useCases.GetVetByIdUseCase(ctx, vetID)

	// Arrange
	s.NoError(err)
	s.Equal(expectedResponse, result)
}
