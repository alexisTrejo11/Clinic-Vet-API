package payment

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/base"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	domainerr "github.com/alexisTrejo11/Clinic-Vet-API/app/core/error"
)

type Payment struct {
	base.Entity
	appointmentID valueobject.AppointmentID
	userID        valueobject.UserID
	amount        valueobject.Money
	currency      string
	paymentMethod enum.PaymentMethod
	status        enum.PaymentStatus
	transactionID *string
	description   *string
	dueDate       *time.Time
	paidAt        *time.Time
	refundedAt    *time.Time
}

// PaymentOption defines the functional option type
type PaymentOption func(*Payment) error

// Functional options
func WithAmount(amount valueobject.Money) PaymentOption {
	return func(p *Payment) error {
		if amount.Amount <= 0 {
			return domainerr.NewValidationError("payment", "amount", "amount must be positive")
		}
		p.amount = amount
		return nil
	}
}

func WithCurrency(currency string) PaymentOption {
	return func(p *Payment) error {
		if currency == "" {
			return domainerr.NewValidationError("payment", "currency", "currency is required")
		}
		if len(currency) != 3 {
			return domainerr.NewValidationError("payment", "currency", "currency must be 3-letter code")
		}
		p.currency = currency
		return nil
	}
}

func WithPaymentMethod(method enum.PaymentMethod) PaymentOption {
	return func(p *Payment) error {
		if !method.IsValid() {
			return domainerr.NewValidationError("payment", "paymentMethod", "invalid payment method")
		}
		p.paymentMethod = method
		return nil
	}
}

func WithStatus(status enum.PaymentStatus) PaymentOption {
	return func(p *Payment) error {
		if !status.IsValid() {
			return domainerr.NewValidationError("payment", "status", "invalid payment status")
		}
		p.status = status
		return nil
	}
}

func WithTransactionID(transactionID *string) PaymentOption {
	return func(p *Payment) error {
		if transactionID != nil && *transactionID == "" {
			return domainerr.NewValidationError("payment", "transactionID", "transaction ID cannot be empty")
		}
		p.transactionID = transactionID
		return nil
	}
}

func WithDescription(description *string) PaymentOption {
	return func(p *Payment) error {
		if description != nil && len(*description) > 500 {
			return domainerr.NewValidationError("payment", "description", "description too long")
		}
		p.description = description
		return nil
	}
}

func WithDueDate(dueDate *time.Time) PaymentOption {
	return func(p *Payment) error {
		if dueDate != nil && dueDate.Before(time.Now()) {
			return domainerr.NewValidationError("payment", "dueDate", "due date cannot be in the past")
		}
		p.dueDate = dueDate
		return nil
	}
}

func WithPaidAt(paidAt *time.Time) PaymentOption {
	return func(p *Payment) error {
		if paidAt != nil && paidAt.After(time.Now()) {
			return domainerr.NewValidationError("payment", "paidAt", "paid date cannot be in the future")
		}
		p.paidAt = paidAt
		return nil
	}
}

func WithRefundedAt(refundedAt *time.Time) PaymentOption {
	return func(p *Payment) error {
		if refundedAt != nil && refundedAt.After(time.Now()) {
			return domainerr.NewValidationError("payment", "refundedAt", "refunded date cannot be in the future")
		}
		p.refundedAt = refundedAt
		return nil
	}
}

// NewPayment creates a new Payment with functional options
func NewPayment(
	id valueobject.PaymentID,
	appointmentID valueobject.AppointmentID,
	userID valueobject.UserID,
	opts ...PaymentOption,
) (*Payment, error) {
	if id.GetValue() == 0 {
		return nil, domainerr.NewValidationError("payment", "id", "payment ID is required")
	}
	if appointmentID.GetValue() == 0 {
		return nil, domainerr.NewValidationError("payment", "appointmentID", "appointment ID is required")
	}
	if userID.GetValue() == 0 {
		return nil, domainerr.NewValidationError("payment", "userID", "user ID is required")
	}

	payment := &Payment{
		Entity:        base.NewEntity(id),
		appointmentID: appointmentID,
		userID:        userID,
		status:        enum.PaymentStatusPending, // Default status
		currency:      "USD",                     // Default currency
	}

	// Apply all options
	for _, opt := range opts {
		if err := opt(payment); err != nil {
			return nil, err
		}
	}

	// Final validation
	if err := payment.validate(); err != nil {
		return nil, err
	}

	return payment, nil
}

