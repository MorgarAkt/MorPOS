-- +goose Up
CREATE TABLE products (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    image BLOB,
    price REAL NOT NULL
);

-- +goose Down
DROP TABLE products;
