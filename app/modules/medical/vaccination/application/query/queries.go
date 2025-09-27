package vaccinequery

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/page"
	"time"
)

type FindVaccinationByIDQuery struct {
	ID            valueobject.VaccinationID
	OptPetID      *valueobject.PetID
	OptEmployeeID *valueobject.EmployeeID
}

type FindVaccinationsByPetQuery struct {
	PetID         valueobject.PetID
	OptEmployeeID *valueobject.EmployeeID
	Pagination    page.PaginationRequest
}

type FindVaccinationsByEmployeeQuery struct {
	EmployeeID valueobject.EmployeeID
	Pagination page.PaginationRequest
}

type FindVaccinationsByDateRangeQuery struct {
	StartDate  time.Time
	EndDate    time.Time
	Pagination page.PaginationRequest
}
