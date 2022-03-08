CREATE TABLE IF NOT EXISTS categories (
    id INT GENERATED ALWAYS AS IDENTITY,
    name VARCHAR,
    description TEXT,
    PRIMARY KEY (id)
);