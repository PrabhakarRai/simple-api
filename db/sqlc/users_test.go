package db

import (
	"context"
	"testing"

	"github.com/PrabhakarRai/simple-api/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) int32 {
	arg := CreateUserParams{
		Username: utils.RandomUsername(),
		Name:     utils.RandomName(),
	}

	id, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, id)

	return id
}

func TestCreateUser(t *testing.T) {
	_ = createRandomUser(t)
}

func TestGetUserByID(t *testing.T) {
	id := createRandomUser(t)
	usr, err := testQueries.GetUserByID(context.Background(), id)
	require.NoError(t, err)
	require.NotEmpty(t, usr)
	require.NotEmpty(t, usr.ID)
	require.NotEmpty(t, usr.Name)
	require.NotEmpty(t, usr.Username)
	require.Equal(t, usr.ID, id)
}

func TestGetUserByUsername(t *testing.T) {
	id := createRandomUser(t)
	usr, err := testQueries.GetUserByID(context.Background(), id)
	require.NoError(t, err)
	require.NotEmpty(t, usr)
	require.Equal(t, usr.ID, id)
	usr2, err := testQueries.GetUserByUsername(context.Background(), usr.Username)
	require.NoError(t, err)
	require.NotEmpty(t, usr2)

	require.Equal(t, usr.ID, usr2.ID)
	require.Equal(t, usr.Name, usr2.Name)
	require.Equal(t, usr.Username, usr2.Username)
}

func TestDeleteUser(t *testing.T) {
	id := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), id)
	require.NoError(t, err)
}

func TestUpdateUserName(t *testing.T) {
	id := createRandomUser(t)
	usr, _ := testQueries.GetUserByID(context.Background(), id)
	arg := UpdateUserNameParams{
		ID:   id,
		Name: utils.RandomName(),
	}
	err := testQueries.UpdateUserName(context.Background(), arg)
	require.NoError(t, err)
	newusr, err := testQueries.GetUserByID(context.Background(), arg.ID)
	require.NoError(t, err)
	require.NotEqual(t, newusr.Name, usr.Name)
	require.Equal(t, newusr.ID, usr.ID)
}
