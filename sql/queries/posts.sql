-- name: CreatePost :one
INSERT INTO posts(id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING *;

-- name: GetUserPosts :many
SELECT * FROM posts
WHERE feed_id in (
	SELECT f.id FROM feeds as f
	INNER JOIN users as u ON f.user_id = $1
)
ORDER BY published_at
LIMIT $2;