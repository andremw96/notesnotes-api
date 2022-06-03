package db

import (
	db "andre/notesnotes-api/db/sqlc"
	"andre/notesnotes-api/util"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) db.User {
	firstName := util.RandomString(10)
	lastName := sql.NullString{String: util.RandomString(10), Valid: true}
	fullName := firstName + " " + lastName.String
	arg := db.CreateUsersParams{
		FullName:  fullName,
		LastName:  lastName,
		FirstName: firstName,
		Username:  util.RandomString(10),
		Email:     util.RandomString(20),
		Password:  util.RandomString(6),
	}

	user, err := testQueries.CreateUsers(context.Background(), arg)
	require.NoError(t, err) // check error must be null
	require.NotEmpty(t, user)

	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetNewtUser(t *testing.T) {
	newUser := createRandomUser(t)

	user, err := testQueries.GetUser(context.Background(), newUser.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, newUser.ID, user.ID)
	require.Equal(t, newUser.FullName, user.FullName)
	require.Equal(t, newUser.LastName, user.LastName)
	require.Equal(t, newUser.FirstName, user.FirstName)
	require.Equal(t, newUser.Username, user.Username)
	require.Equal(t, newUser.Email, user.Email)
	require.Equal(t, newUser.Password, user.Password)
	require.WithinDuration(t, newUser.CreatedAt, user.CreatedAt, time.Second)
	require.WithinDuration(t, newUser.UpdatedAt, user.UpdatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	newUser := createRandomUser(t)

	updatedFirstName := util.RandomString(10)
	updatedLastName := sql.NullString{String: util.RandomString(10), Valid: true}
	updatedFullName := updatedFirstName + " " + updatedLastName.String
	arg := db.UpdateUserParams{
		ID:        newUser.ID,
		FirstName: updatedFirstName,
		LastName:  updatedLastName,
		FullName:  updatedFullName,
		Email:     util.RandomString(20),
		Password:  util.RandomString(6),
	}

	updatedNewUser, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedNewUser)

	require.Equal(t, newUser.ID, updatedNewUser.ID)
	require.Equal(t, arg.FullName, updatedNewUser.FullName)
	require.Equal(t, arg.LastName, updatedNewUser.LastName)
	require.Equal(t, arg.FirstName, updatedNewUser.FirstName)
	require.Equal(t, newUser.Username, updatedNewUser.Username)
	require.Equal(t, arg.Email, updatedNewUser.Email)
	require.Equal(t, arg.Password, updatedNewUser.Password)
	require.WithinDuration(t, newUser.CreatedAt, updatedNewUser.CreatedAt, time.Second)
	require.WithinDuration(t, newUser.UpdatedAt, updatedNewUser.UpdatedAt, time.Second)
}

func TestDeleteUser(t *testing.T) {
	newUser := createRandomUser(t)

	deletedUser, err := testQueries.DeleteUser(context.Background(), newUser.ID)
	require.NoError(t, err)
	require.NotEmpty(t, deletedUser)

	require.Equal(t, deletedUser.ID, newUser.ID)
	require.Equal(t, true, deletedUser.IsDeleted)
	require.WithinDuration(t, newUser.UpdatedAt, deletedUser.UpdatedAt, time.Second)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := db.ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}
