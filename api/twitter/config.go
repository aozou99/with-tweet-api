package twitter

import "os"

type TwitterConfig struct {
	ApiKey      string
	ApiSecret   string
	BearerToken string
	Endpoint    string
}

func NewTwitterConfig() (*TwitterConfig, error) {
	return &TwitterConfig{
		ApiKey:      os.Getenv("TWITTER_API_KEY"),
		ApiSecret:   os.Getenv("TWITTER_API_SECRET"),
		BearerToken: os.Getenv("TWITTER_BEARER_TOKEN"),
		Endpoint:    "https://api.twitter.com/2/",
	}, nil
}
