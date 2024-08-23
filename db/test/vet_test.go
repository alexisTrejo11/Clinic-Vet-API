package db_test

import (
	"context"
	"testing"
	"time"

	"example.com/at/backend/api-vet/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createDummyVeterinarian(t *testing.T) sqlc.Veterinarian {
	arg := sqlc.CreateVeterinarianParams{
		Name:      "Test Veterinarian",
		Photo:     pgtype.Text{String: "test.jpg", Valid: true},
		Specialty: pgtype.Text{String: "test.jpg", Valid: true},
		UserID:    pgtype.Int4{Int32: 1, Valid: true},
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
	Veterinarian1 := createDummyVeterinarian(t)
	Veterinarian2, err := testQueries.GetVeterinarianByID(context.Background(), Veterinarian1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, Veterinarian2)

	require.Equal(t, Veterinarian1.Name, Veterinarian2.Name)
	require.Equal(t, Veterinarian1.UserID, Veterinarian2.UserID)
	require.WithinDuration(t, Veterinarian1.CreatedAt.Time, Veterinarian1.CreatedAt.Time, time.Second)
}

func TestUpdateVeterinarian(t *testing.T) {
	VeterinarianCreated := createDummyVeterinarian(t)

	updateValues := sqlc.UpdateVeterinarianParams{
		ID:        VeterinarianCreated.ID,
		Name:      "Test Veterinarian2",
		Photo:     pgtype.Text{String: "test1.jpg", Valid: true},
		Specialty: pgtype.Text{String: "Zootecnian", Valid: true},
		UserID:    pgtype.Int4{Int32: 2, Valid: true},
	}

	err := testQueries.UpdateVeterinarian(context.Background(), updateValues)
	require.NoError(t, err)

	updatedVeterinarian, err := testQueries.GetVeterinarianByID(context.Background(), VeterinarianCreated.ID)
	require.NoError(t, err)
	require.NotNil(t, updatedVeterinarian)

	require.Equal(t, updateValues.Name, updatedVeterinarian.Name)
	require.Equal(t, updateValues.Specialty, updatedVeterinarian.Specialty)
	require.Equal(t, updateValues.Photo, updatedVeterinarian.Photo)
	require.Equal(t, VeterinarianCreated.UserID, updatedVeterinarian.UserID)
	require.Equal(t, VeterinarianCreated.CreatedAt, updatedVeterinarian.CreatedAt)

	require.True(t, updatedVeterinarian.UpdatedAt.Time.After(VeterinarianCreated.UpdatedAt.Time))
}

func TestDeleteVeterinarian(t *testing.T) {
	VeterinarianCreated := createDummyVeterinarian(t)

	err := testQueries.DeleteVeterinarian(context.Background(), VeterinarianCreated.ID)
	require.NoError(t, err)

	_, err = testQueries.GetVeterinarianByID(context.Background(), VeterinarianCreated.ID)
	require.Error(t, err)

}