// Validation
func (p *Payment) validate() error {
	if p.amount.Amount <= 0 {
		return domainerr.NewValidationError("payment", "amount", "amount must be positive")
	}
	if p.currency == "" {
		return domainerr.NewValidationError("payment", "currency", "currency is required")
	}
	if !p.paymentMethod.IsValid() {
		return domainerr.NewValidationError("payment", "paymentMethod", "payment method is required")
	}
	if !p.status.IsValid() {
		return domainerr.NewValidationError("payment", "status", "status is required")
	}
	return nil
}

// Getters
func (p *Payment) ID() valueobject.PaymentID {
	return p.ID()
}

func (p *Payment) AppointmentID() valueobject.AppointmentID {
	return p.appointmentID
}

func (p *Payment) UserID() valueobject.UserID {
	return p.userID
}

func (p *Payment) Amount() valueobject.Money {
	return p.amount
}

func (p *Payment) Currency() string {
	return p.currency
}

func (p *Payment) PaymentMethod() enum.PaymentMethod {
	return p.paymentMethod
}

func (p *Payment) Status() enum.PaymentStatus {
	return p.status
}

func (p *Payment) TransactionID() *string {
	return p.transactionID
}

func (p *Payment) Description() *string {
	return p.description
}

func (p *Payment) DueDate() *time.Time {
	return p.dueDate
}

func (p *Payment) PaidAt() *time.Time {
	return p.paidAt
}

func (p *Payment) RefundedAt() *time.Time {
	return p.refundedAt
}

// Business logic methods
func (p *Payment) Update(amount *valueobject.Money, paymentMethod *enum.PaymentMethod, description *string, dueDate *time.Time) error {
	if amount != nil {
		if amount.Amount <= 0 {
			return domainerr.NewValidationError("payment", "amount", "amount must be positive")
		}
		p.amount = *amount
	}

	if paymentMethod != nil {
		if !paymentMethod.IsValid() {
			return domainerr.NewValidationError("payment", "paymentMethod", "invalid payment method")
		}
		p.paymentMethod = *paymentMethod
	}

	if description != nil {
		if len(*description) > 500 {
			return domainerr.NewValidationError("payment", "description", "description too long")
		}
		p.description = description
	}

	if dueDate != nil {
		if dueDate.Before(time.Now()) {
			return domainerr.NewValidationError("payment", "dueDate", "due date cannot be in the past")
		}
		p.dueDate = dueDate
	}

	p.IncrementVersion()
	return nil
}

func (p *Payment) Cancel(reason string) error {
	allowedStatuses := []enum.PaymentStatus{
		enum.PaymentStatusPending,
		enum.PaymentStatusPaid,
		enum.PaymentStatusFailed,
	}

	if !p.canTransitionTo(enum.PaymentStatusCancelled, allowedStatuses) {
		return domainerr.NewBusinessRuleError("payment", "cancel", "payment cannot be cancelled in current status")
	}

	if reason == "" {
		return domainerr.NewValidationError("payment", "reason", "cancellation reason is required")
	}

	p.status = enum.PaymentStatusCancelled
	p.IncrementVersion()
	return nil
}

