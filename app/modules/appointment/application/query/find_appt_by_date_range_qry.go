package query

import (
	"clinic-vet-api/app/modules/core/domain/specification"
	apperror "clinic-vet-api/app/shared/error/application"
	"clinic-vet-api/app/shared/page"
	"time"
)

type FindApptsByDateRangeQuery struct {
	startDate  time.Time
	endDate    time.Time
	pagination specification.Pagination
}

func NewFindApptsByDateRangeQuery(startDate, endDate time.Time, pagInput page.PaginationRequest) (FindApptsByDateRangeQuery, error) {
	qry := &FindApptsByDateRangeQuery{
		startDate:  startDate,
		endDate:    endDate,
		pagination: pagInput.ToSpecPagination(),
	}

	if startDate.IsZero() {
		return FindApptsByDateRangeQuery{}, apperror.FieldValidationError("startDate", "zero", "startDate can't be zero")
	}

	if endDate.IsZero() {
		return FindApptsByDateRangeQuery{}, apperror.FieldValidationError("endDate", "zero", "endDate can't be zero")
	}

	if startDate.Before(endDate) {
		return FindApptsByDateRangeQuery{}, apperror.FieldValidationError("date-range", "", "startDate can't be before endDate")
	}

	return *qry, nil
}

func (q FindApptsByDateRangeQuery) StartDate() time.Time                 { return q.startDate }
func (q FindApptsByDateRangeQuery) EndDate() time.Time                   { return q.endDate }
func (q FindApptsByDateRangeQuery) Pagination() specification.Pagination { return q.pagination }
