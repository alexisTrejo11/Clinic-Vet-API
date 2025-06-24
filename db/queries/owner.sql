-- name: CreateOwner :one
INSERT INTO owners (
    id, photo, firstName, lastName, phoneNumber, gender, address, user_id, isActive, date_of_birth, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
)
RETURNING id, photo, firstName, lastName, phoneNumber, gender, address, user_id, isActive, date_of_birth, created_at, updated_at, deleted_at;

-- name: GetOwnerByID :one
SELECT id, photo, firstName, lastName, phoneNumber, gender, address, user_id, isActive, date_of_birth, created_at, updated_at, deleted_at
FROM owners
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetOwnerByUserID :one
SELECT id, photo, firstName, lastName, phoneNumber, gender, address, user_id, isActive, date_of_birth, created_at, updated_at, deleted_at
FROM owners
WHERE user_id = $1 AND deleted_at IS NULL;

-- name: ListOwners :many
SELECT id, photo, firstName, lastName, phoneNumber, gender, address, user_id, isActive, date_of_birth, created_at, updated_at, deleted_at
FROM owners
WHERE deleted_at IS NULL
ORDER BY id;

-- name: UpdateOwner :exec
UPDATE owners
SET 
    photo = $2, 
    firstName = $3, 
    lastName = $4, 
    phoneNumber = $5, 
    gender = $6, 
    address = $7, 
    user_id = $8, 
    isActive = $9, 
    date_of_birth = $10,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted_at IS NULL;

-- name: DeleteOwner :exec
UPDATE owners
SET 
    isActive = FALSE,
    deleted_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;