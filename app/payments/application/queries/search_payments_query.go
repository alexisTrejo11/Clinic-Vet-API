package paymentQuery

import (
	"context"
	"time"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type SearchPaymentsQuery struct {
	UserId        *int                         `json:"owner_id,omitempty"`
	AppointmentId *int                         `json:"appointment_id,omitempty"`
	Status        *paymentDomain.PaymentStatus `json:"status,omitempty"`
	PaymentMethod *paymentDomain.PaymentMethod `json:"payment_method,omitempty"`
	MinAmount     *float64                     `json:"min_amount,omitempty"`
	MaxAmount     *float64                     `json:"max_amount,omitempty"`
	Currency      *string                      `json:"currency,omitempty"`
	StartDate     *time.Time                   `json:"start_date,omitempty"`
	EndDate       *time.Time                   `json:"end_date,omitempty"`
	Pagination    page.PageData                `json:"pagination"`
}

type SearchPaymentsQueryHandler interface {
	Handle(ctx context.Context, query SearchPaymentsQuery) (page.Page[[]PaymentResponse], error)
}

type SearchPaymentsHandler struct {
	repository paymentDomain.PaymentRepository
}

func NewSearchPaymentsHandler(repository paymentDomain.PaymentRepository) *SearchPaymentsHandler {
	return &SearchPaymentsHandler{repository: repository}
}

func (h *SearchPaymentsHandler) Handle(ctx context.Context, query SearchPaymentsQuery) (*page.Page[[]PaymentResponse], error) {
	criteria := toSearchCriteria(query)
	paymentsPage, err := h.repository.Search(context.Background(), query.Pagination, criteria)
	if err != nil {
		return nil, err
	}

	var responses []PaymentResponse
	for _, payment := range paymentsPage.Data {
		responses = append(responses, NewPaymentResponse(&payment))
	}

	return page.NewPage(responses, paymentsPage.Metadata), nil
}

func toSearchCriteria(query SearchPaymentsQuery) map[string]interface{} {
	criteria := make(map[string]interface{})
	if query.UserId != nil {
		criteria["owner_id"] = *query.UserId
	}
	if query.AppointmentId != nil {
		criteria["appointment_id"] = *query.AppointmentId
	}
	if query.Status != nil {
		criteria["status"] = *query.Status
	}
	if query.PaymentMethod != nil {
		criteria["payment_method"] = *query.PaymentMethod
	}
	if query.MinAmount != nil {
		criteria["min_amount"] = *query.MinAmount
	}
	if query.MaxAmount != nil {
		criteria["max_amount"] = *query.MaxAmount
	}
	if query.Currency != nil {
		criteria["currency"] = *query.Currency
	}
	if query.StartDate != nil {
		criteria["start_date"] = *query.StartDate
	}
	if query.EndDate != nil {
		criteria["end_date"] = *query.EndDate
	}
	return criteria
}
