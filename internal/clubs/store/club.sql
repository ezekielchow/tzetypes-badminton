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
SELECT * FROM clubs WHERE owner_id = @owner_id::uuid limit 1;

-- name: FindPlayerInClub :one
SELECT * FROM player_clubs WHERE club_id = @club_id::uuid
AND player_id = @player_id::uuid limit 1;