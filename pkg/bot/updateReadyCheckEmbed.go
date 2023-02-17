package bot

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

func (bot *Bot) UpdateReadyCheckEmbed() error {
	playerStr := ""
	for k := range bot.Ready {
		user, err := bot.Session.User(k)
		if err != nil {
			return errors.Wrap(err, "Cannot get user with id: "+k)
		}
		playerStr += user.String()
	}

	_, err := bot.Session.ChannelMessageEditComplex(
		&discordgo.MessageEdit{
			Channel: bot.Message.ChannelID,
			ID:      bot.Message.ID,
			Embed: &discordgo.MessageEmbed{
				Title:       "Ready Check",
				Description: strconv.Itoa(len(bot.Ready)) + "/5\n\n" + playerStr,
			},
		},
	)
	if err != nil {
		return errors.Wrap(err, "Cannot edit ready check message.")
	}

	return nil
}
