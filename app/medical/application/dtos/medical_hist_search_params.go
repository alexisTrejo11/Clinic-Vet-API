package mhDTOs

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type MedHistSearchParams struct {
	Page       page.PageData
	SortBy     string
	VetId      int
	OwnerId    int
	PetId      int
	BeforeAt   time.Time
	AfterAt    time.Time
	ReasonLike string
}
