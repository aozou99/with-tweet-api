package twitter

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

type args struct {
	ids []string
}

type mockServer struct {
	ReplyStatus  int
	MatchHeaders map[string]string
	MatchParams  map[string]string
	ReplyJSON    map[string]interface{}
}

type test_case struct {
	title string
	args  args
	mockServer
	want    *TweetsLookupResponse
	wantErr bool
}

func Test_TweetsLookup(t *testing.T) {
	os.Setenv("TWITTER_BEARER_TOKEN", "bearer_token")
	defer os.Clearenv()

	twitterConfig, _ := NewTwitterConfig()
	twitterClient, _ := NewTwitterClient(twitterConfig)

	tests := []test_case{
		{
			title: "200 success with multiple ids",
			args:  args{ids: []string{"1", "2"}},
			mockServer: mockServer{
				MatchHeaders: map[string]string{
					"Authorization": "^Bearer bearer_token$",
				},
				MatchParams: map[string]string{
					"ids": "1,2",
				},
				ReplyStatus: 200,
				ReplyJSON: map[string]interface{}{
					"data": []map[string]string{
						{
							"id":   "1",
							"text": "a",
						},
						{
							"id":   "2",
							"text": "b",
						},
					},
				},
			},
			want: &TweetsLookupResponse{
				Tweets: []*Tweet{
					{ID: "1", Text: "a"},
					{ID: "2", Text: "b"},
				},
			},
			wantErr: false,
		},
		{
			title: "200 success with single ids",
			args:  args{ids: []string{"1"}},
			mockServer: mockServer{
				MatchHeaders: map[string]string{
					"Authorization": "^Bearer bearer_token$",
				},
				MatchParams: map[string]string{
					"ids": "1",
				},
				ReplyStatus: 200,
				ReplyJSON: map[string]interface{}{
					"data": []map[string]string{
						{
							"id":   "1",
							"text": "a",
						},
					},
				},
			},
			want: &TweetsLookupResponse{
				Tweets: []*Tweet{
					{ID: "1", Text: "a"},
				},
			},
			wantErr: false,
		},
		{
			title: "200 success with no hit ids",
			args:  args{ids: []string{"1"}},
			mockServer: mockServer{
				MatchHeaders: map[string]string{
					"Authorization": "^Bearer bearer_token$",
				},
				MatchParams: map[string]string{
					"ids": "1",
				},
				ReplyStatus: 200,
				ReplyJSON: map[string]interface{}{
					"errors": []map[string]string{
						{
							"value":         "1",
							"detail":        "Could not find tweet with ids: [1].",
							"title":         "Not Found Error",
							"resource_type": "tweet",
							"parameter":     "ids",
							"resource_id":   "1",
							"type":          "https://api.twitter.com/2/problems/resource-not-found",
						},
					},
				},
			},
			want: &TweetsLookupResponse{
				Tweets: nil,
			},
			wantErr: false,
		},
		{
			title: "401 unauthorized",
			args:  args{ids: []string{"1"}},
			mockServer: mockServer{
				MatchHeaders: map[string]string{
					"Authorization": "^Bearer bearer_token$",
				},
				MatchParams: map[string]string{
					"ids": "1",
				},
				ReplyStatus: 401,
				ReplyJSON: map[string]interface{}{
					"title":  "Unauthorized",
					"type":   "about:blank",
					"status": 401,
					"detail": "Unauthorized",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.title, func(t *testing.T) {
			assert := assert.New(t)
			defer gock.Off()
			gock.New(twitterConfig.Endpoint).
				MatchHeaders(tt.MatchHeaders).
				MatchParams(tt.MatchParams).
				Get("/tweets").
				Reply(tt.ReplyStatus).
				JSON(tt.ReplyJSON)
			res, err := twitterClient.TweetsLookup(tt.args.ids)
			assert.Equal(tt.wantErr, err != nil)
			assert.Equal(tt.want, res)
		})
	}
}
