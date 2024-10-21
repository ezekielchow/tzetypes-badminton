-- name: CreatePlayer :one
INSERT INTO players (
  user_id, name
) VALUES (
  $1,$2
) RETURNING *;

-- name: ListPlayers :many
SELECT
  p.*,
  COUNT(*) OVER() AS total_count
FROM
  players AS p
JOIN
  player_clubs AS pc ON p.id = pc.player_id
JOIN 
  clubs AS c ON pc.club_id = c.id 
WHERE
  (@owner_id::uuid IS NULL OR c.owner_id = @owner_id::uuid) -- Optional filtering by owner_id
ORDER BY
  CASE WHEN @sort_arrangement::text = 'name_asc' THEN p.name END ASC,
  CASE WHEN @sort_arrangement::text = 'name_desc' THEN p.name END DESC
LIMIT @limit_count
OFFSET @offset_count;

-- name: FindPlayerWithName :one
SELECT * FROM players WHERE name=(@name::text);

-- name: AllPlayers :many
SELECT * FROM players;

-- name: UpdatePlayer :one
UPDATE players SET 
name = @name::text,
updated_at = @updated_at
WHERE id = @id::uuid
RETURNING *;

-- name: GetPlayerWithId :one
SELECT * FROM players
WHERE id = @id::uuid limit 1;