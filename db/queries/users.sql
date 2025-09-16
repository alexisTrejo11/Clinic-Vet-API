-- name: CreateUser :one
INSERT INTO users (
    email, 
    phone_number, 
    password, 
    status, 
    role, 
    profile_id,
    customer_id,
    employee_id,
    created_at,
    updated_at
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8,
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
    profile_id = $7,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: FindUserByID :one
SELECT *
FROM users
WHERE id = $1 
AND deleted_at IS NULL;

-- name: FindUserByEmail :one
SELECT *
FROM users
WHERE email = $1
AND deleted_at IS NULL;

-- name: FindUserByPhoneNumber :one
SELECT *
FROM users
WHERE phone_number = $1
AND deleted_at IS NULL;

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
FROM users
WHERE customer_id = $1
AND deleted_at IS NULL;

-- name: ExistsUserByEmployeeID :one
SELECT COUNT(*) > 0
FROM users
WHERE employee_id = $1
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
SELECT *
FROM users
WHERE customer_id = $1
AND deleted_at IS NULL;

-- name: FindUserByEmployeeID :one
SELECT *
FROM users
WHERE employee_id = $1
AND deleted_at IS NULL;

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

-- name: FindUsersBySpecification :many
SELECT *
FROM users
WHERE deleted_at IS NULL
AND ($1::text IS NULL OR email ILIKE '%' || $1 || '%')
AND ($2::text IS NULL OR phone_number ILIKE '%' || $2 || '%')
AND ($3::text IS NULL OR role = $3)
AND ($4::text IS NULL OR status = $4)
AND ($5::timestamptz IS NULL OR last_login >= $5)
AND ($6::timestamptz IS NULL OR created_at >= $6)
ORDER BY created_at DESC
LIMIT $7 OFFSET $8;

-- name: CountUsersBySpecification :one
SELECT COUNT(*)
FROM users
WHERE deleted_at IS NULL
AND ($1::text IS NULL OR email ILIKE '%' || $1 || '%')
AND ($2::text IS NULL OR phone_number ILIKE '%' || $2 || '%')
AND ($3::text IS NULL OR role = $3)
AND ($4::text IS NULL OR status = $4)
AND ($5::timestamptz IS NULL OR last_login >= $5)
AND ($6::timestamptz IS NULL OR created_at >= $6);

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