package command

import (
	"context"
	"fmt"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type MarkOverduePaymentsCommand struct {
	context context.Context
}

type MarkOverduePaymentsHandler struct {
	paymentRepo repository.PaymentRepository
}

func NewMarkOverduePaymentsHandler(paymentRepo repository.PaymentRepository) cqrs.CommandHandler {
	return &MarkOverduePaymentsHandler{
		paymentRepo: paymentRepo,
	}
}

func (h *MarkOverduePaymentsHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command := cmd.(MarkOverduePaymentsCommand)

	searchCriteria := map[string]any{
		"status": enum.PENDING,
	}

	pagination := page.PageInput{
		PageSize:   100,
		PageNumber: 1,
	}

	var updatedCount int
	for {
		paymentsPage, err := h.paymentRepo.Search(command.context, pagination, searchCriteria)
		if err != nil {
			return cqrs.FailureResult("failed to search payments", err)
		}

		payments := paymentsPage.Data
		if h.IsEmptyList(payments) {
			break
		}

		for _, payment := range payments {
			if err := h.UpdatePaymentOverdued(command.context, &payment); err != nil {
				fmt.Printf("Error updating payment %d: %v\n", payment.GetID(), err)
				continue
			}
			updatedCount++
		}

		pagination.PageNumber++

		if h.IsLastPage(pagination, paymentsPage.Metadata.TotalPages) {
			break
		}
	}

	return cqrs.SuccessResult("", fmt.Sprintf("Updated %d overdue payments", updatedCount))
}

func (h *MarkOverduePaymentsHandler) UpdatePaymentOverdued(ctx context.Context, payment *entity.Payment) error {
	if err := payment.Overdue(); err != nil {
		return err
	}

	if err := h.paymentRepo.Save(ctx, payment); err != nil {
		return err
	}

	return nil
}

func (h *MarkOverduePaymentsHandler) IsLastPage(pagination page.PageInput, totalPages int) bool {
	return pagination.PageNumber >= totalPages
}

func (h *MarkOverduePaymentsHandler) IsEmptyList(payments []entity.Payment) bool {
	return len(payments) == 0
}
