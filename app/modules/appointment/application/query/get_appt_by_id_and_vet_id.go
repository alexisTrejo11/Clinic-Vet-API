package query

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type GetApptByIDAndVetIDQuery struct {
	apptID valueobject.AppointmentID
	vetID  valueobject.VetID
	ctx    context.Context
}

func NewGetApptByIDAndVetIDQuery(ctx context.Context, apptID uint, vetID uint) *GetApptByIDAndVetIDQuery {
	return &GetApptByIDAndVetIDQuery{
		apptID: valueobject.NewAppointmentID(apptID),
		vetID:  valueobject.NewVetID(vetID),
	}
}

type GetApptByIDAndVetIDHandler struct {
	appointmentRepo repository.AppointmentRepository
	vetRepo         repository.VetRepository
}

func NewGetApptByIDAndVetIDHandler(
	appointmentRepo repository.AppointmentRepository,
	vetRepo repository.VetRepository,
) cqrs.QueryHandler[ApptResponse] {
	return &GetApptByIDAndVetIDHandler{
		appointmentRepo: appointmentRepo,
		vetRepo:         vetRepo,
	}
}

func (h *GetApptByIDAndVetIDHandler) Handle(q cqrs.Query) (ApptResponse, error) {
	query, valid := q.(GetApptByIDAndVetIDQuery)
	if !valid {
		return ApptResponse{}, errors.New("invalid query type")
	}

	if err := h.validateExistingVet(query.ctx, query.vetID); err != nil {
		return ApptResponse{}, err
	}

	appointment, err := h.appointmentRepo.GetByIDAndVetID(query.ctx, query.apptID, query.vetID)
	if err != nil {
		return ApptResponse{}, err
	}

	return NewApptResponse(&appointment), nil
}

func (h *GetApptByIDAndVetIDHandler) validateExistingVet(ctx context.Context, vetID valueobject.VetID) error {
	exists, err := h.vetRepo.Exists(ctx, vetID)
	if err != nil {
		return err
	}

	if !exists {
		return apperror.EntityValidationError("owner", "id", vetID.String())
	}

	return nil
}
