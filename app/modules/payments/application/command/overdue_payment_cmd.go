package command

import (
	"clinic-vet-api/app/modules/core/domain/entity/payment"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/shared/cqrs"
	"clinic-vet-api/app/shared/page"
	"context"
	"fmt"
)

type MarkOverduePaymentsCommand struct {
}

func NewMarkOverduePaymentsCommand() *MarkOverduePaymentsCommand {
	return &MarkOverduePaymentsCommand{}
}

func (h *paymentCommandHandler) MarkOverduePayments(ctx context.Context, cmd MarkOverduePaymentsCommand) cqrs.CommandResult {
	pagination := page.PaginationRequest{
		PageSize: DefaultPageSize,
		Page:     InitialPage,
	}

	var updatedCount int
	for {
		paymentsPage, err := h.paymentRepository.FindByStatus(ctx, enum.PaymentStatusOverdue, pagination)
		if err != nil {
			return *cqrs.FailureResult(ErrFailedSearchPayments, err)
		}

		payments := paymentsPage.Items
		if h.IsEmptyList(payments) {
			break
		}

		for _, payment := range payments {
			if err := h.UpdatePaymentOverdued(ctx, &payment); err != nil {
				fmt.Printf("Error updating payment %d: %v\n", payment.ID(), err)
				continue
			}
			updatedCount++
		}

		pagination.Page++

		if h.IsLastPage(pagination, paymentsPage.Metadata.TotalPages) {
			break
		}
	}

	return *cqrs.SuccessResult(fmt.Sprintf(MsgOverduePayments, updatedCount))
}

func (h *paymentCommandHandler) UpdatePaymentOverdued(ctx context.Context, payment *payment.Payment) error {
	if err := payment.MarkAsOverdue(ctx); err != nil {
		return err
	}

	if err := h.paymentRepository.Save(ctx, payment); err != nil {
		return err
	}

	return nil
}

// IsLastPage checks if the current pagination page is the last page
// Parameters:
//   - pagination: Current pagination request
//   - totalPages: Total number of pages available
//
// Returns:
//   - bool: True if current page is the last page
func (h *paymentCommandHandler) IsLastPage(pagination page.PaginationRequest, totalPages int32) bool {
	return pagination.Page >= totalPages
}

// IsEmptyList checks if the payments slice is empty
// Parameters:
//   - payments: Slice of payments to check
//
// Returns:
//   - bool: True if the slice is empty
func (h *paymentCommandHandler) IsEmptyList(payments []payment.Payment) bool {
	return len(payments) == 0
}
