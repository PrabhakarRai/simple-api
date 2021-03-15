package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions for execution of queries and TX
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTX executes database queries within a transaction
func (store *Store) execTX(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type StorageItemDownloadParams struct {
	ItemKey string `json:"item_key"`
	APIKey  string `json:"api_key"`
}

// StorageItemDownload for downloading an item in storage using key and API key
func (store *Store) StorageItemDownload(
	ctx context.Context,
	arg StorageItemDownloadParams) (
	Storage, error) {
	var result Storage

	err := store.execTX(ctx, func(q *Queries) error {
		var err error
		apiKeyDetails, err := q.GetAPIKeyDetailsByKey(ctx, arg.APIKey)
		if err != nil {
			return err
		}
		_ = q.UpdateAPIKeyHits(ctx, apiKeyDetails.Key)
		if apiKeyDetails.Enabled {
			result, err := q.GetStorageItemByKey(ctx, arg.ItemKey)
			if err != nil {
				_ = q.UpdateAPIKeyErrors(ctx, apiKeyDetails.Key)
				result = Storage{}
				return err
			}
			if result.CreatedBy != apiKeyDetails.Owner {
				result = Storage{}
				_ = q.UpdateAPIKeyErrors(ctx, apiKeyDetails.Key)
				_ = q.UpdateStorageErrors(ctx, arg.ItemKey)
				return fmt.Errorf("wrong API key for storage item")
			}
			_ = q.UpdateStorageDownload(ctx, arg.ItemKey)
			return nil
		}
		_ = q.UpdateAPIKeyErrors(ctx, apiKeyDetails.Key)
		return fmt.Errorf("API key disabled")
	})
	return result, err
}
