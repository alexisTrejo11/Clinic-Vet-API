package db_test

import (
	"context"
	"testing"
	"time"

	"clinic-vet-api/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

// TODO: FIX TO MAKE PGTYPES NULLABLES
func createRandomcustomer(t *testing.T) sqlc.CreatecustomerRow {
	arg := sqlc.CreatecustomerParams{
		FirstName:   randomString(10),
		LastName:    randomString(10),
		UserID:      pgtype.Int4{Int32: int32(randomInt(1, 1000)), Valid: true},
		DateOfBirth: pgtype.Date{Time: time.Now().AddDate(-30, 0, 0), Valid: true},
		Gender:      sqlc.PersonGenderFemale,
	}

	customer, err := testQueries.Createcustomer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customer)

	require.Equal(t, arg.FirstName, customer.FirstName)
	require.Equal(t, arg.LastName, customer.LastName)
	require.Equal(t, arg.UserID, customer.UserID)

	require.NotZero(t, customer.ID)
	require.NotZero(t, customer.CreatedAt)

	return customer
}

func TestCreatecustomer(t *testing.T) {
	t.Run("CreatecustomerWithRequiredFields", func(t *testing.T) {
		customer := createRandomcustomer(t)
		defer deleteTestcustomer(t, customer.ID)

		require.NotZero(t, customer.ID)
		require.Equal(t, customer.CreatedAt, customer.UpdatedAt)
	})

	t.Run("CreatecustomerWithAllFields", func(t *testing.T) {
		arg := sqlc.CreatecustomerParams{
			Photo:       "test-foto.com",
			FirstName:   "Test",
			LastName:    "customer",
			UserID:      pgtype.Int4{Int32: 123, Valid: true},
			DateOfBirth: pgtype.Date{Time: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
			Gender:      sqlc.PersonGenderMale,
		}

		customer, err := testQueries.Createcustomer(context.Background(), arg)
		require.NoError(t, err)
		defer deleteTestcustomer(t, customer.ID)

		require.Equal(t, arg.Photo, customer.Photo)
		require.Equal(t, arg.DateOfBirth, customer.DateOfBirth)
		require.Equal(t, arg.Gender, customer.Gender)
	})
}

func TestDeletecustomer(t *testing.T) {
	t.Run("DeleteExistingcustomer", func(t *testing.T) {
		customer := createRandomcustomer(t)

		err := testQueries.Deletecustomer(context.Background(), customer.ID)
		require.NoError(t, err)

		_, err = testQueries.GetcustomerByID(context.Background(), customer.ID)
		require.Error(t, err)
	})

	t.Run("DeleteNonExistentcustomer", func(t *testing.T) {
		err := testQueries.Deletecustomer(context.Background(), 999999)
		require.NoError(t, err) // DELETE no devuelve error si no encuentra el registro
	})
}

func deleteTestcustomer(t *testing.T, id int32) {
	err := testQueries.Deletecustomer(context.Background(), id)
	require.NoError(t, err)
}
