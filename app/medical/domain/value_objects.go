package mhDomain

import (
	"fmt"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type MedHistoryId struct {
	id shared.IntegerId
}

func NewMedHistoryId(value any) (MedHistoryId, error) {
	id, err := shared.NewIntegerId(value)
	if err != nil {
		return MedHistoryId{}, fmt.Errorf("invalid MedHistoryId: %w", err)
	}

	return MedHistoryId{id: id}, nil
}

func (m MedHistoryId) GetValue() int {
	return m.id.GetValue()
}

func (m MedHistoryId) String() string {
	return m.id.String()
}

func (m MedHistoryId) Equals(other MedHistoryId) bool {
	return m.id.GetValue() == other.id.GetValue()
}
