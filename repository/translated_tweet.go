package repository

import (
	"aozou99/with-tweet-api/model"

	"github.com/rs/zerolog/log"
	"github.com/spiegel-im-spiegel/errs"
)

type TranslatedTweetRepository struct {
	Repository
}

var translated_tweet_repository = &TranslatedTweetRepository{}

func NewTranslatedTweetRepository() *TranslatedTweetRepository {
	translated_tweet_repository.init()
	return translated_tweet_repository
}

func (r *TranslatedTweetRepository) Create(m *model.TranslatedTweet) bool {
	result := r.db.Create(&m)
	if result.Error != nil {
		log.Error().Interface("error", errs.Wrap(result.Error)).Send()
	}
	return result.Error != nil
}
