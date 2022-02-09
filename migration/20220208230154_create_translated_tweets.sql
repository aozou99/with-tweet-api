-- +goose Up
-- +goose StatementBegin
CREATE TABLE translated_tweets(
	tweet_id INT,
	origin_text VARCHAR(200) not null,
	translated_text VARCHAR(500) not null,
	created_at TIMESTAMP,
	primary key(tweet_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE translated_tweets;
-- +goose StatementEnd
