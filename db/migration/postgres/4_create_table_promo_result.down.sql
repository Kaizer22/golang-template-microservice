CREATE TABLE IF NOT EXISTS promotion_result (
    id INT GENERATED ALWAYS AS IDENTITY,
    winner_id INT,
    description TEXT,
    category_id INT,
    PRIMARY KEY (id),
    CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES categories(id)
);
