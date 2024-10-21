// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: player.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const allPlayers = `-- name: AllPlayers :many
SELECT id, user_id, name, created_at, updated_at FROM players
`

func (q *Queries) AllPlayers(ctx context.Context) ([]Player, error) {
	rows, err := q.db.Query(ctx, allPlayers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Player
	for rows.Next() {
		var i Player
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const createPlayer = `-- name: CreatePlayer :one
INSERT INTO players (
  user_id, name
) VALUES (
  $1,$2
) RETURNING id, user_id, name, created_at, updated_at
`

type CreatePlayerParams struct {
	UserID pgtype.UUID
	Name   string
}

func (q *Queries) CreatePlayer(ctx context.Context, arg CreatePlayerParams) (Player, error) {
	row := q.db.QueryRow(ctx, createPlayer, arg.UserID, arg.Name)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findPlayerWithName = `-- name: FindPlayerWithName :one
SELECT id, user_id, name, created_at, updated_at FROM players WHERE name=($1::text)
`

func (q *Queries) FindPlayerWithName(ctx context.Context, name string) (Player, error) {
	row := q.db.QueryRow(ctx, findPlayerWithName, name)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPlayerWithId = `-- name: GetPlayerWithId :one
SELECT id, user_id, name, created_at, updated_at FROM players
WHERE id = $1::uuid limit 1
`

func (q *Queries) GetPlayerWithId(ctx context.Context, id pgtype.UUID) (Player, error) {
	row := q.db.QueryRow(ctx, getPlayerWithId, id)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listPlayers = `-- name: ListPlayers :many
SELECT
  p.id, p.user_id, p.name, p.created_at, p.updated_at,
  COUNT(*) OVER() AS total_count
FROM
  players AS p
JOIN
  player_clubs AS pc ON p.id = pc.player_id
JOIN 
  clubs AS c ON pc.club_id = c.id 
WHERE
  ($1::uuid IS NULL OR c.owner_id = $1::uuid) -- Optional filtering by owner_id
ORDER BY
  CASE WHEN $2::text = 'name_asc' THEN p.name END ASC,
  CASE WHEN $2::text = 'name_desc' THEN p.name END DESC
LIMIT $4
OFFSET $3
`

type ListPlayersParams struct {
	OwnerID         pgtype.UUID
	SortArrangement string
	OffsetCount     int32
	LimitCount      int32
}

type ListPlayersRow struct {
	ID         pgtype.UUID
	UserID     pgtype.UUID
	Name       string
	CreatedAt  pgtype.Timestamp
	UpdatedAt  pgtype.Timestamp
	TotalCount int64
}

func (q *Queries) ListPlayers(ctx context.Context, arg ListPlayersParams) ([]ListPlayersRow, error) {
	rows, err := q.db.Query(ctx, listPlayers,
		arg.OwnerID,
		arg.SortArrangement,
		arg.OffsetCount,
		arg.LimitCount,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListPlayersRow
	for rows.Next() {
		var i ListPlayersRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.TotalCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePlayer = `-- name: UpdatePlayer :one
UPDATE players SET 
name = $1::text,
updated_at = $2
WHERE id = $3::uuid
RETURNING id, user_id, name, created_at, updated_at
`

type UpdatePlayerParams struct {
	Name      string
	UpdatedAt pgtype.Timestamp
	ID        pgtype.UUID
}

func (q *Queries) UpdatePlayer(ctx context.Context, arg UpdatePlayerParams) (Player, error) {
	row := q.db.QueryRow(ctx, updatePlayer, arg.Name, arg.UpdatedAt, arg.ID)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
