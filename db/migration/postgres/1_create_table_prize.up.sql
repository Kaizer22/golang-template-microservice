CREATE TABLE IF NOT EXISTS prize (
    id INT GENERATED ALWAYS AS IDENTITY,
    description TEXT,
    PRIMARY KEY (id)

--     name VARCHAR,
--     category_id INT,
--     PRIMARY KEY (id),
--     CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES categories(id)
);
