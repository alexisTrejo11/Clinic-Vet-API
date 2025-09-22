package event

import (
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/service"
	"context"
	"sync"
	"time"

	"go.uber.org/zap"
)

type userEventHandler struct {
	service  service.UserAccountService
	wg       sync.WaitGroup
	eventLog *EventLogger
}

func NewUserEventHandler(service service.UserAccountService) *userEventHandler {
	return &userEventHandler{
		eventLog: NewEventLogger(),
		service:  service,
	}
}

func (h *userEventHandler) Registered(event UserRegisteredEvent) {
	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		h.processUserRegistration(event)
	}()
}

func (h *userEventHandler) processUserRegistration(event UserRegisteredEvent) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, ctxKey("user_id"), event.UserID)
	ctx = context.WithValue(ctx, ctxKey("user_email"), event.Email)
	ctx = context.WithValue(ctx, ctxKey("user_role"), event.Role)

	ctx = h.eventLog.LogEvent(ctx, "user_registered", "process_registration",
		zap.String("user_id", event.UserID.String()),
		zap.String("user_email", event.Email.String()),
		zap.String("user_role", event.Role.String()),
	)

	defer h.logProcessingTime(ctx, time.Now())

	switch event.Role {
	case enum.UserRoleCustomer:
		h.handleCustomerRegistration(ctx, event)
	case enum.UserRoleVeterinarian, enum.UserRoleReceptionist:
		h.handleEmployeeRegistration(ctx, event)
	default:
		h.eventLog.LogOperationError(ctx, nil, "Unknown user role",
			zap.String("received_role", event.Role.String()),
			zap.String("sub_operation", "role_validation"))
	}
}

func (h *userEventHandler) handleCustomerRegistration(ctx context.Context, event UserRegisteredEvent) {
	ctx = h.eventLog.LogEvent(ctx, "customer_registration", "process_customer",
		zap.String("parent_event_id", getContextString(ctx, "event_id")),
	)

	defer h.logProcessingTime(ctx, time.Now())

	if event.PersonalData != nil {
		customerID, err := h.service.CreateCustomer(ctx, event.UserID, *event.PersonalData)
		if err != nil {
			h.eventLog.LogOperationError(ctx, err, "Create customer failed")
			return
		}

		h.eventLog.LogOperationSuccess(ctx, "Customer created",
			zap.String("customer_id", customerID.String()))

		if err := h.service.AttachCustomerToUser(ctx, event.UserID, customerID); err != nil {
			h.eventLog.LogOperationError(ctx, err, "Attach customer failed",
				zap.String("customer_id", customerID.String()))
			return
		}
		h.eventLog.LogOperationSuccess(ctx, "Customer attached",
			zap.String("customer_id", customerID.String()))
	} else {
		h.eventLog.LogOperationWarning(ctx, "No personal data provided for customer creation")
	}

	if err := h.service.SendActivationEmail(ctx, event.Email, event.Name.FirstName); err != nil {
		h.eventLog.LogOperationWarning(ctx, "Send activation email failed",
			zap.String("email", event.Email.String()),
			zap.Error(err))
	} else {
		h.eventLog.LogOperationSuccess(ctx, "Activation email sent",
			zap.String("email", event.Email.String()))
	}

	h.eventLog.LogOperationSuccess(ctx, "Customer registration completed")
}

func (h *userEventHandler) handleEmployeeRegistration(ctx context.Context, event UserRegisteredEvent) {
	ctx = h.eventLog.LogEvent(ctx, "employee_registration", "process_employee",
		zap.String("parent_event_id", getContextString(ctx, "event_id")),
	)

	defer h.logProcessingTime(ctx, time.Now())

	if err := h.service.SendWelcomeEmail(ctx, event.Email, event.Name.FirstName); err != nil {
		h.eventLog.LogOperationWarning(ctx, "Send welcome email failed",
			zap.String("email", event.Email.String()),
			zap.Error(err))
	} else {
		h.eventLog.LogOperationSuccess(ctx, "Welcome email sent",
			zap.String("email", event.Email.String()))
	}

	h.eventLog.LogOperationSuccess(ctx, "Employee registration completed")
}

func (h *userEventHandler) logProcessingTime(ctx context.Context, startTime time.Time) {
	duration := time.Since(startTime)
	h.eventLog.LogOperationSuccess(ctx, "Processing completed",
		zap.Duration("processing_time", duration),
		zap.String("stage", "completed"))
}

func (h *userEventHandler) Wait() {
	h.eventLog.LogOperationSuccess(context.TODO(), "Waiting for event processing completion")
	h.wg.Wait()
	h.eventLog.LogOperationSuccess(context.TODO(), "All event processing completed")
}
