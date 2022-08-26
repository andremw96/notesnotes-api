// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: users.sql

package db

import (
	"context"
	"database/sql"
)

const createUsers = `-- name: CreateUsers :one
INSERT INTO users (
  username, email, password
) VALUES (
  $1, $2, $3
) RETURNING id, full_name, first_name, last_name, username, email, password, created_at, updated_at, is_deleted, notes_count
`

type CreateUsersParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) CreateUsers(ctx context.Context, arg CreateUsersParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUsers, arg.Username, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
		&i.NotesCount,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :one
UPDATE users 
SET is_deleted = TRUE, updated_at = now() 
WHERE id = $1 
RETURNING id, full_name, first_name, last_name, username, email, password, created_at, updated_at, is_deleted, notes_count
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, deleteUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
		&i.NotesCount,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, full_name, first_name, last_name, username, email, password, created_at, updated_at, is_deleted, notes_count FROM users
WHERE id = $1 AND is_deleted = FALSE LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
		&i.NotesCount,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, full_name, first_name, last_name, username, email, password, created_at, updated_at, is_deleted, notes_count FROM users
WHERE username = $1 AND is_deleted = FALSE LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
		&i.NotesCount,
	)
	return i, err
}

const getUserForUpdate = `-- name: GetUserForUpdate :one
SELECT id, full_name, first_name, last_name, username, email, password, created_at, updated_at, is_deleted, notes_count FROM users
WHERE id = $1 AND is_deleted = FALSE 
LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetUserForUpdate(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserForUpdate, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
		&i.NotesCount,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, full_name, first_name, last_name, username, email, password, created_at, updated_at, is_deleted, notes_count FROM users
WHERE is_deleted = FALSE
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FullName,
			&i.FirstName,
			&i.LastName,
			&i.Username,
			&i.Email,
			&i.Password,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.IsDeleted,
			&i.NotesCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
UPDATE users 
SET first_name = $2, last_name = $3, full_name = $4, email = $5, password = $6, updated_at = now()
WHERE id = $1 AND is_deleted = FALSE
RETURNING id, full_name, first_name, last_name, username, email, password, created_at, updated_at, is_deleted, notes_count
`

type UpdateUserParams struct {
	ID        int32          `json:"id"`
	FirstName sql.NullString `json:"first_name"`
	LastName  sql.NullString `json:"last_name"`
	FullName  sql.NullString `json:"full_name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.FullName,
		arg.Email,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
		&i.NotesCount,
	)
	return i, err
}

const updateUserNotesCountPlusOne = `-- name: UpdateUserNotesCountPlusOne :one
UPDATE users
SET notes_count = notes_count + 1
WHERE id = $1 AND is_deleted = FALSE
RETURNING id, full_name, first_name, last_name, username, email, password, created_at, updated_at, is_deleted, notes_count
`

func (q *Queries) UpdateUserNotesCountPlusOne(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserNotesCountPlusOne, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
		&i.NotesCount,
	)
	return i, err
}
