-- name: FindAppointmentByID :one
SELECT * FROM appointments 
WHERE id = $1 
AND deleted_at IS NULL;

-- name: FindAppointmentByIDAndCustomerID :one
SELECT * FROM appointments 
WHERE id = $1 
AND customer_id = $2
AND deleted_at IS NULL;

-- name: FindAppointmentByIDAndEmployeeID :one
SELECT * FROM appointments 
WHERE id = $1 
AND employee_id = $2
AND deleted_at IS NULL;

-- name: FindAppointmentsByCustomerID :many
SELECT * FROM appointments 
WHERE customer_id = $1
AND deleted_at IS NULL
ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: CountAppointmentsByCustomerID :one
SELECT COUNT(*) FROM appointments
WHERE customer_id = $1
AND deleted_at IS NULL;


-- name: FindAppointmentsByEmployeeID :many
SELECT * FROM appointments
WHERE employee_id = $1
AND deleted_at IS NULL
ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: CountAppointmentsByEmployeeID :one
SELECT COUNT(*) FROM appointments
WHERE employee_id = $1
AND deleted_at IS NULL;

-- name: FindAppointmentsByPetID :many
SELECT * FROM appointments
WHERE pet_id = $1
AND deleted_at IS NULL
ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: CountAppointmentsByPetID :one
SELECT COUNT(*) FROM appointments
WHERE pet_id = $1
AND deleted_at IS NULL;

-- name: FindAppointments :many
SELECT * FROM appointments 
WHERE deleted_at IS NULL 
ORDER BY created_at DESC LIMIT $1 OFFSET $2;

-- name: CountAppointments :one
SELECT COUNT(*) FROM appointments
WHERE deleted_at IS NULL;


-- name: FindAppointmentsByDateRange :many
SELECT * FROM appointments
WHERE schedule_date BETWEEN $1 AND $2
AND deleted_at IS NULL
ORDER BY schedule_date DESC LIMIT $3 OFFSET $4;

-- name: CountAppointmentsByDateRange :one
SELECT COUNT(*) FROM appointments
WHERE schedule_date BETWEEN $1 AND $2
AND deleted_at IS NULL;

-- name: FindAppointmentsByStatus :many
SELECT * FROM appointments
WHERE status = $1
AND deleted_at IS NULL
ORDER BY created_at DESC LIMIT $2 OFFSET $3;

-- name: CountAppointmentsByStatus :one
SELECT COUNT(*) FROM appointments
WHERE status = $1
AND deleted_at IS NULL;


-- name: CreateAppointment :one
INSERT INTO appointments (
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

-- name: UpdateAppointment :one
UPDATE appointments SET
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

-- name: DeleteAppointment :exec
UPDATE appointments SET
    deleted_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: ExistsAppointmentID :one
SELECT COUNT(*) > 0 FROM appointments
WHERE id = $1 AND deleted_at IS NULL;


-- name: ExistsConflictingAppointment :one
SELECT COUNT(*) > 0 FROM appointments
WHERE (schedule_date BETWEEN $1 AND $2) AND employee_id = $3;
