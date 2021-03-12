-- name: CreateStorageItem :one
INSERT INTO storage (key, value, created_by) VALUES ($1, $2, $3) RETURNING id, key, created_by;

-- name: GetStorageItemByKey :one
SELECT * FROM storage
WHERE key = $1 LIMIT 1;

-- name: GetStorageItemsByUserID :many
SELECT (id, key) FROM storage
WHERE created_by = $1;

-- name: GetStorageItemsByUsername :many
SELECT (id, key) FROM storage
WHERE created_by = (SELECT id FROM users WHERE users.username = $1 LIMIT 1);

-- name: DeleteStorageItemByKey :exec
DELETE FROM storage WHERE key = $1;

-- name: DeleteStorageItemsByUserID :exec
DELETE FROM storage WHERE created_by = $1;

-- name: UpdateStorageValue :exec
UPDATE storage SET value = $2 WHERE key = $1;

-- name: UpdateStorageDownload :exec
UPDATE storage SET downloads = downloads+1 WHERE key = $1;

-- name: UpdateStorageErrors :exec
UPDATE storage SET errors = errors+1 WHERE key = $1;

-- name: UpdateStorageAvailable :exec
UPDATE storage SET available = $2 WHERE key = $1;

-- name: UpdateStorageAvailableByUserID :exec
UPDATE storage SET available = $2 WHERE created_by = $1;
