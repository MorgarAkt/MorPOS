-- +goose Up
CREATE TABLE salons (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE tables (
    id TEXT PRIMARY KEY,
    number INTEGER NOT NULL,
    salon_id TEXT NOT NULL,
    is_full INTEGER NOT NULL DEFAULT 0,
    total_bill REAL DEFAULT 0,
    FOREIGN KEY (salon_id) REFERENCES salons(id)
);

-- +goose Down
DROP TABLE tables;
DROP TABLE salons;
