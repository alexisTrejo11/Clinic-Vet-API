-- name: CreateAppointment :one
INSERT INTO appointments (pet_id, vet_id, service, date, created_at, updated_at)
VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, pet_id, vet_id, service, date, created_at, updated_at;

-- name: GetAppointmentByID :one
SELECT id, pet_id, vet_id, service, date, created_at, updated_at
FROM appointments
WHERE id = $1;

-- name: ListAppointments :many
SELECT id, pet_id, vet_id, service, date, created_at, updated_at
FROM appointments
ORDER BY id;

-- name: UpdateAppointment :exec
UPDATE appointments
SET pet_id = $2, vet_id = $3, service = $4, date = $5, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeleteAppointment :exec
DELETE FROM appointments
WHERE id = $1;
