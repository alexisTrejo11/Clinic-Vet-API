package query

import (
	"time"

	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/page"
)

type FindDewormByIDQuery struct {
	ID            valueobject.DewormID
	OptPetID      *valueobject.PetID
	OptEmployeeID *valueobject.EmployeeID
}

type FindDewormsByPetQuery struct {
	PetID         valueobject.PetID
	OptEmployeeID *valueobject.EmployeeID
	Pagination    page.PaginationRequest
}

type FindDewormsByEmployeeQuery struct {
	EmployeeID valueobject.EmployeeID
	Pagination page.PaginationRequest
}

type FindDewormsByDateRangeQuery struct {
	StartDate  time.Time
	EndDate    time.Time
	Pagination page.PaginationRequest
}
