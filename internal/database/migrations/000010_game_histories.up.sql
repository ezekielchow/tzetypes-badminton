CREATE TABLE IF NOT EXISTS game_histories(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL,
    game_id uuid NOT NULL,
    player_position text NOT NULL,
    is_game_won int NOT NULL,
    game_started_at timestamp NOT NULL,
    game_won_by text NOT NULL,
    total_points int NOT NULL,
    points_won int NOT NULL,
    points_lost int NOT NULL,
    average_time_per_point_seconds int NOT NULL,
    average_time_per_point_won_seconds int NOT NULL,
    average_time_per_point_lost_seconds int NOT NULL,
    longest_rally_seconds int NOT NULL,
    longest_rally_is_won int NOT NULL,
    shortest_rally_seconds int NOT NULL,
    shortest_rally_is_won int NOT NULL,
    total_game_time_seconds int NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp 
);

ALTER TABLE game_histories
ADD CONSTRAINT game_histories_user_id_game_id UNIQUE (user_id,game_id);
