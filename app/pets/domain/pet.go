package domain

import "time"

type Pet struct {
	ID        int
	Name      string
	Photo     *string
	Species   string
	Breed     *string
	Age       *int
	OwnerID   int
	CreatedAt time.Time
	UpdatedAt time.Time
}
