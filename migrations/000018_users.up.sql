CREATE TABLE users
(
    id UUID PRIMARY KEY,
    phone TEXT,
    password TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
)