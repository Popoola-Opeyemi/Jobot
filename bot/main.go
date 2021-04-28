package main

import "jbot/handlers"

func main() {

	handlers.InitScraper()

	handlers.SearchText("fashanu olakunle", "", 10)
	// handlers.GetTrends()
	// scraper := twitterscraper.New()

	// scraper.SetSearchMode(twitterscraper.SearchLatest)
	// var tweetArr []*twitterscraper.TweetResult
	// for tweet := range scraper.SearchTweets(context.Background(),
	// 	"vue -filter:retweets", 10) {
	// 	if tweet.Error != nil {
	// 		panic(tweet.Error)
	// 	}
	// 	fmt.Println(tweet.Text, time.Unix(tweet.Timestamp, 0))
	// 	tweetArr = append(tweetArr, tweet)
	// }

	// fmt.Println(len(tweetArr))
}
