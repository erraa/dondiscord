package scraper

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type RedditStruct struct {
	Token     string
	Url       string
	AccessUrl string
}

func (reddit RedditStruct) Authenticate() string {
	// This needs to be urlencoded probably why it doesn't work
	b := []byte(`{"grant_type": "password", "username": "", "password": ""}`)
	buf := bytes.NewBuffer(b)
	client := &http.Client{}
	fmt.Println(buf)
	req, err := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token", buf)

	// Get this from config file
	req.SetBasicAuth("", "")
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
