package shared

import "fmt"

type IntegerId struct {
	value int
}

func NewIntegerId(value any) (IntegerId, error) {
	idInt, err := parseToInt(value)
	if err != nil {
		return IntegerId{}, fmt.Errorf("invalid IntegerId: %w", err)
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
	v, ok := value.(int)

	if !ok {
		return 0, fmt.Errorf("invalid type for IntegerId: expected int, got %T", value)
	}

	if v <= 0 {
		return 0, fmt.Errorf("invalid IntegerId: %d", v)
	}
	return v, nil
}
