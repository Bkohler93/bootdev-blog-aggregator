-- +goose Up
CREATE TABLE posts(
	id UUID PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	title VARCHAR(64) NOT NULL,
	url VARCHAR(500) NOT NULL UNIQUE,
	description VARCHAR(2000),
	published_at TIMESTAMP NOT NULL,
	feed_id UUID NOT NULL,
	FOREIGN KEY(feed_id) REFERENCES feeds(id)
);


-- +goose Down
DROP TABLE posts;