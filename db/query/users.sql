-- name: CreateUsers :one
INSERT INTO users (
  full_name, first_name, last_name, username, email, password
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 AND is_deleted = FALSE LIMIT 1;

-- name: GetUserForUpdate :one
SELECT * FROM users
WHERE id = $1 AND is_deleted = FALSE 
LIMIT 1
FOR NO KEY UPDATE;

-- name: ListUsers :many
SELECT * FROM users
WHERE is_deleted = FALSE
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users 
SET first_name = $2, last_name = $3, full_name = $4, email = $5, password = $6, updated_at = now()
WHERE id = $1 AND is_deleted = FALSE
RETURNING *;

-- name: UpdateUserNotesCountPlusOne :one
UPDATE users
SET notes_count = notes_count + 1
WHERE id = $1 AND is_deleted = FALSE
RETURNING *;

-- name: DeleteUser :one
UPDATE users 
SET is_deleted = TRUE, updated_at = now() 
WHERE id = $1 
RETURNING *;