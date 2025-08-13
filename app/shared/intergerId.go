package shared

import (
	"fmt"

	domainErr "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/domain"
)

type IntegerId struct {
	value int
}

func invalidIntegerIdError(value any) error {
	return domainErr.NewValidationError("invalid IntegerId", fmt.Sprintf("%T", value), fmt.Sprintf("expected int, got %T", value))
}

func NilIntegerId() IntegerId {
	return IntegerId{value: 0}
}

func NewIntegerId(value any) (IntegerId, error) {
	idInt, err := parseToInt(value)
	if err != nil {
		return IntegerId{}, invalidIntegerIdError(value)
	}

	return IntegerId{value: idInt}, nil
}

func (m IntegerId) GetValue() int {
	return m.value
}

func (m IntegerId) String() string {
	return fmt.Sprintf("%d", m.value)
}

func (m IntegerId) Equals(other IntegerId) bool {
	return m.value == other.value
}

func parseToInt(value any) (int, error) {
	v32, ok := value.(int32)
	if ok {
		return int(v32), nil
	}

	v64, ok := value.(int64)
	if ok {
		return int(v64), nil
	}

	vUint, ok := value.(uint)
	if ok {
		return int(vUint), nil
	}

	vUint32, ok := value.(uint32)
	if ok {
		return int(vUint32), nil
	}

	vUint64, ok := value.(uint64)
	if ok {
		return int(vUint64), nil
	}

	v, ok := value.(int)
	if ok {
		return v, nil
	}

	return 0, invalidIntegerIdError(value)
}
