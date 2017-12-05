package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token          string
	BotPrefix      string
	config         *configStruct
	RedditUsername string
	RedditPassword string
	RedditUrl      string
)

type configStruct struct {
	Token          string `json:"Token"`
	BotPrefix      string `json:"BotPrefix"`
	RedditUsername string `json:"RedditUsername"`
	RedditPassword string `json:"RedditPassword"`
	RedditUrl      string `json:"RedditUrl"`
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
	RedditUsername = config.RedditUsername
	RedditPassword = config.RedditPassword
	RedditUrl = config.RedditUrl
	return nil
}
