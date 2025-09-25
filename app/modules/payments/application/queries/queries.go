package query

import (
	"context"
	"time"

	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	apperror "clinic-vet-api/app/shared/error/application"
	"clinic-vet-api/app/shared/page"
)

type FindPaymentByIDQuery struct {
	id  valueobject.PaymentID
	ctx context.Context
}

type FindPaymentsByDateRangeQuery struct {
	startDate time.Time
	endDate   time.Time
	ctx       context.Context
	page.PaginationRequest
}

func NewFindPaymentByIDQuery(idInt uint) (*FindPaymentByIDQuery, error) {
	paymentID := valueobject.NewPaymentID(idInt)
	if paymentID.IsZero() {
		return nil, apperror.FieldValidationError("id", "0", "invalid payment id")
	}

	return &FindPaymentByIDQuery{id: paymentID}, nil
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
	ctx        context.Context
}

func NewFindPaymentsByStatusQuery(ctx context.Context, status string, pagination page.PaginationRequest) (*FindPaymentsByStatusQuery, error) {
	statusEnum, err := enum.ParsePaymentStatus(status)
	if err != nil {
		return nil, err
	}

	return &FindPaymentsByStatusQuery{
		status:     statusEnum,
		pagination: pagination,
		ctx:        ctx,
	}, nil
}

type FindOverduePaymentsQuery struct {
	pagination page.PaginationRequest
	ctx        context.Context
}

func NewFindOverduePaymentsQuery(ctx context.Context, pagination page.PaginationRequest) FindOverduePaymentsQuery {
	return FindOverduePaymentsQuery{
		pagination: pagination,
		ctx:        ctx,
	}
}

type FindPaymentsByCustomerQuery struct {
	customerID valueobject.CustomerID
	pagination page.PaginationRequest
	ctx        context.Context
}

func NewFindPaymentsByCustomerQuery(ctx context.Context, customerIDInt uint, pagination page.PaginationRequest) FindPaymentsByCustomerQuery {
	customerID := valueobject.NewCustomerID(customerIDInt)

	return FindPaymentsByCustomerQuery{
		customerID: customerID,
		pagination: pagination,
		ctx:        ctx,
	}
}

type FindPaymentsBySpecification struct {
	ctx  context.Context
	spec specification.PaymentSpecification
}

func NewFindPaymentsBySpecification(ctx context.Context, spec specification.PaymentSpecification) *FindPaymentsBySpecification {
	return &FindPaymentsBySpecification{
		ctx:  ctx,
		spec: spec,
	}
}
