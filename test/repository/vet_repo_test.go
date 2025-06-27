package sqlcVetRepo_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	vetRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/repositories"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
	sqlcVetRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/infrastructure/persistence/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/alexisTrejo11/Clinic-Vet-API/test/mock"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type SqlcVetRepositoryTestSuite struct {
	suite.Suite
	ctrl        *gomock.Controller
	mockQueries *mock.MockQuerier
	repo        vetRepo.VeterinarianRepository
}

func TestSqlcVetRepositorySuite(t *testing.T) {
	suite.Run(t, new(SqlcVetRepositoryTestSuite))
}

func (s *SqlcVetRepositoryTestSuite) SetupTest() {

	s.ctrl = gomock.NewController(s.T())
	s.mockQueries = mock.NewMockQuerier(s.ctrl)
	s.repo = sqlcVetRepo.NewSqlcVetRepository(s.mockQueries)
}

func (s *SqlcVetRepositoryTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *SqlcVetRepositoryTestSuite) TestList_Success() {
	ctx := context.Background()
	now := time.Now()

	name, _ := shared.NewPersonName("John", "Doe")
	expectedVets := []vetDomain.Veterinarian{
		{
			ID:               1,
			Name:             name,
			LicenseNumber:    "VET123",
			Specialty:        vetDomain.GeneralPracticeSpecialty,
			WorkDaysSchedule: []vetDomain.WorkDaySchedule{},
			YearsExperience:  5,
			IsActive:         true,
			CreatedAt:        now,
			UpdatedAt:        now,
		},
	}

	mockReturn := []sqlc.Veterinarian{
		{
			ID:                1,
			FirstName:         "John",
			LastName:          "Doe",
			LicenseNumber:     "VET123",
			Speciality:        sqlc.VeterinarianSpecialityGeneralPractice,
			YearsOfExperience: 5,
			IsActive:          pgtype.Bool{Bool: true, Valid: true},
			CreatedAt:         pgtype.Timestamptz{Time: now, Valid: true},
			UpdatedAt:         pgtype.Timestamptz{Time: now, Valid: true},
		},
	}

	s.mockQueries.EXPECT().
		ListVeterinarians(ctx, gomock.Any()).
		Return(mockReturn, nil)

	result, err := s.repo.List(ctx, nil)

	s.NoError(err)
	s.Equal(expectedVets, result)
}

