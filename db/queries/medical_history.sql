-- name: CreateMedicalHistory :one
INSERT INTO medical_history (
    pet_id, 
    owner_id,
    veterinarian_id,
    visit_date,
    visit_type,
    diagnosis, 
    treatment,
    notes,
    condition 
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: UpdateMedicalHistory :one
UPDATE medical_history
SET 
    pet_id = $2, 
    owner_id = $3,
    veterinarian_id = $4,
    visit_date = $5, 
    diagnosis = $6, 
    visit_type = $7,
    notes = $8,
    condition = $9, 
    treatment = $10, 
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: SoftDeleteMedicalHistory :exec
UPDATE medical_history
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: HardDeleteMedicalHistory :exec
DELETE FROM medical_history
WHERE id = $1;

-- name: ListMedicalHistoryByPet :many
SELECT * FROM medical_history
WHERE pet_id = $1 AND deleted_at IS NULL
ORDER BY visit_date DESC;

-- name: GetMedicalHistoryByOwnerID :many
SELECT *
FROM medical_history
WHERE owner_id = $1 AND deleted_at IS NULL
ORDER BY $2 DESC
OFFSET $3
LIMIT $4;

-- name: ListMedicalHistoryByVet :many
SELECT * FROM medical_history
WHERE veterinarian_id = $1 AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: GetMedicalHistoryByID :one
SELECT *
FROM medical_history
WHERE id = $1 AND deleted_at IS NULL;

-- name: SearchMedicalHistory :many
SELECT * FROM medical_history
WHERE deleted_at IS NULL
ORDER BY $3 DESC
OFFSET $1
LIMIT $2;
