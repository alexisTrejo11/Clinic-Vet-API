package valueobject

import (
	domainerr "clinic-vet-api/app/core/error"
	"context"
	"fmt"
)

type Money struct {
	amount   int64 // Amount in smallest currency unit (cents)
	currency string
}

func (m Money) Amount() int64 {
	return m.amount
}

func (m Money) Currency() string {
	return m.currency
}

func NewMoney(amount float64, currency string) Money {
	return Money{
		amount:   int64(amount * 100), // Convert to cents
		currency: currency,
	}
}

func (m Money) ToFloat() float64 {
	return float64(m.amount) / 100.0
}

func (m Money) FormatWithCurrency(currency string) string {
	amount := m.ToFloat()
	switch currency {
	case "USD":
		return fmt.Sprintf("$%.2f", amount)
	case "EUR":
		return fmt.Sprintf("€%.2f", amount)
	case "MXN":
		return fmt.Sprintf("$%.2f MXN", amount)
	default:
		return fmt.Sprintf("%.2f %s", amount, currency)
	}
}

func (m Money) Add(other Money) (Money, error) {
	if m.currency != other.currency {
		return Money{}, domainerr.InvalidFieldFormat(context.Background(), "money", "currency mismatch", fmt.Sprintf("cannot add different currencies: %s and %s", m.currency, other.Currency()), "create money")
	}
	return Money{
		amount:   m.amount + other.Amount(),
		currency: m.currency,
	}, nil
}

func (m Money) Subtract(other Money) (Money, error) {
	if m.currency != other.Currency() {
		return Money{}, domainerr.InvalidFieldFormat(context.Background(), "money", "currency mismatch", fmt.Sprintf("cannot subtract different currencies: %s and %s", m.currency, other.Currency()), "create money")
	}
	return Money{
		amount:   m.amount - other.Amount(),
		currency: m.currency,
	}, nil
}

func (m Money) IsZero() bool {
	return m.amount == 0
}

func (m Money) IsPositive() bool {
	return m.amount > 0
}

func (m Money) IsNegative() bool {
	return m.amount < 0
}
