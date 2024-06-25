-- name: GetUser :one
SELECT * FROM users WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users ORDER BY id;

-- name: CreateUser :one
INSERT INTO users (id, full_name, username, password, role)
VALUES (?, ?, ?, ?, ?) RETURNING *;

-- name: UpdateUser :exec
UPDATE users
SET full_name = ?, username = ?, password = ?, role = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;
