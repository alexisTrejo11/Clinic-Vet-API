package domain

import "time"

type Veterinarian struct {
	ID        int
	Name      string
	Photo     *string
	Specialty *string
	UserID    *int
	CreatedAt time.Time
	UpdatedAt time.Time
}
