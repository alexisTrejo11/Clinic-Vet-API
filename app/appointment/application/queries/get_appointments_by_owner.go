package appointmentQuery

import (
	"context"
	"strconv"

	appointmentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	ownerRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/repository"
	appError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/application"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type GetAppointmentsByOwnerQuery struct {
	OwnerId   int `json:"owner_id"`
	PageInput page.PageData
}

func NewGetAppointmentsByOwnerQuery(ownerId, pageNumber, pageSize int) GetAppointmentsByOwnerQuery {
	return GetAppointmentsByOwnerQuery{
		OwnerId: ownerId,
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
	appointmentRepo appointmentDomain.AppointmentRepository
	ownerRepo       ownerRepository.OwnerRepository
}

func NewGetAppointmentsByOwnerHandler(
	appointmentRepo appointmentDomain.AppointmentRepository,
	ownerRepo ownerRepository.OwnerRepository,
) GetAppointmentsByOwnerHandler {
	return &getAppointmentsByOwnerHandler{
		appointmentRepo: appointmentRepo,
		ownerRepo:       ownerRepo,
	}
}

func (h *getAppointmentsByOwnerHandler) Handle(ctx context.Context, query GetAppointmentsByOwnerQuery) (page.Page[[]AppointmentResponse], error) {
	if err := h.validateExistingOwner(ctx, query.OwnerId); err != nil {
		return page.Page[[]AppointmentResponse]{}, err
	}

	appointmentsPage, err := h.appointmentRepo.ListByOwnerId(ctx, query.OwnerId, query.PageInput)
	if err != nil {
		return page.Page[[]AppointmentResponse]{}, err
	}

	return *page.NewPage(
		mapAppointmentsToResponses(appointmentsPage.Data),
		appointmentsPage.Metadata,
	), nil
}

func (h *getAppointmentsByOwnerHandler) validateExistingOwner(ctx context.Context, ownerId int) error {
	if exists, err := h.ownerRepo.ExistsByID(ctx, ownerId); err != nil {
		return err
	} else if !exists {
		return appError.NewEntityNotFoundError("owner", strconv.Itoa(ownerId))
	} else {
		return nil
	}
}
