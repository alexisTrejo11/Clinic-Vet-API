-- name: CreateUser :one
INSERT INTO users (
    email, 
    phone_number, 
    password, 
    status, 
    role, 
    created_at,
    updated_at
)
VALUES (
    $1, $2, $3, $4, $5,
    CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET 
    email = $2,
    phone_number = $3,
    password = $4,
    status = $5,
    role = $6,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: FindUserByID :one
SELECT 
    u.*,
    c.id as customer_id,
    e.id as employee_id,
    CASE 
        WHEN c.id IS NOT NULL THEN 'customer'
        WHEN e.id IS NOT NULL THEN 'employee'
        ELSE 'user'
    END as user_type
FROM users u
LEFT JOIN customers c ON u.id = c.user_id AND c.deleted_at IS NULL
LEFT JOIN employees e ON u.id = e.user_id AND e.deleted_at IS NULL
WHERE u.id = $1 
AND u.deleted_at IS NULL;

-- name: FindUserByEmail :one
SELECT 
    u.*,
    c.id as customer_id,
    e.id as employee_id,
    CASE 
        WHEN c.id IS NOT NULL THEN 'customer'
        WHEN e.id IS NOT NULL THEN 'employee'
        ELSE 'user'
    END as user_type
FROM users u
LEFT JOIN customers c ON u.id = c.user_id AND c.deleted_at IS NULL
LEFT JOIN employees e ON u.id = e.user_id AND e.deleted_at IS NULL
WHERE u.email = $1
AND u.deleted_at IS NULL;

-- name: FindUserByPhoneNumber :one
SELECT 
    u.*,
    c.id as customer_id,
    e.id as employee_id,
    CASE 
        WHEN c.id IS NOT NULL THEN 'customer'
        WHEN e.id IS NOT NULL THEN 'employee'
        ELSE 'user'
    END as user_type
FROM users u
LEFT JOIN customers c ON u.id = c.user_id AND c.deleted_at IS NULL
LEFT JOIN employees e ON u.id = e.user_id AND e.deleted_at IS NULL
WHERE u.phone_number = $1
AND u.deleted_at IS NULL;

-- name: SoftDeleteUser :exec
UPDATE users
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: HardDeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: RestoreUser :exec
UPDATE users
SET deleted_at = NULL
WHERE id = $1;

-- name: UpdateUserLastLogin :exec
UPDATE users
SET last_login = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: ExistsUserByEmail :one
SELECT COUNT(*) > 0
FROM users
WHERE email = $1;

-- name: ExistsUserByPhoneNumber :one
SELECT COUNT(*) > 0
FROM users
WHERE phone_number = $1;

-- name: ExistsUserByID :one
SELECT COUNT(*) > 0
FROM users
WHERE id = $1;

-- name: FindUsersByRole :many
SELECT *
FROM users
WHERE role = $1
AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: CountUsersByRole :one
SELECT COUNT(*)
FROM users
WHERE role = $1
AND deleted_at IS NULL;

-- name: CountActiveUsers :one
SELECT COUNT(*)
FROM users
WHERE status = 'active' 
AND deleted_at IS NULL;

-- name: CountAllUsers :one
SELECT COUNT(*)
FROM users
WHERE deleted_at IS NULL;

-- name: CountUsersByStatus :one
SELECT COUNT(*)
FROM users
WHERE status = $1
AND deleted_at IS NULL;

-- name: ExistsUserByCustomerID :one
SELECT COUNT(*) > 0
FROM customers
WHERE id = $1
AND deleted_at IS NULL;

-- name: ExistsUserByEmployeeID :one
SELECT COUNT(*) > 0
FROM employees
WHERE id = $1
AND deleted_at IS NULL;

-- name: FindActiveUsers :many
SELECT *
FROM users
WHERE status = 'active'
AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: FindAllUsers :many
SELECT *
FROM users
WHERE deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: FindUserByCustomerID :one
SELECT u.*
FROM users u
INNER JOIN customers c ON u.id = c.user_id
WHERE c.id = $1
AND u.deleted_at IS NULL;

-- name: FindUserByEmployeeID :one
SELECT u.*
FROM users u
INNER JOIN employees e ON u.id = e.user_id
WHERE e.id = $1
AND u.deleted_at IS NULL;

-- name: FindInactiveUsers :many
SELECT *
FROM users
WHERE status != 'active'
AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: FindRecentlyLoggedInUsers :many
SELECT *
FROM users
WHERE last_login >= $1
AND deleted_at IS NULL
ORDER BY last_login DESC
LIMIT $2 OFFSET $3;



-- name: UpdateUserPassword :exec
UPDATE users
SET 
    password = $2,
    updated_at = CURRENT_TIMESTAMP,
    password_changed_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: UpdateUserStatus :exec
UPDATE users
SET 
    status = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;


-- name: GetCustomerIDByUserID :one
SELECT c.id
FROM customers c
WHERE c.user_id = $1
AND c.deleted_at IS NULL;

-- name: GetEmployeeIDByUserID :one
SELECT e.id
FROM employees e
WHERE e.user_id = $1
AND e.deleted_at IS NULL;

-- name: GetUserCustomerProfile :one
SELECT u.*, c.*
FROM users u
JOIN customers c ON u.id = c.user_id
WHERE u.id = $1
AND u.deleted_at IS NULL;

-- name: GetUserEmployeeProfile :one
SELECT u.*, e.*
FROM users u
JOIN employees e ON u.id = e.user_id
WHERE u.id = $1
AND u.deleted_at IS NULL;