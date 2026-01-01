-- name: FindUsername :one
SELECT id, username, password FROM users
WHERE username = $1 LIMIT 1;

-- name: SetToken :one
INSERT INTO sessions (user_id, token)
VALUES ($1, $2)
RETURNING token;

-- name: SearchToken :one
SELECT token FROM sessions
WHERE token = $1;

-- name: DeleteSessionByToken :exec
DELETE FROM sessions
WHERE token = $1;

-- name: SelectUserBySession :one
SELECT user_id FROM sessions
WHERE token = $1;

-- name: CreateProducts :exec
INSERT INTO products (nama_products, price)
VALUES ($1, $2);

-- name: 