package valueobject

import "errors"

type VaccineType string

const (
	VaccineTypeCore       VaccineType = "core"
	VaccineTypeNonCore    VaccineType = "non_core"
	VaccineTypeLegallyReq VaccineType = "legally_req"
)

type VaccineName struct {
	value string
}

func ParseVaccineName(name string) (VaccineName, error) {
	if name == "" {
		return VaccineName{}, errors.New("vaccine name cannot be empty")
	}
	if len(name) > 100 {
		return VaccineName{}, errors.New("vaccine name too long")
	}
	return VaccineName{value: name}, nil
}

func (v VaccineName) Value() string { return v.value }
