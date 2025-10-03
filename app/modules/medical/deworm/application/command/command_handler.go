package command

import (
	"context"

	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
)

type DewormCommandHandler struct {
	dewormRepo   repository.DewormRepository
	employeeRepo repository.EmployeeRepository
	petRepo      repository.PetRepository
}

func NewDewormCommandHandler(
	dewormRepo repository.DewormRepository,
	employeeRepo repository.EmployeeRepository,
	petRepo repository.PetRepository,
) *DewormCommandHandler {
	return &DewormCommandHandler{
		dewormRepo:   dewormRepo,
		employeeRepo: employeeRepo,
		petRepo:      petRepo,
	}
}

// Helper Methods
func (h *DewormCommandHandler) validateDewormExistence(ctx context.Context, id valueobject.DewormID, petID *valueobject.PetID) error {
	if petID != nil {
		if _, err := h.dewormRepo.FindByIDAndPetID(ctx, id, *petID); err != nil {
			return err
		}
	}
	_, err := h.dewormRepo.FindByID(ctx, id)
	return err
}
