package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store -> to provides all functions to execute db queries and transactions
type Store struct {
	// embedding Queries struct into Store, this is called COMPOSITION
	// prefrered way to extend struct in golang instead inheritance
	// by embedding Queries inside Store all individual query by Queries will be available for Store also
	*Queries
	db *sql.DB
}

// create new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
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
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
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

type InsertNoteTxResult struct {
	Note Note `json:"note"`
	User User `json:"user"`
}

func (store *Store) InsertNewNote(ctx context.Context, arg InsertNoteTxParams) (InsertNoteTxResult, error) {
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

		// TODO: UPDATE USER COUNT NOTES

		return nil
	})

	return result, err
}
