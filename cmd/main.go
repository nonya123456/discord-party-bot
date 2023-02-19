package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nonya123456/discord-party-bot/internal/config"
	"github.com/nonya123456/discord-party-bot/pkg/bot"
)

func main() {
	conf := config.New()

	b, err := bot.New(conf.Token, conf)
	if err != nil {
		panic(err)
	}

	readyCheckMessage, err := b.FindReadyCheckEmbedMessage()
	if err != nil {
		panic(err)
	}

	if readyCheckMessage == nil {
		readyCheckMessage, err = b.SendReadyCheckEmbed()
		if err != nil {
			panic(err)
		}
	}

	b.Message = readyCheckMessage
	b.Reset()
	b.UpdateReadyCheckEmbed()

	b.AddHandler(func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate,
	) {
		if i.MessageComponentData().CustomID == string(bot.Ready) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
			})

			_, ok := b.Ready[i.Member.User.ID]
			if ok {
				return
			}

			if b.CurrentTime == nil {
				b.StartTicker()
			}

			b.Ready[i.Member.User.ID] = bot.Ready

			if len(b.Ready) >= 5 {
				b.SendReadyEmbed()
				b.Reset()
			}

			b.UpdateReadyCheckEmbed()
		} else if b.Ready[i.Member.User.ID] == bot.Ready {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
			})

			_, ok := b.Ready[i.Member.User.ID]
			if !ok {
				return
			}

			delete(b.Ready, i.Member.User.ID)

			if len(b.Ready) == 0 {
				b.Reset()
			}

			b.UpdateReadyCheckEmbed()
		}
	})

	go func() {
		for {
			select {
			case <-b.ResetTicker.C:
				b.Reset()
				b.UpdateReadyCheckEmbed()
			case <-b.UpdateEmbedTicker.C:
				if b.CurrentTime == nil {
					continue
				}

				if *b.CurrentTime < b.UpdateEmbedPeriod {
					*b.CurrentTime = 0
				} else {
					*b.CurrentTime -= b.UpdateEmbedPeriod
				}

				b.UpdateReadyCheckEmbed()
			}
		}
	}()

	<-make(chan struct{})
}
