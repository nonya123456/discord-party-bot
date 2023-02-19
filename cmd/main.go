package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nonya123456/discord-party-bot/internal/config"
	"github.com/nonya123456/discord-party-bot/pkg/bot"
)

func main() {
	conf := config.New()

	bot, err := bot.New(conf.Token, conf)
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
	bot.Reset()
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

			_, ok := bot.Ready[i.Member.User.ID]
			if ok {
				return
			}

			if bot.CurrentTime == nil {
				bot.StartTicker()
			}

			bot.Ready[i.Member.User.ID] = exists

			if len(bot.Ready) >= 5 {
				bot.SendReadyEmbed()
				bot.Reset()
			}

			bot.UpdateReadyCheckEmbed()
		} else if bot.Ready[i.Member.User.ID] == exists {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
			})

			_, ok := bot.Ready[i.Member.User.ID]
			if !ok {
				return
			}

			delete(bot.Ready, i.Member.User.ID)

			if len(bot.Ready) == 0 {
				bot.Reset()
			}

			bot.UpdateReadyCheckEmbed()
		}
	})

	go func() {
		for {
			select {
			case <-bot.ResetTicker.C:
				bot.Reset()
				bot.UpdateReadyCheckEmbed()
			case <-bot.UpdateEmbedTicker.C:
				if bot.CurrentTime == nil {
					continue
				}

				if *bot.CurrentTime < bot.UpdateEmbedPeriod {
					*bot.CurrentTime = 0
				} else {
					*bot.CurrentTime -= bot.UpdateEmbedPeriod
				}

				bot.UpdateReadyCheckEmbed()
			}
		}
	}()

	<-make(chan struct{})
}
