package main

import (
	"aozou99/with-tweet-api/api/deepl"
	"aozou99/with-tweet-api/api/twitter"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	t_config, _ := twitter.NewTwitterConfig()
	twitterClient, _ := twitter.NewTwitterClient(t_config)
	tweets, err := twitterClient.TweetsLookup([]string{
		"1489729243810787328",
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

	for _, v := range res.Translations {
		fmt.Println(v)
	}

}
