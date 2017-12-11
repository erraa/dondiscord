package scraper

import (
	"testing"

	"github.com/erraa/dondiscord/config"
)

func TestScraper(t *testing.T) {
	config.ReadConfig()
	reddit := InitReddit(config.RedditUrl, false)
	reddit.GetPicture()
}
