-- name: CreatePlayer :one
INSERT INTO players (
  user_id, name
) VALUES (
  $1,$2
) RETURNING *;
