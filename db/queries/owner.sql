-- name: CreateOwner :one
INSERT INTO owners (photo, name, last_name, user_id, birthday, genre, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, photo, name, last_name, user_id, birthday, genre, created_at, updated_at;

-- name: GetOwnerByID :one
SELECT id, photo, name, last_name, user_id, birthday, genre, created_at, updated_at
FROM owners
WHERE id = $1;

-- name: GetOwnerByUserID :one
SELECT id, photo, name, last_name, user_id, birthday, genre, created_at, updated_at
FROM owners
WHERE user_id = $1;

-- name: ListOwners :many
SELECT id, photo, name, last_name, user_id, birthday, genre, created_at, updated_at
FROM owners
ORDER BY id;

-- name: UpdateOwner :exec
UPDATE owners
SET photo = $2, name = $3, last_name = $4, user_id = $5, birthday = $6, genre = $7, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeleteOwner :exec
DELETE FROM owners
WHERE id = $1;