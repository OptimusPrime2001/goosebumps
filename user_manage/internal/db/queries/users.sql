-- name: CreateUser :one
INSERT INTO users(
  user_email,
  user_password,
  user_full_name,
  user_age,
  user_status,
  user_level
)
VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;


-- name: UpdateUser :one
UPDATE users
SET
  user_password = COALESCE(sqlc.narg(user_password), user_password),
  user_full_name = COALESCE(sqlc.narg(user_full_name), user_full_name),
  user_age = COALESCE(sqlc.narg(user_age), user_age),
  user_status = COALESCE(sqlc.narg(user_status), user_status),
  user_level = COALESCE(sqlc.narg(user_level), user_level)
WHERE
  user_id = sqlc.arg(user_id)
  AND user_deleted_at IS NULL
RETURNING *;

-- name: SoftDeleteUser :one
UPDATE users
SET
  user_deleted_at = NOW()
WHERE
  user_id = sqlc.arg(user_id)::uuid
  AND user_deleted_at IS NULL
RETURNING *;

-- name: RestoreUser :one
UPDATE users
SET
  user_deleted_at = NULL
WHERE
  user_id = sqlc.arg(user_id)::uuid
  AND user_deleted_at IS NOT NULL
RETURNING *;

-- name: TrashUsers :one
DELETE FROM users
WHERE
  user_deleted_at IS NOT NULL
RETURNING *;
