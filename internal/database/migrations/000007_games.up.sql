CREATE TABLE IF NOT EXISTS games(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    club_id uuid NOT NULL,
    left_odd_player_name text,
    left_even_player_name text NOT NULL,
    right_odd_player_name text,
    right_even_player_name text NOT NULL,
    game_type text NOT NULL,
    serving_side text NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp 
);

