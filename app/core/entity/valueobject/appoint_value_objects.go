package valueobject

import "github.com/alexisTrejo11/Clinic-Vet-API/app/shared"

type AppointmentID struct {
	shared.IntegerId
}

func NewAppointmentID(value any) (AppointmentID, error) {
	id, err := shared.NewIntegerId(value)
	if err != nil {
		return AppointmentID{IntegerId: shared.NilIntegerId()}, err
	}
	return AppointmentID{IntegerId: id}, nil
}

func (a AppointmentID) GetValue() int {
	return a.IntegerId.GetValue()
}

func (a AppointmentID) String() string {
	return a.IntegerId.String()
}
