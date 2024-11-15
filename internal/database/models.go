// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Feed struct {
	ID            uuid.UUID
	UserID        uuid.NullUUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Url           string
	Name          string
	LastFetchedAt sql.NullTime
}

type FeedFollow struct {
	ID        uuid.UUID
	UserID    uuid.NullUUID
	FeedID    uuid.NullUUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Post struct {
	ID          uuid.UUID
	FeedID      uuid.NullUUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Url         string
	Description sql.NullString
	PublishedAt time.Time
}

type User struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
