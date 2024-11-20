CREATE TABLE IF NOT EXISTS game_statistics(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    game_id uuid NOT NULL,
    total_game_time_seconds int, 
    right_consecutive_points int,
    left_consecutive_points int,
    left_longest_point_seconds int,
    left_shortest_point_seconds int,
    right_longest_point_seconds int,
    right_shortest_point_seconds int,
    average_time_per_point_seconds int,
    left_average_time_per_point_seconds int,
    right_average_time_per_point_seconds int,
    created_at timestamp DEFAULT now(),
    updated_at timestamp 
)