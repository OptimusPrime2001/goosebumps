-- name: GetUserByID :one
SELECT * FROM users WHERE user_id = $1;

-- name: CreateUser :one
INSERT INTO users (user_id, username, email, age, password, status, level)
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;
