package paymentQuery

import (
	"time"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	appError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/application"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListPaymentsByDateRangeQuery struct {
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	page.PageData `json:"page_data"`
}

type ListPaymentsByDateRangeQueryHandler interface {
	Handle(query ListPaymentsByDateRangeQuery) (*page.Page[[]PaymentResponse], error)
}

type listPaymentsByDateRangeHandler struct {
	paymentRepository paymentDomain.PaymentRepository
}

func NewListPaymentsByDateRangeHandler(paymentRepository paymentDomain.PaymentRepository) ListPaymentsByDateRangeQueryHandler {
	return &listPaymentsByDateRangeHandler{
		paymentRepository: paymentRepository,
	}
}

func (h *listPaymentsByDateRangeHandler) Handle(query ListPaymentsByDateRangeQuery) (*page.Page[[]PaymentResponse], error) {
	if query.StartDate.After(query.EndDate) {
		return nil, PaymentRangeDateErr(query.StartDate, query.EndDate)
	}

	paymentsPage, err := h.paymentRepository.ListPaymentsByDateRange(query.StartDate, query.EndDate)
	if err != nil {
		return nil, err
	}

	response := mapPaymentsToResponses(paymentsPage.Data)
	return page.NewPage(response, paymentsPage.Metadata), nil
}

func PaymentRangeDateErr(startDate, endDate time.Time) error {
	return appError.NewValidationError("start_date",
		"startdate:"+startDate.String()+"- enddate"+endDate.String(),
		"Start date cannot be after end date")
}
