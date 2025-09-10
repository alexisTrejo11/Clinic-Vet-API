package dto

import (
	"context"
	"fmt"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/application/command"
)

type RequestApptmentData struct {
	petID    int
	service  string
	datetime time.Time
	status   string
	reason   string
	notes    *string
}
type RescheduleApptRequest struct {
	AppointmentID uint
	Datetime      CustomDate
	Reason        *string
}

func (r *RescheduleApptRequest) ToCommand(ctx context.Context) (command.RescheduleApptCommand, error) {
	if r.Datetime.IsZero() {
		return command.RescheduleApptCommand{}, fmt.Errorf("date time cannot be zero")
	}

	if r.Datetime.Before(time.Now()) {
		return command.RescheduleApptCommand{}, fmt.Errorf("date time cannot be in the past")
	}

	rescheduleCommand := command.NewRescheduleApptCommand(ctx, r.AppointmentID, nil, r.Datetime.Time, r.Reason)
	return *rescheduleCommand, nil
}
