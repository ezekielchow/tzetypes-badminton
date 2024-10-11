-- name: CreateClub :one
INSERT INTO clubs (
  owner_id, name
) VALUES (
  $1,$2
) RETURNING *;
