-- name: CreatePet :one
INSERT INTO pets (name, photo, species, breed, age, owner_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, name, photo, species, breed, age, owner_id, created_at, updated_at;

-- name: GetPetByID :one
SELECT id, name, photo, species, breed, age, owner_id, created_at, updated_at
FROM pets
WHERE id = $1;

-- name: ListPets :many
SELECT id, name, photo, species, breed, age, owner_id, created_at, updated_at
FROM pets
ORDER BY id;

-- name: ListPetsByOwnerByID :many
SELECT id, name, photo, species, breed, age, owner_id, created_at, updated_at
FROM pets
WHERE owner_id = $1;

-- name: UpdatePet :exec
UPDATE pets
SET name = $2, photo = $3, species = $4, breed = $5, age = $6, owner_id = $7, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeletePet :exec
DELETE FROM pets
WHERE id = $1;
