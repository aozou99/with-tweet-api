package twitter

import (
	"os"
	"testing"
)

func TestNewTwitterConfig(t *testing.T) {
	os.Setenv("TWITTER_API_KEY", "api_key")
	os.Setenv("TWITTER_API_SECRET", "api_secret")
	os.Setenv("TWITTER_BEARER_TOKEN", "bearer_token")
	defer os.Clearenv()
	got := NewTwitterConfig()
	if got.apiKey != os.Getenv("TWITTER_API_KEY") || got.apiSecret != os.Getenv("TWITTER_API_SECRET") || got.bearerToken != os.Getenv("TWITTER_BEARER_TOKEN") {
		t.Errorf("ApiKey %v, ApiSecret %v, BearerToken %v", got.apiKey, got.apiSecret, got.bearerToken)
	}
}
