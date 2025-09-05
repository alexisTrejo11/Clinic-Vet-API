// Package dto contains all the data structures for HTTP requests
package dto

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/application/command"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/application/query"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListAppointmentsByDateRangeRequest struct {
	StartDate CustomDate `json:"start_date"  form:"start_date" binding:"required"`
	EndDate   CustomDate `json:"end_date"  form:"end_date" binding:"required"`
	page.PageInput
}

func (r *ListAppointmentsByDateRangeRequest) ToQuery() (query.ListAppointmentsByDateRangeQuery, error) {
	qry, err := query.NewListAppointmentsByDateRangeQuery(r.StartDate.Time, r.EndDate.Time, r.PageInput)
	if err != nil {
		return query.ListAppointmentsByDateRangeQuery{}, err
	}

	return qry, nil
}

type CustomDate struct {
	time.Time
}

func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return fmt.Errorf("invalid date format, expected YYYY-MM-DD: %w", err)
	}
	cd.Time = t
	return nil
}

func (cd *CustomDate) UnmarshalText(text []byte) error {
	t, err := time.Parse("2006-01-02", string(text))
	if err != nil {
		return fmt.Errorf("invalid date format, expected YYYY-MM-DD: %w", err)
	}
	cd.Time = t
	return nil
}

type UpdateAppointmentRequest struct {
	AppointmentID int
	VetID         *int
	Status        string
	Reason        *string
	Notes         *string
	Service       *string
}

func (r *UpdateAppointmentRequest) ToCommand(ctx context.Context) (command.UpdateAppointmentCommand, error) {
	updateCommand, err := command.NewUpdateAppointmentCommand(
		ctx,
		r.AppointmentID,
		r.VetID,
		r.Status,
		r.Reason,
		r.Notes,
		r.Service)
	if err != nil {
		return command.UpdateAppointmentCommand{}, err
	}

	return *updateCommand, nil
}

type RescheduleAppointmentRequest struct {
	AppointmentID int
	Datetime      CustomDate
	Reason        *string
}

func (r *RescheduleAppointmentRequest) ToCommand(ctx context.Context) (command.RescheduleAppointmentCommand, error) {
	rescheduleCommand, err := command.NewRescheduleAppointmentCommand(ctx, r.AppointmentID, r.Datetime.Time, r.Reason)
	if err != nil {
		return command.RescheduleAppointmentCommand{}, err
	}

	return rescheduleCommand, nil
}

type CompleteAppointmentRequest struct {
	Notes *string
}

func (r *CompleteAppointmentRequest) ToCommand(ctx context.Context, appointmentID int) (command.CompleteAppointmentCommand, error) {
	cmd, err := command.NewCompleteAppointmenCommand(ctx, appointmentID, r.Notes)
	if err != nil {
		return command.CompleteAppointmentCommand{}, err
	}

	return *cmd, nil
}
