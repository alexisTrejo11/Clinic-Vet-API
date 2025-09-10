-- name: GetPaymentByID :one
SELECT *
FROM payments
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetPaymentByTransactionID :one
SELECT *
FROM payments
WHERE transaction_id = $1 AND deleted_at IS NULL;

-- name: ListPaymentsByCustomerID :many
SELECT *
FROM payments
WHERE paid_from_customer  = $1 AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: CountPaymentsByCustomerID :one
SELECT COUNT(*)
FROM payments
WHERE paid_from_customer  = $1 AND deleted_at IS NULL;

-- name: ListPaymentsByStatus :many
SELECT *
FROM payments
WHERE status = $1 AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: CountPaymentsByStatus :one
SELECT COUNT(*)
FROM payments
WHERE status = $1 AND deleted_at IS NULL;

-- name: ListPaymentsByDateRange :many
SELECT *
FROM payments
WHERE created_at BETWEEN $1 AND $2 AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $3 OFFSET $4;

-- name: CountPaymentsByDateRange :one
SELECT COUNT(*)
FROM payments
WHERE created_at BETWEEN $1 AND $2 AND deleted_at IS NULL;

-- name: ListOverduePayments :many
SELECT *
FROM payments
WHERE duedate < CURRENT_TIMESTAMP AND status != 'paid' AND deleted_at IS NULL
ORDER BY duedate ASC
LIMIT $1 OFFSET $2;

-- name: CountOverduePayments :one
SELECT COUNT (*)
WHERE duedate < CURRENT_TIMESTAMP AND status != 'paid' AND deleted_at IS NULL;

-- name: CreatePayment :one
INSERT INTO payments (
    amount,
    currency,
    status,
    method,
    transaction_id,
    description,
    duedate,
    paid_at,
    refunded_at,
    paid_from_customer,
    paid_to_employee,
    is_active,
    created_at,
    updated_at,
    deleted_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    TRUE,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    NULL
) RETURNING *;


-- name: UpdatePayment :one
UPDATE payments
SET
    amount = $1,
    currency = $2,
    status = $3,
    method = $4,
    transaction_id = $5,
    description = $6,
    duedate = $7,
    paid_at = $8,
    refunded_at = $9,
    is_active = $10,
    paid_from_customer = $11,
    paid_to_employee = $12,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $12
RETURNING *;

-- name: SoftDeletePayment :exec
UPDATE payments
SET
    is_active = FALSE,
    deleted_at = CURRENT_TIMESTAMP
WHERE id = $1 AND is_active = TRUE;
