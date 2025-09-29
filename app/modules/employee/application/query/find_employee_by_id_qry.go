package query

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	apperror "clinic-vet-api/app/shared/error/application"
)

type FindEmployeeByIDQuery struct {
	id valueobject.EmployeeID
}

func NewFindEmployeeByIDQuery(id uint) (FindEmployeeByIDQuery, error) {
	if id == 0 {
		return FindEmployeeByIDQuery{}, apperror.FieldValidationError("id", "0", "id is required for FindEmployeeByIDQuery")
	}

	cmd := &FindEmployeeByIDQuery{id: valueobject.NewEmployeeID(id)}
	return *cmd, nil
}

func (q FindEmployeeByIDQuery) ID() valueobject.EmployeeID {
	return q.id
}
