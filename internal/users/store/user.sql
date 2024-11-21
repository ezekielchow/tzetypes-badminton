-- name: CreateUser :one
INSERT INTO users (
  email,password_hash,user_type
) VALUES (
  $1,$2,$3
) RETURNING *;

-- name: FindUserWithEmail :one
SELECT * FROM users WHERE email=$1;

-- name: FindUserWithID :one
SELECT * FROM users WHERE id=$1;
