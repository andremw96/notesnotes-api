// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"database/sql"
)

type Note struct {
	ID          int32          `json:"id"`
	UserID      int32          `json:"user_id"`
	Title       sql.NullString `json:"title"`
	Description sql.NullString `json:"description"`
	CreatedAt   sql.NullTime   `json:"created_at"`
}

type User struct {
	ID        int32          `json:"id"`
	FullName  sql.NullString `json:"full_name"`
	FirstName sql.NullString `json:"first_name"`
	LastName  sql.NullString `json:"last_name"`
	Username  sql.NullString `json:"username"`
	Email     sql.NullString `json:"email"`
	Password  sql.NullString `json:"password"`
	CreatedAt sql.NullTime   `json:"created_at"`
}
