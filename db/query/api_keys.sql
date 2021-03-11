-- name: CreateAPIKey :one
INSERT INTO api_keys (key, owner) VALUES ($1, $2) RETURNING id;

-- name: GetAPIKeyByUsername :one
SELECT * FROM api_keys
WHERE api_keys.owner = (SELECT id FROM users WHERE users.username = $1 LIMIT 1);

-- name: GetAPIKeyByUserID :one
SELECT * FROM api_keys
WHERE api_keys.owner = $1;

-- name: DeleteAPIKeysByUserID :exec
DELETE FROM api_keys WHERE owner = $1;

-- name: DeleteAPIKeysByUsername :exec
DELETE FROM api_keys WHERE api_keys.owner = (SELECT id FROM users WHERE users.username = $1 LIMIT 1);