-- name: CreateFeed :one
INSERT INTO feeds (id, url, name, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetAllFeeds :many
SELECT * FROM feeds;

-- name: MarkFeedFetched :one
UPDATE feeds
SET last_fetched_at=$1, updated_at=$2
WHERE id=$3
RETURNING *;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds
ORDER BY last_fetched_at NULLS FIRST
LIMIT $1;