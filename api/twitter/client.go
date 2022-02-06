package twitter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strings"
)

type TweetsLookupResponse struct {
	Tweets []*Tweet `json:"data"`
}

type Tweet struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type TwitterClient struct {
	config *TwitterConfig
	client *http.Client
}

func buildUrl(base, p string, param map[string]string) (*url.URL, error) {
	u, err := url.Parse(base)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "tweets")

	q := url.Values{}
	for k, v := range param {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	return u, nil
}

func (tc *TwitterClient) genGetRequest(target string, param map[string]string) (*http.Request, error) {
	url, err := buildUrl(tc.config.Endpoint, target, param)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tc.config.BearerToken))
	return req, nil
}

func NewTwitterClient(config *TwitterConfig) (*TwitterClient, error) {
	return &TwitterClient{config: config, client: &http.Client{}}, nil
}

func (tc *TwitterClient) TweetsLookup(ids []string) (*TweetsLookupResponse, error) {
	req, err := tc.genGetRequest("tweets", map[string]string{
		"ids": strings.Join(ids, ","),
	})
	if err != nil {
		return nil, err
	}
	res, err := tc.client.Do(req)
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

	var result TweetsLookupResponse
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
