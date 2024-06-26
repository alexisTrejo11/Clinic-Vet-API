-- name: CreateMedicalHistory :one
INSERT INTO medical_histories (pet_id, date, description, vet_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, pet_id, date, description, vet_id, created_at, updated_at;

-- name: GetMedicalHistoryByID :one
SELECT id, pet_id, date, description, vet_id, created_at, updated_at
FROM medical_histories
WHERE id = $1;

-- name: ListMedicalHistories :many
SELECT id, pet_id, date, description, vet_id, created_at, updated_at
FROM medical_histories
ORDER BY id;

-- name: UpdateMedicalHistory :exec
UPDATE medical_histories
SET pet_id = $2, date = $3, description = $4, vet_id = $5, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeleteMedicalHistory :exec
DELETE FROM medical_histories
WHERE id = $1;
