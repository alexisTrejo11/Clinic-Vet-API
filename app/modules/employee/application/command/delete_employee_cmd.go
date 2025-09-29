package command

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	apperror "clinic-vet-api/app/shared/error/application"
)

type DeleteEmployeeCommand struct {
	id valueobject.EmployeeID
}

func NewDeleteEmployeeCommand(id uint) (DeleteEmployeeCommand, error) {
	if id == 0 {
		return DeleteEmployeeCommand{}, apperror.FieldValidationError("ID", "0", "id is required")
	}
	cmd := &DeleteEmployeeCommand{
		id: valueobject.NewEmployeeID(id),
	}

	return *cmd, nil
}

func (c DeleteEmployeeCommand) ID() valueobject.EmployeeID { return c.id }
