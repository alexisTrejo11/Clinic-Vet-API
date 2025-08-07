package paymentDomain

type PaymentStatus string

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
	CASH          PaymentMethod = "cash"
	CREDIT_CARD   PaymentMethod = "credit_card"
	DEBIT_CARD    PaymentMethod = "debit_card"
	BANK_TRANSFER PaymentMethod = "bank_transfer"
	PAYPAL        PaymentMethod = "paypal"
	STRIPE        PaymentMethod = "stripe"
	CHECK         PaymentMethod = "check"
)

func (pm PaymentMethod) IsValid() bool {
	switch pm {
	case CASH, CREDIT_CARD, DEBIT_CARD, BANK_TRANSFER, PAYPAL, STRIPE, CHECK:
		return true
	default:
		return false
	}
}

func (pm PaymentMethod) RequiresOnlineProcessing() bool {
	switch pm {
	case CREDIT_CARD, DEBIT_CARD, PAYPAL, STRIPE:
		return true
	default:
		return false
	}
}
