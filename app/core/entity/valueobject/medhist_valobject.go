package valueobject

import (
	"fmt"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type MedHistoryID struct {
	id shared.IntegerId
}

func NewMedHistoryID(value any) (MedHistoryID, error) {
	id, err := shared.NewIntegerId(value)
	if err != nil {
		return MedHistoryID{}, fmt.Errorf("invalid MedHistoryID: %w", err)
	}

	return MedHistoryID{id: id}, nil
}

func (m MedHistoryID) GetValue() int {
	return m.id.GetValue()
}

func (m MedHistoryID) String() string {
	return m.id.String()
}

func (m MedHistoryID) Equals(other MedHistoryID) bool {
	return m.id.GetValue() == other.id.GetValue()
}
