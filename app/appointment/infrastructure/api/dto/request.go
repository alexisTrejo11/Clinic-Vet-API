package appControllerDTO

import (
	"encoding/json"
	"fmt"
	"time"
)

type GetAppointmentsByDateRangeRequest struct {
	StartDate  CustomDate `json:"start_date" binding:"required"`
	EndDate    CustomDate `json:"end_date" binding:"required"`
	PageNumber int        `json:"page_number" validate:"min=1"`
	PageSize   int        `json:"page_size" validate:"min=1,max=100"`
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
