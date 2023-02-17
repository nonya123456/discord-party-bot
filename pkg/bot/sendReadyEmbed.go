package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

func (bot *Bot) SendReadyEmbed() error {
	var i = 0
	fields := make([]*discordgo.MessageEmbedField, len(bot.Ready))
	for k := range bot.Ready {
		fields[i] = &discordgo.MessageEmbedField{
			Value: "<@" + k + ">",
		}
		i += 1
	}

	embed := &discordgo.MessageEmbed{
		Title:  "Ready!",
		Fields: fields,
		Color:  1752220,
	}

	_, err := bot.Session.ChannelMessageSendEmbed(bot.ReadyChannel, embed)
	if err != nil {
		return errors.Wrap(err, "Fail to send message")
	}

	return nil
}
