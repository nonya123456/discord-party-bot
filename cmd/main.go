package main

import (
	"github.com/nonya123456/discord-party-bot/internal/config"
	"github.com/nonya123456/discord-party-bot/pkg/bot"
)

func main() {
	conf := config.New()

	bot, err := bot.New(conf.Token, conf.Channel)
	if err != nil {
		panic(err)
	}

	bot.SendReadyCheckEmbed()

	<-make(chan struct{})
}
