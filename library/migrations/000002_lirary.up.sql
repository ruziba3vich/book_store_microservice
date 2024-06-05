CREATE TABLE IF NOT EXISTS books(
    ID SERIAL PRIMARY KEY,
    author_id bigint REFERENCES(authors(id)),
    title VARCHAR(255),
    price DECIMAL(8,2),
    page INT
)