-- name: GetAppoinmentByID :one
SELECT * FROM appoinments 
WHERE id = $1 
AND deleted_at IS NULL;

-- name: GetAppointmentByIDAndCustomerID :one
SELECT * FROM appoinments 
WHERE id = $1 
AND customer_id = $2
AND deleted_at IS NULL;

-- name: GetAppointmentByIDAndEmployeeID :one
SELECT * FROM appoinments 
WHERE id = $1 
AND employee_id = $2
AND deleted_at IS NULL;

-- name: ListAppoinmentsByCustomerID :many
SELECT * FROM appoinments 
WHERE customer_id = $1
AND deleted_at IS NULL
ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: CountAppoinmentsByCustomerID :one
SELECT COUNT(*) FROM appoinments
WHERE customer_id = $1
AND deleted_at IS NULL;


-- name: ListAppoinmentsByEmployeeID :many
SELECT * FROM appoinments
WHERE employee_id = $1
AND deleted_at IS NULL
ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: CountAppoinmentsByEmployeeID :one
SELECT COUNT(*) FROM appoinments
WHERE employee_id = $1
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


-- name: CreateAppoinment :one
INSERT INTO appoinments (
    clinic_service, 
    schedule_date, 
    status, 
    reason, 
    notes, 
    customer_id, 
    employee_id,
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
    customer_id = $7,
    employee_id = $8,
    pet_id = $9,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteAppoinment :exec
UPDATE appoinments SET
    deleted_at = CURRENT_TIMESTAMP
WHERE id = $1;
