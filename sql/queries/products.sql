-- name: GetProduct :one
SELECT * FROM products WHERE id = ? LIMIT 1;

-- name: ListProducts :many
SELECT * FROM products ORDER BY id;

-- name: CreateProduct :one
INSERT INTO products (id, name, image, price)
VALUES (?, ?, ?, ?) RETURNING *;

-- name: UpdateProduct :exec
UPDATE products
SET name = ?, image = ?, price = ?
WHERE id = ?;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = ?;
