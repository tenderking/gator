-- name: CreateUserFeed :one
INSERT INTO feeds (id, user_id, created_at, updated_at, url, name)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
)
RETURNING *;

-- name: GetUserFeeds :many
SELECT 
    f.user_id,
    f.url,
    f.name,
    f.id
FROM
    feeds f;

-- name: GetFeedByURL :one
SELECT * FROM feeds
WHERE url = $1;

-- name: MarkFeedAsFetched :exec
UPDATE feeds
SET 
    last_fetched_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;