-- name: GetVeterinarianById :one
SELECT 
    id, first_name, last_name, photo, license_number, speciality, years_of_experience, 
    is_active, user_id, schedule_json, created_at, updated_at, deleted_at
FROM veterinarians
WHERE id = $1 AND deleted_at IS NULL;


-- name: ListVeterinarians :many
SELECT 
    id, first_name, last_name, photo, license_number, speciality, years_of_experience, 
    is_active, user_id, schedule_json, created_at, updated_at, deleted_at
FROM veterinarians
WHERE 
    (first_name ILIKE $1 OR last_name ILIKE $2)
    AND (license_number ILIKE $3 OR $3 = '')
    AND (speciality = $4 OR $4 = '')
    AND (years_of_experience >= $5 OR $5 = 0)
    AND (years_of_experience <= $6 OR $6 = 0)
    AND (is_active = $7 OR $7 IS NULL)
    AND deleted_at IS NULL
ORDER BY
    CASE WHEN $8 THEN first_name END ASC NULLS LAST,
    CASE WHEN $9 THEN first_name END DESC NULLS LAST,
    CASE WHEN $10 THEN speciality END ASC NULLS LAST,
    CASE WHEN $11 THEN speciality END DESC NULLS LAST,
    CASE WHEN $12 THEN years_of_experience END ASC NULLS LAST,
    CASE WHEN $13 THEN years_of_experience END DESC NULLS LAST,
    CASE WHEN $14 THEN created_at END ASC NULLS LAST,
    CASE WHEN $15 THEN created_at END DESC NULLS LAST
LIMIT $16 OFFSET $17;


-- name: CreateVeterinarian :one
INSERT INTO veterinarians(
    first_name, last_name, photo, license_number, speciality,
    years_of_experience, is_active, schedule_json, created_at, updated_at
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8::jsonb, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
)
RETURNING
    id, first_name, last_name, photo, license_number, speciality, years_of_experience,
    is_active, user_id, schedule_json, created_at, updated_at, deleted_at;


-- name: UpdateVeterinarian :one
UPDATE veterinarians
SET
    first_name = COALESCE($1, first_name),
    last_name = COALESCE($2, last_name),  
    photo = COALESCE($3, photo),          
    license_number = COALESCE($4, license_number),
    speciality = COALESCE($5, speciality),
    years_of_experience = COALESCE($6, years_of_experience),
    is_active = COALESCE($7, is_active),
    schedule_json = COALESCE($8::jsonb, schedule_json),
    updated_at = CURRENT_TIMESTAMP        
WHERE
    id = $9 AND deleted_at IS NULL       
RETURNING
    id, first_name, last_name, photo, license_number, speciality, years_of_experience,
    is_active, user_id, schedule_json, created_at, updated_at, deleted_at;


-- name: SoftDeleteVeterinarian :exec
UPDATE veterinarians
SET
    deleted_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = $1 AND deleted_at IS NULL;


-- name: GetVeterinariansWithSchedule :many
SELECT * FROM veterinarians
WHERE 
    deleted_at IS NULL
    AND schedule_json @> '{"work_days": [{"day": 1, "start_hour": 9}]}';