func (p *Payment) Pay(transactionID string) error {
	allowedStatuses := []enum.PaymentStatus{
		enum.PaymentStatusFailed,
		enum.PaymentStatusPending,
	}

	if !p.canTransitionTo(enum.PaymentStatusPaid, allowedStatuses) {
		return domainerr.NewBusinessRuleError("payment", "pay", "payment cannot be paid in current status")
	}

	if transactionID == "" {
		return domainerr.NewValidationError("payment", "transactionID", "transaction ID is required")
	}

	now := time.Now()
	p.status = enum.PaymentStatusPaid
	p.transactionID = &transactionID
	p.paidAt = &now
	p.IncrementVersion()

	return nil
}

func (p *Payment) Refund() error {
	allowedStatuses := []enum.PaymentStatus{enum.PaymentStatusPaid}

	if !p.canTransitionTo(enum.PaymentStatusRefunded, allowedStatuses) {
		return domainerr.NewBusinessRuleError("payment", "refund", "payment cannot be refunded in current status")
	}

	now := time.Now()
	p.status = enum.PaymentStatusRefunded
	p.refundedAt = &now
	p.IncrementVersion()

	return nil
}

func (p *Payment) MarkAsOverdue() error {
	allowedStatuses := []enum.PaymentStatus{
		enum.PaymentStatusFailed,
		enum.PaymentStatusPending,
	}

	if !p.canTransitionTo(enum.PaymentStatusOverdue, allowedStatuses) {
		return domainerr.NewBusinessRuleError("payment", "overdue", "payment cannot be marked as overdue in current status")
	}

	if p.dueDate != nil && time.Now().Before(*p.dueDate) {
		return domainerr.NewBusinessRuleError("payment", "overdue", "cannot mark as overdue before due date")
	}

	p.status = enum.PaymentStatusOverdue
	p.IncrementVersion()
	return nil
}

func (p *Payment) MarkAsFailed() error {
	allowedStatuses := []enum.PaymentStatus{enum.PaymentStatusPending}

	if !p.canTransitionTo(enum.PaymentStatusFailed, allowedStatuses) {
		return domainerr.NewBusinessRuleError("payment", "failed", "payment cannot be marked as failed in current status")
	}

	p.status = enum.PaymentStatusFailed
	p.IncrementVersion()
	return nil
}

// Helper methods
func (p *Payment) IsOverdue() bool {
	if p.status == enum.PaymentStatusOverdue {
		return true
	}

	if p.paidAt != nil || p.status.IsFinal() {
		return false
	}

	if p.dueDate != nil && time.Now().After(*p.dueDate) {
		return true
	}

	return false
}

func (p *Payment) IsPaid() bool {
	return p.status == enum.PaymentStatusPaid
}

func (p *Payment) IsRefunded() bool {
	return p.status == enum.PaymentStatusRefunded
}

func (p *Payment) IsCancelled() bool {
	return p.status == enum.PaymentStatusCancelled
}

func (p *Payment) CanRetry() bool {
	return p.status == enum.PaymentStatusFailed || p.status == enum.PaymentStatusPending
}

func (p *Payment) AmountDue() valueobject.Money {
	if p.IsPaid() || p.IsRefunded() || p.IsCancelled() {
		return valueobject.NewMoney(0, p.currency)
	}
	return p.amount
}

// Private helper
func (p *Payment) canTransitionTo(newStatus enum.PaymentStatus, allowedFrom []enum.PaymentStatus) bool {
	for _, status := range allowedFrom {
		if p.status == status {
			return true
		}
	}
	return false
}

// Additional business logic
func (p *Payment) RequiresPaymentProcessing() bool {
	return p.status == enum.PaymentStatusPending &&
		p.paymentMethod.RequiresOnlineProcessing()
}

func (p *Payment) CanBeRefunded() bool {
	return p.status == enum.PaymentStatusPaid &&
		p.paidAt != nil &&
		p.paidAt.After(time.Now().AddDate(0, -6, 0)) // Within 6 months
}

func (p *Payment) DaysOverdue() int {
	if !p.IsOverdue() || p.dueDate == nil {
		return 0
	}
	return int(time.Since(*p.dueDate).Hours() / 24)
}
