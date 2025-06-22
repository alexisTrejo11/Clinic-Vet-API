package db_test

import (
	"context"
	"testing"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

// TODO: FIX TO MAKE PGTYPES NULLABLES
func createRandomOwner(t *testing.T) sqlc.Owner {
	arg := sqlc.CreateOwnerParams{
		Name:     randomString(10),
		LastName: randomString(10),
		UserID:   pgtype.Int4{Int32: int32(randomInt(1, 1000)), Valid: true},
		Birthday: pgtype.Date{Time: time.Now().AddDate(-30, 0, 0), Valid: true},
		Genre:    pgtype.Text{String: "male", Valid: true},
	}

	owner, err := testQueries.CreateOwner(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, owner)

	require.Equal(t, arg.Name, owner.Name)
	require.Equal(t, arg.LastName, owner.LastName)
	require.Equal(t, arg.UserID, owner.UserID)

	require.NotZero(t, owner.ID)
	require.NotZero(t, owner.CreatedAt)

	return owner
}

func TestCreateOwner(t *testing.T) {
	t.Run("CreateOwnerWithRequiredFields", func(t *testing.T) {
		owner := createRandomOwner(t)
		defer deleteTestOwner(t, owner.ID)

		require.NotZero(t, owner.ID)
		require.Equal(t, owner.CreatedAt, owner.UpdatedAt)
	})

	t.Run("CreateOwnerWithAllFields", func(t *testing.T) {
		arg := sqlc.CreateOwnerParams{
			Photo:    pgtype.Text{String: "photo.jpg", Valid: true},
			Name:     "Test",
			LastName: "Owner",
			UserID:   pgtype.Int4{Int32: 123, Valid: true},
			Birthday: pgtype.Date{Time: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC), Valid: true},
			Genre:    pgtype.Text{String: "female", Valid: true},
		}

		owner, err := testQueries.CreateOwner(context.Background(), arg)
		require.NoError(t, err)
		defer deleteTestOwner(t, owner.ID)

		require.Equal(t, arg.Photo, owner.Photo)
		require.Equal(t, arg.Birthday, owner.Birthday)
		require.Equal(t, arg.Genre, owner.Genre)
	})
}

func TestDeleteOwner(t *testing.T) {
	t.Run("DeleteExistingOwner", func(t *testing.T) {
		owner := createRandomOwner(t)

		err := testQueries.DeleteOwner(context.Background(), owner.ID)
		require.NoError(t, err)

		_, err = testQueries.GetOwnerByID(context.Background(), owner.ID)
		require.Error(t, err)
	})

	t.Run("DeleteNonExistentOwner", func(t *testing.T) {
		err := testQueries.DeleteOwner(context.Background(), 999999)
		require.NoError(t, err) // DELETE no devuelve error si no encuentra el registro
	})
}

func deleteTestOwner(t *testing.T, id int32) {
	err := testQueries.DeleteOwner(context.Background(), id)
	require.NoError(t, err)
}
