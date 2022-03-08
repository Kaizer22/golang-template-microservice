CREATE TABLE IF NOT EXISTS products (
    id INT GENERATED ALWAYS AS IDENTITY,
    name VARCHAR,
    description TEXT,
    category_id INT,
    PRIMARY KEY (id),
    CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES categories(id)
);
