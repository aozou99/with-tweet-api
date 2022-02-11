package deepl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type TransLatingTextResponse struct {
	Translations []*Translasion `json:"translations"`
}

type Translasion struct {
	DetectedSourceLanguage string `json:"detected_source_language"`
	Text                   string `json:"text"`
}

type DeepLClient struct {
	config *deepLConfig
	client *http.Client
}

func NewDeepLClient(c deepLConfig) *DeepLClient {
	return &DeepLClient{
		config: &c,
		client: &http.Client{},
	}
}

func (dc DeepLClient) TranslateText(tl []string, tlang string) (*TransLatingTextResponse, error) {
	ps := url.Values{}
	ps.Add("auth_key", dc.config.apiAuthKey)
	for _, v := range tl {
		ps.Add("text", v)
	}
	ps.Add("target_lang", tlang)

	res, err := http.PostForm(dc.config.endpoint+"translate", ps)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		b, err := httputil.DumpResponse(res, true)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf(string(b))
	}

	var result TransLatingTextResponse
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
