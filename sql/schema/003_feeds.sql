-- +goose Up
CREATE TABLE feeds(
	id UUID PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	name VARCHAR(124) NOT NULL,
	url VARCHAR(500) UNIQUE NOT NULL,
	user_id UUID NOT NULL,
	FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);


-- +goose Down
DROP TABLE feeds;