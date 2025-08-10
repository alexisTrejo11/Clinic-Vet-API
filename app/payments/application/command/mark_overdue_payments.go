package paymentCmd

import (
	"context"
	"errors"
	"fmt"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type MarkOverduePaymentsCommand struct {
}

type MarkOverduePaymentsHandler interface {
	Handle(ctx context.Context, command MarkOverduePaymentsCommand) shared.CommandResult
}

type markOverduePaymentsHandler struct {
	paymentRepo paymentDomain.PaymentRepository
}

func NewMarkOverduePaymentsHandler(paymentRepo paymentDomain.PaymentRepository) MarkOverduePaymentsHandler {
	return &markOverduePaymentsHandler{
		paymentRepo: paymentRepo,
	}
}

func (h *markOverduePaymentsHandler) Handle(ctx context.Context, command MarkOverduePaymentsCommand) shared.CommandResult {
	searchCriteria := map[string]interface{}{
		"status": paymentDomain.PENDING,
	}

	pagination := page.PageData{
		PageSize:   100,
		PageNumber: 1,
	}

	var updatedCount int
	for {
		paymentsPage, err := h.paymentRepo.Search(ctx, pagination, searchCriteria)
		if err != nil {
			return shared.FailureResult("failed to search payments", err)
		}

		payments := paymentsPage.Data
		if h.IsEmptyList(payments) {
			break
		}

		for _, payment := range payments {
			if err := h.UpdatePaymentOverdued(ctx, &payment); err != nil {
				fmt.Printf("Error updating payment %d: %v\n", payment.Id, err)
				continue
			}
			updatedCount++
		}

		pagination.PageNumber++

		if h.IsLastPage(pagination, paymentsPage.Metadata.TotalPages) {
			break
		}
	}

	return shared.SuccesResult("", fmt.Sprintf("Updated %d overdue payments", updatedCount))
}

func (h *markOverduePaymentsHandler) UpdatePaymentOverdued(ctx context.Context, payment *paymentDomain.Payment) error {
	if !payment.IsOverdue() {
		return errors.New("payment is not overdue")
	}
	payment.MarkAsOverdue()

	if err := h.paymentRepo.Save(ctx, payment); err != nil {
		return err
	}

	return nil
}

func (h *markOverduePaymentsHandler) IsLastPage(pagination page.PageData, totalPages int) bool {
	return pagination.PageNumber >= totalPages
}
func (h *markOverduePaymentsHandler) IsEmptyList(payments []paymentDomain.Payment) bool {
	return len(payments) == 0
}
