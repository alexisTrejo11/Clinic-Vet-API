package query

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListApptByOwnerQuery struct {
	ownerID   valueobject.OwnerID
	ctx       context.Context
	pageInput page.PageInput
}

func NewListApptByOwnerQuery(ctx context.Context, id uint, pageInput page.PageInput) *ListApptByOwnerQuery {
	return &ListApptByOwnerQuery{
		ownerID:   valueobject.NewOwnerID(id),
		pageInput: pageInput,
		ctx:       ctx,
	}
}

type ListApptByOwnerHandler struct {
	appointmentRepo repository.AppointmentRepository
	ownerRepo       repository.OwnerRepository
}

func NewListApptsByOwnerHandler(
	appointmentRepo repository.AppointmentRepository,
	ownerRepo repository.OwnerRepository,
) cqrs.QueryHandler[(page.Page[[]ApptResponse])] {
	return &ListApptByOwnerHandler{
		appointmentRepo: appointmentRepo,
		ownerRepo:       ownerRepo,
	}
}

func (h *ListApptByOwnerHandler) Handle(q cqrs.Query) (page.Page[[]ApptResponse], error) {
	query, valid := q.(ListApptByOwnerQuery)
	if !valid {
		return page.Page[[]ApptResponse]{}, errors.New("invalid query type")
	}

	if err := h.validateExistingOwner(query.ctx, query.ownerID); err != nil {
		return page.Page[[]ApptResponse]{}, err
	}

	appointmentsPage, err := h.appointmentRepo.ListByOwnerID(query.ctx, query.ownerID, query.pageInput)
	if err != nil {
		return page.Page[[]ApptResponse]{}, err
	}

	return page.NewPage(
		mapApptsToResponse(appointmentsPage.Data),
		appointmentsPage.Metadata,
	), nil
}

func (h *ListApptByOwnerHandler) validateExistingOwner(ctx context.Context, ownerID valueobject.OwnerID) error {
	exists, err := h.ownerRepo.ExistsByID(ctx, ownerID)
	if err != nil {
		return err
	}

	if !exists {
		return apperror.EntityValidationError("owner", "id", ownerID.String())
	}

	return nil
}
