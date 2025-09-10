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

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1 
AND deleted_at IS NULL;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1
AND deleted_at IS NULL;

-- name: GetUserByPhoneNumber :one
SELECT *
FROM users
WHERE phone_number = $1
AND deleted_at IS NULL;

-- name: ListUsers :many
SELECT *
FROM users
WHERE deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $1
OFFSET $2;

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

-- name: ListUsersByRole :many
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
