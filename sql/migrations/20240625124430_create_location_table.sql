-- +goose Up
CREATE TABLE locations (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE locations;
