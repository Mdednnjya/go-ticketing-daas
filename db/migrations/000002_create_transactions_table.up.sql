CREATE TABLE transactions (
    ID SERIAL PRIMARY KEY,
    ticket_id INT REFERENCES tickets(id),
    buyer_email VARCHAR(255),
    quantity INT,
    total_price INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)