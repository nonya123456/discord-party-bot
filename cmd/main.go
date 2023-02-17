package main

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/nonya123456/discord-party-bot/internal/config"
	"github.com/nonya123456/discord-party-bot/pkg/bot"
)

func main() {
	conf := config.New()

	bot, err := bot.New(conf.Token, conf.ReadyCheckChannel, conf.ReadyChannel)
	if err != nil {
		panic(err)
	}

	readyCheckMessage, err := bot.FindReadyCheckEmbedMessage()
	if err != nil {
		panic(err)
	}

	if readyCheckMessage == nil {
		readyCheckMessage, err = bot.SendReadyCheckEmbed()
		if err != nil {
			panic(err)
		}
	}

	bot.Message = readyCheckMessage
	bot.UpdateReadyCheckEmbed()

	bot.AddHandler(func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate,
	) {
		var exists = struct{}{}
		if i.MessageComponentData().CustomID == "ready" {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
			})

			bot.Ready[i.Member.User.ID] = exists

			if len(bot.Ready) >= 5 {
				bot.SendReadyEmbed()
				bot.ResetReady()
			}

			bot.UpdateReadyCheckEmbed()
		} else if bot.Ready[i.Member.User.ID] == exists {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
			})

			delete(bot.Ready, i.Member.User.ID)

			bot.UpdateReadyCheckEmbed()
		}
	})

	ticker := time.NewTicker(30 * time.Minute)
	go func() {
		for {
			<-ticker.C
			bot.ResetReady()
			bot.UpdateReadyCheckEmbed()
		}
	}()

	<-make(chan struct{})
}
