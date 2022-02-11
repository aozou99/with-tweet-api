package model

import (
	"aozou99/with-tweet-api/pkg/entity"
)

type TranslatedTweet struct {
	TweetID        string `json:"tweet_id"`
	OriginText     string `json:"origin_text"`
	TranslatedText string `json:"translated_text"`
}

func NewTranslatedTweetFromEntity(e *entity.TranslatedTweet) *TranslatedTweet {
	return &TranslatedTweet{
		TweetID:        e.ID,
		OriginText:     e.OriginText,
		TranslatedText: e.TranslatedText,
	}
}
