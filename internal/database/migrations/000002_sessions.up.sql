CREATE TABLE IF NOT EXISTS sessions(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL,
    session_token uuid NOT NULL DEFAULT uuid_generate_v4(),
    refresh_token uuid NOT NULL DEFAULT uuid_generate_v4(),
    session_token_expires_at timestamp NOT NULL,
    refresh_token_expires_at timestamp NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp 
);