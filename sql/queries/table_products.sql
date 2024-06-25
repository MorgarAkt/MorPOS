-- name: AddProductToTable :exec
INSERT INTO table_products (table_id, product_id)
VALUES (?, ?);

-- name: RemoveProductFromTable :exec
DELETE FROM table_products WHERE table_id = ? AND product_id = ?;

-- name: ListProductsByTable :many
SELECT p.* FROM products p
INNER JOIN table_products tp ON tp.product_id = p.id
WHERE tp.table_id = ?;
