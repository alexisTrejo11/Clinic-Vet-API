package paymentCmd

import (
	"context"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
)

type PaymentCommandFactory struct {
	paymentRepo paymentDomain.PaymentRepository
	options     *FactoryOptions
}

type FactoryOptions struct {
	EnableValidation bool
	EnableAuditing   bool
	EnableMetrics    bool
	BatchSize        int
}

func DefaultFactoryOptions() *FactoryOptions {
	return &FactoryOptions{
		EnableValidation: true,
		EnableAuditing:   false,
		EnableMetrics:    false,
		BatchSize:        100,
	}
}

func NewPaymentCommandFactory(paymentRepo paymentDomain.PaymentRepository) *PaymentCommandFactory {
	return &PaymentCommandFactory{
		paymentRepo: paymentRepo,
		options:     DefaultFactoryOptions(),
	}
}

func NewPaymentCommandFactoryWithOptions(paymentRepo paymentDomain.PaymentRepository, options *FactoryOptions) *PaymentCommandFactory {
	if options == nil {
		options = DefaultFactoryOptions()
	}
	return &PaymentCommandFactory{
		paymentRepo: paymentRepo,
		options:     options,
	}
}

func (f *PaymentCommandFactory) SetOptions(options *FactoryOptions) {
	if options != nil {
		f.options = options
	}
}

func (f *PaymentCommandFactory) CreateCommandBus() CommandBus {
	return NewPaymentCommandBus(f.paymentRepo)
}

func (f *PaymentCommandFactory) CreateCommandService() paymentDomain.PaymentService {
	commandBus := f.CreateCommandBus()
	return NewPaymentCommandService(commandBus, f.paymentRepo)
}

func (f *PaymentCommandFactory) CreateHandlers() *PaymentHandlers {
	return &PaymentHandlers{
		CreatePayment:       NewCreatePaymentHandler(f.paymentRepo),
		ProcessPayment:      NewProcessPaymentHandler(f.paymentRepo),
		RefundPayment:       NewRefundPaymentHandler(f.paymentRepo),
		CancelPayment:       NewCancelPaymentHandler(f.paymentRepo),
		UpdatePayment:       NewUpdatePaymentHandler(f.paymentRepo),
		DeletePayment:       NewDeletePaymentHandler(f.paymentRepo),
		MarkOverduePayments: NewMarkOverduePaymentsHandler(f.paymentRepo),
	}
}

type PaymentHandlers struct {
	CreatePayment       CreatePaymentHander
	ProcessPayment      ProcessPaymentHandler
	RefundPayment       RefundPaymentHandler
	CancelPayment       CancelPaymentHandler
	UpdatePayment       UpdatePaymentHandler
	DeletePayment       DeletePaymentHandler
	MarkOverduePayments MarkOverduePaymentsHandler
}

func (h *PaymentHandlers) ExecuteCreatePayment(ctx context.Context, cmd CreatePaymentCommand) (*paymentDomain.Payment, error) {
	return h.CreatePayment.Handle(cmd)
}

// executes a process payment command
func (h *PaymentHandlers) ExecuteProcessPayment(ctx context.Context, cmd ProcessPaymentCommand) error {
	cmd.CTX = ctx
	return h.ProcessPayment.Handle(cmd)
}

func (h *PaymentHandlers) ExecuteRefundPayment(ctx context.Context, cmd RefundPaymentCommand) error {
	cmd.CTX = ctx
	return h.RefundPayment.Handle(cmd)
}

func (h *PaymentHandlers) ExecuteCancelPayment(ctx context.Context, cmd CancelPaymentCommand) error {
	cmd.CTX = ctx
	return h.CancelPayment.Handle(cmd)
}

func (h *PaymentHandlers) ExecuteUpdatePayment(ctx context.Context, cmd UpdatePaymentCommand) (*paymentDomain.Payment, error) {
	cmd.CTX = ctx
	return h.UpdatePayment.Handle(cmd)
}

func (h *PaymentHandlers) ExecuteDeletePayment(ctx context.Context, cmd DeletePaymentCommand) error {
	cmd.CTX = ctx
	return h.DeletePayment.Handle(cmd)
}

func (h *PaymentHandlers) ExecuteMarkOverduePayments(ctx context.Context) (int, error) {
	cmd := MarkOverduePaymentsCommand{CTX: ctx}
	return h.MarkOverduePayments.Handle(cmd)
}

type PaymentCommandFacade struct {
	handlers   *PaymentHandlers
	commandBus CommandBus
}

func (f *PaymentCommandFactory) CreateFacade() *PaymentCommandFacade {
	return &PaymentCommandFacade{
		handlers:   f.CreateHandlers(),
		commandBus: f.CreateCommandBus(),
	}
}

func (facade *PaymentCommandFacade) CreatePayment(ctx context.Context, cmd CreatePaymentCommand) (*paymentDomain.Payment, error) {
	return facade.handlers.ExecuteCreatePayment(ctx, cmd)
}

