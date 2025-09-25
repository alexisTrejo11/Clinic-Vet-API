package query

import (
	"time"

	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/page"
)

type FindPaymentByIDQuery struct {
	id            valueobject.PaymentID
	optCustomerID *valueobject.CustomerID
}

type FindPaymentsByDateRangeQuery struct {
	startDate time.Time
	endDate   time.Time
	page.PaginationRequest
}

func NewFindPaymentByIDQuery(idInt uint, optCustomerID *uint) *FindPaymentByIDQuery {
	paymentID := valueobject.NewPaymentID(idInt)

	var customerID *valueobject.CustomerID
	if optCustomerID != nil {
		customerIDObj := valueobject.NewCustomerID(*optCustomerID)
		customerID = &customerIDObj
	}

	return &FindPaymentByIDQuery{id: paymentID, optCustomerID: customerID}
}

func NewFindPaymentsByDateRangeQuery(startDate, endDate time.Time, pageData page.PaginationRequest) *FindPaymentsByDateRangeQuery {
	return &FindPaymentsByDateRangeQuery{
		startDate:         startDate,
		endDate:           endDate,
		PaginationRequest: pageData,
	}
}

type FindPaymentsByStatusQuery struct {
	status     enum.PaymentStatus
	pagination page.PaginationRequest
}

func NewFindPaymentsByStatusQuery(status string, pagination page.PaginationRequest) *FindPaymentsByStatusQuery {
	return &FindPaymentsByStatusQuery{
		status:     enum.PaymentStatus(status),
		pagination: pagination,
	}
}

type FindOverduePaymentsQuery struct {
	pagination page.PaginationRequest
}

func NewFindOverduePaymentsQuery(pagination page.PaginationRequest) FindOverduePaymentsQuery {
	return FindOverduePaymentsQuery{
		pagination: pagination,
	}
}

type FindPaymentsByCustomerQuery struct {
	customerID valueobject.CustomerID
	pagination page.PaginationRequest
}

func NewFindPaymentsByCustomerQuery(customerIDInt uint, pagination page.PaginationRequest) FindPaymentsByCustomerQuery {
	customerID := valueobject.NewCustomerID(customerIDInt)
	return FindPaymentsByCustomerQuery{
		customerID: customerID,
		pagination: pagination,
	}
}

type FindPaymentsBySpecification struct {
	spec specification.PaymentSpecification
}

func NewFindPaymentsBySpecification(spec specification.PaymentSpecification) *FindPaymentsBySpecification {
	return &FindPaymentsBySpecification{
		spec: spec,
	}
}
