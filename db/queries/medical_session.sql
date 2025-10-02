-- name: FindMedicalSessionByID :one
SELECT * FROM medical_sessions
WHERE id = $1 AND deleted_at IS NULL;

-- name: FindMedicalSessionByIDAndPetID :one
SELECT * FROM medical_sessions
WHERE id = $1 AND pet_id = $2 AND deleted_at IS NULL;

-- name: FindMedicalSessionByIDAndCustomerID :one
SELECT * FROM medical_sessions
WHERE id = $1 AND customer_id = $2 AND deleted_at IS NULL;

-- name: FindMedicalSessionByIDAndEmployeeID :one
SELECT * FROM medical_sessions
WHERE id = $1 AND employee_id = $2 AND deleted_at IS NULL;

-- name: FindAllMedicalSession :many
SELECT * FROM medical_sessions
WHERE deleted_at IS NULL
ORDER BY visit_date DESC
LIMIT $1 OFFSET $2;

-- name: CountAllMedicalSession :one
SELECT COUNT(*) FROM medical_sessions
WHERE deleted_at IS NULL;

-- name: FindMedicalSessionByEmployeeID :many
SELECT * FROM medical_sessions
WHERE employee_id = $1 AND deleted_at IS NULL
ORDER BY visit_date DESC
LIMIT $2 OFFSET $3;


-- name: FindMedicalSessionByPetID :many
SELECT * FROM medical_sessions
WHERE pet_id = $1 AND deleted_at IS NULL
ORDER BY visit_date DESC
LIMIT $2 OFFSET $3;


-- name: FindMedicalSessionByCustomerID :many
SELECT * FROM medical_sessions
WHERE customer_id = $1 AND deleted_at IS NULL
ORDER BY visit_date DESC
LIMIT $2 OFFSET $3;


-- name: FindRecentMedicalSessionByPetID :many
SELECT * FROM medical_sessions
WHERE pet_id = $1 AND deleted_at IS NULL
ORDER BY visit_date DESC
LIMIT $2;

-- name: FindMedicalSessionByDateRange :many
SELECT * FROM medical_sessions
WHERE visit_date BETWEEN $1 AND $2
AND deleted_at IS NULL
ORDER BY visit_date DESC
LIMIT $3 OFFSET $4;

-- name: FindMedicalSessionByPetAndDateRange :many
SELECT * FROM medical_sessions
WHERE pet_id = $1
AND visit_date BETWEEN $2 AND $3
AND deleted_at IS NULL
ORDER BY visit_date DESC;

-- name: FindMedicalSessionByDiagnosis :many
SELECT * FROM medical_sessions
WHERE diagnosis ILIKE '%' || $1 || '%'
AND deleted_at IS NULL
ORDER BY visit_date DESC
LIMIT $2 OFFSET $3;

-- name: CountMedicalSessionByDiagnosis :one
SELECT COUNT(*) FROM medical_sessions
WHERE diagnosis ILIKE '%' || $1 || '%'
AND deleted_at IS NULL;

-- name: ExistsMedicalSessionByID :one
SELECT COUNT(*) > 0 FROM medical_sessions
WHERE id = $1 AND deleted_at IS NULL;

-- name: ExistsMedicalSessionByPetAndDate :one
SELECT COUNT(*) > 0 FROM medical_sessions
WHERE pet_id = $1
AND DATE(visit_date) = DATE($2)
AND deleted_at IS NULL;

-- name: SaveMedicalSession :one
INSERT INTO medical_sessions (
    pet_id, 
    customer_id,
    employee_id,
    visit_date,
    visit_type,
    diagnosis, 
    clinic_service,
    treatment,
    notes,
    condition,
    weight,
    temperature,
    heart_rate,
    respiratory_rate
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
)
RETURNING *;

-- name: UpdateMedicalSession :one
UPDATE medical_sessions
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
    clinic_service = $15,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SoftDeleteMedicalSession :exec
UPDATE medical_sessions
SET 
    deleted_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: HardDeleteMedicalSession :exec
DELETE FROM medical_sessions
WHERE id = $1;

-- name: CountMedicalSessionByPetID :one
SELECT COUNT(*) FROM medical_sessions
WHERE pet_id = $1 AND deleted_at IS NULL;

-- name: CountMedicalSessionByEmployeeID :one
SELECT COUNT(*) FROM medical_sessions
WHERE employee_id = $1 AND deleted_at IS NULL;

-- name: CountMedicalSessionByCustomerID :one
SELECT COUNT(*) FROM medical_sessions
WHERE customer_id = $1 AND deleted_at IS NULL;

-- name: CountMedicalSessionByDateRange :one
SELECT COUNT(*) FROM medical_sessions
WHERE visit_date BETWEEN $1 AND $2
AND deleted_at IS NULL;