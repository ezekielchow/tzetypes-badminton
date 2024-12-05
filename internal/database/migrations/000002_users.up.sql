CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    firebase_uid TEXT UNIQUE NOT NULL,
    email varchar(255) NOT NULL,
    account_tier text NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp 
);