-- name: CreateAPIKey :one
INSERT INTO api_keys (key, owner) VALUES ($1, $2) RETURNING id, key, owner;

-- name: GetAPIKeyDetailsByKey :one
SELECT * FROM api_keys
WHERE key = $1 LIMIT 1;

-- name: GetAPIKeysByOwner :many
SELECT * FROM api_keys
WHERE owner = $1;

-- name: GetAPIKeysByUsername :many
SELECT * FROM api_keys
WHERE api_keys.owner = (SELECT id FROM users WHERE users.username = $1 LIMIT 1);

-- name: DeleteAPIKeyByAPIKey :exec
DELETE FROM api_keys WHERE key = $1;

-- name: DeleteAPIKeysByUserID :exec
DELETE FROM api_keys WHERE owner = $1;

-- name: DeleteAPIKeysByUsername :exec
DELETE FROM api_keys WHERE api_keys.owner = (SELECT id FROM users WHERE users.username = $1);

-- name: UpdateAPIKeyEnabled :exec
UPDATE api_keys SET enabled = $2 WHERE key = $1;

-- name: UpdateAPIKeyHits :exec
UPDATE api_keys SET hits = hits+1 WHERE key = $1;

-- name: UpdateAPIKeyErrors :exec
UPDATE api_keys SET errors = errors+1 WHERE key = $1;
