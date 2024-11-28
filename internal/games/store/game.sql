-- name: CreateGame :one
INSERT INTO games (
    club_id,
    left_odd_player_name,
    left_even_player_name,
    right_odd_player_name,
    right_even_player_name,
    game_type,
    serving_side,
    created_at
) VALUES (
    @club_id::uuid,
    @left_odd_player_name::text,
    @left_even_player_name::text,
    @right_odd_player_name::text,
    @right_even_player_name::text,
    @game_type::text,
    @serving_side::text,
    @created_at
) RETURNING *;

-- name: CreateGameStep :one
INSERT INTO game_steps (
    game_id,
    team_left_score,
    team_right_score,
    score_at,
    step_num,
    current_server,
    left_odd_player_name,
    left_even_player_name,
    right_odd_player_name,
    right_even_player_name,
    sync_id
) VALUES (
    @game_id::uuid,
    @team_left_score,
    @team_right_score,
    @score_at,
    @step_num,
    @current_server,
    @left_odd_player_name::text,                         
    @left_even_player_name::text,
    @right_odd_player_name::text,
    @right_even_player_name::text,
    @sync_id
) ON CONFLICT (game_id, step_num) DO NOTHING 
RETURNING *;

-- name: DeleteGameStep :exec
DELETE FROM game_steps where id = @id::uuid;

-- name: EndGame :exec
UPDATE games SET is_ended = @is_ended 
WHERE id = @id;

-- name: GetGameWithID :one
SELECT * FROM games WHERE id = @id::uuid limit 1;

-- name: GetGameStepsWithGameID :many
SELECT * FROM game_steps WHERE game_id = @game_id::uuid
ORDER BY step_num ASC;

-- name: GetGameStatisticsWithGameID :one
SELECT * FROM game_statistics WHERE game_id = @game_id::uuid LIMIT 1;

-- name: CreateGameStatistic :one
INSERT INTO game_statistics(
    game_id,
    total_game_time_seconds, 
    right_consecutive_points,
    left_consecutive_points,
    left_longest_point_seconds,
    left_shortest_point_seconds,
    right_longest_point_seconds,
    right_shortest_point_seconds,
    average_time_per_point_seconds,
    right_average_time_per_point_seconds,
    left_average_time_per_point_seconds
) VALUES (
    @game_id::uuid,
    @total_game_time_seconds::int,
    @right_consecutive_points::int,
    @left_consecutive_points::int,
    @left_longest_point_seconds::int,
    @left_shortest_point_seconds::int,
    @right_longest_point_seconds::int,
    @right_shortest_point_seconds::int,
    @average_time_per_point_seconds::int,
    @right_average_time_per_point_seconds::int,
    @left_average_time_per_point_seconds::int
) RETURNING *;

-- name: CreateOrUpdateGameHistory :one
INSERT INTO game_histories(
    user_id,
    game_id,
    player_position,
    game_started_at,
    game_won_by,
    total_points,
    points_won,
    points_lost,
    average_time_per_point_seconds,
    average_time_per_point_won_seconds,
    average_time_per_point_lost_seconds,
    longest_rally_seconds,
    longest_rally_is_won,
    shortest_rally_seconds,
    shortest_rally_is_won,
    is_game_won
) VALUES (
    @user_id::uuid,
    @game_id::uuid,
    @player_position::text,
    @game_started_at,
    @game_won_by::text,
    @total_points::int,
    @points_won::int,
    @points_lost::int,
    @average_time_per_point_seconds::int,
    @average_time_per_point_won_seconds::int,
    @average_time_per_point_lost_seconds::int,
    @longest_rally_seconds::int,
    @longest_rally_is_won::int,
    @shortest_rally_seconds::int,
    @shortest_rally_is_won::int,
    @is_game_won::int
) 
ON CONFLICT (user_id,game_id) DO UPDATE
    SET 
    player_position = EXCLUDED.player_position,
    game_started_at = EXCLUDED.game_started_at,
    game_won_by = EXCLUDED.game_won_by,
    total_points = EXCLUDED.total_points,
    points_won = EXCLUDED.points_won,
    points_lost = EXCLUDED.points_lost,
    average_time_per_point_seconds = EXCLUDED.average_time_per_point_seconds,
    average_time_per_point_won_seconds = EXCLUDED.average_time_per_point_won_seconds,
    average_time_per_point_lost_seconds = EXCLUDED.average_time_per_point_lost_seconds,
    longest_rally_seconds = EXCLUDED.longest_rally_seconds,
    longest_rally_is_won = EXCLUDED.longest_rally_is_won,
    shortest_rally_seconds = EXCLUDED.shortest_rally_seconds,
    shortest_rally_is_won = EXCLUDED.shortest_rally_is_won,
    is_game_won = EXCLUDED.is_game_won,
    updated_at = now()
RETURNING *;

-- name: GetGameHistoryGivenUserIdAndGameId :one
SELECT * FROM game_histories WHERE game_id = @game_id::uuid AND user_id = @user_id::uuid limit 1;

-- name: CreateOrUpdateGameRecentStatistic :one
INSERT INTO game_recent_statistics(
    user_id,
    game_count,
    wins,
    losses,
    total_points,
    points_won,
    average_time_per_point_seconds,
    average_time_per_point_won_seconds,
    average_time_per_point_lost_seconds,
    longest_rally_seconds,
    longest_rally_is_won,
    shortest_rally_seconds,
    shortest_rally_is_won,
    needs_regenerating
) VALUES (
    @user_id::uuid,
    @game_count::int,
    @wins::int,
    @losses::int,
    @total_points::int,
    @points_won::int,
    @average_time_per_point_seconds::int,
    @average_time_per_point_won_seconds::int,
    @average_time_per_point_lost_seconds::int,
    @longest_rally_seconds::int,
    @longest_rally_is_won::int,
    @shortest_rally_seconds::int,
    @shortest_rally_is_won::int,
    @needs_regenerating::int
) 
ON CONFLICT (user_id) DO UPDATE
    SET 
    game_count = EXCLUDED.game_count,
    wins = EXCLUDED.wins,
    losses = EXCLUDED.losses,
    total_points = EXCLUDED.total_points,
    points_won = EXCLUDED.points_won,
    average_time_per_point_seconds = EXCLUDED.average_time_per_point_seconds, 
    average_time_per_point_won_seconds = EXCLUDED.average_time_per_point_won_seconds,
    average_time_per_point_lost_seconds = EXCLUDED.average_time_per_point_lost_seconds,
    longest_rally_seconds = EXCLUDED.longest_rally_seconds,
    longest_rally_is_won = EXCLUDED.longest_rally_is_won,
    shortest_rally_seconds = EXCLUDED.shortest_rally_seconds,
    shortest_rally_is_won = EXCLUDED.shortest_rally_is_won,
    needs_regenerating = EXCLUDED.needs_regenerating,
    updated_at = now()
RETURNING *;

-- name: GetGameRecentStatisticWithUserId :one
SELECT * FROM game_recent_statistics WHERE user_id = @user_id::uuid limit 1;

-- name: GetGameRecentStatisticThatNeedsRegeneration :many
SELECT * FROM game_recent_statistics WHERE needs_regenerating = 1 limit 10;

-- name: GetMostRecentGameHistories :many
SELECT * from game_histories WHERE user_id = @user_id::uuid ORDER BY game_started_at DESC limit 12;

-- name: GetGameStepsGivenGameIds :many
SELECT * from game_steps WHERE game_id = ANY(@game_ids::uuid[]);