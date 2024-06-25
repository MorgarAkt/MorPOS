-- +goose Up
CREATE TABLE table_products (
    table_id TEXT NOT NULL,
    product_id TEXT NOT NULL,
    PRIMARY KEY (table_id, product_id),
    FOREIGN KEY (table_id) REFERENCES tables(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

-- +goose Down
DROP TABLE table_products;
