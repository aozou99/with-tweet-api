package repository

import (
	"aozou99/with-tweet-api/pkg/entity"

	"github.com/rs/zerolog/log"
	"github.com/spiegel-im-spiegel/errs"
)

type TranslatedTweetRepository struct {
	Repository
}

func (r *TranslatedTweetRepository) Create(m *entity.TranslatedTweet) int64 {
	result := r.db.FirstOrCreate(&m)
	if result.Error != nil {
		log.Error().Interface("error", errs.Wrap(result.Error)).Send()
	}
	return result.RowsAffected
}

func (r *TranslatedTweetRepository) Latest(limit int) []*entity.TranslatedTweet {
	var results []*entity.TranslatedTweet
	r.db.Limit(limit).Order("created_at desc").Find(&results)
	return results
}

func (r *TranslatedTweetRepository) Find(id string) *entity.TranslatedTweet {
	var result entity.TranslatedTweet
	r.db.First(&result, id)
	return &result
}