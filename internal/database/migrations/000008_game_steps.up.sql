CREATE TABLE IF NOT EXISTS game_steps(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    game_id uuid NOT NULL,
    team_left_score integer NOT NULL,
    team_right_score integer NOT NULL,
    score_at timestamp NOT NULL,
    step_num integer NOT NULL,
    current_server text NOT NULL,
    left_odd_player_name text,
    left_even_player_name text NOT NULL,
    right_odd_player_name text,
    right_even_player_name text NOT NULL,
    is_paused integer NOT NULL DEFAULT 0,
    sync_id text NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp 
);

ALTER TABLE game_steps
ADD CONSTRAINT game_steps_game_id_step_num_key UNIQUE (game_id,step_num);
