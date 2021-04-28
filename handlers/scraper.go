package handlers

import (
	"context"
	"fmt"
	"time"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

// var tweetArr []*twitterscraper.TweetResult

func SearchText(searchTerm, Options string, limit int) (result []*twitterscraper.TweetResult) {
	scraper := Application.Scraper

	if limit == 0 {
		limit = 10
	}

	// scraper.SetSearchMode(twitterscraper.SearchLatest)
	query := fmt.Sprintf("%s", searchTerm)

	for tweet := range scraper.SearchTweets(context.Background(), query, 10) {
		if tweet.Error != nil {
			panic(tweet.Error)
		}
		fmt.Println(tweet.Text, time.Unix(tweet.Timestamp, 0))
		result = append(result, tweet)
	}
	return
}

func SearchLattest(searchTerm, Options string, limit int) {
	scraper := Application.Scraper

	scraper.SetSearchMode(twitterscraper.SearchLatest)

	SearchText(searchTerm, Options, limit)

}

func SearchPhotos(searchTerm, Options string, limit int) {
	scraper := Application.Scraper

	scraper.SetSearchMode(twitterscraper.SearchPhotos)

	SearchText(searchTerm, Options, limit)

}
func SearchVideos(searchTerm, Options string, limit int) {
	scraper := Application.Scraper

	scraper.SetSearchMode(twitterscraper.SearchVideos)

	SearchText(searchTerm, Options, limit)

}
func SearchUsers(searchTerm, Options string, limit int) {
	scraper := Application.Scraper

	scraper.SetSearchMode(twitterscraper.SearchUsers)

	SearchText(searchTerm, Options, limit)

}

func GetTrends() {
	scraper := Application.Scraper

	trends, err := scraper.GetTrends()
	if err != nil {
		panic(err)
	}
	fmt.Println(trends)

}
