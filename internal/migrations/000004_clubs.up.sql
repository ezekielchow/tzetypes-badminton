CREATE TABLE IF NOT EXISTS clubs(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    owner_id uuid NOT NULL,
    name text NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp 
);