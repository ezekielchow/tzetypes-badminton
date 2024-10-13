-- name: CreateClub :one
INSERT INTO clubs (
  owner_id, name
) VALUES (
  $1,$2
) RETURNING *;

-- name: AddPlayerToClub :exec
INSERT INTO player_clubs (
  player_id, club_id
) VALUES (
  @player_id::uuid, @club_id::uuid
);

-- name: GetClubGivenOwnerId :one
SELECT * FROM clubs where owner_id = @owner_id::uuid limit 1;