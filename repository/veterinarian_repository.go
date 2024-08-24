package repository

import (
	"context"

	"example.com/at/backend/api-vet/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type VeterinarianRepository interface {
	CreateVeterinarian(ctx context.Context, arg sqlc.CreateVeterinarianParams) (sqlc.Veterinarian, error)
	GetVeterinarianByID(ctx context.Context, id int32) (sqlc.Veterinarian, error)
	UpdateVeterinarian(ctx context.Context, arg sqlc.UpdateVeterinarianParams) error
	AddUserIdToExisitngVet(veterinarianId, userId int32)
	DeleteVeterinarian(ctx context.Context, id int32) error
	ValidateExistingVeterinarian(ctx context.Context, VeterinarianId int32) bool
	GetVeterinarianByUserID(userId int32) (*sqlc.Veterinarian, error)
}

type VeterinarianRepositoryImpl struct {
	queries *sqlc.Queries
}

func NewVeterinarianRepository(queries *sqlc.Queries) VeterinarianRepository {
	return &VeterinarianRepositoryImpl{queries: queries}
}

func (r *VeterinarianRepositoryImpl) CreateVeterinarian(ctx context.Context, arg sqlc.CreateVeterinarianParams) (sqlc.Veterinarian, error) {
	veterinarian, err := r.queries.CreateVeterinarian(ctx, arg)
	if err != nil {
		return sqlc.Veterinarian{}, err
	}
	return veterinarian, nil
}

func (r *VeterinarianRepositoryImpl) GetVeterinarianByID(ctx context.Context, VeterinarianId int32) (sqlc.Veterinarian, error) {
	veterinarian, err := r.queries.GetVeterinarianByID(ctx, VeterinarianId)
	if err != nil {
		return sqlc.Veterinarian{}, err
	}
	return veterinarian, nil
}

func (r *VeterinarianRepositoryImpl) GetVeterinarianByUserID(userId int32) (*sqlc.Veterinarian, error) {
	userIdPgType := pgtype.Int4{Int32: userId, Valid: true}
	veterinarian, err := r.queries.GetVeterinarianByUserID(context.Background(), userIdPgType)
	if err != nil {
		return nil, err
	}
	return &veterinarian, nil
}

func (r *VeterinarianRepositoryImpl) UpdateVeterinarian(ctx context.Context, arg sqlc.UpdateVeterinarianParams) error {
	err := r.queries.UpdateVeterinarian(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}

func (r *VeterinarianRepositoryImpl) DeleteVeterinarian(ctx context.Context, VeterinarianId int32) error {
	err := r.queries.DeleteVeterinarian(ctx, VeterinarianId)
	if err != nil {
		return err
	}
	return nil
}

func (r *VeterinarianRepositoryImpl) GetAppointmentsByVeterinarian(ctx context.Context, VeterinarianID int32) ([]sqlc.Appointment, error) {
	return nil, nil
}

func (r *VeterinarianRepositoryImpl) ValidateExistingVeterinarian(ctx context.Context, VeterinarianId int32) bool {
	_, err := r.queries.GetVeterinarianByID(ctx, VeterinarianId)
	return err == nil
}

func (r *VeterinarianRepositoryImpl) AddUserIdToExisitngVet(veterinarianId, userId int32) {
	args := sqlc.UpdateVeterinarianUserIDParams{
		ID:     veterinarianId,
		UserID: pgtype.Int4{Int32: userId, Valid: true},
	}

	r.queries.UpdateVeterinarianUserID(context.Background(), args)
}
