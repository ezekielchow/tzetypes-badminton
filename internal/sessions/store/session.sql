-- name: CreateSession :one
INSERT INTO sessions (
  user_id,session_token_expires_at,refresh_token_expires_at
) VALUES (
  $1,$2,$3
) RETURNING *;

-- name: FindSessionWithSessionID :one
SELECT * FROM sessions WHERE session_token=$1;

-- name: FindSessionWithRefreshToken :one
SELECT * FROM sessions WHERE refresh_token=$1;

-- name: UpdateSessionWithRefreshToken :one
UPDATE sessions
SET session_token = gen_random_uuid(), session_token_expires_at=$2, updated_at=now() 
WHERE refresh_token=$1
RETURNING *;

-- name: DeleteSession :exec
DELETE from sessions WHERE id=$1;
