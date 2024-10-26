-- name: CreateGame :one
INSERT INTO games (
    club_id,
    left_odd_player_name,
    left_even_player_name,
    right_odd_player_name,
    right_even_player_name,
    game_type,
    serving_side
) VALUES (
    @club_id::uuid,
    @left_odd_player_name::text,
    @left_even_player_name::text,
    @right_odd_player_name::text,
    @right_even_player_name::text,
    @game_type::text,
    @serving_side::text
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