func (facade *PaymentCommandFacade) ProcessPayment(ctx context.Context, paymentId int, transactionId string) error {
	cmd := ProcessPaymentCommand{
		PaymentId:     paymentId,
		TransactionId: transactionId,
		CTX:           ctx,
	}
	return facade.handlers.ExecuteProcessPayment(ctx, cmd)
}

func (facade *PaymentCommandFacade) RefundPayment(ctx context.Context, paymentId int, reason string) error {
	cmd := RefundPaymentCommand{
		PaymentId: paymentId,
		Reason:    reason,
		CTX:       ctx,
	}
	return facade.handlers.ExecuteRefundPayment(ctx, cmd)
}

func (facade *PaymentCommandFacade) CancelPayment(ctx context.Context, paymentId int, reason string) error {
	cmd := CancelPaymentCommand{
		PaymentId: paymentId,
		Reason:    reason,
		CTX:       ctx,
	}
	return facade.handlers.ExecuteCancelPayment(ctx, cmd)
}

func (facade *PaymentCommandFacade) UpdatePayment(ctx context.Context, cmd UpdatePaymentCommand) (*paymentDomain.Payment, error) {
	return facade.handlers.ExecuteUpdatePayment(ctx, cmd)
}

func (facade *PaymentCommandFacade) DeletePayment(ctx context.Context, paymentId int) error {
	cmd := DeletePaymentCommand{
		PaymentId: paymentId,
		CTX:       ctx,
	}
	return facade.handlers.ExecuteDeletePayment(ctx, cmd)
}

func (facade *PaymentCommandFacade) MarkOverduePayments(ctx context.Context) (int, error) {
	return facade.handlers.ExecuteMarkOverduePayments(ctx)
}

func (facade *PaymentCommandFacade) ExecuteCommand(ctx context.Context, command Command) (interface{}, error) {
	return facade.commandBus.Execute(ctx, command)
}

func (f *PaymentCommandFactory) Validate() error {
	if f.paymentRepo == nil {
		return paymentDomain.NewPaymentError("INVALID_FACTORY", "payment repository cannot be nil", 0, "")
	}

	if f.options == nil {
		return paymentDomain.NewPaymentError("INVALID_FACTORY", "factory options cannot be nil", 0, "")
	}

	if f.options.BatchSize <= 0 {
		return paymentDomain.NewPaymentError("INVALID_FACTORY", "batch size must be greater than 0", 0, "")
	}

	return nil
}

func (f *PaymentCommandFactory) GetRepository() paymentDomain.PaymentRepository {
	return f.paymentRepo
}

func (f *PaymentCommandFactory) GetOptions() *FactoryOptions {
	return f.options
}

type PaymentCommandBuilder struct {
	factory *PaymentCommandFactory
}

func (f *PaymentCommandFactory) NewBuilder() *PaymentCommandBuilder {
	return &PaymentCommandBuilder{
		factory: f,
	}
}

func (b *PaymentCommandBuilder) BuildCreatePaymentCommand() *CreatePaymentCommandBuilder {
	return &CreatePaymentCommandBuilder{}
}

type CreatePaymentCommandBuilder struct {
	cmd CreatePaymentCommand
}

func (b *CreatePaymentCommandBuilder) WithAppointmentId(appointmentId int) *CreatePaymentCommandBuilder {
	b.cmd.AppointmentId = appointmentId
	return b
}

func (b *CreatePaymentCommandBuilder) WithOwnerId(ownerId int) *CreatePaymentCommandBuilder {
	b.cmd.OwnerId = ownerId
	return b
}

func (b *CreatePaymentCommandBuilder) WithAmount(amount float64, currency string) *CreatePaymentCommandBuilder {
	b.cmd.Amount = amount
	b.cmd.Currency = currency
	return b
}

func (b *CreatePaymentCommandBuilder) WithPaymentMethod(method paymentDomain.PaymentMethod) *CreatePaymentCommandBuilder {
	b.cmd.PaymentMethod = method
	return b
}

func (b *CreatePaymentCommandBuilder) WithDescription(description string) *CreatePaymentCommandBuilder {
	b.cmd.Description = &description
	return b
}

func (b *CreatePaymentCommandBuilder) WithTransactionId(transactionId string) *CreatePaymentCommandBuilder {
	b.cmd.TransactionId = &transactionId
	return b
}

func (b *CreatePaymentCommandBuilder) Build() CreatePaymentCommand {
	return b.cmd
}

func (b *CreatePaymentCommandBuilder) Validate() error {
	if b.cmd.AppointmentId <= 0 {
		return paymentDomain.NewPaymentError("INVALID_COMMAND", "appointment ID must be greater than 0", 0, "")
	}

	if b.cmd.OwnerId <= 0 {
		return paymentDomain.NewPaymentError("INVALID_COMMAND", "owner ID must be greater than 0", 0, "")
	}

	if b.cmd.Amount <= 0 {
		return paymentDomain.ErrInvalidAmount
	}

	if b.cmd.Currency == "" {
		return paymentDomain.ErrInvalidCurrency
	}

	if !b.cmd.PaymentMethod.IsValid() {
		return paymentDomain.ErrInvalidPaymentMethod
	}

	return nil
}
