package query

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type SearchPaymentsQuery struct {
	UserID        *int                `json:"owner_id,omitempty"`
	AppointmentID *int                `json:"appointment_id,omitempty"`
	Status        *enum.PaymentStatus `json:"status,omitempty"`
	PaymentMethod *enum.PaymentMethod `json:"payment_method,omitempty"`
	MinAmount     *float64            `json:"min_amount,omitempty"`
	MaxAmount     *float64            `json:"max_amount,omitempty"`
	Currency      *string             `json:"currency,omitempty"`
	StartDate     *time.Time          `json:"start_date,omitempty"`
	EndDate       *time.Time          `json:"end_date,omitempty"`
	Pagination    page.PageInput      `json:"pagination"`
	ctx           context.Context
}

type SearchPaymentsHandler struct {
	repository repository.PaymentRepository
}

func NewSearchPaymentsHandler(repository repository.PaymentRepository) cqrs.QueryHandler[page.Page[[]PaymentResponse]] {
	return &SearchPaymentsHandler{repository: repository}
}

func (h *SearchPaymentsHandler) Handle(q cqrs.Query) (page.Page[[]PaymentResponse], error) {
	query := q.(SearchPaymentsQuery)

	criteria := toSearchCriteria(query)
	paymentsPage, err := h.repository.Search(context.Background(), query.Pagination, criteria)
	if err != nil {
		return page.Page[[]PaymentResponse]{}, err
	}

	var responses []PaymentResponse
	for _, payment := range paymentsPage.Data {
		responses = append(responses, NewPaymentResponse(&payment))
	}

	return page.NewPage(responses, paymentsPage.Metadata), nil
}

func toSearchCriteria(query SearchPaymentsQuery) map[string]any {
	criteria := make(map[string]any)
	if query.UserID != nil {
		criteria["owner_id"] = *query.UserID
	}
	if query.AppointmentID != nil {
		criteria["appointment_id"] = *query.AppointmentID
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
