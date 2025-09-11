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

type FindApptsByCustomerIDQuery struct {
	ownerID   valueobject.CustomerID
	ctx       context.Context
	pageInput page.PageInput
}

func NewFindApptsByCustomerIDQuery(ctx context.Context, id uint, pageInput page.PageInput) *FindApptsByCustomerIDQuery {
	return &FindApptsByCustomerIDQuery{
		ownerID:   valueobject.NewCustomerID(id),
		pageInput: pageInput,
		ctx:       ctx,
	}
}

type FindApptsByCustomerIDHandler struct {
	appointmentRepo repository.AppointmentRepository
	ownerRepo       repository.CustomerRepository
}

func NewFindApptsByCustomerIDHandler(
	appointmentRepo repository.AppointmentRepository,
	ownerRepo repository.CustomerRepository,
) cqrs.QueryHandler[(page.Page[ApptResponse])] {
	return &FindApptsByCustomerIDHandler{
		appointmentRepo: appointmentRepo,
		ownerRepo:       ownerRepo,
	}
}

func (h *FindApptsByCustomerIDHandler) Handle(q cqrs.Query) (page.Page[ApptResponse], error) {
	query, valid := q.(FindApptsByCustomerIDQuery)
	if !valid {
		return page.Page[ApptResponse]{}, errors.New("invalid query type")
	}

	if err := h.validateExistingCustomer(query.ctx, query.ownerID); err != nil {
		return page.Page[ApptResponse]{}, err
	}

	appointmentsPage, err := h.appointmentRepo.FindByCustomerID(query.ctx, query.ownerID, query.pageInput)
	if err != nil {
		return page.Page[ApptResponse]{}, err
	}

	return page.NewPage(
		mapApptsToResponse(appointmentsPage.Items),
		appointmentsPage.Metadata,
	), nil
}

func (h *FindApptsByCustomerIDHandler) validateExistingCustomer(ctx context.Context, ownerID valueobject.CustomerID) error {
	exists, err := h.ownerRepo.ExistsByID(ctx, ownerID)
	if err != nil {
		return err
	}

	if !exists {
		return apperror.EntityValidationError("owner", "id", ownerID.String())
	}

	return nil
}
