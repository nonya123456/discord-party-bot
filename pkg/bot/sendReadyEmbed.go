package bot

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

func (bot *Bot) SendReadyEmbed(l []string) error {
	fields := make([]*discordgo.MessageEmbedField, len(l))
	for i, userId := range l {
		fields[i] = &discordgo.MessageEmbedField{
			Value: "<@" + userId + ">",
		}
	}

	embed := &discordgo.MessageEmbed{
		Title:  "Ready! (" + strconv.Itoa(len(l)) + "-stack)",
		Fields: fields,
		Color:  1752220,
	}

	_, err := bot.Session.ChannelMessageSendEmbed(bot.ReadyChannel, embed)
	if err != nil {
		return errors.Wrap(err, "Fail to send message")
	}

	return nil
}
