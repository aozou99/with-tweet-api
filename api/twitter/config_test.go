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
	got, err := NewTwitterConfig()
	if err != nil {
		t.Errorf("fail to NewTwitterConfig")
	}
	if got.ApiKey != os.Getenv("TWITTER_API_KEY") || got.ApiSecret != os.Getenv("TWITTER_API_SECRET") || got.BearerToken != os.Getenv("TWITTER_BEARER_TOKEN") {
		t.Errorf("ApiKey %v, ApiSecret %v, BearerToken %v", got.ApiKey, got.ApiSecret, got.BearerToken)
	}
}
