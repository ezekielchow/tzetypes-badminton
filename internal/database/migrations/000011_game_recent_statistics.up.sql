CREATE TABLE IF NOT EXISTS game_recent_statistics(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL,
    game_count int,
    wins int,
    losses int,
    total_points int,
    points_won int,
    average_time_per_point_seconds int,
    average_time_per_point_won_seconds int,
    average_time_per_point_lost_seconds int,
    longest_rally_seconds int,
    longest_rally_is_won int,
    shortest_rally_seconds int,
    shortest_rally_is_won int,
    needs_regenerating int,
    created_at timestamp DEFAULT now(),
    updated_at timestamp 
)