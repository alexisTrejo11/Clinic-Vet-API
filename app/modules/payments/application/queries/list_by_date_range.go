package query

import (
	"context"
	"errors"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type FindPaymentsByDateRangeQuery struct {
	startDate time.Time
	endDate   time.Time
	ctx       context.Context
	page.PageInput
}

func NewFindPaymentsByDateRangeQuery(startDate, endDate time.Time, pageData page.PageInput) *FindPaymentsByDateRangeQuery {
	return &FindPaymentsByDateRangeQuery{
		startDate: startDate,
		endDate:   endDate,
		PageInput: pageData,
	}
}

type FindPaymentsByDateRangeHandler struct {
	paymentRepository repository.PaymentRepository
}

func NewFindPaymentsByDateRangeHandler(paymentRepository repository.PaymentRepository) cqrs.QueryHandler[page.Page[PaymentResponse]] {
	return &FindPaymentsByDateRangeHandler{
		paymentRepository: paymentRepository,
	}
}

func (h *FindPaymentsByDateRangeHandler) Handle(q cqrs.Query) (page.Page[PaymentResponse], error) {
	query, valid := q.(FindPaymentsByDateRangeQuery)
	if !valid {
		return page.Page[PaymentResponse]{}, errors.New("invalid query type")
	}
	if query.startDate.After(query.endDate) {
		return page.Page[PaymentResponse]{}, PaymentRangeDateErr(query.startDate, query.endDate)
	}

	paymentsPage, err := h.paymentRepository.FindByDateRange(query.ctx, query.startDate, query.endDate, query.PageInput)
	if err != nil {
		return page.Page[PaymentResponse]{}, err
	}

	response := mapPaymentsToResponses(paymentsPage.Items)
	return page.NewPage(response, paymentsPage.Metadata), nil
}

func PaymentRangeDateErr(startDate, endDate time.Time) error {
	return apperror.FieldValidationError("start_date",
		"start date:"+startDate.String()+"- end date"+endDate.String(),
		"Start date cannot be after end date")
}
