CREATE TABLE transactions
(
    id uuid uuid DEFAULT uuid_generate_v4(),
    account_id VARCHAR NOT NULL,
    amount DECIMAL,
    status VARCHAR NOT NULL,
    error_message VARCHAR,
    created_at VARCHAR,
    updated_at VARCHAR
);