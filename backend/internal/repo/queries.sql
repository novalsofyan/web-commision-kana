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
INSERT INTO products (nama_products, price, user_id)
VALUES ($1, $2, $3);

-- name: DeleteProduct :execresult
DELETE FROM products
WHERE id = $1;

-- name: GetProductAdmin :many
SELECT id, nama_products, price 
FROM products 
WHERE user_id = $1;