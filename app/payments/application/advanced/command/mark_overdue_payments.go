package paymentCmd

import (
	"context"
	"errors"
	"fmt"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type MarkOverduePaymentsCommand struct {
	CTX context.Context `json:"-"`
}

type MarkOverduePaymentsHandler interface {
	Handle(command MarkOverduePaymentsCommand) (int, error)
}

type markOverduePaymentsHandler struct {
	paymentRepo paymentDomain.PaymentRepository
}

func NewMarkOverduePaymentsHandler(paymentRepo paymentDomain.PaymentRepository) MarkOverduePaymentsHandler {
	return &markOverduePaymentsHandler{
		paymentRepo: paymentRepo,
	}
}

func (h *markOverduePaymentsHandler) Handle(command MarkOverduePaymentsCommand) (int, error) {
	searchCriteria := map[string]interface{}{
		"status": paymentDomain.PENDING,
	}

	pagination := page.PageData{
		PageSize:   100,
		PageNumber: 1,
	}

	var updatedCount int
	for {
		paymentsPage, err := h.paymentRepo.Search(command.CTX, pagination, searchCriteria)
		if err != nil {
			return updatedCount, err
		}

		payments := paymentsPage.Data
		if h.IsEmptyList(payments) {
			break
		}

		for _, payment := range payments {
			if err := h.UpdatePaymentOverdued(&payment); err != nil {
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

	return updatedCount, nil
}

func (h *markOverduePaymentsHandler) UpdatePaymentOverdued(payment *paymentDomain.Payment) error {
	if !payment.IsOverdue() {
		return errors.New("payment is not overdue")
	}
	payment.MarkAsOverdue()

	if err := h.paymentRepo.Save(payment); err != nil {
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
