-- name: CreateVeterinarian :one
INSERT INTO veterinarians (name, photo, email, specialty, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, name, photo, email, specialty, user_id, created_at, updated_at;

-- name: GetVeterinarianByID :one
SELECT id, name, photo, email, specialty, user_id, created_at, updated_at
FROM veterinarians
WHERE id = $1;

-- name: ListVeterinarians :many
SELECT id, name, photo, email, specialty, user_id, created_at, updated_at
FROM veterinarians
ORDER BY id;

-- name: UpdateVeterinarian :exec
UPDATE veterinarians
SET name = $2, photo = $3, email = $4, specialty = $5, user_id = $6, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeleteVeterinarian :exec
DELETE FROM veterinarians
WHERE id = $1;
