package main

import (
	"github.com/nonya123456/discord-party-bot/internal/config"
	"github.com/nonya123456/discord-party-bot/pkg/bot"
)

func main() {
	conf := config.New()

	bot, err := bot.New(conf.Token)
	if err != nil {
		panic(err)
	}

	bot.SendEmbed()

	<-make(chan struct{})
}
