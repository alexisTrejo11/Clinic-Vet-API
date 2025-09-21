package valueobject

import (
	domainerr "clinic-vet-api/app/core/error"
	"context"
	"fmt"
)

type Money struct {
	amount   Decimal
	currency string
}

func (m Money) Amount() Decimal {
	return m.amount
}

func (m Money) Currency() string {
	return m.currency
}

func NewMoney(amount Decimal, currency string) Money {
	return Money{
		amount:   amount,
		currency: currency,
	}
}

func (m Money) FormatWithCurrency(currency string) string {
	switch currency {
	case "USD":
		return fmt.Sprintf("$%s", m.amount.String())
	case "EUR":
		return fmt.Sprintf("€%s", m.amount.String())
	case "MXN":
		return fmt.Sprintf("$%s MXN", m.Amount().String())
	default:
		return fmt.Sprintf("%s %s", m.amount.String(), currency)
	}
}

func (m Money) Add(other Money) (Money, error) {
	if m.currency != other.currency {
		return Money{}, domainerr.InvalidFieldFormat(context.Background(), "money", "currency mismatch", fmt.Sprintf("cannot add different currencies: %s and %s", m.currency, other.Currency()), "create money")
	}
	return Money{
		amount:   m.amount.Add(other.amount),
		currency: m.currency,
	}, nil
}

func (m Money) Subtract(other Money) (Money, error) {
	if m.currency != other.Currency() {
		return Money{}, domainerr.InvalidFieldFormat(context.Background(), "money", "currency mismatch", fmt.Sprintf("cannot subtract different currencies: %s and %s", m.currency, other.Currency()), "create money")
	}
	return Money{
		amount:   m.amount.Sub(other.amount),
		currency: m.currency,
	}, nil
}

func (m Money) IsZero() bool {
	return m.amount.value == 0
}

func (m Money) IsPositive() bool {
	return m.amount.value > 0
}

func (m Money) IsNegative() bool {
	return m.amount.value < 0
}
