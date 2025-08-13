package appControllerDTO

import (
	"encoding/json"
	"fmt"
	"time"
)

type GetAppointmentsByDateRangeRequest struct {
	StartDate CustomDate `json:"start_date"  form:"start_date" binding:"required"`
	EndDate   CustomDate `json:"end_date"  form:"end_date" binding:"required"`
	PaginationRequest
}

type PaginationRequest struct {
	PageNumber int `json:"page_number" form:"page_number" validate:"min=1"`
	PageSize   int `json:"page_size" form:"page_size" validate:"min=1,max=100"`
}

func (r *PaginationRequest) SetDefaultsIfNotProvided() {
	if r.PageNumber <= 0 {
		r.PageNumber = 1
	}

	if r.PageSize <= 0 {
		r.PageSize = 10
	}

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
