package main

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/nonya123456/discord-party-bot/internal/config"
	"github.com/nonya123456/discord-party-bot/pkg/bot"
)

func main() {
	conf := config.New()

	bot, err := bot.New(conf.Token, conf.Channel)
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

	bot.AddHandler(func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate,
	) {
		var exists = struct{}{}
		if i.MessageComponentData().CustomID == "ready" {
			bot.Ready[i.Member.User.ID] = exists

			s.ChannelMessageEditComplex(
				&discordgo.MessageEdit{
					Channel: bot.Message.ChannelID,
					ID:      bot.Message.ID,
					Embed: &discordgo.MessageEmbed{
						Title:       "Ready Check",
						Description: strconv.Itoa(len(bot.Ready)) + "/5",
					},
				},
			)

			if len(bot.Ready) >= 5 {
				bot.SendReadyEmbed()
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
			})
		} else if bot.Ready[i.Member.User.ID] == exists {
			delete(bot.Ready, i.Member.User.ID)

			s.ChannelMessageEditComplex(
				&discordgo.MessageEdit{
					Channel: bot.Message.ChannelID,
					ID:      bot.Message.ID,
					Embed: &discordgo.MessageEmbed{
						Title:       "Ready Check",
						Description: strconv.Itoa(len(bot.Ready)) + "/5",
					},
				},
			)

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
			})
		}
	})

	<-make(chan struct{})
}
