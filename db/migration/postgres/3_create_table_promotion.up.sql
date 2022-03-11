CREATE TABLE IF NOT EXISTS promotion (
    id INT GENERATED ALWAYS AS IDENTITY,
    name VARCHAR,
    description TEXT,
    prizes INT[],
    participants INT[],
    PRIMARY KEY (id)
);