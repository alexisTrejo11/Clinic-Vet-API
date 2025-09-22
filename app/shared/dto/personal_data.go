package commondto

import (
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"time"
)

type PersonalData struct {
	DateOfBirth time.Time              `json:"date_of_birth,omitempty"`
	Name        valueobject.PersonName `json:"name,omitempty"`
	Gender      enum.PersonGender      `json:"gender,omitempty"`
	Location    string                 `json:"location,omitempty"`
}
