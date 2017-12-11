package scraper

import (
	"fmt"
	"testing"

	"github.com/erraa/dondiscord/config"
)

func TestScraper(t *testing.T) {
	config.ReadConfig()
	reddit := InitReddit(config.RedditUrl, false)
	fmt.Println(reddit.GetPicture())
}
