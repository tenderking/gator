-- name: CreateFeedFollow :one
WITH new_feed_follow AS (
    INSERT INTO feed_follows (id, user_id, feed_id, created_at, updated_at)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT 
    new_feed_follow.*,
    feeds.name as feed_name,
    users.name as user_name
FROM
    new_feed_follow
INNER JOIN
    feeds ON new_feed_follow.feed_id = feeds.id
INNER JOIN
    users ON new_feed_follow.user_id = users.id;

-- name: GetFeedFollowsForUser :many
SELECT feed_follows.*, feeds.name AS feed_name, users.name AS user_name
FROM feed_follows
INNER JOIN feeds ON feed_follows.feed_id = feeds.id
INNER JOIN users ON feeds.user_id = users.id
WHERE feed_follows.user_id = $1;
    
-- name: UnfollowFeed :exec
DELETE FROM feed_follows as ff 
WHERE ff.feed_id = (SELECT feeds.id FROM feeds WHERE feeds.url = $1) 
  AND ff.user_id = (SELECT users.id FROM users WHERE users.id = $2);