package appointDomain

import "github.com/alexisTrejo11/Clinic-Vet-API/app/shared"

type AppointmentId struct {
	shared.IntegerId
}

func NewAppointmentId(value any) (AppointmentId, error) {
	id, err := shared.NewIntegerId(value)
	if err != nil {
		return AppointmentId{IntegerId: shared.NilIntegerId()}, err
	}
	return AppointmentId{IntegerId: id}, nil
}

func (a AppointmentId) GetValue() int {
	return a.IntegerId.GetValue()
}

func (a AppointmentId) String() string {
	return a.IntegerId.String()
}
