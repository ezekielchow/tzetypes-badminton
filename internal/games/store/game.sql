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
    step_num
) VALUES (
    @game_id::uuid,
    @team_left_score,
    @team_right_score,
    @score_at,
    @step_num
) RETURNING *;