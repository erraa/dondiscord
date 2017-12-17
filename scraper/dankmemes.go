package scraper

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/erraa/dondiscord/config"
)

type RedditStruct struct {
	Token     string
	Url       string
	AccessUrl string
}

func InitReddit(url string, auth bool) RedditStruct {
	reddit := RedditStruct{}
	if auth {
		reddit.Url = config.RedditUrl
		reddit.AccessUrl = reddit.Url + "access_token"
		reddit.Authenticate()
	}
	return reddit
}

func (reddit RedditStruct) Authenticate() string {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	client := &http.Client{}
	req, err := http.NewRequest("POST",
		reddit.Url,
		strings.NewReader(data.Encode()),
	)

	// Get this from config file
	req.SetBasicAuth(config.RedditAuthUsername, config.RedditAuthPassword)
	req.Header.Add("User-Agent", config.UserAgent)
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
	Preview imageData `json:"preview"`
}

type imageData struct {
	Images []imageStruct `json:"images"`
}

type imageStruct struct {
	Source sourceStruct `json:"source"`
}

type sourceStruct struct {
	Url string `json:"url"`
}

// GetPicture Returns a random URL to a picture from reddit
func (reddit RedditStruct) GetPicture() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	filename := dir + "/redditdata.txt"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		touchfile(filename)
	}

	resp, err := http.Get(config.MemeUrl)
	if err != nil {
		log.Fatal(err)
	}

	var bodyText []byte
	if resp.StatusCode == 429 {
		bodyText = readfile(filename)
	} else {
		bodyText, err = ioutil.ReadAll(resp.Body)
		touchfile(filename)
		tofile(filename, bodyText)
	}

	var nestedData redditData
	err = json.Unmarshal(bodyText, &nestedData)
	if err != nil {
		log.Fatal(err)
	}
	if len(nestedData.Data.Children) == 0 {
		return "Something went wrong"
	}
	resp.Body.Close()
	randomNumber := random(1, len(nestedData.Data.Children))
	randomchild := nestedData.Data.Children[randomNumber]
	randomUrl := randomchild.Data.Preview.Images[0].Source.Url

	return randomUrl
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
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

func touchfile(filename string) {
	_, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
}
