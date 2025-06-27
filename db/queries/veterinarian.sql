-- name: GetVeterinarianById :one
SELECT 
    id, first_name, last_name, photo, license_number, speciality, years_of_experience, 
    is_active, user_id, created_at, updated_at, deleted_at
FROM veterinarians
WHERE id = $1 AND deleted_at IS NULL;


-- name: ListVeterinarians :many
SELECT
    id, first_name, last_name, photo, license_number, speciality, years_of_experience, 
    is_active, user_id, created_at, updated_at, deleted_at
FROM
    veterinarians
WHERE
    deleted_at IS NULL
ORDER BY
    CASE WHEN @order_by = 'id' AND @order_direction = 'ASC' THEN id END ASC,
    CASE WHEN @order_by = 'id' AND @order_direction = 'DESC' THEN id END DESC,
    CASE WHEN @order_by = 'created_at' AND @order_direction = 'ASC' THEN created_at END ASC,
    CASE WHEN @order_by = 'created_at' AND @order_direction = 'DESC' THEN created_at END DESC,
    id ASC -- Fallback order if no match
LIMIT $1 OFFSET $2;


-- name: CreateVeterinarian :one
INSERT INTO veterinarians(
    first_name, last_name, photo, license_number, speciality,
    years_of_experience, is_active, created_at, updated_at, deleted_at
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL
)
RETURNING
    id, first_name, last_name, photo, license_number, speciality, years_of_experience,
    is_active, user_id, created_at, updated_at, deleted_at;


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
    updated_at = CURRENT_TIMESTAMP        
WHERE
    id = $8 AND deleted_at IS NULL       
RETURNING
    id, first_name, last_name, photo, license_number, speciality, years_of_experience,
    is_active, created_at, updated_at, deleted_at;


-- name: SoftDeleteVeterinarian :exec
UPDATE veterinarians
SET
    deleted_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = $1 AND deleted_at IS NULL;

