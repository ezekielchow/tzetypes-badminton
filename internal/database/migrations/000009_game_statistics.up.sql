CREATE TABLE IF NOT EXISTS game_statistics(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    game_id uuid NOT NULL,
    total_game_time_seconds int, 
    right_consecutive_points int,
    left_consecutive_points int,
    longest_point_seconds int,
    longest_point_team text,
    shortest_point_seconds int,
    shortest_point_team text,
    average_time_per_point_seconds int,
    left_average_time_per_point_seconds int,
    right_average_time_per_point_seconds int,
    created_at timestamp DEFAULT now(),
    updated_at timestamp 
)