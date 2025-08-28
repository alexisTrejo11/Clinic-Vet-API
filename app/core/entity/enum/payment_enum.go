package enum

import "errors"

type PaymentStatus string

func NewPaymentStatus(status string) (PaymentStatus, error) {
	ps := PaymentStatus(status)
	if !ps.IsValid() {
		return "", errors.New("invalid payment status: " + string(ps))
	}
	return ps, nil
}

const (
	PENDING   PaymentStatus = "pending"
	PAID      PaymentStatus = "paid"
	FAILED    PaymentStatus = "failed"
	CANCELLED PaymentStatus = "cancelled"
	REFUNDED  PaymentStatus = "refunded"
	OVERDUE   PaymentStatus = "overdue"
)

func (ps PaymentStatus) IsValid() bool {
	switch ps {
	case PENDING, PAID, FAILED, CANCELLED, REFUNDED, OVERDUE:
		return true
	default:
		return false
	}
}

type PaymentMethod string

const (
	CASH         PaymentMethod = "cash"
	CREDITCARD   PaymentMethod = "credit_card"
	DEBITCARD    PaymentMethod = "debit_card"
	BANKTRANSFER PaymentMethod = "bank_transfer"
	PAYPAL       PaymentMethod = "paypal"
	STRIPE       PaymentMethod = "stripe"
	CHECK        PaymentMethod = "check"
)

func NewPaymentMethod(method string) (PaymentMethod, error) {
	pm := PaymentMethod(method)
	if !pm.IsValid() {
		return "", errors.New("invalid payment method: " + string(pm))
	}
	return pm, nil
}

func (pm PaymentMethod) IsValid() bool {
	switch pm {
	case CASH, CREDITCARD, DEBITCARD, BANKTRANSFER, PAYPAL, STRIPE, CHECK:
		return true
	default:
		return false
	}
}

func (pm PaymentMethod) RequiresOnlineProcessing() bool {
	switch pm {
	case CREDITCARD, DEBITCARD, PAYPAL, STRIPE:
		return true
	default:
		return false
	}
}
