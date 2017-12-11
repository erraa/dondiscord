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

type redditData struct {
	Data subRedditData `json:"data"`
}

type subRedditData struct {
	Children []subRedditChildren `json:"children"`
}

type subRedditChildren struct {
	Data childData `json:"data"`
}

type childData struct {
	Domain string `json:"domain"`
}

func (reddit RedditStruct) GetPicture() string {
	filename := "./redditdata.txt"
	resp, err := http.Get(config.MemeUrl)
	if err != nil {
		log.Fatal(err)
	}

	var bodyText []byte
	if resp.StatusCode == 429 {
		bodyText = readfile(filename)
	} else {
		bodyText, err = ioutil.ReadAll(resp.Body)
		tofile(filename, bodyText)
	}

	var nestedData redditData
	err = json.Unmarshal(bodyText, &nestedData)
	resp.Body.Close()
	fmt.Printf("%+v", nestedData)
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

func readfile(filename string) []byte {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return b
}

func tofile(filename string, data []byte) {
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		panic(err)
	}
}
