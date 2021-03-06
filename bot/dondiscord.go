package dondiscord

import (
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/erraa/dondiscord/config"
	"github.com/erraa/dondiscord/scraper"
)

var botid string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	// Logging
	logfilename := config.LogFile
	f, err := os.OpenFile(logfilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("Couldn't open logfile " + logfilename)
	}
	log.SetOutput(f)

	if err != nil {
		log.Println(err.Error())
		return
	}

	user, err := goBot.User("@me")

	if err != nil {
		log.Println(err.Error())
	}

	botid = user.ID

	goBot.AddHandler(messageHandler)
	err = goBot.Open()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Bot is running with ID", botid)

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !strings.HasPrefix(m.Content, config.BotPrefix) {
		return
	}
	if m.Author.ID == botid {
		return
	}
	log.Println(m.Content)
	if m.Content == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "Ping Successfull")
	}
	if m.Content == "!memes" {
		reddit := scraper.InitReddit(config.RedditUrl, false)
		url := reddit.GetPicture()
		s.ChannelMessageSend(m.ChannelID, url)
	}
}
