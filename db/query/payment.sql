-- name: CreatePayment :one
INSERT INTO payments (appointment_id, amount, payment_method, created_at, updated_at)
VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, appointment_id, amount, payment_method, created_at, updated_at;

-- name: GetPaymentByID :one
SELECT id, appointment_id, amount, payment_method, created_at, updated_at
FROM payments
WHERE id = $1;

-- name: ListPayments :many
SELECT id, appointment_id, amount, payment_method, created_at, updated_at
FROM payments
ORDER BY id;

-- name: UpdatePayment :exec
UPDATE payments
SET appointment_id = $2, amount = $3, payment_method = $4, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeletePayment :exec
DELETE FROM payments
WHERE id = $1;
