package specification

import (
	"fmt"
	"strings"
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/payment"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

// PaymentSpecification implementa la interfaz Specification para Payment
type PaymentSpecification struct {
	conditions []Condition
}

type Condition struct {
	Field    string
	Operator string
	Value    interface{}
}

// Constructores para diferentes tipos de filtros
func NewPaymentSpecification() *PaymentSpecification {
	return &PaymentSpecification{
		conditions: make([]Condition, 0),
	}
}

// Métodos para agregar condiciones específicas de Payment
func (ps *PaymentSpecification) WithID(id valueobject.PaymentID) *PaymentSpecification {
	ps.conditions = append(ps.conditions, Condition{
		Field:    "id",
		Operator: "=",
		Value:    id,
	})
	return ps
}

func (ps *PaymentSpecification) WithAppointmentID(appointmentID valueobject.AppointmentID) *PaymentSpecification {
	ps.conditions = append(ps.conditions, Condition{
		Field:    "appointment_id",
		Operator: "=",
		Value:    appointmentID,
	})
	return ps
}

func (ps *PaymentSpecification) WithUserID(userID valueobject.UserID) *PaymentSpecification {
	ps.conditions = append(ps.conditions, Condition{
		Field:    "user_id",
		Operator: "=",
		Value:    userID,
	})
	return ps
}

func (ps *PaymentSpecification) WithStatus(status enum.PaymentStatus) *PaymentSpecification {
	ps.conditions = append(ps.conditions, Condition{
		Field:    "status",
		Operator: "=",
		Value:    status,
	})
	return ps
}

func (ps *PaymentSpecification) WithStatusIn(statuses []enum.PaymentStatus) *PaymentSpecification {
	ps.conditions = append(ps.conditions, Condition{
		Field:    "status",
		Operator: "IN",
		Value:    statuses,
	})
	return ps
}

func (ps *PaymentSpecification) WithPaymentMethod(method enum.PaymentMethod) *PaymentSpecification {
	ps.conditions = append(ps.conditions, Condition{
		Field:    "payment_method",
		Operator: "=",
		Value:    method,
	})
	return ps
}

func (ps *PaymentSpecification) WithAmountRange(minAmount, maxAmount valueobject.Money) *PaymentSpecification {
	if minAmount.Amount().IsPositive() {
		ps.conditions = append(ps.conditions, Condition{
			Field:    "amount",
			Operator: ">=",
			Value:    minAmount,
		})
	}
	if maxAmount.Amount().IsPositive() {
		ps.conditions = append(ps.conditions, Condition{
			Field:    "amount",
			Operator: "<=",
			Value:    maxAmount,
		})
	}
	return ps
}

func (ps *PaymentSpecification) WithCurrency(currency string) *PaymentSpecification {
	ps.conditions = append(ps.conditions, Condition{
		Field:    "currency",
		Operator: "=",
		Value:    currency,
	})
	return ps
}

func (ps *PaymentSpecification) WithTransactionID(transactionID string) *PaymentSpecification {
	ps.conditions = append(ps.conditions, Condition{
		Field:    "transaction_id",
		Operator: "=",
		Value:    transactionID,
	})
	return ps
}

func (ps *PaymentSpecification) WithDueDateRange(from, to time.Time) *PaymentSpecification {
	if !from.IsZero() {
		ps.conditions = append(ps.conditions, Condition{
			Field:    "due_date",
			Operator: ">=",
			Value:    from,
		})
	}
	if !to.IsZero() {
		ps.conditions = append(ps.conditions, Condition{
			Field:    "due_date",
			Operator: "<=",
			Value:    to,
		})
	}
	return ps
}

func (ps *PaymentSpecification) WithPaidAtRange(from, to time.Time) *PaymentSpecification {
	if !from.IsZero() {
		ps.conditions = append(ps.conditions, Condition{
			Field:    "paid_at",
			Operator: ">=",
			Value:    from,
		})
	}
	if !to.IsZero() {
		ps.conditions = append(ps.conditions, Condition{
			Field:    "paid_at",
			Operator: "<=",
			Value:    to,
		})
	}
	return ps
}

func (ps *PaymentSpecification) WithRefundedAtRange(from, to time.Time) *PaymentSpecification {
	if !from.IsZero() {
		ps.conditions = append(ps.conditions, Condition{
			Field:    "refunded_at",
			Operator: ">=",
			Value:    from,
		})
	}
	if !to.IsZero() {
		ps.conditions = append(ps.conditions, Condition{
			Field:    "refunded_at",
			Operator: "<=",
			Value:    to,
		})
	}
	return ps
}

func (ps *PaymentSpecification) WithDescriptionLike(description string) *PaymentSpecification {
	ps.conditions = append(ps.conditions, Condition{
		Field:    "description",
		Operator: "LIKE",
		Value:    "%" + description + "%",
	})
	return ps
}

// Métodos para condiciones de nullabilidad
func (ps *PaymentSpecification) WithoutTransactionID() *PaymentSpecification {
	ps.conditions = append(ps.conditions, Condition{
		Field:    "transaction_id",
		Operator: "IS NULL",
		Value:    nil,
	})
	return ps
}

func (ps *PaymentSpecification) WithTransactionIDNotNull() *PaymentSpecification {
	ps.conditions = append(ps.conditions, Condition{
		Field:    "transaction_id",
		Operator: "IS NOT NULL",
		Value:    nil,
	})
	return ps
}

func (ps *PaymentSpecification) WithoutPaidAt() *PaymentSpecification {
	ps.conditions = append(ps.conditions, Condition{
		Field:    "paid_at",
		Operator: "IS NULL",
		Value:    nil,
	})
	return ps
}

func (ps *PaymentSpecification) WithPaidAtNotNull() *PaymentSpecification {
	ps.conditions = append(ps.conditions, Condition{
		Field:    "paid_at",
		Operator: "IS NOT NULL",
		Value:    nil,
	})
	return ps
}

// Implementación de la interfaz Specification
func (ps *PaymentSpecification) IsSatisfiedBy(entity any) bool {
	payment, ok := entity.(*payment.Payment)
	if !ok {
		return false
	}

	for _, condition := range ps.conditions {
		if !ps.evaluateCondition(payment, condition) {
			return false
		}
	}
	return true
}

func (ps *PaymentSpecification) ToSQL() (string, []any) {
	if len(ps.conditions) == 0 {
		return "1=1", []any{}
	}

	var clauses []string
	var params []any
	paramIndex := 1

	for _, condition := range ps.conditions {
		clause, param := ps.buildSQLClause(condition, &paramIndex)
		clauses = append(clauses, clause)
		if param != nil {
			params = append(params, param...)
		}
	}

	return strings.Join(clauses, " AND "), params
}

func (ps *PaymentSpecification) buildSQLClause(condition Condition, paramIndex *int) (string, []any) {
	switch condition.Operator {
	case "=", "!=", "<", "<=", ">", ">=", "LIKE":
		placeholder := fmt.Sprintf("$%d", *paramIndex)
		*paramIndex++
		return fmt.Sprintf("%s %s %s", condition.Field, condition.Operator, placeholder), []any{condition.Value}

	case "IN":
		values, ok := condition.Value.([]enum.PaymentStatus)
		if !ok {
			// Manejar otros tipos de slice si es necesario
			return "1=1", nil
		}

		var placeholders []string
		var params []any
		for _, v := range values {
			placeholders = append(placeholders, fmt.Sprintf("$%d", *paramIndex))
			params = append(params, v)
			*paramIndex++
		}

		return fmt.Sprintf("%s IN (%s)", condition.Field, strings.Join(placeholders, ",")), params

	case "IS NULL", "IS NOT NULL":
		return fmt.Sprintf("%s %s", condition.Field, condition.Operator), nil

	default:
		return "1=1", nil
	}
}

func (ps *PaymentSpecification) evaluateCondition(payment *payment.Payment, condition Condition) bool {
	switch condition.Field {
	case "id":
		return ps.compareValues(payment.ID(), condition.Value, condition.Operator)
	case "appointment_id":
		return ps.compareValues(payment.AppointmentID(), condition.Value, condition.Operator)
	case "customer_id":
		return ps.compareValues(payment.PaidFromCustomer(), condition.Value, condition.Operator)
	case "amount":
		return ps.compareValues(payment.Amount(), condition.Value, condition.Operator)
	case "currency":
		return ps.compareValues(payment.Currency(), condition.Value, condition.Operator)
	case "payment_method":
		return ps.compareValues(payment.Method(), condition.Value, condition.Operator)
	case "status":
		if condition.Operator == "IN" {
			statuses, ok := condition.Value.([]enum.PaymentStatus)
			if !ok {
				return false
			}
			for _, status := range statuses {
				if payment.Status() == status {
					return true
				}
			}
			return false
		}
		return ps.compareValues(payment.Status(), condition.Value, condition.Operator)
	case "transaction_id":
		return ps.comparePointerValues(payment.TransactionID(), condition.Value, condition.Operator)
	case "description":
		return ps.comparePointerValues(payment.Description(), condition.Value, condition.Operator)
	case "due_date":
		return ps.compareTimePointerValues(payment.DueDate(), condition.Value, condition.Operator)
	case "paid_at":
		return ps.compareTimePointerValues(payment.PaidAt(), condition.Value, condition.Operator)
	case "refunded_at":
		return ps.compareTimePointerValues(payment.RefundedAt(), condition.Value, condition.Operator)
	}
	return false
}

func (ps *PaymentSpecification) compareValues(actual, expected any, operator string) bool {
	switch operator {
	case "=":
		return actual == expected
	case "!=":
		return actual != expected
	// Agregar más operadores según necesidad
	default:
		return false
	}
}

func (ps *PaymentSpecification) comparePointerValues(actual *string, expected any, operator string) bool {
	switch operator {
	case "IS NULL":
		return actual == nil
	case "IS NOT NULL":
		return actual != nil
	case "=":
		if actual == nil {
			return expected == nil
		}
		expectedStr, ok := expected.(string)
		if !ok {
			return false
		}
		return *actual == expectedStr
	case "LIKE":
		if actual == nil {
			return false
		}
		expectedStr, ok := expected.(string)
		if !ok {
			return false
		}
		return strings.Contains(*actual, strings.Trim(expectedStr, "%"))
	}
	return false
}

func (ps *PaymentSpecification) compareTimePointerValues(actual *time.Time, expected any, operator string) bool {
	switch operator {
	case "IS NULL":
		return actual == nil
	case "IS NOT NULL":
		return actual != nil
	case "=":
		if actual == nil {
			return expected == nil
		}
		expectedTime, ok := expected.(time.Time)
		if !ok {
			return false
		}
		return actual.Equal(expectedTime)
	case ">=":
		if actual == nil {
			return false
		}
		expectedTime, ok := expected.(time.Time)
		if !ok {
			return false
		}
		return actual.After(expectedTime) || actual.Equal(expectedTime)
	case "<=":
		if actual == nil {
			return false
		}
		expectedTime, ok := expected.(time.Time)
		if !ok {
			return false
		}
		return actual.Before(expectedTime) || actual.Equal(expectedTime)
	}
	return false
}

// PaymentQuery combina specification con pagination
type PaymentQuery struct {
	Specification *PaymentSpecification
	Pagination    *Pagination
}

func NewPaymentQuery() *PaymentQuery {
	return &PaymentQuery{
		Specification: NewPaymentSpecification(),
		Pagination:    nil,
	}
}

func (pq *PaymentQuery) WithSpecification(spec *PaymentSpecification) *PaymentQuery {
	pq.Specification = spec
	return pq
}

func (pq *PaymentQuery) WithPagination(pagination *Pagination) *PaymentQuery {
	pq.Pagination = pagination
	return pq
}

func (pq *PaymentQuery) ToSQL() (string, []any, string) {
	whereClause, params := pq.Specification.ToSQL()

	orderClause := ""
	if pq.Pagination != nil && pq.Pagination.OrderBy != "" {
		sortDir := "ASC"
		if pq.Pagination.SortDir == "DESC" {
			sortDir = "DESC"
		}
		orderClause = fmt.Sprintf("ORDER BY %s %s", pq.Pagination.OrderBy, sortDir)

		if pq.Pagination.PageSize > 0 {
			orderClause += fmt.Sprintf(" LIMIT %d OFFSET %d",
				pq.Pagination.GetLimit(),
				pq.Pagination.GetOffset())
		}
	}

	return whereClause, params, orderClause
}
