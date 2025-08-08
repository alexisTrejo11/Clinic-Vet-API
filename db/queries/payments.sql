-- name: createPayment: one
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
    is_active
    created_at,
    updated_at
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
    TRUE,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
)RETURNING *;

-- name: updatePayment: one
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
        updated_at = CURRENT_TIMESTAMP
    WHERE id = $11
RETURNING *;

-- name: softDeletePayment: exec
UPDATE payments
SET
    is_active = FALSE,
    deleted_at = CURRENT_TIMESTAMP
WHERE id = $1

-- name: getPaymentById: one
SELECT *
FROM payments
WHERE id = $1 AND deleted_at IS NULL;

-- name: getPaymentsByUserId: many
SELECT *
FROM payments
WHERE user_id = $1 AND is_active = TRUE AND deleted_at IS NULL;
