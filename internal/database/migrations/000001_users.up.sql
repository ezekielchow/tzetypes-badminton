CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    email varchar(255) NOT NULL,
    password_hash text,
    created_at timestamp DEFAULT now(),
    updated_at timestamp 
);