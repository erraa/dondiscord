package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token              string
	BotPrefix          string
	config             *configStruct
	RedditUrl          string
	RedditAuthUsername string
	RedditAuthPassword string
	MemeUrl            string
	UserAgent          string
)

type configStruct struct {
	Token              string `json:"Token"`
	BotPrefix          string `json:"BotPrefix"`
	RedditUrl          string `json:"RedditUrl"`
	RedditAuthUsername string `json:"RedditAuthUsername`
	RedditAuthPassword string `json:"RedditAuthPassword`
	UserAgent          string `json:"User-Agent"`
	MemeUrl            string `json:"MemeUrl"`
}

// ReadConfig read
func ReadConfig() error {
	fmt.Println("Reading new file...")
	file, err := ioutil.ReadFile("/home/erra/go/src/github.com/erraa/dondiscord/config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	Token = config.Token
	BotPrefix = config.BotPrefix
	RedditAuthUsername = config.RedditAuthUsername
	RedditAuthPassword = config.RedditAuthPassword
	RedditUrl = config.RedditUrl
	MemeUrl = config.MemeUrl
	UserAgent = config.UserAgent
	return nil
}
