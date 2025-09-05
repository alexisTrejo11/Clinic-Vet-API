package query

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListAppointmentsByOwnerQuery struct {
	ownerID   valueobject.OwnerID
	ctx       context.Context
	pageInput page.PageInput
}

func NewListAppointmentsByOwnerQuery(id int, pageInput page.PageInput) (ListAppointmentsByOwnerQuery, error) {
	ownerID, err := valueobject.NewOwnerID(id)
	if err != nil {
		return ListAppointmentsByOwnerQuery{}, err
	}

	qry := &ListAppointmentsByOwnerQuery{
		ownerID:   ownerID,
		pageInput: pageInput,
	}

	return *qry, nil
}

type ListAppointmentsByOwnerHandler struct {
	appointmentRepo repository.AppointmentRepository
	ownerRepo       repository.OwnerRepository
}

func NewListAppointmentsByOwnerHandler(appointmentRepo repository.AppointmentRepository, ownerRepo repository.OwnerRepository) cqrs.QueryHandler[(page.Page[[]AppointmentResponse])] {
	return &ListAppointmentsByOwnerHandler{
		appointmentRepo: appointmentRepo,
		ownerRepo:       ownerRepo,
	}
}

func (h *ListAppointmentsByOwnerHandler) Handle(q cqrs.Query) (page.Page[[]AppointmentResponse], error) {
	query := q.(ListAppointmentsByOwnerQuery)

	if err := h.validateExistingOwner(query.ctx, query.ownerID); err != nil {
		return page.Page[[]AppointmentResponse]{}, err
	}

	appointmentsPage, err := h.appointmentRepo.ListByOwnerID(query.ctx, query.ownerID, query.pageInput)
	if err != nil {
		return page.Page[[]AppointmentResponse]{}, err
	}

	return page.NewPage(
		mapAppointmentsToResponses(appointmentsPage.Data),
		appointmentsPage.Metadata,
	), nil
}

func (h *ListAppointmentsByOwnerHandler) validateExistingOwner(ctx context.Context, ownerID valueobject.OwnerID) error {
	exists, err := h.ownerRepo.ExistsByID(ctx, ownerID)
	if err != nil {
		return err
	}

	if !exists {
		return apperror.EntityValidationError("owner", "id", ownerID.String())
	}

	return nil
}
