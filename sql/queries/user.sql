-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY first_name;

-- name: CreateUser :one
INSERT INTO users (
  first_name, last_name, email, password
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
  SET first_name = $2, last_name = $3, email = $4, password = $5
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
