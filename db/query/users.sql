-- name: CreateUsers :one
INSERT INTO users (
  full_name, first_name, last_name, username, email, password
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users 
SET first_name = $2, last_name = $3, full_name = $4, email = $5, password = $6
WHERE id = $1 
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;