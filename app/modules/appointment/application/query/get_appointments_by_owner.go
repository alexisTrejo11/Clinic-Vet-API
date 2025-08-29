package query

import (
	"context"
	"strconv"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	appError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/application"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type GetAppointmentsByOwnerQuery struct {
	OwnerID   int `json:"owner_id"`
	PageInput page.PageData
}

func NewGetAppointmentsByOwnerQuery(ownerID, pageNumber, pageSize int) GetAppointmentsByOwnerQuery {
	return GetAppointmentsByOwnerQuery{
		OwnerID: ownerID,
		PageInput: page.PageData{
			PageNumber: pageNumber,
			PageSize:   pageSize,
		},
	}
}

type GetAppointmentsByOwnerHandler interface {
	Handle(ctx context.Context, query GetAppointmentsByOwnerQuery) (page.Page[[]AppointmentResponse], error)
}

type getAppointmentsByOwnerHandler struct {
	appointmentRepo repository.AppointmentRepository
	ownerRepo       repository.OwnerRepository
}

func NewGetAppointmentsByOwnerHandler(
	appointmentRepo repository.AppointmentRepository,
	ownerRepo repository.OwnerRepository,
) GetAppointmentsByOwnerHandler {
	return &getAppointmentsByOwnerHandler{
		appointmentRepo: appointmentRepo,
		ownerRepo:       ownerRepo,
	}
}

func (h *getAppointmentsByOwnerHandler) Handle(ctx context.Context, query GetAppointmentsByOwnerQuery) (page.Page[[]AppointmentResponse], error) {
	if err := h.validateExistingOwner(ctx, query.OwnerID); err != nil {
		return page.Page[[]AppointmentResponse]{}, err
	}

	appointmentsPage, err := h.appointmentRepo.ListByOwnerID(ctx, query.OwnerID, query.PageInput)
	if err != nil {
		return page.Page[[]AppointmentResponse]{}, err
	}

	return page.NewPage(
		mapAppointmentsToResponses(appointmentsPage.Data),
		appointmentsPage.Metadata,
	), nil
}

func (h *getAppointmentsByOwnerHandler) validateExistingOwner(ctx context.Context, ownerID int) error {
	if exists, err := h.ownerRepo.ExistsByID(ctx, ownerID); err != nil {
		return err
	} else if !exists {
		return appError.NewEntityNotFoundError("owner", strconv.Itoa(ownerID))
	} else {
		return nil
	}
}
