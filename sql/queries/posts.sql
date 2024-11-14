-- name: CreatePost :exec
INSERT INTO posts (id,created_at, updated_at, title, url, description, published_at, feed_id)
VALUES(
  $1, $2, $3, $4, $5, $6, $7, $8
);

-- name: GetUserPosts :many
WITH feeds_by_user AS (
  SELECT id FROM feeds
  WHERE user_id = $1 
)
SELECT p.* 
FROM posts p
JOIN feeds_by_user fbu ON fbu.id = p.feed_id
ORDER BY p.published_at
LIMIT $2;