package db

import (
	"context"
	"testing"

	"github.com/PrabhakarRai/simple-api/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAPIKey(t *testing.T) CreateAPIKeyRow {
	arg := CreateAPIKeyParams{
		Key:   utils.RandomAPIKey(),
		Owner: createRandomUser(t),
	}

	row, err := testQueries.CreateAPIKey(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, row)
	require.Equal(t, row.Key, arg.Key)
	require.Equal(t, row.Owner, arg.Owner)
	return row
}

func TestCreateAPIKey(t *testing.T) {
	_ = createRandomAPIKey(t)
}

func TestDeleteAPIKeyByAPIKey(t *testing.T) {
	data := createRandomAPIKey(t)
	err := testQueries.DeleteAPIKeyByAPIKey(context.Background(), data.Key)
	require.NoError(t, err)
}

func TestDeleteAPIKeysByUserID(t *testing.T) {
	data := createRandomAPIKey(t)
	err := testQueries.DeleteAPIKeysByUserID(context.Background(), data.Owner)
	require.NoError(t, err)
}

func TestDeleteAPIKeysByUsername(t *testing.T) {
	data := createRandomAPIKey(t)
	usr, err := testQueries.GetUserByID(context.Background(), data.Owner)
	require.NoError(t, err)
	require.NotEmpty(t, usr)
	err = testQueries.DeleteAPIKeysByUsername(context.Background(), usr.Username)
	require.NoError(t, err)
}

func TestGetAPIKeyDetailsByKey(t *testing.T) {
	data := createRandomAPIKey(t)
	res, err := testQueries.GetAPIKeyDetailsByKey(context.Background(), data.Key)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, res.ID, data.ID)
	require.Equal(t, res.Key, data.Key)
	require.Equal(t, res.Owner, data.Owner)
	require.Equal(t, res.Enabled, true)
	require.Equal(t, res.Hits, int32(0))
	require.Equal(t, res.Errors, int32(0))
}

func TestGetAPIKeysByOwner(t *testing.T) {
	data := createRandomAPIKey(t)
	res, err := testQueries.GetAPIKeysByOwner(context.Background(), data.Owner)
	require.NoError(t, err)
	require.NotEmpty(t, res)
}

func TestGetAPIKeysByUsername(t *testing.T) {
	data := createRandomAPIKey(t)
	usr, _ := testQueries.GetUserByID(context.Background(), data.Owner)
	res, err := testQueries.GetAPIKeysByUsername(context.Background(), usr.Username)
	require.NoError(t, err)
	require.NotEmpty(t, res)
}

func TestUpdateAPIKeyEnabled(t *testing.T) {
	data := createRandomAPIKey(t)
	arg := UpdateAPIKeyEnabledParams{
		Key:     data.Key,
		Enabled: false,
	}
	err := testQueries.UpdateAPIKeyEnabled(context.Background(), arg)
	require.NoError(t, err)
}

func TestUpdateAPIKeyErrors(t *testing.T) {
	data := createRandomAPIKey(t)
	err := testQueries.UpdateAPIKeyErrors(context.Background(), data.Key)
	require.NoError(t, err)
	res, err := testQueries.GetAPIKeyDetailsByKey(context.Background(), data.Key)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, res.Errors, int32(1))
}

func TestUpdateAPIKeyHits(t *testing.T) {
	data := createRandomAPIKey(t)
	err := testQueries.UpdateAPIKeyHits(context.Background(), data.Key)
	require.NoError(t, err)
	res, err := testQueries.GetAPIKeyDetailsByKey(context.Background(), data.Key)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, res.Hits, int32(1))
}
