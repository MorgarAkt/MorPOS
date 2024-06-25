-- name: GetSalon :one
SELECT * FROM salons WHERE id = ? LIMIT 1;

-- name: ListSalons :many
SELECT * FROM salons ORDER BY id;

-- name: CreateSalon :one
INSERT INTO salons (id, name)
VALUES (?, ?) RETURNING *;

-- name: UpdateSalon :exec
UPDATE salons
SET name = ?
WHERE id = ?;

-- name: DeleteSalon :exec
DELETE FROM salons WHERE id = ?;