func (s *SqlcVetRepositoryTestSuite) TestGetByID_Success() {
	ctx := context.Background()
	now := time.Now()
	vetID := uint(1)

	// Expected domain model
	name, _ := shared.NewPersonName("John", "Doe")

	expected := vetDomain.Veterinarian{
		ID:               vetID,
		Name:             name,
		LicenseNumber:    "VET123",
		WorkDaysSchedule: []vetDomain.WorkDaySchedule{},
		Specialty:        vetDomain.GeneralPracticeSpecialty,
		YearsExperience:  5,
		IsActive:         true,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	// Mock database return
	mockReturn := sqlc.Veterinarian{
		ID:                1,
		FirstName:         "John",
		LastName:          "Doe",
		LicenseNumber:     "VET123",
		Speciality:        sqlc.VeterinarianSpecialityGeneralPractice,
		YearsOfExperience: 5,
		IsActive:          pgtype.Bool{Bool: true, Valid: true},
		CreatedAt:         pgtype.Timestamptz{Time: now, Valid: true},
		UpdatedAt:         pgtype.Timestamptz{Time: now, Valid: true},
	}

	// Set expectations
	s.mockQueries.EXPECT().
		GetVeterinarianById(ctx, int32(vetID)).
		Return(mockReturn, nil)

	// Call the method
	result, err := s.repo.GetByID(ctx, vetID)

	// Assertions
	s.NoError(err)
	s.Equal(expected, result)
}

func (s *SqlcVetRepositoryTestSuite) TestSave_Create_Success() {
	ctx := context.Background()
	now := time.Now()

	name, _ := shared.NewPersonName("John", "Doe")
	vet := &vetDomain.Veterinarian{
		Name:            name,
		LicenseNumber:   "VET123",
		Specialty:       vetDomain.GeneralPracticeSpecialty,
		YearsExperience: 5,
		IsActive:        true,
	}

	createParams := sqlc.CreateVeterinarianParams{
		FirstName:         "John",
		LastName:          "Doe",
		LicenseNumber:     "VET123",
		Speciality:        sqlc.VeterinarianSpeciality("general_practice"),
		YearsOfExperience: 5,
		IsActive:          pgtype.Bool{Bool: true, Valid: true},
	}

	mockVet := sqlc.Veterinarian{
		ID:                1,
		FirstName:         "John",
		LastName:          "Doe",
		LicenseNumber:     "VET123",
		Speciality:        sqlc.VeterinarianSpeciality("general_practice"),
		YearsOfExperience: 5,
		IsActive:          pgtype.Bool{Bool: true, Valid: true},
		CreatedAt:         pgtype.Timestamptz{Time: now, Valid: true},
		UpdatedAt:         pgtype.Timestamptz{Time: now, Valid: true},
	}

	s.mockQueries.EXPECT().
		CreateVeterinarian(ctx, createParams).
		Return(mockVet, nil)

	err := s.repo.Save(ctx, vet)

	s.NoError(err)
	s.Equal(uint(1), vet.ID)
	s.Equal(now, vet.CreatedAt)
	s.Equal(now, vet.UpdatedAt)
}

func (s *SqlcVetRepositoryTestSuite) TestSave_Update_Success() {
	ctx := context.Background()
	now := time.Now()

	name, _ := shared.NewPersonName("John", "Doe")
	vet := &vetDomain.Veterinarian{
		ID:              1,
		Name:            name,
		LicenseNumber:   "VET123",
		Specialty:       vetDomain.GeneralPracticeSpecialty,
		YearsExperience: 5,
		IsActive:        true,
		CreatedAt:       now.Add(-24 * time.Hour),
	}

	updateParams := sqlc.UpdateVeterinarianParams{
		ID:                1,
		FirstName:         "John",
		LastName:          "Doe",
		LicenseNumber:     "VET123",
		Speciality:        sqlc.VeterinarianSpeciality("general_practice"),
		YearsOfExperience: 5,
		IsActive:          pgtype.Bool{Bool: true, Valid: true},
	}

	mockReturn := sqlc.UpdateVeterinarianRow{
		ID:            1,
		FirstName:     "John",
		LastName:      "Doe",
		LicenseNumber: "VET123",
		Speciality:    sqlc.VeterinarianSpecialityGeneralPractice,
		UpdatedAt:     pgtype.Timestamptz{Time: now, Valid: true},
	}

	s.mockQueries.EXPECT().
		UpdateVeterinarian(ctx, updateParams).
		Return(mockReturn, nil)

	err := s.repo.Save(ctx, vet)

	s.NoError(err)
	s.Equal(now, vet.UpdatedAt)
}

func (s *SqlcVetRepositoryTestSuite) TestExists_True() {
	ctx := context.Background()
	id := uint(1)

	mockVet := sqlc.Veterinarian{
		ID: 1,
	}

	s.mockQueries.EXPECT().
		GetVeterinarianById(ctx, int32(id)).
		Return(mockVet, nil)

	exists, err := s.repo.Exists(ctx, id)

	s.NoError(err)
	s.True(exists)
}

func (s *SqlcVetRepositoryTestSuite) TestExists_False() {
	ctx := context.Background()
	id := uint(999)

	s.mockQueries.EXPECT().
		GetVeterinarianById(ctx, int32(id)).
		Return(sqlc.Veterinarian{}, sql.ErrNoRows)

	exists, err := s.repo.Exists(ctx, id)

	s.NoError(err)
	s.False(exists)
}
