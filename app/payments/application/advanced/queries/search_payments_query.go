package paymentQuery

import (
	"context"

	paymentDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/payments/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type SearchPaymentsQuery struct {
	UserId        *int          `json:"user_id,omitempty"`
	AppointmentId *int          `json:"appointment_id,omitempty"`
	PaymentMethod *string       `json:"payment_method,omitempty"`
	Status        *string       `json:"status,omitempty"`
	Pagination    page.PageData `json:"pagination"`
}

type SearchPaymentsQueryHandler interface {
	Handle(query SearchPaymentsQuery) (page.Page[[]PaymentResponse], error)
}

type SearchPaymentsHandler struct {
	repository paymentDomain.PaymentRepository
}

func NewSearchPaymentsHandler(repository paymentDomain.PaymentRepository) *SearchPaymentsHandler {
	return &SearchPaymentsHandler{repository: repository}
}

func (h *SearchPaymentsHandler) Handle(query SearchPaymentsQuery) (*page.Page[[]PaymentResponse], error) {
	criteria := make(map[string]interface{})

	if query.UserId != nil {
		criteria["owner_id"] = *query.UserId
	}
	if query.AppointmentId != nil {
		criteria["appointment_id"] = *query.AppointmentId
	}
	if query.PaymentMethod != nil {
		criteria["payment_method"] = *query.PaymentMethod
	}
	if query.Status != nil {
		criteria["status"] = *query.Status
	}

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
