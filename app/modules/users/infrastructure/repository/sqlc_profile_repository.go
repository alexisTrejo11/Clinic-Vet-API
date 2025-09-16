package repositoryimpl

import (
	"context"

	"clinic-vet-api/app/core/domain/entity/user/profile"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

type SQLCProfileRepository struct {
	queries *sqlc.Queries
}

func NewSQLCProfileRepository(queries *sqlc.Queries) repository.ProfileRepository {
	return &SQLCProfileRepository{
		queries: queries,
	}
}

func (r *SQLCProfileRepository) GetByUserID(ctx context.Context, userID valueobject.UserID) (profile.Profile, error) {
	sqlRow, err := r.queries.GetUserProfile(ctx, pgtype.Int4{Int32: int32(userID.Value()), Valid: true})
	if err != nil {
		return profile.Profile{}, err
	}

	return profile.Profile{
		UserID:   userID,
		PhotoURL: sqlRow.ProfilePic.String,
		Bio:      sqlRow.Bio.String,
		JoinedAt: sqlRow.CreatedAt.Time,
	}, nil
}

func (r *SQLCProfileRepository) Create(ctx context.Context, profile *profile.Profile) error {
	_, err := r.queries.CreateProfile(ctx, sqlc.CreateProfileParams{
		UserID:     pgtype.Int4{Int32: int32(profile.UserID.Value()), Valid: true},
		Bio:        pgtype.Text{String: profile.Bio, Valid: true},
		ProfilePic: pgtype.Text{String: profile.PhotoURL, Valid: profile.PhotoURL != ""},
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLCProfileRepository) Update(ctx context.Context, profile *profile.Profile) error {
	_, err := r.queries.UpdateUserProfile(ctx, sqlc.UpdateUserProfileParams{
		UserID:     pgtype.Int4{Int32: int32(profile.UserID.Value()), Valid: !profile.UserID.IsZero()},
		Bio:        pgtype.Text{String: profile.Bio, Valid: true},
		ProfilePic: pgtype.Text{String: profile.PhotoURL, Valid: profile.PhotoURL != ""},
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLCProfileRepository) DeleteByUserID(ctx context.Context, userID valueobject.UserID) error {
	err := r.queries.DeleteUserProfile(ctx, pgtype.Int4{Int32: int32(userID.Value()), Valid: true})
	if err != nil {
		return err
	}

	return nil
}
