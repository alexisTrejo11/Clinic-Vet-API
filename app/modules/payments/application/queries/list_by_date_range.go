package query

import (
	"context"
	"time"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListPaymentsByDateRangeQuery struct {
	startDate time.Time
	endDate   time.Time
	page.PageInput
}

func NewListPaymentsByDateRangeQuery(startDate, endDate time.Time, pageData page.PageInput) ListPaymentsByDateRangeQuery {
	return ListPaymentsByDateRangeQuery{
		startDate: startDate,
		endDate:   endDate,
		PageInput: pageData,
	}
}

type ListPaymentsByDateRangeQueryHandler interface {
	Handle(ctx context.Context, query ListPaymentsByDateRangeQuery) (page.Page[[]PaymentResponse], error)
}

type ListPaymentsByDateRangeHandler struct {
	paymentRepository repository.PaymentRepository
}

func NewListPaymentsByDateRangeHandler(paymentRepository repository.PaymentRepository) ListPaymentsByDateRangeQueryHandler {
	return &ListPaymentsByDateRangeHandler{
		paymentRepository: paymentRepository,
	}
}

func (h *ListPaymentsByDateRangeHandler) Handle(ctx context.Context, query ListPaymentsByDateRangeQuery) (page.Page[[]PaymentResponse], error) {
	if query.startDate.After(query.endDate) {
		return page.Page[[]PaymentResponse]{}, PaymentRangeDateErr(query.startDate, query.endDate)
	}

	paymentsPage, err := h.paymentRepository.ListPaymentsByDateRange(ctx, query.startDate, query.endDate, query.PageInput)
	if err != nil {
		return page.Page[[]PaymentResponse]{}, err
	}

	response := mapPaymentsToResponses(paymentsPage.Data)
	return page.NewPage(response, paymentsPage.Metadata), nil
}

func PaymentRangeDateErr(startDate, endDate time.Time) error {
	return apperror.FieldValidationError("start_date",
		"startdate:"+startDate.String()+"- enddate"+endDate.String(),
		"Start date cannot be after end date")
}
