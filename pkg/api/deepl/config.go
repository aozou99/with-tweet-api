package deepl

import "os"

type deepLConfig struct {
	apiAuthKey string
	endpoint   string
}

func NewDeepLConfig() *deepLConfig {
	return &deepLConfig{
		apiAuthKey: os.Getenv("DEEPL_AUTH_KEY"),
		endpoint:   "https://api-free.deepl.com/v2/",
	}
}
