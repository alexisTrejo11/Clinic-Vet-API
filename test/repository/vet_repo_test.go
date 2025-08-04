package sqlcVetRepo_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	vetDtos "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/dtos"
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

func (s *SqlcVetRepositoryTestSuite) TestList_WithFiltersAndSorting() {
	ctx := context.Background()
	now := time.Now()

	// Configurar datos de prueba
	name, _ := shared.NewPersonName("John", "Doe")
	expectedVet := vetDomain.Veterinarian{
		ID:              1,
		Name:            name,
		LicenseNumber:   "VET123",
		Specialty:       vetDomain.GeneralPracticeSpecialty,
		Schedule:        nil,
		YearsExperience: 5,
		IsActive:        true,
		CreatedAt:       now,
		UpdatedAt:       now,
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

	// Configurar parámetros de búsqueda
	searchParams := vetDtos.VetSearchParams{
		PageData: page.PageData{
			PageSize:      10,
			PageNumber:    0,
			SortDirection: page.DESC,
		},
		OrderBy: "years_experience",
		Filters: vetDtos.VetFilters{
			Name:      StringPtr("John"),
			IsActive:  BoolPtr(true),
			Specialty: VetSpecialtyPtr(vetDomain.GeneralPracticeSpecialty),
			YearsExperience: &struct {
				Min *int `json:"min"`
				Max *int `json:"max"`
			}{
				Min: intPtr(3),
				Max: intPtr(10),
			},
		},
	}

	// Configurar expectativas del mock
	expectedParams := sqlc.ListVeterinariansParams{
		FirstName:           "%John%",
		LastName:            "%John%",
		LicenseNumber:       "%",
		Speciality:          sqlc.VeterinarianSpecialityGeneralPractice,
		YearsOfExperience:   3,
		YearsOfExperience_2: 10,
		IsActive:            pgtype.Bool{Bool: true, Valid: true},
		Column8:             false, // name asc
		Column9:             false, // name desc
		Column10:            false, // specialty asc
		Column11:            false, // specialty desc
		Column12:            false, // years asc
		Column13:            true,  // years desc (nuestro caso)
		Column14:            false, // created_at asc
		Column15:            false, // created_at desc
		Limit:               10,
		Offset:              0,
	}

	s.mockQueries.EXPECT().
		ListVeterinarians(ctx, expectedParams).
		Return(mockReturn, nil)

	// Ejecutar
	result, err := s.repo.List(ctx, searchParams)

	// Verificar
	s.NoError(err)
	s.Len(result, 1)
	s.Equal(expectedVet, result[0])
}

func (s *SqlcVetRepositoryTestSuite) TestList_NoFilters() {
	ctx := context.Background()
	now := time.Now()

	// Configurar datos de prueba
	name, _ := shared.NewPersonName("John", "Doe")
	expectedVet := vetDomain.Veterinarian{
		ID:              1,
		Name:            name,
		LicenseNumber:   "VET123",
		Specialty:       vetDomain.GeneralPracticeSpecialty,
		Schedule:        nil,
		YearsExperience: 5,
		IsActive:        true,
		CreatedAt:       now,
		UpdatedAt:       now,
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

	// Parámetros vacíos
	searchParams := vetDtos.VetSearchParams{
		PageData: page.PageData{
			PageSize:   10,
			PageNumber: 0,
		},
	}

	// Configurar expectativas del mock
	expectedParams := sqlc.ListVeterinariansParams{
		FirstName:           "%",
		LastName:            "%",
		LicenseNumber:       "%",
		Speciality:          "",
		YearsOfExperience:   0,
		YearsOfExperience_2: 0,
		IsActive:            pgtype.Bool{Bool: false, Valid: false},
		Column8:             false,
		Column9:             false,
		Column10:            false,
		Column11:            false,
		Column12:            false,
		Column13:            false,
		Column14:            false,
		Column15:            true, // Orden por defecto (created_at DESC)
		Limit:               10,
		Offset:              0,
	}

	s.mockQueries.EXPECT().
		ListVeterinarians(ctx, expectedParams).
		Return(mockReturn, nil)

	// Ejecutar
	result, err := s.repo.List(ctx, searchParams)

	// Verificar
	s.NoError(err)
	s.Len(result, 1)
	s.Equal(expectedVet, result[0])
}

func (s *SqlcVetRepositoryTestSuite) TestGetByID_Success() {
	ctx := context.Background()
	now := time.Now()
	vetID := int(1)

	// Expected domain model
	name, _ := shared.NewPersonName("John", "Doe")

	expected := vetDomain.Veterinarian{
		ID:              vetID,
		Name:            name,
		LicenseNumber:   "VET123",
		Schedule:        nil,
		Specialty:       vetDomain.GeneralPracticeSpecialty,
		YearsExperience: 5,
		IsActive:        true,
		CreatedAt:       now,
		UpdatedAt:       now,
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
	result, err := s.repo.GetById(ctx, vetID)

	// Assertions
	s.NoError(err)
	s.Equal(expected, result)
}

func (s *SqlcVetRepositoryTestSuite) TestList_NamePartialSearch() {
	ctx := context.Background()
	now := time.Now()

	// Configurar mock
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

	// Configurar parámetros de búsqueda
	searchParams := vetDtos.VetSearchParams{
		Filters: vetDtos.VetFilters{
			Name: StringPtr("Joh"),
		},
	}

	// Configurar expectativas del mock
	expectedParams := sqlc.ListVeterinariansParams{
		FirstName:           "%Joh%",
		LastName:            "%Joh%",
		LicenseNumber:       "%",
		Speciality:          "",
		YearsOfExperience:   0,
		YearsOfExperience_2: 0,
		IsActive:            pgtype.Bool{Bool: false, Valid: false},
		Column15:            true, // Orden por defecto
		Limit:               10,
		Offset:              0,
	}

	s.mockQueries.EXPECT().
		ListVeterinarians(ctx, gomock.Any()).
		DoAndReturn(func(_ context.Context, p sqlc.ListVeterinariansParams) ([]sqlc.Veterinarian, error) {
			s.Equal(expectedParams.FirstName, p.FirstName)
			s.Equal(expectedParams.LastName, p.LastName)
			return mockReturn, nil
		})

	// Ejecutar
	result, err := s.repo.List(ctx, searchParams)

	// Verificar
	s.NoError(err)
	s.Len(result, 1)
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
	s.Equal(int(1), vet.ID)
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

	mockReturn := sqlc.Veterinarian{
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
	id := int(1)

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
	id := int(999)

	s.mockQueries.EXPECT().
		GetVeterinarianById(ctx, int32(id)).
		Return(sqlc.Veterinarian{}, sql.ErrNoRows)

	exists, err := s.repo.Exists(ctx, id)

	s.NoError(err)
	s.False(exists)
}

func StringPtr(s string) *string {
	return &s
}

func BoolPtr(b bool) *bool {
	return &b
}

func intPtr(u int) *int {
	return &u
}

func VetSpecialtyPtr(s vetDomain.VetSpecialty) *vetDomain.VetSpecialty {
	return &s
}
