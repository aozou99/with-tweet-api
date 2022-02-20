package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"aozou99/with-tweet-api/graph/generated"
	"aozou99/with-tweet-api/graph/model"
	"context"
)

func (r *queryResolver) LatestTweets(ctx context.Context) ([]*model.TranslatedTweet, error) {
	entities := r.Repository.TranslatedTweet().Latest(5)
	var models []*model.TranslatedTweet
	for _, v := range entities {
		models = append(models, model.NewTranslatedTweetFromEntity(v))
	}
	return models, nil
}

func (r *queryResolver) Tweet(ctx context.Context, id string) (*model.TranslatedTweet, error) {
	entity := r.Repository.TranslatedTweet().Find(id)
	return model.NewTranslatedTweetFromEntity(entity), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
