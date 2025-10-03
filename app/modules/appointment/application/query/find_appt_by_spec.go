package query

import (
	"clinic-vet-api/app/modules/core/domain/specification"
)

type FindApptsBySpecQuery struct {
	spec specification.ApptSearchSpecification
}

func NewFindApptsBySpecQuery(spec specification.ApptSearchSpecification) FindApptsBySpecQuery {
	return FindApptsBySpecQuery{spec: spec}
}

func (q FindApptsBySpecQuery) Spec() specification.ApptSearchSpecification {
	return q.spec
}
