-- name: CreateProfile :one
INSERT INTO profiles (user_id, bio, owner_id, veterinarian_id, profile_pic, created_at, last_update)
VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING *;

-- name: UpdateUserProfile :one
UPDATE profiles
SET 
    bio = $2, 
    profile_pic = $3, 
    owner_id = $4,
    veterinarian_id = $5,
    last_update = CURRENT_TIMESTAMP
WHERE 
    user_id = $1
RETURNING *;

-- name: DeleteUserProfile :exec
DELETE FROM profiles
WHERE user_id = $1;

-- name: GetUserProfile :one
SELECT * FROM profiles
WHERE user_id = $1;