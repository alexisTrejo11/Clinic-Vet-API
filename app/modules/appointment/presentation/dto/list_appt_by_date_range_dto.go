// Package dto contains all the data structures for HTTP requests
package dto

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/application/query"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListApptByDateRangeRequest struct {
	StartDate CustomDate `json:"start_date"  form:"start_date" binding:"required"`
	EndDate   CustomDate `json:"end_date"  form:"end_date" binding:"required"`
	page.PageInput
}

func (r *ListApptByDateRangeRequest) ToQuery() (query.ListApptByDateRangeQuery, error) {
	qry, err := query.NewListApptByDateRangeQuery(r.StartDate.Time, r.EndDate.Time, r.PageInput)
	if err != nil {
		return query.ListApptByDateRangeQuery{}, err
	}

	return qry, nil
}

type CustomDate struct {
	time.Time
}

func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return fmt.Errorf("invalid date format, expected YYYY-MM-DD: %w", err)
	}
	cd.Time = t
	return nil
}

func (cd *CustomDate) UnmarshalText(text []byte) error {
	t, err := time.Parse("2006-01-02", string(text))
	if err != nil {
		return fmt.Errorf("invalid date format, expected YYYY-MM-DD: %w", err)
	}
	cd.Time = t
	return nil
}
