package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// Store provides all functions to execute db queries and transactions
type Store interface {
	Querier
	InsertNewNote(ctx context.Context, arg InsertNoteTxParams) (InsertNoteTxResult, error)
}

// SQLStore -> to provides all functions to execute SQL queries and transactions
type SQLStore struct {
	// embedding Queries struct into Store, this is called COMPOSITION
	// prefrered way to extend struct in golang instead inheritance
	// by embedding Queries inside Store all individual query by Queries will be available for Store also
	*Queries
	db *sql.DB
}

// create new Store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// add function to the Store, to execute generic database transaction
// takes context and callback function as input
// it will start new db transaction
// create new Queries object with that transaction
// and call callback function with the created Queries
// finally commit or rollback transaction
// executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type InsertNoteTxParams struct {
	UserID      int32  `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type userResponse struct {
	FullName   string         `json:"full_name"`
	FirstName  string         `json:"first_name"`
	LastName   sql.NullString `json:"last_name"`
	Username   string         `json:"username"`
	Email      string         `json:"email"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	NotesCount int32          `json:"notes_count"`
}

func newUserResponse(user User) userResponse {
	return userResponse{
		FullName:   user.FullName,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Username:   user.Username,
		Email:      user.Email,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		NotesCount: user.NotesCount,
	}
}

type InsertNoteTxResult struct {
	Note Note         `json:"note"`
	User userResponse `json:"user"`
}

func (store *SQLStore) InsertNewNote(ctx context.Context, arg InsertNoteTxParams) (InsertNoteTxResult, error) {
	var result InsertNoteTxResult

	// create and run db transaction
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Note, err = q.CreateNote(ctx, CreateNoteParams{
			UserID:      arg.UserID,
			Title:       arg.Title,
			Description: sql.NullString{String: arg.Description, Valid: true},
		})
		if err != nil {
			return err
		}

		// get account -> update count
		user, err := q.GetUserForUpdate(ctx, arg.UserID)
		if err != nil {
			return err
		}

		resultUser, err := q.UpdateUserNotesCountPlusOne(ctx, user.ID)
		if err != nil {
			return err
		}

		result.User = newUserResponse(resultUser)

		return nil
	})

	return result, err
}
