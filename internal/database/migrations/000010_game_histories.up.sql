CREATE TABLE IF NOT EXISTS game_histories(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL,
    game_id uuid NOT NULL,
    player_position text NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp 
);

ALTER TABLE game_histories
ADD CONSTRAINT game_histories_user_id_game_id UNIQUE (user_id,game_id);
