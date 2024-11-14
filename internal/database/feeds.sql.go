// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: feeds.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUserFeed = `-- name: CreateUserFeed :one
INSERT INTO feeds (id, user_id, created_at, updated_at, url, name)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
)
RETURNING id, user_id, created_at, updated_at, url, name
`

type CreateUserFeedParams struct {
	ID        uuid.UUID
	UserID    uuid.NullUUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Url       string
	Name      string
}

func (q *Queries) CreateUserFeed(ctx context.Context, arg CreateUserFeedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, createUserFeed,
		arg.ID,
		arg.UserID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Url,
		arg.Name,
	)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Url,
		&i.Name,
	)
	return i, err
}

const getFeedByURL = `-- name: GetFeedByURL :one
SELECT id, user_id, created_at, updated_at, url, name FROM feeds
WHERE url = $1
`

func (q *Queries) GetFeedByURL(ctx context.Context, url string) (Feed, error) {
	row := q.db.QueryRowContext(ctx, getFeedByURL, url)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Url,
		&i.Name,
	)
	return i, err
}

const getUserFeeds = `-- name: GetUserFeeds :many
SELECT 
    f.user_id,
    f.url,
    f.name,
    f.id
FROM
    feeds f
`

type GetUserFeedsRow struct {
	UserID uuid.NullUUID
	Url    string
	Name   string
	ID     uuid.UUID
}

func (q *Queries) GetUserFeeds(ctx context.Context) ([]GetUserFeedsRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserFeeds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserFeedsRow
	for rows.Next() {
		var i GetUserFeedsRow
		if err := rows.Scan(
			&i.UserID,
			&i.Url,
			&i.Name,
			&i.ID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}