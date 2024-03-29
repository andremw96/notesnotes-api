// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"time"
)

type Note struct {
	ID          int32          `json:"id"`
	UserID      int32          `json:"user_id"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	IsDeleted   bool           `json:"is_deleted"`
}

type User struct {
	ID         int32          `json:"id"`
	FullName   sql.NullString `json:"full_name"`
	FirstName  sql.NullString `json:"first_name"`
	LastName   sql.NullString `json:"last_name"`
	Username   string         `json:"username"`
	Email      string         `json:"email"`
	Password   string         `json:"password"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	IsDeleted  bool           `json:"is_deleted"`
	NotesCount int32          `json:"notes_count"`
}
