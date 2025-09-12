package query

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/specification"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

type FindCustomerByIDQuery struct {
	ID  valueobject.CustomerID
	CTX context.Context
}

func NewFindCustomerByIDQuery(ctx context.Context, customerID uint) *FindCustomerByIDQuery {
	return &FindCustomerByIDQuery{
		CTX: ctx,
		ID:  valueobject.NewCustomerID(customerID),
	}
}

type FindCustomerBySpecificationQuery struct {
	querySpect specification.CustomerSpecification
	CTX        context.Context
}

func NewFindCustomerBySpecificationQuery(ctx context.Context, spec specification.CustomerSpecification) *FindCustomerBySpecificationQuery {
	return &FindCustomerBySpecificationQuery{
		CTX:        ctx,
		querySpect: spec,
	}
}
