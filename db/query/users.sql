-- name: CreateUser :one
INSERT INTO users (username, name) VALUES ($1, $2) RETURNING id;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;