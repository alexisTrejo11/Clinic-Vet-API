package query

import (
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

type FindCustomerByIDQuery struct {
	ID valueobject.CustomerID
}

func NewFindCustomerByIDQuery(customerID uint) *FindCustomerByIDQuery {
	return &FindCustomerByIDQuery{
		ID: valueobject.NewCustomerID(customerID),
	}
}

type FindCustomerBySpecificationQuery struct {
	querySpect specification.CustomerSpecification
}

func NewFindCustomerBySpecificationQuery(spec specification.CustomerSpecification) *FindCustomerBySpecificationQuery {
	return &FindCustomerBySpecificationQuery{
		querySpect: spec,
	}
}
