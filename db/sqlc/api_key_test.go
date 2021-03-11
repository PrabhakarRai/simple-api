package db

import (
	"context"
	"testing"

	"github.com/PrabhakarRai/simple-api/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAPIKey(t *testing.T) (key int32) {
	arg := CreateAPIKeyParams{
		Key:   utils.RandomAPIKey(),
		Owner: createRandomUser(t),
	}

	key, err := testQueries.CreateAPIKey(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, key)
	return
}

func TestCreateAPIKey(t *testing.T) {
	_ = createRandomAPIKey(t)
}

func TestGetAPIKeyByUsername(t *testing.T) {

}

func TestGetAPIKeyByUserID(t *testing.T) {

}

func TestDeleteAPIKeysByUserID(t *testing.T) {

}

func TestDeleteAPIKeysByUsername(t *testing.T) {

}
