-- name: GetEmployeeById :one
SELECT 
    id, first_name, last_name, photo, license_number, speciality, years_of_experience, 
    is_active, user_id, schedule_json, created_at, updated_at, deleted_at
FROM employees
WHERE id = $1 AND deleted_at IS NULL;


-- name: CreateEmployee :one
INSERT INTO employees(
    first_name, last_name, photo, license_number, speciality,
    years_of_experience, is_active, schedule_json, created_at, updated_at
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8::jsonb, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
)
RETURNING *; 


-- name: UpdateEmployee :one
UPDATE employees
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


-- name: SoftDeleteEmployee :exec
UPDATE employees
SET
    deleted_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = $1 AND deleted_at IS NULL;


-- name: GetEmployeesWithSchedule :many
SELECT * FROM employees
WHERE 
    deleted_at IS NULL
    AND schedule_json @> '{"work_days": [{"day": 1, "start_hour": 9}]}';
