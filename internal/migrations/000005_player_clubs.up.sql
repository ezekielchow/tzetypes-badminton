CREATE TABLE IF NOT EXISTS player_clubs(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    player_id uuid NOT NULL,
    club_id uuid NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp 
);