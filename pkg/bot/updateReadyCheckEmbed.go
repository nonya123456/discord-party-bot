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
		playerStr += user.String() + "\n"
	}

	_, err := bot.Session.ChannelMessageEditComplex(
		&discordgo.MessageEdit{
			Channel: bot.Message.ChannelID,
			ID:      bot.Message.ID,
			Embed: &discordgo.MessageEmbed{
				Title:       "Ready Check (reset every 30 mins)",
				Description: ":white_check_mark:\t" + "**" + strconv.Itoa(len(bot.Ready)) + "/5**\n\n" + playerStr,
				Color:       1752220,
			},
		},
	)
	if err != nil {
		return errors.Wrap(err, "Cannot edit ready check message.")
	}

	return nil
}
