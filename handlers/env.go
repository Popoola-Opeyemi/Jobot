package handlers

import twitterscraper "github.com/n0madic/twitter-scraper"

type AppEnvironment struct {
	Scraper *twitterscraper.Scraper
}

var Application *AppEnvironment

func init() {
	if Application == nil {
		Application = new(AppEnvironment)
	}
}
