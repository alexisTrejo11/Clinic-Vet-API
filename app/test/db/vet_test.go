package db_test

import (
	"context"
	"testing"
	"time"

	"clinic-vet-api/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createDummyVeterinarian(t *testing.T) sqlc.Veterinarian {
	arg := sqlc.CreateVeterinarianParams{
		FirstName:         "Test Veterinarian FirstName",
		LastName:          "Test Veterinarian LastName",
		Photo:             "htpp://photo-test/vet-1235.com",
		Speciality:        sqlc.VeterinarianSpecialityAnesthesiology,
		LicenseNumber:     "1234567890",
		YearsOfExperience: 8,
		IsActive:          pgtype.Bool{Bool: true, Valid: true},
	}

	veterinarian, err := testQueries.CreateVeterinarian(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, veterinarian)

	require.NotZero(t, veterinarian.ID)
	require.NotZero(t, veterinarian.CreatedAt)

	return veterinarian

}

func TestCreateAccount(t *testing.T) {
	createDummyVeterinarian(t)
}

func TestGetVeterinarian(t *testing.T) {
	veterinarianCreated := createDummyVeterinarian(t)
	VeterinarianRetrieved, err := testQueries.GetVeterinarianById(context.Background(), veterinarianCreated.ID)
	require.NoError(t, err)
	require.NotEmpty(t, VeterinarianRetrieved)

	require.Equal(t, veterinarianCreated.ID, VeterinarianRetrieved.ID)
	require.Equal(t, veterinarianCreated.FirstName, VeterinarianRetrieved.FirstName)
	require.Equal(t, veterinarianCreated.LastName, VeterinarianRetrieved.LastName)
	require.Equal(t, veterinarianCreated.LicenseNumber, VeterinarianRetrieved.LicenseNumber)
	require.Equal(t, veterinarianCreated.Speciality, VeterinarianRetrieved.Speciality)
	require.Equal(t, veterinarianCreated.Photo, VeterinarianRetrieved.Photo)
	require.Equal(t, veterinarianCreated.YearsOfExperience, VeterinarianRetrieved.YearsOfExperience)
	require.WithinDuration(t, veterinarianCreated.CreatedAt.Time, VeterinarianRetrieved.CreatedAt.Time, time.Second)
}

func TestUpdateVeterinarian(t *testing.T) {
	veterinarianCreated := createDummyVeterinarian(t)

	updateValues := sqlc.UpdateVeterinarianParams{
		ID:                veterinarianCreated.ID,
		FirstName:         "Test Veterinarian FirstName UPDATE",
		LastName:          "Test Veterinarian LastName UPDATE",
		Photo:             "htpp://photo-test/vet-updated-1235.com",
		Speciality:        sqlc.VeterinarianSpecialityDentistry,
		LicenseNumber:     "0987654321",
		YearsOfExperience: 2,
		IsActive:          pgtype.Bool{Bool: true, Valid: true},
	}

	_, err := testQueries.UpdateVeterinarian(context.Background(), updateValues)
	require.NoError(t, err)

	updatedVeterinarian, err := testQueries.GetVeterinarianById(context.Background(), veterinarianCreated.ID)
	require.NoError(t, err)
	require.NotNil(t, updatedVeterinarian)

	require.Equal(t, veterinarianCreated.ID, updatedVeterinarian.ID)
	require.NotEqual(t, veterinarianCreated.FirstName, updatedVeterinarian.FirstName)
	require.NotEqual(t, veterinarianCreated.LastName, updatedVeterinarian.LastName)
	require.NotEqual(t, veterinarianCreated.LicenseNumber, updatedVeterinarian.LicenseNumber)
	require.NotEqual(t, veterinarianCreated.Speciality, updatedVeterinarian.Speciality)
	require.NotEqual(t, veterinarianCreated.Photo, updatedVeterinarian.Photo)
	require.True(t, updatedVeterinarian.UpdatedAt.Time.After(veterinarianCreated.UpdatedAt.Time))
}

func TestDeleteVeterinarian(t *testing.T) {
	veterinarianCreated := createDummyVeterinarian(t)

	err := testQueries.SoftDeleteVeterinarian(context.Background(), veterinarianCreated.ID)
	require.NoError(t, err)

	_, err = testQueries.GetVeterinarianById(context.Background(), veterinarianCreated.ID)
	require.Error(t, err)
}
