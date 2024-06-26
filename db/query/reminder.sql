-- name: CreateReminder :one
INSERT INTO reminders (appointment_id, method, time_before, created_at, updated_at)
VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, appointment_id, method, time_before, created_at, updated_at;

-- name: GetReminderByID :one
SELECT id, appointment_id, method, time_before, created_at, updated_at
FROM reminders
WHERE id = $1;

-- name: ListReminders :many
SELECT id, appointment_id, method, time_before, created_at, updated_at
FROM reminders
ORDER BY id;

-- name: UpdateReminder :exec
UPDATE reminders
SET appointment_id = $2, method = $3, time_before = $4, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeleteReminder :exec
DELETE FROM reminders
WHERE id = $1;
