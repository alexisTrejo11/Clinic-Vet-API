-- name: GetPetByID :one
SELECT * FROM pets
WHERE id = $1;

-- name: GetPetByIDAndCustomerID :one
SELECT * FROM pets
WHERE id = $1 AND customer_id = $2;

-- name: GetPetsByCustomerID :many
SELECT * FROM pets
WHERE customer_id = $1
ORDER BY id;

-- name: ListPets :many
SELECT * FROM pets
ORDER BY id;

-- name: CreatePet :one
INSERT INTO pets (
    name, 
    photo, 
    species, 
    breed, 
    age, 
    gender, 
    weight, 
    color, 
    microchip, 
    is_neutered, 
    customer_id, 
    allergies, 
    current_medications, 
    special_needs, 
    is_active,
    created_at,
    updated_at
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15,
    CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
)
RETURNING *; 

-- name: UpdatePet :exec
UPDATE pets
SET 
    name = $2,
    photo = $3,
    species = $4,
    breed = $5,
    age = $6,
    gender = $7,
    weight = $8,
    color = $9,
    microchip = $10,
    is_neutered = $11,
    customer_id = $12,
    allergies = $13,
    current_medications = $14,
    special_needs = $15,
    is_active = $16,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeletePet :exec
DELETE FROM pets
WHERE id = $1;
