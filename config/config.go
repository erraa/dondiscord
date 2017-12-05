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
)

type configStruct struct {
	Token              string `json:"Token"`
	BotPrefix          string `json:"BotPrefix"`
	RedditUrl          string `json:"RedditUrl"`
	RedditAuthUsername string
	RedditAuthPassword string
}

// ReadConfig read
func ReadConfig() error {
	fmt.Println("Reading new file...")
	file, err := ioutil.ReadFile("/home/erra/go/src/github.com/erraa/dondiscord/config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))
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
	return nil
}
