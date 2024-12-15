CREATE TABLE IF NOT EXISTS instagram_feeds(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    media_id text,
    media_type text,
    media_url text,
    permalink text,
    posted_at timestamp,
    created_at timestamp DEFAULT now(),
    updated_at timestamp,
    UNIQUE (media_id)
)