package query

import (
	"context"
	"time"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	ApplicationError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/application"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListPaymentsByDateRangeQuery struct {
	startDate time.Time
	endDate   time.Time
	page.PageData
}

func NewListPaymentsByDateRangeQuery(startDate, endDate time.Time, pageData page.PageData) ListPaymentsByDateRangeQuery {
	return ListPaymentsByDateRangeQuery{
		startDate: startDate,
		endDate:   endDate,
		PageData:  pageData,
	}
}

type ListPaymentsByDateRangeQueryHandler interface {
	Handle(ctx context.Context, query ListPaymentsByDateRangeQuery) (page.Page[[]PaymentResponse], error)
}

type listPaymentsByDateRangeHandler struct {
	paymentRepository repository.PaymentRepository
}

func NewListPaymentsByDateRangeHandler(paymentRepository repository.PaymentRepository) ListPaymentsByDateRangeQueryHandler {
	return &listPaymentsByDateRangeHandler{
		paymentRepository: paymentRepository,
	}
}

func (h *listPaymentsByDateRangeHandler) Handle(ctx context.Context, query ListPaymentsByDateRangeQuery) (page.Page[[]PaymentResponse], error) {
	if query.startDate.After(query.endDate) {
		return page.Page[[]PaymentResponse]{}, PaymentRangeDateErr(query.startDate, query.endDate)
	}

	paymentsPage, err := h.paymentRepository.ListPaymentsByDateRange(ctx, query.startDate, query.endDate, query.PageData)
	if err != nil {
		return page.Page[[]PaymentResponse]{}, err
	}

	response := mapPaymentsToResponses(paymentsPage.Data)
	return page.NewPage(response, paymentsPage.Metadata), nil
}

func PaymentRangeDateErr(startDate, endDate time.Time) error {
	return ApplicationError.NewValidationError("start_date",
		"startdate:"+startDate.String()+"- enddate"+endDate.String(),
		"Start date cannot be after end date")
}
