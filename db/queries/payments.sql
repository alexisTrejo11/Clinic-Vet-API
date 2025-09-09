-- name: GetPaymentById :one
SELECT *
FROM payments
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetPaymentByTransactionId :one
SELECT *
FROM payments
WHERE transaction_id = $1 AND deleted_at IS NULL;

-- name: ListPaymentsByUserId :many
SELECT *
FROM payments
WHERE user_id = $1 AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: CountPaymentsByUserId :one
SELECT COUNT(*)
FROM payments
WHERE user_id = $1 AND deleted_at IS NULL;

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
    user_id,
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
    user_id = $11,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $12
RETURNING *;

-- name: SoftDeletePayment :exec
UPDATE payments
SET
    is_active = FALSE,
    deleted_at = CURRENT_TIMESTAMP
WHERE id = $1 AND is_active = TRUE;
