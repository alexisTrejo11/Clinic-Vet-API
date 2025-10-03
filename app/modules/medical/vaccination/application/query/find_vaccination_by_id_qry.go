package query

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

type FindVaccinationByIDQuery struct {
	id            valueobject.VaccinationID
	optPetID      *valueobject.PetID
	optEmployeeID *valueobject.EmployeeID
}

func NewFindVaccinationByIDQuery(id uint, optPetID *uint, optEmployeeID *uint) (FindVaccinationByIDQuery, error) {
	cmd := FindVaccinationByIDQuery{
		id:            valueobject.NewVaccinationID(id),
		optPetID:      valueobject.NewOptPetID(optPetID),
		optEmployeeID: valueobject.NewOptEmployeeID(optEmployeeID),
	}

	if err := cmd.Validate(); err != nil {
		return FindVaccinationByIDQuery{}, nil
	}

	return cmd, nil
}

func (q FindVaccinationByIDQuery) Validate() error {
	return nil
}

func (q FindVaccinationByIDQuery) ID() valueobject.VaccinationID          { return q.id }
func (q FindVaccinationByIDQuery) OptPetID() *valueobject.PetID           { return q.optPetID }
func (q FindVaccinationByIDQuery) OptEmployeeID() *valueobject.EmployeeID { return q.optEmployeeID }
