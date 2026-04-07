CREATE TABLE tickets (
    id SERIAL PRIMARY KEY,
    event_name VARCHAR(255) NOT NULL,
    price INT NOT NULL
)