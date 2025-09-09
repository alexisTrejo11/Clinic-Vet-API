package query

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type GetApptByIDAndOwnerIDQuery struct {
	apptID  valueobject.AppointmentID
	ownerID valueobject.OwnerID
	ctx     context.Context
}

func NewGetApptByIDAndOwnerIDQuery(ctx context.Context, apptID uint, ownerID uint) *GetApptByIDAndOwnerIDQuery {
	return &GetApptByIDAndOwnerIDQuery{
		apptID:  valueobject.NewAppointmentID(apptID),
		ownerID: valueobject.NewOwnerID(ownerID),
	}
}

type GetApptByIDAndOwnerIDHandler struct {
	appointmentRepo repository.AppointmentRepository
	ownerRepo       repository.OwnerRepository
}

func NewGetApptByIDAndOwnerIDHandler(
	appointmentRepo repository.AppointmentRepository,
	ownerRepo repository.OwnerRepository,
) cqrs.QueryHandler[ApptResponse] {
	return &GetApptByIDAndOwnerIDHandler{
		appointmentRepo: appointmentRepo,
		ownerRepo:       ownerRepo,
	}
}

func (h *GetApptByIDAndOwnerIDHandler) Handle(q cqrs.Query) (ApptResponse, error) {
	query, valid := q.(GetApptByIDAndOwnerIDQuery)
	if !valid {
		return ApptResponse{}, errors.New("invalid query type")
	}

	if err := h.validateExistingOwner(query.ctx, query.ownerID); err != nil {
		return ApptResponse{}, err
	}

	appointment, err := h.appointmentRepo.GetByIDAndOwnerID(query.ctx, query.apptID, query.ownerID)
	if err != nil {
		return ApptResponse{}, err
	}

	return NewApptResponse(&appointment), nil
}

func (h *GetApptByIDAndOwnerIDHandler) validateExistingOwner(ctx context.Context, ownerID valueobject.OwnerID) error {
	exists, err := h.ownerRepo.ExistsByID(ctx, ownerID)
	if err != nil {
		return err
	}

	if !exists {
		return apperror.EntityValidationError("owner", "id", ownerID.String())
	}

	return nil
}
