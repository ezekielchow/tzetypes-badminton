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
SET session_token = @session_token::uuid, session_token_expires_at = @session_token_expires_at, updated_at=now() 
WHERE refresh_token = @refresh_token::uuid
RETURNING *;

-- name: DeleteSession :exec
DELETE from sessions WHERE id=$1;
