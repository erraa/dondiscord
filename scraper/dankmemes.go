package scraper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/erraa/dondiscord/config"
)

type RedditStruct struct {
	Token     string
	Url       string
	AccessUrl string
}

func (reddit RedditStruct) Authenticate() string {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	client := &http.Client{}
	req, err := http.NewRequest("POST",
		"https://www.reddit.com/api/v1/access_token",
		strings.NewReader(data.Encode()),
	)

	// Get this from config file
	req.SetBasicAuth(config.RedditAuthUsername, config.RedditAuthPassword)
	req.Header.Add("User-Agent", "Kungerra")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Request)
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	return s
}

func InitReddit(url string) RedditStruct {
	reddit := RedditStruct{}
	reddit.Url = "https://www.reddit.com/api/v1/"
	reddit.AccessUrl = reddit.Url + "access_token"
	return reddit
}
