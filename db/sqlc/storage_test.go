package db

import (
	"context"
	"testing"

	"github.com/PrabhakarRai/simple-api/utils"
	"github.com/stretchr/testify/require"
)

func createRandomStorageItem(t *testing.T) CreateStorageItemRow {
	arg := CreateStorageItemParams{
		Key:       utils.RandomKey(),
		Value:     utils.RandomValue(),
		CreatedBy: createRandomUser(t),
	}

	data, err := testQueries.CreateStorageItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, data)
	require.Equal(t, arg.Key, data.Key)
	require.Equal(t, arg.CreatedBy, data.CreatedBy)
	return data
}

func TestCreateStorageItem(t *testing.T) {
	_ = createRandomStorageItem(t)
}

func TestGetStorageItemByKey(t *testing.T) {
	data := createRandomStorageItem(t)
	result, err := testQueries.GetStorageItemByKey(context.Background(), data.Key)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, data.ID, result.ID)
}

func TestGetStorageItemsByUserID(t *testing.T) {
	data := createRandomStorageItem(t)
	user, err := testQueries.GetUserByID(context.Background(), data.CreatedBy)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	result, err := testQueries.GetStorageItemsByUserID(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestGetStorageItemsByUsername(t *testing.T) {
	data := createRandomStorageItem(t)
	user, err := testQueries.GetUserByID(context.Background(), data.CreatedBy)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	result, err := testQueries.GetStorageItemsByUsername(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestDeleteStorageItemByKey(t *testing.T) {
	data := createRandomStorageItem(t)
	err := testQueries.DeleteStorageItemByKey(context.Background(), data.Key)
	require.NoError(t, err)
	result, err := testQueries.GetStorageItemByKey(context.Background(), data.Key)
	require.Error(t, err)
	require.Empty(t, result)
}

func TestDeleteStorageItemsByUserID(t *testing.T) {
	data := createRandomStorageItem(t)
	err := testQueries.DeleteStorageItemsByUserID(context.Background(), data.CreatedBy)
	require.NoError(t, err)
	result, err := testQueries.GetStorageItemsByUserID(context.Background(), data.CreatedBy)
	require.Error(t, err)
	require.Empty(t, result)
}
