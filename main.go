package main

import (
	"fmt"

	"github.com/erraa/dondiscord/bot"
	"github.com/erraa/dondiscord/config"
)

var BotID string

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	dondiscord.Start()

	<-make(chan struct{})

	return
}
