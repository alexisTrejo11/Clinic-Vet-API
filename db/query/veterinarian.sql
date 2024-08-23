-- name: CreateVeterinarian :one
INSERT INTO veterinarians (name, photo, specialty, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, name, photo, specialty, user_id, created_at, updated_at;

-- name: GetVeterinarianByID :one
SELECT id, name, photo, specialty, user_id, created_at, updated_at
FROM veterinarians
WHERE id = $1;

-- name: ListVeterinarians :many
SELECT id, name, photo, specialty, user_id, created_at, updated_at
FROM veterinarians
ORDER BY id;

-- name: UpdateVeterinarian :exec
UPDATE veterinarians
SET name = $2, photo = $3, specialty = $4, user_id = $5, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: UpdateVeterinarianUserID :exec
UPDATE veterinarians
SET user_id = $2, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeleteVeterinarian :exec
DELETE FROM veterinarians
WHERE id = $1;
