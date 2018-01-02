package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	Config             string
	Token              string
	BotPrefix          string
	config             *configStruct
	RedditUrl          string
	RedditAuthUsername string
	RedditAuthPassword string
	MemeUrl            string
	UserAgent          string
	LogFile            string
)

type configStruct struct {
	Token              string `json:"Token"`
	BotPrefix          string `json:"BotPrefix"`
	RedditUrl          string `json:"RedditUrl"`
	RedditAuthUsername string `json:"RedditAuthUsername`
	RedditAuthPassword string `json:"RedditAuthPassword`
	UserAgent          string `json:"User-Agent"`
	MemeUrl            string `json:"MemeUrl"`
	LogFile            string `json:"LogFile"`
}

// ReadConfig read
func ReadConfig() error {
	fmt.Println("Reading new file...")
	Config = os.Getenv("HOME") + "/dondiscord.json"
	file, err := ioutil.ReadFile(Config)

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
	LogFile = config.LogFile

	if len(LogFile) < 0 {
		LogFile = "/var/log/dondiscord.log"
	}

	_, err = os.OpenFile(LogFile, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	return nil
}
