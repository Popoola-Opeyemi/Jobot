package handlers

import twitterscraper "github.com/n0madic/twitter-scraper"

func InitScraper() {
	scraper := twitterscraper.New()
	Application.Scraper = scraper
}
