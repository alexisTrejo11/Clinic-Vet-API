package payment

import (
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/base"
	"clinic-vet-api/app/modules/core/domain/enum"
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
)

type Payment struct {
	base.Entity[vo.PaymentID]
	amount         vo.Money
	status         enum.PaymentStatus
	method         enum.PaymentMethod
	transactionID  *string
	description    *string
	dueDate        *time.Time
	paidAt         *time.Time
	refundedAt     *time.Time
	isActive       bool
	paidByCustomer vo.CustomerID
	paidToEmployee vo.EmployeeID
	medSessionID   *vo.MedSessionID
	invoiceID      *string
	refundAmount   *vo.Money
	failureReason  *string
	createdAt      time.Time
	updatedAt      time.Time
	deletedAt      *time.Time
}

type PaymentBuilder struct{ payment *Payment }

func NewPaymentBuilder() *PaymentBuilder {
	return &PaymentBuilder{payment: &Payment{}}
}

func (b *PaymentBuilder) Build() *Payment { return b.payment }

func (b *PaymentBuilder) WithAmount(amount vo.Money) *PaymentBuilder {
	b.payment.amount = amount
	return b
}

func (b *PaymentBuilder) WithID(id vo.PaymentID) *PaymentBuilder {
	b.payment.Entity.SetID(id)
	return b
}

func (b *PaymentBuilder) WithPaymentMethod(method enum.PaymentMethod) *PaymentBuilder {
	b.payment.method = method
	return b
}

func (b *PaymentBuilder) WithStatus(status enum.PaymentStatus) *PaymentBuilder {
	b.payment.status = status
	return b
}

func (b *PaymentBuilder) WithTransactionID(transactionID string) *PaymentBuilder {
	b.payment.transactionID = &transactionID
	return b
}

func (b *PaymentBuilder) WithDescription(description *string) *PaymentBuilder {
	b.payment.description = description
	return b
}

func (b *PaymentBuilder) WithDueDate(dueDate *time.Time) *PaymentBuilder {
	b.payment.dueDate = dueDate
	return b
}

func (b *PaymentBuilder) WithPaidAt(paidAt time.Time) *PaymentBuilder {
	b.payment.paidAt = &paidAt
	return b
}

func (b *PaymentBuilder) WithRefundedAt(refundedAt time.Time) *PaymentBuilder {
	b.payment.refundedAt = &refundedAt
	return b
}

func (b *PaymentBuilder) WithPaidByCustomer(customerID vo.CustomerID) *PaymentBuilder {
	b.payment.paidByCustomer = customerID
	return b
}

func (b *PaymentBuilder) WithMedSessionID(medSessionID vo.MedSessionID) *PaymentBuilder {
	b.payment.medSessionID = &medSessionID
	return b
}

func (b *PaymentBuilder) WithInvoiceID(invoiceID string) *PaymentBuilder {
	b.payment.invoiceID = &invoiceID
	return b
}

func (b *PaymentBuilder) WithRefundAmount(refundAmount vo.Money) *PaymentBuilder {
	b.payment.refundAmount = &refundAmount
	return b
}

func (b *PaymentBuilder) WithFailureReason(failureReason string) *PaymentBuilder {
	b.payment.failureReason = &failureReason
	return b
}

func (b *PaymentBuilder) WithIsActive(isActive bool) *PaymentBuilder {
	b.payment.isActive = isActive
	return b
}

func (b *PaymentBuilder) WithTimeStamps(createdAt, updatedAt time.Time) *PaymentBuilder {
	b.payment.Entity.SetTimeStamps(createdAt, updatedAt)
	return b
}

func (p *Payment) ID() vo.PaymentID               { return p.Entity.ID() }
func (p *Payment) Amount() vo.Money               { return p.amount }
func (p *Payment) Currency() string               { return p.amount.Currency() }
func (p *Payment) Status() enum.PaymentStatus     { return p.status }
func (p *Payment) Method() enum.PaymentMethod     { return p.method }
func (p *Payment) TransactionID() *string         { return p.transactionID }
func (p *Payment) Description() *string           { return p.description }
func (p *Payment) DueDate() *time.Time            { return p.dueDate }
func (p *Payment) PaidAt() *time.Time             { return p.paidAt }
func (p *Payment) RefundedAt() *time.Time         { return p.refundedAt }
func (p *Payment) IsActive() bool                 { return p.isActive }
func (p *Payment) PaidByCustomer() vo.CustomerID  { return p.paidByCustomer }
func (p *Payment) PaidToEmployee() vo.EmployeeID  { return p.paidToEmployee }
func (p *Payment) MedSessionID() *vo.MedSessionID { return p.medSessionID }
func (p *Payment) InvoiceID() *string             { return p.invoiceID }
func (p *Payment) RefundAmount() *vo.Money        { return p.refundAmount }
func (p *Payment) FailureReason() *string         { return p.failureReason }
func (p *Payment) CreatedAt() time.Time           { return p.createdAt }
func (p *Payment) UpdatedAt() time.Time           { return p.updatedAt }
func (p *Payment) DeletedAt() *time.Time          { return p.deletedAt }

func (p *Payment) IsPaid() bool      { return p.status == enum.PaymentStatusPaid }
func (p *Payment) IsPending() bool   { return p.status == enum.PaymentStatusPending }
func (p *Payment) IsFailed() bool    { return p.status == enum.PaymentStatusFailed }
func (p *Payment) IsRefunded() bool  { return p.status == enum.PaymentStatusRefunded }
func (p *Payment) IsCancelled() bool { return p.status == enum.PaymentStatusCancelled }

func (p *Payment) IsOverdue() bool {
	if p.status == enum.PaymentStatusOverdue {
		return true
	}

	if p.paidAt != nil || p.status.IsFinal() {
		return false
	}

	now := time.Now()
	return p.dueDate.Before(now) &&
		p.status != enum.PaymentStatusPaid &&
		p.status != enum.PaymentStatusRefunded &&
		p.status != enum.PaymentStatusCancelled
}

func (p *Payment) IsRefundable() bool {
	return p.status == enum.PaymentStatusPaid && p.paidAt != nil && !p.isRefundPeriodExpired()
}

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
