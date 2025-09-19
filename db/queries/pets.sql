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

-- name: FindAllPets :many
SELECT * FROM pets
WHERE deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CountAllPets :one
SELECT COUNT(*) FROM pets
WHERE deleted_at IS NULL;

-- name: FindPetsBySpecification :many
SELECT * FROM pets
WHERE deleted_at IS NULL
AND ($1::text IS NULL OR name ILIKE '%' || $1 || '%')
AND ($2::text IS NULL OR species = $2)
AND ($3::text IS NULL OR breed = $3)
AND ($4::int IS NULL OR customer_id = $4)
AND ($5::bool IS NULL OR is_active = $5)
AND ($6::bool IS NULL OR is_neutered = $6)
ORDER BY created_at DESC
LIMIT $7 OFFSET $8;

-- name: CountPetsBySpecification :one
SELECT COUNT(*) FROM pets
WHERE deleted_at IS NULL
AND ($1::text IS NULL OR name ILIKE '%' || $1 || '%')
AND ($2::text IS NULL OR species = $2)
AND ($3::text IS NULL OR breed = $3)
AND ($4::int IS NULL OR customer_id = $4)
AND ($5::bool IS NULL OR is_active = $5)
AND ($6::bool IS NULL OR is_neutered = $6);

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
    weight, 
    color, 
    microchip, 
    is_neutered, 
    customer_id, 
    allergies, 
    current_medications, 
    special_needs, 
    is_active,
    date_of_birth,
    insurance_info,
    veterinary_contact,
    feeding_instructions,
    behavioral_notes
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15,
    $16, $17, $18, $19, $20
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
    weight = $8,
    color = $9,
    microchip = $10,
    is_neutered = $11,
    customer_id = $12,
    allergies = $13,
    current_medications = $14,
    special_needs = $15,
    is_active = $16,
    date_of_birth = $17,
    insurance_info = $18,
    veterinary_contact = $19,
    feeding_instructions = $20,
    behavioral_notes = $21,
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