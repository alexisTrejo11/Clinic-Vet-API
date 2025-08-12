-- name: CreateAppoinment :one
INSERT INTO appoinments (
    clinic_service, 
    schedule_date, 
    status, 
    reason, 
    notes, 
    owner_id, 
    veterinarian_id,
    pet_id,
    created_at,
    updated_at,
    deleted_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL
) RETURNING *;



-- name: UpdateAppoinment :one
UPDATE appoinments SET
    clinic_service = $2,
    schedule_date = $3,
    status = $4,
    reason = $5,
    notes = $6,
    owner_id = $7,
    veterinarian_id = $8,
    pet_id = $9,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteAppoinment :exec
UPDATE appoinments SET
    deleted_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: ListAppoinmentsByOwnerID :many
SELECT * FROM appoinments 
WHERE owner_id = $1
AND deleted_at IS NULL
ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: CountAppoinmentsByOwnerID :one
SELECT COUNT(*) FROM appoinments
WHERE owner_id = $1
AND deleted_at IS NULL;


-- name: ListAppoinmentsByVeterinarianID :many
SELECT * FROM appoinments
WHERE veterinarian_id = $1
AND deleted_at IS NULL
ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: CountAppoinmentsByVeterinarianID :one
SELECT COUNT(*) FROM appoinments
WHERE veterinarian_id = $1
AND deleted_at IS NULL;

-- name: ListAppoinmentsByPetID :many
SELECT * FROM appoinments
WHERE pet_id = $1
AND deleted_at IS NULL
ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: CountAppoinmentsByPetID :one
SELECT COUNT(*) FROM appoinments
WHERE pet_id = $1
AND deleted_at IS NULL;

-- name: GetAppoinmentByID :one
SELECT * FROM appoinments 
WHERE id = $1 
AND deleted_at IS NULL;

-- name: ListAppoinments :many
SELECT * FROM appoinments 
WHERE deleted_at IS NULL 
ORDER BY created_at DESC LIMIT $1 OFFSET $2;

-- name: CountAppoinments :one
SELECT COUNT(*) FROM appoinments
WHERE deleted_at IS NULL;


-- name: ListAppoinmentsByDateRange :many
SELECT * FROM appoinments
WHERE schedule_date BETWEEN $1 AND $2
AND deleted_at IS NULL
ORDER BY schedule_date DESC LIMIT $3 OFFSET $4;

-- name: CountAppoinmentsByDateRange :one
SELECT COUNT(*) FROM appoinments
WHERE schedule_date BETWEEN $1 AND $2
AND deleted_at IS NULL;

-- name: ListAppoinmentsByStatus :many
SELECT * FROM appoinments
WHERE status = $1
AND deleted_at IS NULL
ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: CountAppoinmentsByStatus :one
SELECT COUNT(*) FROM appoinments
WHERE status = $1
AND deleted_at IS NULL;
