package twitter

import "os"

type twitterConfig struct {
	apiKey      string
	apiSecret   string
	bearerToken string
	endpoint    string
}

func NewTwitterConfig() *twitterConfig {
	return &twitterConfig{
		apiKey:      os.Getenv("TWITTER_API_KEY"),
		apiSecret:   os.Getenv("TWITTER_API_SECRET"),
		bearerToken: os.Getenv("TWITTER_BEARER_TOKEN"),
		endpoint:    "https://api.twitter.com/2/",
	}
}
