package dto

import (
	"context"
	"fmt"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/application/command"
)

type UpdateApptRequest struct {
	AppointmentID uint
	VetID         *uint
	Status        string
	Reason        *string
	Notes         *string
	Service       *string
}

func (r *UpdateApptRequest) ToCommand(ctx context.Context) (command.UpdateApptCommand, error) {
	var clinicService *enum.ClinicService
	if r.Service != nil {
		serviceEnum, err := enum.ParseClinicService(*r.Service)
		if err != nil {
			return command.UpdateApptCommand{}, fmt.Errorf("invalid service: %w", err)
		}
		clinicService = &serviceEnum
	}

	updateCommand := command.NewUpdateApptCommand(
		ctx,
		r.AppointmentID,
		r.VetID,
		r.Status,
		r.Reason,
		r.Notes,
		clinicService)

	return *updateCommand, nil
}
