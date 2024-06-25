-- name: GetLocation :one
SELECT * FROM locations WHERE id = ? LIMIT 1;

-- name: ListLocations :many
SELECT * FROM locations ORDER BY id;

-- name: CreateLocation :one
INSERT INTO locations (id, name)
VALUES (?, ?) RETURNING *;

-- name: UpdateLocation :exec
UPDATE locations
SET name = ?
WHERE id = ?;

-- name: DeleteLocation :exec
DELETE FROM locations WHERE id = ?;
