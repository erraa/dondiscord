package scraper

import (
	"encoding/json"
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
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

	var respData map[string]interface{}
	err = json.Unmarshal([]byte(bodyText), &respData)
	reddit.Token = respData["access_token"].(string)

	// Returning is just made for testing stuff out
	s := string(bodyText)
	return s
}

func (reddit RedditStruct) GetPicture() string {
	fmt.Println(config.MemeUrl)
	resp, err := http.Get(config.MemeUrl)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return string(bodyText)
}

func InitReddit(url string, auth bool) RedditStruct {
	reddit := RedditStruct{}
	if auth {
		reddit.Url = "https://www.reddit.com/api/v1/"
		reddit.AccessUrl = reddit.Url + "access_token"
		reddit.Authenticate()
	}
	return reddit
}
