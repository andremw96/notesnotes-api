package db

import (
	db "andre/notesnotes-api/db/sqlc"

	"andre/notesnotes-api/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsertNewNote(t *testing.T) {
	store := db.NewStore(testDb)

	user := createRandomUser(t)

	// run n concurrent transfer transactions make sure transaction goes well
	n := 5

	// we can't use testify because goroutine
	// we verify it by send them back to the main goroutine our test is running on then check there from there
	// use channels -> connect concurrent goroutine allow them safely share data with each other without explicit locking
	errs := make(chan error)
	results := make(chan db.InsertNoteTxResult)

	for i := 0; i < n; i++ {
		go func() {
			title := util.RandomString(10)
			description := util.RandomString(30)

			result, err := store.InsertNewNote(context.Background(), db.InsertNoteTxParams{
				UserID:      user.ID,
				Title:       title,
				Description: description,
			})

			errs <- err
			results <- result
		}()
	}

	// check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check inserted note
		insertedNote := result.Note
		require.NotEmpty(t, insertedNote)
		require.Equal(t, user.ID, insertedNote.UserID)
		require.NotEmpty(t, insertedNote.Title)
		require.NotEmpty(t, insertedNote.Description)
		require.NotZero(t, insertedNote.ID)
		require.False(t, insertedNote.IsDeleted)
		require.NotZero(t, insertedNote.CreatedAt)
		require.NotZero(t, insertedNote.UpdatedAt)

		_, err = store.GetNote(context.Background(), insertedNote.ID)
		require.NoError(t, err)

		// check user notes count
	}
}
