package valueobject

import (
	domainerr "clinic-vet-api/app/modules/core/error"
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Decimal struct {
	value int64
}

func NewDecimalFromString(s string) (Decimal, error) {
	s = strings.TrimSpace(s)
	s = strings.Replace(s, ",", "", -1)

	parts := strings.Split(s, ".")
	if len(parts) > 2 {
		return Decimal{}, domainerr.ValidationError(context.Background(), "Decimal", s, "invalid decimal format", "creating decimal")
	}

	integerPart := parts[0]
	if integerPart == "" {
		integerPart = "0"
	}

	decimalPart := "00"
	if len(parts) == 2 {
		decimalPart = parts[1]
		if len(decimalPart) > 2 {
			decimalPart = decimalPart[:2]
		} else if len(decimalPart) < 2 {
			decimalPart = decimalPart + strings.Repeat("0", 2-len(decimalPart))

		}
	}

	fullNumberStr := integerPart + decimalPart
	value, err := strconv.ParseInt(fullNumberStr, 10, 64)
	if err != nil {
		return Decimal{}, fmt.Errorf("error parsing decimal: %v", err)
	}

	if strings.HasPrefix(integerPart, "-") {
		if value > 0 {
			value = -value
		}
	}

	return Decimal{value: value}, nil
}

func NewDecimalFromFloat(f float64) Decimal {
	rounded := math.Round(f * 100)
	return Decimal{value: int64(rounded)}
}

func NewDecimalFromInt(i int64) Decimal {
	return Decimal{value: i * 100}
}

func (d Decimal) Float64() float64 {
	return float64(d.value) / 100.0
}

func (d Decimal) String() string {
	absValue := d.value
	if d.value < 0 {
		absValue = -d.value
	}

	integerPart := absValue / 100
	decimalPart := absValue % 100

	if d.value < 0 {
		return fmt.Sprintf("-%d.%02d", integerPart, decimalPart)
	}
	return fmt.Sprintf("%d.%02d", integerPart, decimalPart)
}

func (d Decimal) Add(other Decimal) Decimal {
	return Decimal{value: d.value + other.value}
}

func (d Decimal) Sub(other Decimal) Decimal {
	return Decimal{value: d.value - other.value}
}

func (d Decimal) Mul(other Decimal) Decimal {
	result := (d.value * other.value) / 100
	return Decimal{value: result}
}

func (d Decimal) Value() (driver.Value, error) {
	return d.Float64(), nil
}

func (d *Decimal) Int() int64 {
	return d.value
}

func (d *Decimal) Scan(value interface{}) error {
	switch v := value.(type) {
	case float64:
		*d = NewDecimalFromFloat(v)
	case int64:
		*d = NewDecimalFromInt(v)
	case string:
		decimal, err := NewDecimalFromString(v)
		if err != nil {
			return err
		}
		*d = decimal
	case nil:
		*d = Decimal{value: 0}
	default:
		return fmt.Errorf("tipo no soportado: %T", value)
	}
	return nil
}

func (d Decimal) IsZero() bool {
	return d.value == 0
}

func (d Decimal) IsNegative() bool {
	return d.value < 0
}

func (d Decimal) IsPositive() bool {
	return d.value > 0
}

func (d Decimal) Equal(other Decimal) bool {
	return d.value == other.value
}

func (d Decimal) GreaterThan(other Decimal) bool {
	return d.value > other.value
}

func (d Decimal) LessThan(other Decimal) bool {
	return d.value < other.value
}

// Implementación para JSON
func (d Decimal) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Float64())
}

func (d *Decimal) UnmarshalJSON(data []byte) error {
	var f float64
	if err := json.Unmarshal(data, &f); err != nil {
		var s string
		if err2 := json.Unmarshal(data, &s); err2 != nil {
			return fmt.Errorf("decimal debe ser número o string: %v", err)
		}
		decimal, err2 := NewDecimalFromString(s)
		if err2 != nil {
			return err2
		}
		*d = decimal
		return nil
	}
	*d = NewDecimalFromFloat(f)
	return nil
}
