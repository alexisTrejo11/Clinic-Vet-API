package payment

import (
	"time"

	"clinic-vet-api/app/core/domain/enum"
)

type PaymentReport struct {
	StartDate        time.Time                  `json:"start_date"`
	EndDate          time.Time                  `json:"end_date"`
	TotalPayments    int                        `json:"total_payments"`
	TotalAmount      float64                    `json:"total_amount"`
	TotalCurrency    string                     `json:"total_currency"`
	PaidAmount       float64                    `json:"paid_amount"`
	PendingAmount    float64                    `json:"pending_amount"`
	RefundedAmount   float64                    `json:"refunded_amount"`
	OverdueAmount    float64                    `json:"overdue_amount"`
	PaymentsByMethod map[enum.PaymentMethod]any `json:"payments_by_method"`
	PaymentsByStatus map[enum.PaymentStatus]any `json:"payments_by_status"`
}
