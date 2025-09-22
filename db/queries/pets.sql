-- name: FindPetByID :one
SELECT * FROM pets
WHERE id = $1 AND deleted_at IS NULL;

-- name: FindPetByIDAndCustomerID :one
SELECT * FROM pets
WHERE id = $1 AND customer_id = $2 AND deleted_at IS NULL;

-- name: FindPetsByCustomerID :many
SELECT * FROM pets
WHERE customer_id = $1 AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: FindPetsBySpecies :many
SELECT * FROM pets
WHERE species = $1 AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: CountPetsBySpecies :one
SELECT COUNT(*) FROM pets
WHERE species = $1 AND deleted_at IS NULL;

-- name: CountPetsByCustomerID :one
SELECT COUNT(*) FROM pets
WHERE customer_id = $1 AND deleted_at IS NULL;

-- name: ExistsPetByID :one
SELECT COUNT(*) > 0 FROM pets
WHERE id = $1 AND deleted_at IS NULL;

-- name: ExistsPetByMicrochip :one
SELECT COUNT(*) > 0 FROM pets
WHERE microchip = $1 AND deleted_at IS NULL;

-- name: CreatePet :one
INSERT INTO pets (
    name, 
    photo, 
    species, 
    breed, 
    age, 
    gender,
    color, 
    microchip, 
    tattoo,
    blood_type,
    is_neutered, 
    customer_id,  
    is_active,
    created_at, 
    updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
)
RETURNING *;

-- name: UpdatePet :one
UPDATE pets
SET 
    name = $2,
    photo = $3,
    species = $4,
    breed = $5,
    age = $6,
    gender = $7,
    color = $8,
    microchip = $9,
    is_neutered = $10,
    customer_id = $11,
    tattoo = $12,
    blood_type = $13,
    is_active = $14,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: SoftDeletePet :exec
UPDATE pets
SET 
    deleted_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP,
    is_active = FALSE
WHERE id = $1;

-- name: HardDeletePet :exec
DELETE FROM pets WHERE id = $1;


-- name: RestorePet :exec
UPDATE pets
SET 
    deleted_at = NULL,
    updated_at = CURRENT_TIMESTAMP,
    is_active = TRUE
WHERE id = $1;