-- name: CreateCustomer :one
INSERT INTO customers (
    photo, first_name, last_name, gender, 
    user_id, is_active, date_of_birth
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: GetCustomerByID :one
SELECT *
FROM customers
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetCustomerByPhone :one
SELECT *
FROM customers
WHERE phone_number = $1 AND deleted_at IS NULL;

-- name: GetCustomerByUserID :one
SELECT *
FROM customers
WHERE user_id = $1 AND deleted_at IS NULL;

-- name: UpdateCustomer :exec
UPDATE customers
SET 
    photo = $2, 
    first_name = $3, 
    last_name = $4, 
    gender = $5, 
    user_id = $6, 
    is_active = $7, 
    date_of_birth = $8,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted_at IS NULL;

-- name: SoftDeleteCustomer :exec
UPDATE customers
SET 
    is_active = FALSE,
    deleted_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: HardDeleteCustomer :exec
DELETE FROM customers WHERE id = $1;

-- name: ExistsCustomerByID :one
SELECT COUNT(*) > 0
FROM customers
WHERE id = $1 AND deleted_at IS NULL;

-- name: CountAllCustomers :one
SELECT COUNT(*)
FROM customers
WHERE deleted_at IS NULL;

-- name: CountActiveCustomers :one
SELECT COUNT(*)
FROM customers
WHERE is_active = TRUE AND deleted_at IS NULL;

-- name: FindActiveCustomers :many
SELECT *
FROM customers
WHERE is_active = TRUE AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: DeactivateCustomer :exec
UPDATE customers
SET 
    is_active = false, 
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted_at IS NULL;

-- name: ActivateCustomer :exec
UPDATE customers
SET 
    is_active = true, 
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted_at IS NULL;
