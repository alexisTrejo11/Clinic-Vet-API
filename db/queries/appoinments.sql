-- name: FindAppointmentsBySpec :many
SELECT 
    id, clinic_service, scheduled_date, status, notes,
    customer_id, employee_id, pet_id, created_at, updated_at
FROM appointments 
WHERE 
    ($1::INT IS NULL OR id = $1)
    AND ($2::INT IS NULL OR customer_id = $2)
    AND ($3::INT IS NULL OR employee_id = $3)
    AND ($4::INT IS NULL OR pet_id = $4)
    AND ($5::VARCHAR IS NULL OR clinic_service = $5)
    AND ($6::VARCHAR IS NULL OR status = $6)
    AND ($7::timestamp IS NULL OR scheduled_date >= $7)
    AND ($8::timestamp IS NULL OR scheduled_date <= $8)
    AND ($9::timestamp IS NULL OR scheduled_date = $9)
ORDER BY scheduled_date DESC
LIMIT $10 OFFSET $11;

-- name: CountAppointmentsBySpec :one
SELECT COUNT(*) 
FROM appointments 
WHERE 
    ($1::INT IS NULL OR id = $1)
    AND ($2::INT IS NULL OR customer_id = $2)
    AND ($3::INT IS NULL OR employee_id = $3)
    AND ($4::INT IS NULL OR pet_id = $4)
    AND ($5::VARCHAR IS NULL OR clinic_service = $5)
    AND ($6::VARCHAR IS NULL OR status = $6)
    AND ($7::timestamp IS NULL OR scheduled_date >= $7)
    AND ($8::timestamp IS NULL OR scheduled_date <= $8)
    AND ($9::timestamp IS NULL OR scheduled_date = $9);



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

-- name: CreateAppointment :one
INSERT INTO appointments (
    clinic_service, 
    scheduled_date, 
    status, 
    notes, 
    customer_id, 
    employee_id,
    pet_id,
    created_at,
    updated_at,
    deleted_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL
) RETURNING *;

-- name: UpdateAppointment :one
UPDATE appointments SET
    clinic_service = $2,
    scheduled_date = $3,
    status = $4,
    notes = $5,
    customer_id = $6,
    employee_id = $7,
    pet_id = $8,
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
WHERE (scheduled_date BETWEEN $1 AND $2) AND employee_id = $3;
