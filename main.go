package main

import (
	"aozou99/with-tweet-api/api/deepl"
	"aozou99/with-tweet-api/api/twitter"
	"aozou99/with-tweet-api/model"
	"aozou99/with-tweet-api/repository"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	twitterClient, _ := twitter.NewTwitterClient(twitter.NewTwitterConfig())
	tweets, err := twitterClient.TweetsLookup([]string{
		"1489729243810787328",
		"1491005135471841286",
		"1491019310415822849",
	})
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	txts := make([]string, 0, len(tweets.Tweets))

	for _, t := range tweets.Tweets {
		txts = append(txts, t.Text)
	}
	fmt.Println(len(tweets.Tweets), txts)

	deepLClient := deepl.NewDeepLClient(*deepl.NewDeepLConfig())
	res, err := deepLClient.TranslateText(txts, "EN")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	ttr := repository.NewTranslatedTweetRepository()
	if ttr == nil {
		return
	}
	for i, v := range res.Translations {
		ttr.Create(&model.TranslatedTweet{
			ID:             tweets.Tweets[i].ID,
			OriginText:     tweets.Tweets[i].Text,
			TranslatedText: v.Text,
		})
	}

}
