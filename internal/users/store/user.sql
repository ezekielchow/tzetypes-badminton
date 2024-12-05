-- name: CreateUser :one
INSERT INTO users (
  firebase_uid,email,account_tier
) VALUES (
  @firebase_uid::text,
  @email::text,
  @account_tier::text
) RETURNING *;

-- name: FindUserWithEmail :one
SELECT * FROM users WHERE email=$1;

-- name: FindUserWithID :one
SELECT * FROM users WHERE id=$1;
