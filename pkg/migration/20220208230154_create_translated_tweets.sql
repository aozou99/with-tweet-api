-- +goose Up
-- +goose StatementBegin
CREATE TABLE translated_tweets(
	id VARCHAR(32), -- tweetID
	origin_text VARCHAR(200) not null,
	translated_text VARCHAR(500) not null,
	created_at TIMESTAMP,
  updated_at TIMESTAMP,
	deleted_at TIMESTAMP,
	primary key(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE translated_tweets;
-- +goose StatementEnd
