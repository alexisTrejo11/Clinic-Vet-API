package query

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	apperror "clinic-vet-api/app/shared/error/application"
)

type FindCustomerByIDQuery struct {
	id valueobject.CustomerID
}

func NewFindCustomerByIDQuery(customerID uint) (FindCustomerByIDQuery, error) {
	cmd := FindCustomerByIDQuery{id: valueobject.NewCustomerID(customerID)}

	if err := cmd.Validate(); err != nil {
		return FindCustomerByIDQuery{}, err
	}

	return cmd, nil
}

func (q FindCustomerByIDQuery) Validate() error {
	if q.id.IsZero() {
		return apperror.EntityNotFoundValidationError("FindCustomerByIDQuery", "ID", q.id.String())
	}
	return nil
}

func (q FindCustomerByIDQuery) ID() valueobject.CustomerID { return q.id }
