-- name: CreateNote :one
INSERT INTO notes (
    user_id, title, description
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetNote :one
SELECT * FROM notes
WHERE id = $1 AND is_deleted = FALSE LIMIT 1;

-- name: ListNotes :many
SELECT * FROM notes
WHERE is_deleted = FALSE
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListNotesByUserId :many
SELECT * FROM notes
WHERE is_deleted = FALSE AND user_id = $3
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateNote :one
UPDATE notes
SET title = $2, description = $3, updated_at = now()
WHERE id = $1 AND is_deleted = FALSE
RETURNING *;

-- name: DeleteNote :one
UPDATE notes
SET is_deleted = TRUE, updated_at = now()
WHERE id = $1
RETURNING *;