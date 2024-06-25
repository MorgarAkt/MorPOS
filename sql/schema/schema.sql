-- Table: users
CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    full_name TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role TEXT NOT NULL
);

-- Table: locations
CREATE TABLE IF NOT EXISTS locations (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL
);

-- Table: products
CREATE TABLE IF NOT EXISTS products (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    image BLOB,
    price REAL NOT NULL
);

-- Table: salons
CREATE TABLE IF NOT EXISTS salons (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL
);

-- Table: tables
CREATE TABLE IF NOT EXISTS tables (
    id TEXT PRIMARY KEY,
    number INTEGER NOT NULL,
    salon_id TEXT NOT NULL,
    is_full INTEGER NOT NULL DEFAULT 0,
    total_bill REAL DEFAULT 0,
    FOREIGN KEY (salon_id) REFERENCES salons(id)
);

-- Table: table_products (junction table)
CREATE TABLE IF NOT EXISTS table_products (
    table_id TEXT NOT NULL,
    product_id TEXT NOT NULL,
    PRIMARY KEY (table_id, product_id),
    FOREIGN KEY (table_id) REFERENCES tables(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);
