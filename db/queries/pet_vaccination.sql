-- name: FindVaccinationByID :one
SELECT *
FROM pet_vaccinations
WHERE id = $1;

-- name: FindVaccinationsByDateRange :many
SELECT *
FROM pet_vaccinations
WHERE administered_date BETWEEN $1 AND $2
ORDER BY administered_date DESC
LIMIT $3 OFFSET $4;


-- name: CountVaccinationsByDateRange :one
SELECT COUNT(*)
FROM pet_vaccinations
WHERE administered_date BETWEEN $1 AND $2;

-- name: FindAllVaccinationsByPetID :many
SELECT *
FROM pet_vaccinations
WHERE pet_id = $1
ORDER BY administered_date DESC;


-- name: FindVaccinationsByPetID :many
SELECT *
FROM pet_vaccinations
WHERE pet_id = $1
ORDER BY administered_date DESC
LIMIT $2 OFFSET $3;

-- name: CountVaccinationsByPetID :one
SELECT COUNT(*)
FROM pet_vaccinations
WHERE pet_id = $1;

-- name: FindVaccinationsByPetIDs :many
SELECT *
FROM pet_vaccinations
WHERE pet_id = ANY($1::int[])
ORDER BY administered_date DESC
LIMIT $2 OFFSET $3;

-- name: CountVaccinationsByPetIDs :one
SELECT COUNT(*)
FROM pet_vaccinations
WHERE pet_id = ANY($1::int[]);

-- name: FindByIDAndPetID :one
SELECT *
FROM pet_vaccinations
WHERE id = $1 AND pet_id = $2;


-- name: FindVaccinationByIDAndAdministeredBy :one
SELECT *
FROM pet_vaccinations
WHERE id = $1 AND administered_by = $2;

-- name: FindVaccinationsByAdministeredBy :many
SELECT *
FROM pet_vaccinations
WHERE administered_by = $1
ORDER BY administered_date DESC
LIMIT $2 OFFSET $3;


-- name: CountVaccinationsByAdministeredBy :one
SELECT COUNT(*)
FROM pet_vaccinations
WHERE administered_by = $1;

-- name: FindVaccinationsByIDAndAdministeredBy :one
SELECT *
FROM pet_vaccinations   
WHERE id = $1 AND administered_by = $2; 


-- name: CreatePetVaccination :one
INSERT INTO pet_vaccinations(
    pet_id,
    vaccine_name,
    administered_date,
    administered_by,
    notes,
    batch_number,
    next_due_date,
    vaccine_type,
    created_at,
    updated_at
)
VALUES ($1,$2,$3,$4,$5,$6, $7, $8, NOW(), NOW())
RETURNING *;

-- name: UpdateVaccination :one
UPDATE pet_vaccinations
SET
    pet_id = $2,
    vaccine_name = $3,
    administered_date = $4,
    administered_by = $5,
    notes = $6,
    next_due_date = $7,
    batch_number = $8,
    vaccine_type = $9,
    updated_at = NOW()
WHERE id = $1
RETURNING *;


-- name: DeleteVaccination :exec
DELETE FROM pet_vaccinations
WHERE id = $1;