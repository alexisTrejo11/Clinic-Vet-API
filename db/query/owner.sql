-- name: CreateOwner :one
INSERT INTO owners (photo, name, user_id, created_at, updated_at)
VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, photo, name, user_id, created_at, updated_at;

-- name: GetOwnerByID :one
SELECT id, photo, name, user_id, created_at, updated_at
FROM owners
WHERE id = $1;

-- name: ListOwners :many
SELECT id, photo, name, user_id, created_at, updated_at
FROM owners
ORDER BY id;

-- name: UpdateOwner :exec
UPDATE owners
SET photo = $2, name = $3, user_id = $4, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeleteOwner :exec
DELETE FROM owners
WHERE id = $1;
