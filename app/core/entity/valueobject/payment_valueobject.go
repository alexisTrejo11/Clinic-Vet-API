package valueobject

import "fmt"

type Money struct {
	Amount   int64  `json:"amount" db:"amount"` // Amount in smallest currency unit (cents)
	Currency string `json:"currency" db:"currency"`
}

func NewMoney(amount float64, currency string) Money {
	return Money{
		Amount:   int64(amount * 100), // Convert to cents
		Currency: currency,
	}
}

func (m Money) ToFloat() float64 {
	return float64(m.Amount) / 100.0
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
	if m.Currency != other.Currency {
		return Money{}, fmt.Errorf("cannot add different currencies: %s and %s", m.Currency, other.Currency)
	}
	return Money{
		Amount:   m.Amount + other.Amount,
		Currency: m.Currency,
	}, nil
}

func (m Money) Subtract(other Money) (Money, error) {
	if m.Currency != other.Currency {
		return Money{}, fmt.Errorf("cannot subtract different currencies: %s and %s", m.Currency, other.Currency)
	}
	return Money{
		Amount:   m.Amount - other.Amount,
		Currency: m.Currency,
	}, nil
}

func (m Money) IsZero() bool {
	return m.Amount == 0
}

func (m Money) IsPositive() bool {
	return m.Amount > 0
}

func (m Money) IsNegative() bool {
	return m.Amount < 0
}
