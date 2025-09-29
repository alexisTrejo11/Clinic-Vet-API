package command

import (
	apperror "clinic-vet-api/app/shared/error/application"
)

func updateApptCmdErr(field, issue string) error {
	return apperror.CommandDataValidationError(field, issue, "UpdateApptCommand")
}

func deleteApptCmdErr(field, issue string) error {
	return apperror.CommandDataValidationError(field, issue, "DeleteApptCommand")
}

func completeApptCmdErr(field, issue string) error {
	return apperror.CommandDataValidationError(field, issue, "CompleteApptCommand")
}

func createApptCommand(field, issue string) error {
	return apperror.CommandDataValidationError(field, issue, "CreateApptCommand")
}

func confirmApptCmdErr(field, issue string) error {
	return apperror.CommandDataValidationError(field, issue, "ConfirmApptCommand")
}

func notAttendApptCmdErr(field, issue string) error {
	return apperror.CommandDataValidationError(field, issue, "NotAttendApptCommand")
}

func requestScheduleCmdErr(field, issue string) error {
	return apperror.CommandDataValidationError(field, issue, "RequestScheduleCmd")
}
