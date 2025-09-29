package bus

import (
	cmd "clinic-vet-api/app/modules/appointment/application/command"
	"clinic-vet-api/app/modules/appointment/application/handler"
	icqrs "clinic-vet-api/app/shared/cqrs"
	"context"
)

type ApptCmdBus struct {
	apptHandler handler.ApptCommandHandler
}

func NewApptCmdBus(apptHandler handler.ApptCommandHandler) *ApptCmdBus {
	return &ApptCmdBus{apptHandler: apptHandler}
}

func (b *ApptCmdBus) RequestAppointmentByCustomer(ctx context.Context, cmd cmd.RequestApptByCustomerCommand) icqrs.CommandResult {
	return b.apptHandler.HandleRequestByCustomer(ctx, cmd)
}

func (b *ApptCmdBus) CreateAppointment(ctx context.Context, cmd cmd.CreateApptCommand) icqrs.CommandResult {
	return b.apptHandler.HandleCreate(ctx, cmd)
}

func (b *ApptCmdBus) DeleteAppointment(ctx context.Context, cmd cmd.DeleteApptCommand) icqrs.CommandResult {
	return b.apptHandler.HandleDelete(ctx, cmd)
}

func (b *ApptCmdBus) RescheduleAppointment(ctx context.Context, cmd cmd.RescheduleApptCommand) icqrs.CommandResult {
	return b.apptHandler.HandleReschedule(ctx, cmd)
}

func (b *ApptCmdBus) CancelAppointment(ctx context.Context, cmd cmd.CancelApptCommand) icqrs.CommandResult {
	return b.apptHandler.HandleCancel(ctx, cmd)
}

func (b *ApptCmdBus) ConfirmAppointment(ctx context.Context, cmd cmd.ConfirmApptCommand) icqrs.CommandResult {
	return b.apptHandler.HandleConfirm(ctx, cmd)
}

func (b *ApptCmdBus) MarkAppointmentAsNotAttend(ctx context.Context, cmd cmd.NotAttendApptCommand) icqrs.CommandResult {
	return b.apptHandler.HandleMarkAsNotAttend(ctx, cmd)
}

func (b *ApptCmdBus) CompleteAppointment(ctx context.Context, cmd cmd.CompleteApptCommand) icqrs.CommandResult {
	return b.apptHandler.HandleComplete(ctx, cmd)
}

func (b *ApptCmdBus) UpdateAppointment(ctx context.Context, cmd cmd.UpdateApptCommand) icqrs.CommandResult {
	return b.apptHandler.HandleUpdate(ctx, cmd)
}
