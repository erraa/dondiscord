package dondiscord

import (
	"fmt"

	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/erraa/dondiscord/config"
	"github.com/erraa/dondiscord/scraper"
)

var botid string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	botid = user.ID

	goBot.AddHandler(messageHandler)
	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running with ID", botid)

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !strings.HasPrefix(m.Content, config.BotPrefix) {
		return
	}
	if m.Author.ID == botid {
		return
	}
	fmt.Println(m.Content)
	if m.Content == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "Ping Successfull")
	}
	if m.Content == "!dankmemes" {
		reddit := scraper.InitReddit(config.RedditUrl, false)
		url := reddit.GetPicture()
		s.ChannelMessageSend(m.ChannelID, url)
	}
}
