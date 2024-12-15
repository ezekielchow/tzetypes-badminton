CREATE TABLE IF NOT EXISTS instagram_feeds(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    media_id text NOT NULL,
    media_type text NOT NULL,
    media_url text NOT NULL,
    permalink text NOT NULL,
    posted_at timestamp NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp,
    UNIQUE (media_id)
)