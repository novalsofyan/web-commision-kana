-- name: FindUsername :one
SELECT id, username, password FROM users
WHERE username = $1 LIMIT 1;

-- name: SetToken :one
INSERT INTO sessions (user_id, token, expires_at)
VALUES ($1, $2, NOW() + INTERVAL '30 days')
RETURNING token, expires_at;

-- name: SearchToken :one
SELECT token FROM sessions
WHERE token = $1 AND expires_at > NOW();

-- name: DeleteSessionByToken :exec
DELETE FROM sessions
WHERE token = $1;

-- name: SelectUserBySession :one
SELECT user_id FROM sessions
WHERE token = $1 AND expires_at > NOW();

-- name: CleanupExpiredSessions :exec
DELETE FROM sessions
WHERE expires_at <= NOW();

-- name: CreateProducts :exec
INSERT INTO products (nama_products, price, user_id)
VALUES ($1, $2, $3);

-- name: DeleteProduct :execresult
DELETE FROM products
WHERE id = $1;

-- name: GetProductAdmin :many
SELECT id, nama_products, price 
FROM products 
WHERE user_id = $1;

-- name: UpdateUser :one
UPDATE users
SET 
    username = COALESCE(sqlc.narg('username'), username),
    password = COALESCE(sqlc.narg('password'), password)
WHERE id = sqlc.arg('id')
RETURNING id, username;