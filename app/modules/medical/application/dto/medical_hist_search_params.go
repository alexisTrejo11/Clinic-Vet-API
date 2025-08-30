package dto

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type MedHistSearchParams struct {
	Page       page.PageData
	SortBy     string
	VetID      int
	OwnerID    int
	PetID      int
	BeforeAt   time.Time
	AfterAt    time.Time
	ReasonLike string
}
