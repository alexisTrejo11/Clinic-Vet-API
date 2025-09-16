-- name: FindMedicalHistoryByID :one
SELECT * FROM medical_history
WHERE id = $1 AND deleted_at IS NULL;


-- name: FindAllMedicalHistory :many
SELECT * FROM medical_history
WHERE deleted_at IS NULL
ORDER BY visit_date DESC
LIMIT $1 OFFSET $2;

-- name: CountAllMedicalHistory :one
SELECT COUNT(*) FROM medical_history
WHERE deleted_at IS NULL;

-- name: FindMedicalHistoryByEmployeeID :many
SELECT * FROM medical_history
WHERE employee_id = $1 AND deleted_at IS NULL
ORDER BY visit_date DESC
LIMIT $2 OFFSET $3;


-- name: FindMedicalHistoryByPetID :many
SELECT * FROM medical_history
WHERE pet_id = $1 AND deleted_at IS NULL
ORDER BY visit_date DESC
LIMIT $2 OFFSET $3;


-- name: FindMedicalHistoryByCustomerID :many
SELECT * FROM medical_history
WHERE customer_id = $1 AND deleted_at IS NULL
ORDER BY visit_date DESC
LIMIT $2 OFFSET $3;


-- name: FindRecentMedicalHistoryByPetID :many
SELECT * FROM medical_history
WHERE pet_id = $1 AND deleted_at IS NULL
ORDER BY visit_date DESC
LIMIT $2;

-- name: FindMedicalHistoryByDateRange :many
SELECT * FROM medical_history
WHERE visit_date BETWEEN $1 AND $2
AND deleted_at IS NULL
ORDER BY visit_date DESC
LIMIT $3 OFFSET $4;

-- name: FindMedicalHistoryByPetAndDateRange :many
SELECT * FROM medical_history
WHERE pet_id = $1
AND visit_date BETWEEN $2 AND $3
AND deleted_at IS NULL
ORDER BY visit_date DESC;

-- name: FindMedicalHistoryByDiagnosis :many
SELECT * FROM medical_history
WHERE diagnosis ILIKE '%' || $1 || '%'
AND deleted_at IS NULL
ORDER BY visit_date DESC
LIMIT $2 OFFSET $3;

-- name: CountMedicalHistoryByDiagnosis :one
SELECT COUNT(*) FROM medical_history
WHERE diagnosis ILIKE '%' || $1 || '%'
AND deleted_at IS NULL;

-- name: ExistsMedicalHistoryByID :one
SELECT COUNT(*) > 0 FROM medical_history
WHERE id = $1 AND deleted_at IS NULL;

-- name: ExistsMedicalHistoryByPetAndDate :one
SELECT COUNT(*) > 0 FROM medical_history
WHERE pet_id = $1
AND DATE(visit_date) = DATE($2)
AND deleted_at IS NULL;

-- name: SaveMedicalHistory :one
INSERT INTO medical_history (
    pet_id, 
    customer_id,
    employee_id,
    visit_date,
    visit_type,
    diagnosis, 
    treatment,
    notes,
    condition,
    weight,
    temperature,
    heart_rate,
    respiratory_rate
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
)
RETURNING *;

-- name: UpdateMedicalHistory :one
UPDATE medical_history
SET 
    pet_id = $2, 
    customer_id = $3,
    employee_id = $4,
    visit_date = $5, 
    visit_type = $6,
    diagnosis = $7, 
    treatment = $8,
    notes = $9,
    condition = $10,
    weight = $11,
    temperature = $12,
    heart_rate = $13,
    respiratory_rate = $14,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SoftDeleteMedicalHistory :exec
UPDATE medical_history
SET 
    deleted_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: HardDeleteMedicalHistory :exec
DELETE FROM medical_history
WHERE id = $1;

-- name: CountMedicalHistoryByPetID :one
SELECT COUNT(*) FROM medical_history
WHERE pet_id = $1 AND deleted_at IS NULL;

-- name: CountMedicalHistoryByEmployeeID :one
SELECT COUNT(*) FROM medical_history
WHERE employee_id = $1 AND deleted_at IS NULL;

-- name: CountMedicalHistoryByCustomerID :one
SELECT COUNT(*) FROM medical_history
WHERE customer_id = $1 AND deleted_at IS NULL;

-- name: CountMedicalHistoryByDateRange :one
SELECT COUNT(*) FROM medical_history
WHERE visit_date BETWEEN $1 AND $2
AND deleted_at IS NULL;