-- name: FindPaymentByID :one
SELECT * FROM payments
WHERE id = $1 AND deleted_at IS NULL;

-- name: FindPaymentByTransactionID :one
SELECT * FROM payments
WHERE transaction_id = $1 AND deleted_at IS NULL;

-- name: FindPaymentByIDAndCustomerID :one
SELECT * FROM payments
WHERE id = $1 AND paid_by_customer_id = $2 AND deleted_at IS NULL;

-- name: FindPaymentsByCustomerID :many
SELECT * FROM payments
WHERE paid_by_customer_id = $1 AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: CountPaymentsByCustomerID :one
SELECT COUNT(*) FROM payments
WHERE paid_by_customer_id = $1 AND deleted_at IS NULL;

-- name: FindPaymentsByStatus :many
SELECT * FROM payments
WHERE status = $1 AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: CountPaymentsByStatus :one
SELECT COUNT(*) FROM payments
WHERE status = $1 AND deleted_at IS NULL;

-- name: FindOverduePayments :many
SELECT * FROM payments
WHERE due_date < CURRENT_TIMESTAMP 
AND status NOT IN ('paid', 'refunded', 'cancelled') 
AND deleted_at IS NULL
ORDER BY due_date ASC
LIMIT $1 OFFSET $2;

-- name: CountOverduePayments :one
SELECT COUNT(*) FROM payments
WHERE due_date < CURRENT_TIMESTAMP 
AND status NOT IN ('paid', 'refunded', 'cancelled') 
AND deleted_at IS NULL;

-- name: FindPaymentsByDateRange :many
SELECT * FROM payments
WHERE created_at BETWEEN $1 AND $2 AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $3 OFFSET $4;

-- name: CountPaymentsByDateRange :one
SELECT COUNT(*) FROM payments
WHERE created_at BETWEEN $1 AND $2 AND deleted_at IS NULL;

-- name: FindRecentPaymentsByCustomerID :many
SELECT * FROM payments
WHERE paid_by_customer_id = $1 AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $2;

-- name: FindPendingPayments :many
SELECT * FROM payments
WHERE status = 'pending' AND deleted_at IS NULL
ORDER BY due_date ASC
LIMIT $1 OFFSET $2;

-- name: FindSuccessfulPayments :many
SELECT * FROM payments
WHERE status = 'paid' AND deleted_at IS NULL
ORDER BY paid_at DESC
LIMIT $1 OFFSET $2;

-- name: ExistsPaymentByID :one
SELECT COUNT(*) > 0 FROM payments
WHERE id = $1 AND deleted_at IS NULL;

-- name: ExistsPaymentByTransactionID :one
SELECT COUNT(*) > 0 FROM payments
WHERE transaction_id = $1 AND deleted_at IS NULL;

-- name: ExistsPendingPaymentByCustomerID :one
SELECT COUNT(*) > 0 FROM payments
WHERE paid_by_customer_id = $1 
AND status = 'pending' 
AND deleted_at IS NULL;

-- name: CreatePayment :one
INSERT INTO payments (
    amount, currency, status, method, transaction_id, description,
    due_date, paid_at, refunded_at, paid_by_customer_id,
    med_session_id, invoice_id, refund_amount, failure_reason
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
) RETURNING *;

-- name: UpdatePayment :one
UPDATE payments SET
    amount = $2,
    currency = $3,
    status = $4,
    method = $5,
    transaction_id = $6,
    description = $7,
    due_date = $8,
    paid_at = $9,
    refunded_at = $10,
    paid_by_customer_id = $11,
    med_session_id = $12,
    invoice_id = $13,
    refund_amount = $14,
    failure_reason = $15,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SoftDeletePayment :exec
UPDATE payments SET
    is_active = FALSE,
    deleted_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: HardDeletePayment :exec
DELETE FROM payments WHERE id = $1;

-- name: TotalRevenueByDateRange :one
SELECT COALESCE(SUM(amount), 0) FROM payments
WHERE status = 'paid'
AND paid_at BETWEEN $1 AND $2
AND deleted_at IS NULL;

-- name: UpdatePaymentStatus :exec
UPDATE payments SET
    status = $2,
    paid_at = CASE WHEN $2 = 'paid' THEN COALESCE(paid_at, CURRENT_TIMESTAMP) ELSE paid_at END,
    refunded_at = CASE WHEN $2 = 'refunded' THEN COALESCE(refunded_at, CURRENT_TIMESTAMP) ELSE refunded_at END,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted_at IS NULL;

-- name: FindPaymentsByAppointmentID :one
SELECT * FROM payments
WHERE med_session_id = $1 AND deleted_at IS NULL;

-- name: FindPaymentsByInvoiceID :one
SELECT * FROM payments
WHERE invoice_id = $1 AND deleted_at IS NULL;