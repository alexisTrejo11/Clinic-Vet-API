-- name: CreateUser :one
INSERT INTO users (name, email, phone_number, password, role, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, name, email, password, role, created_at, updated_at;

-- name: GetUserByID :one
SELECT id, name, email, phone_number, password, role, created_at, updated_at
FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT id, name, email, phone_number, password, role, created_at, updated_at
FROM users
ORDER BY id;

-- name: UpdateUser :exec
UPDATE users
SET name = $2, email = $3, phone_number = $4, password = $5, role = $6, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: CheckEmailExists :one
SELECT EXISTS (
    SELECT 1 FROM users WHERE email = $1
) AS exists;

-- name: CheckPhoneNumberExists :one
SELECT EXISTS (
    SELECT 1 FROM users WHERE phone_number = $1
) AS exists;

-- name: UpdateLastLogin :exec
UPDATE users
SET updated_at = CURRENT_TIMESTAMP
WHERE id = $1;