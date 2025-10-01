-- name: FindDewormingByID :one
SELECT 
    id, pet_id, medication_name, administered_date, next_due_date,
    administered_by, notes, created_at
FROM pet_deworming 
WHERE id = $1;

-- name: FindDewormingByIDAndPetID :one
SELECT 
    id, pet_id, medication_name, administered_date, next_due_date,
    administered_by, notes, created_at
FROM pet_deworming 
WHERE id = $1 AND pet_id = $2;

-- name: FindDewormingByIDAndEmployeeID :one
SELECT 
    id, pet_id, medication_name, administered_date, next_due_date,
    administered_by, notes, created_at
FROM pet_deworming 
WHERE id = $1 AND administered_by = $2;

-- name: FindDewormingsByPetID :many
SELECT 
    id, pet_id, medication_name, administered_date, next_due_date,
    administered_by, notes, created_at
FROM pet_deworming 
WHERE pet_id = $1
ORDER BY administered_date DESC
LIMIT $2 OFFSET $3;

-- name: CountDewormingsByPetID :one
SELECT COUNT(*) FROM pet_deworming WHERE pet_id = $1;

-- name: FindDewormingsByPetIDs :many
SELECT 
    id, pet_id, medication_name, administered_date, next_due_date,
    administered_by, notes, created_at
FROM pet_deworming 
WHERE pet_id = ANY($1::int[])
ORDER BY pet_id, administered_date DESC
LIMIT $2 OFFSET $3;

-- name: CountDewormingsByPetIDs :one
SELECT COUNT(*) FROM pet_deworming WHERE pet_id = ANY($1::int[]);

-- name: FindDewormingsByEmployeeID :many
SELECT 
    id, pet_id, medication_name, administered_date, next_due_date,
    administered_by, notes, created_at
FROM pet_deworming 
WHERE administered_by = $1
ORDER BY administered_date DESC
LIMIT $2 OFFSET $3;

-- name: CountDewormingsByEmployeeID :one
SELECT COUNT(*) FROM pet_deworming WHERE administered_by = $1;

-- name: FindDewormingsByDateRange :many
SELECT 
    id, pet_id, medication_name, administered_date, next_due_date,
    administered_by, notes, created_at
FROM pet_deworming 
WHERE administered_date BETWEEN $1 AND $2
ORDER BY administered_date DESC
LIMIT $3 OFFSET $4;

-- name: CountDewormingsByDateRange :one
SELECT COUNT(*) FROM pet_deworming 
WHERE administered_date BETWEEN $1 AND $2;

-- name: CreateDeworming :one
INSERT INTO pet_deworming (
    pet_id, medication_name, administered_date, next_due_date,
    administered_by, notes
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: UpdateDeworming :one
UPDATE pet_deworming 
SET 
    medication_name = $2,
    administered_date = $3,
    next_due_date = $4,
    administered_by = $5,
    notes = $6
WHERE id = $1
RETURNING *;

-- name: SoftDeleteDeworming :exec
UPDATE pet_deworming 
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: HardDeleteDeworming :exec
DELETE FROM pet_deworming 
WHERE id = $1;

-- name: FindDewormingsBySpec :many
SELECT 
    id, pet_id, medication_name, administered_date, next_due_date,
    administered_by, notes, created_at
FROM pet_deworming 
WHERE 
    (@id::INT IS NULL OR id = @id)
    AND (@pet_id::INT IS NULL OR pet_id = @pet_id)
    AND (@administered_by::INT IS NULL OR administered_by = @administered_by)
    AND (@medication_name::VARCHAR IS NULL OR medication_name ILIKE '%' || @medication_name || '%')
    AND (@administered_date_from::DATE IS NULL OR administered_date >= @administered_date_from)
    AND (@administered_date_to::DATE IS NULL OR administered_date <= @administered_date_to)
    AND (@administered_date_exact::DATE IS NULL OR administered_date = @administered_date_exact)
    AND (@next_due_date_from::DATE IS NULL OR next_due_date >= @next_due_date_from)
    AND (@next_due_date_to::DATE IS NULL OR next_due_date <= @next_due_date_to)
    AND (@next_due_date_exact::DATE IS NULL OR next_due_date = @next_due_date_exact)
    AND (@created_at_from::TIMESTAMPTZ IS NULL OR created_at >= @created_at_from)
    AND (@created_at_to::TIMESTAMPTZ IS NULL OR created_at <= @created_at_to)
ORDER BY 
    CASE WHEN @sort_by = 'administered_date_desc' THEN administered_date END DESC,
    CASE WHEN @sort_by = 'administered_date_asc' THEN administered_date END ASC,
    CASE WHEN @sort_by = 'next_due_date_desc' THEN next_due_date END DESC,
    CASE WHEN @sort_by = 'next_due_date_asc' THEN next_due_date END ASC,
    CASE WHEN @sort_by = 'created_at_desc' THEN created_at END DESC,
    CASE WHEN @sort_by = 'created_at_asc' THEN created_at END ASC,
    administered_date DESC -- default ordering
LIMIT @limit_val OFFSET @offset_val;

-- name: CountDewormingsBySpec :one
SELECT COUNT(*)
FROM pet_deworming 
WHERE 
    (@id::INT IS NULL OR id = @id)
    AND (@pet_id::INT IS NULL OR pet_id = @pet_id)
    AND (@administered_by::INT IS NULL OR administered_by = @administered_by)
    AND (@medication_name::VARCHAR IS NULL OR medication_name ILIKE '%' || @medication_name || '%')
    AND (@administered_date_from::DATE IS NULL OR administered_date >= @administered_date_from)
    AND (@administered_date_to::DATE IS NULL OR administered_date <= @administered_date_to)
    AND (@administered_date_exact::DATE IS NULL OR administered_date = @administered_date_exact)
    AND (@next_due_date_from::DATE IS NULL OR next_due_date >= @next_due_date_from)
    AND (@next_due_date_to::DATE IS NULL OR next_due_date <= @next_due_date_to)
    AND (@next_due_date_exact::DATE IS NULL OR next_due_date = @next_due_date_exact)
    AND (@created_at_from::TIMESTAMPTZ IS NULL OR created_at >= @created_at_from)
    AND (@created_at_to::TIMESTAMPTZ IS NULL OR created_at <= @created_at_to);