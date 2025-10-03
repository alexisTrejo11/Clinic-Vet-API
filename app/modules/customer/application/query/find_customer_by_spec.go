package query

import (
	"clinic-vet-api/app/modules/core/domain/specification"
)

type FindCustomerBySpecificationQuery struct {
	querySpect specification.CustomerSpecification
}

func NewFindCustomerBySpecificationQuery(spec specification.CustomerSpecification) (FindCustomerBySpecificationQuery, error) {
	cmd := FindCustomerBySpecificationQuery{querySpect: spec}
	return cmd, nil
}

func (q FindCustomerBySpecificationQuery) QuerySpecification() specification.CustomerSpecification {
	return q.querySpect
}
