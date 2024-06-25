-- name: GetTable :one
SELECT * FROM tables WHERE id = ? LIMIT 1;

-- name: ListTables :many
SELECT * FROM tables ORDER BY id;

-- name: CreateTable :one
INSERT INTO tables (id, number, salon_id, is_full, total_bill)
VALUES (?, ?, ?, ?, ?) RETURNING *;

-- name: UpdateTable :exec
UPDATE tables
SET number = ?, salon_id = ?, is_full = ?, total_bill = ?
WHERE id = ?;

-- name: DeleteTable :exec
DELETE FROM tables WHERE id = ?;
