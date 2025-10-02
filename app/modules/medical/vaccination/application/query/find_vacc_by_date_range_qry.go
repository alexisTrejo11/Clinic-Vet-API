package query

import (
	apperror "clinic-vet-api/app/shared/error/application"
	p "clinic-vet-api/app/shared/page"
	"time"
)

type FindVaccinationsByDateRangeQuery struct {
	startDate  time.Time
	endDate    time.Time
	pagination p.PaginationRequest
}

func NewFindVaccinationsByDateRangeQuery(
	startDate time.Time,
	endDate time.Time,
	pagination p.PaginationRequest,
) FindVaccinationsByDateRangeQuery {
	return FindVaccinationsByDateRangeQuery{
		startDate:  startDate,
		endDate:    endDate,
		pagination: pagination,
	}
}

func (q FindVaccinationsByDateRangeQuery) StartDate() time.Time            { return q.startDate }
func (q FindVaccinationsByDateRangeQuery) EndDate() time.Time              { return q.endDate }
func (q FindVaccinationsByDateRangeQuery) Pagination() p.PaginationRequest { return q.pagination }

func (q *FindVaccinationsByDateRangeQuery) Validate() error {
	if q.startDate.IsZero() {
		return FindVaccineByDateRangeErr("startDate", "is required")
	}

	if q.endDate.IsZero() {
		return FindVaccineByDateRangeErr("endDate", "is required")
	}

	if q.startDate.Before(q.endDate) {
		return FindVaccineByDateRangeErr("range  date", "start date cannot be before end date")
	}

	return nil
}

func FindVaccineByDateRangeErr(field, issue string) error {
	return apperror.QueryDataValidationError(field, issue, "FindVaccinationsByDateRangeQuery")
}
