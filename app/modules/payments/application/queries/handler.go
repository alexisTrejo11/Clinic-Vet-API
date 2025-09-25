package query

import (
	"time"

	"clinic-vet-api/app/modules/core/repository"
	apperror "clinic-vet-api/app/shared/error/application"
	"clinic-vet-api/app/shared/page"
)

type PaymentQueryHandler interface {
	FindByID(qry FindPaymentByIDQuery) (PaymentResult, error)
	FindOverdues(query FindOverduePaymentsQuery) (page.Page[PaymentResult], error)
	FindByStatus(query FindPaymentsByStatusQuery) (page.Page[PaymentResult], error)
	FindByCustomer(query FindPaymentsByCustomerQuery) (page.Page[PaymentResult], error)
	FindByDateRange(query FindPaymentsByDateRangeQuery) (page.Page[PaymentResult], error)
	FindBySpecification(query FindPaymentsBySpecification) (page.Page[PaymentResult], error)
}

type paymentQueryHandler struct {
	paymentRepository repository.PaymentRepository
}

func NewPaymentQueryHandler(paymentRepository repository.PaymentRepository) PaymentQueryHandler {
	return &paymentQueryHandler{
		paymentRepository: paymentRepository,
	}
}

func (h *paymentQueryHandler) FindByID(query FindPaymentByIDQuery) (PaymentResult, error) {
	payment, err := h.paymentRepository.FindByID(query.ctx, query.id)
	if err != nil {
		return PaymentResult{}, err
	}

	return entityToResult(payment), nil
}

func (h *paymentQueryHandler) FindOverdues(query FindOverduePaymentsQuery) (page.Page[PaymentResult], error) {
	paymentsPage, err := h.paymentRepository.FindOverdue(query.ctx, query.pagination)
	if err != nil {
		return page.Page[PaymentResult]{}, err
	}

	return page.MapItems(paymentsPage, entityToResult), nil
}

func (h *paymentQueryHandler) FindByStatus(query FindPaymentsByStatusQuery) (page.Page[PaymentResult], error) {
	paymentPage, err := h.paymentRepository.FindByStatus(query.ctx, query.status, query.pagination)
	if err != nil {
		return page.Page[PaymentResult]{}, err
	}

	return page.MapItems(paymentPage, entityToResult), nil
}

func (h *paymentQueryHandler) FindByCustomer(query FindPaymentsByCustomerQuery) (page.Page[PaymentResult], error) {
	paymentsPage, err := h.paymentRepository.FindByCustomerID(query.ctx, query.customerID, query.pagination)
	if err != nil {
		return page.Page[PaymentResult]{}, err
	}

	return page.MapItems(paymentsPage, entityToResult), nil
}

func (h *paymentQueryHandler) FindByDateRange(query FindPaymentsByDateRangeQuery) (page.Page[PaymentResult], error) {
	if query.startDate.After(query.endDate) {
		return page.Page[PaymentResult]{}, PaymentRangeDateErr(query.startDate, query.endDate)
	}

	paymentsPage, err := h.paymentRepository.FindByDateRange(query.ctx, query.startDate, query.endDate, query.PaginationRequest)
	if err != nil {
		return page.Page[PaymentResult]{}, err
	}

	return page.MapItems(paymentsPage, entityToResult), nil
}

func (h *paymentQueryHandler) FindBySpecification(query FindPaymentsBySpecification) (page.Page[PaymentResult], error) {
	paymentsPage, err := h.paymentRepository.FindBySpecification(query.ctx, query.spec)
	if err != nil {
		return page.Page[PaymentResult]{}, err
	}

	return page.MapItems(paymentsPage, entityToResult), nil
}

func PaymentRangeDateErr(startDate, endDate time.Time) error {
	return apperror.FieldValidationError("start_date",
		"start date:"+startDate.String()+"- end date"+endDate.String(),
		"Start date cannot be after end date")
}
