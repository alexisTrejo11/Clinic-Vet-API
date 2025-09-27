-- name: CreatePetDeworming :one
INSERT INTO pet_deworming (
    pet_id, medication_name, administered_date, next_due_date, 
    administered_by, notes
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING id, pet_id, medication_name, administered_date, next_due_date, 
           administered_by, notes, created_at;

-- name: GetPetDewormingByID :one
SELECT id, pet_id, medication_name, administered_date, next_due_date, 
       administered_by, notes, created_at
FROM pet_deworming 
WHERE id = $1;

-- name: UpdatePetDeworming :one
UPDATE pet_deworming 
SET 
    medication_name = COALESCE($2, medication_name),
    administered_date = COALESCE($3, administered_date),
    next_due_date = COALESCE($4, next_due_date),
    administered_by = COALESCE($5, administered_by),
    notes = COALESCE($6, notes)
WHERE id = $1
RETURNING id, pet_id, medication_name, administered_date, next_due_date, 
          administered_by, notes, created_at;

-- name: DeletePetDeworming :exec
DELETE FROM pet_deworming WHERE id = $1;

-- name: GetPetDewormingsByPetID :many
SELECT id, pet_id, medication_name, administered_date, next_due_date, 
       administered_by, notes, created_at
FROM pet_deworming 
WHERE pet_id = $1
ORDER BY administered_date DESC;

-- name: GetUpcomingDewormings :many
SELECT id, pet_id, medication_name, administered_date, next_due_date, 
       administered_by, notes, created_at
FROM pet_deworming 
WHERE next_due_date IS NOT NULL 
  AND next_due_date <= $1
ORDER BY next_due_date ASC;


-- name: FindPetDewormingsBySpec :many
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

-- name: CountPetDewormingsBySpec :one
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