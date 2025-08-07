package sqlcUserRepo

import (
	"context"

	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type SQLCProfileRepository struct {
	queries *sqlc.Queries
}

func NewSQLCProfileRepository(queries *sqlc.Queries) *SQLCProfileRepository {
	return &SQLCProfileRepository{
		queries: queries,
	}
}

func (r *SQLCProfileRepository) GetByUserId(ctx context.Context, userId int) (*user.Profile, error) {
	sqlRow, err := r.queries.GetUserProfile(ctx, pgtype.Int4{Int32: int32(userId), Valid: true})
	if err != nil {
		return nil, err
	}

	return &user.Profile{
		UserId: int(sqlRow.UserID.Int32),
		OwnerId: func() *int {
			if sqlRow.OwnerID.Valid {
				v := int(sqlRow.OwnerID.Int32)
				return &v
			}
			return nil
		}(),
		VeterinarianId: func() *int {
			if sqlRow.VeterinarianID.Valid {
				v := int(sqlRow.VeterinarianID.Int32)
				return &v
			}
			return nil
		}(),
		Name:     user.PersonName{FirstName: "Jhon", LastName: "Doe"},
		PhotoURL: sqlRow.ProfilePic.String,
		Bio:      sqlRow.Bio.String,
		Location: "",
		Gender:   "MALE",
		JoinedAt: sqlRow.CreatedAt.Time,
	}, nil
}

func (r *SQLCProfileRepository) Create(ctx context.Context, profile *user.Profile) error {
	_, err := r.queries.CreateProfile(ctx, sqlc.CreateProfileParams{
		UserID:         pgtype.Int4{Int32: int32(profile.UserId), Valid: true},
		Bio:            pgtype.Text{String: profile.Bio, Valid: true},
		ProfilePic:     pgtype.Text{String: profile.PhotoURL, Valid: profile.PhotoURL != ""},
		VeterinarianID: pgtype.Int4{Int32: int32(*profile.VeterinarianId), Valid: profile.VeterinarianId != nil},
		OwnerID:        pgtype.Int4{Int32: int32(*profile.OwnerId), Valid: profile.OwnerId != nil},
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLCProfileRepository) Update(ctx context.Context, profile *user.Profile) error {
	_, err := r.queries.UpdateUserProfile(ctx, sqlc.UpdateUserProfileParams{
		UserID:         pgtype.Int4{Int32: int32(profile.UserId), Valid: true},
		Bio:            pgtype.Text{String: profile.Bio, Valid: true},
		ProfilePic:     pgtype.Text{String: profile.PhotoURL, Valid: profile.PhotoURL != ""},
		VeterinarianID: pgtype.Int4{Int32: int32(*profile.VeterinarianId), Valid: profile.VeterinarianId != nil},
		OwnerID:        pgtype.Int4{Int32: int32(*profile.OwnerId), Valid: profile.OwnerId != nil},
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLCProfileRepository) Delete(ctx context.Context, id int) error {
	err := r.queries.DeleteUserProfile(ctx, pgtype.Int4{Int32: int32(id), Valid: true})
	if err != nil {
		return err
	}

	return nil
}
