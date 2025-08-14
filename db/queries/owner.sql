-- name: CreateOwner :one
INSERT INTO owners (
    photo, first_name, last_name, phone_number, gender, address, user_id, is_active, date_of_birth, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
)
RETURNING *;

-- name: GetOwnerByID :one
SELECT *
FROM owners
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetOwnerByPhone :one
SELECT *
FROM owners
WHERE phone_number = $1 AND deleted_at IS NULL;

-- name: GetOwnerByUserID :one
SELECT *
FROM owners
WHERE user_id = $1 AND deleted_at IS NULL;

-- name: ListOwners :many
SELECT * FROM owners
WHERE deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: UpdateOwner :exec
UPDATE owners
SET 
    photo = $2, 
    first_name = $3, 
    last_name = $4, 
    phone_number = $5, 
    gender = $6, 
    address = $7, 
    user_id = $8, 
    is_active = $9, 
    date_of_birth = $10,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SoftDeleteOwner :exec
UPDATE owners
SET 
    is_active = FALSE,
    deleted_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;


-- name: ExistByID :one
SELECT COUNT(*) > 0
FROM owners
WHERE id = $1;

-- name: ExistByPhoneNumber :one
SELECT COUNT(*) > 0
FROM owners
WHERE phone_number = $1;

-- name: DeactivateUser :exec
UPDATE owners
SET 
    is_active = false, 
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: ActivateUser :exec
UPDATE owners
SET 
    is_active = true, 
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;