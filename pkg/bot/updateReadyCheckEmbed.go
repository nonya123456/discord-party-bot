package bot

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

func (bot *Bot) UpdateReadyCheckEmbed() error {
	timeRemainingStr := ""
	if bot.CurrentTime != nil {
		timeRemainingStr = ":hourglass:\t" + "**Time Remaining: ~" + strconv.Itoa(int(*bot.CurrentTime/60)) + " minutes**"
	}

	countStr := ":white_check_mark:\t" + "**Count: " + strconv.Itoa(len(bot.Ready)) + "/5**"

	playerStr := ""
	for k, rt := range bot.Ready {
		user, err := bot.Session.User(k)
		if err != nil {
			return errors.Wrap(err, "Cannot get user with id: "+k)
		}

		if rt == Ready {
			playerStr += user.String() + "\n"
		} else if rt == ReadyFiveStack {
			playerStr += user.String() + " (5-stack)\n"
		}
	}

	_, err := bot.Session.ChannelMessageEditComplex(
		&discordgo.MessageEdit{
			Channel: bot.Message.ChannelID,
			ID:      bot.Message.ID,
			Embed: &discordgo.MessageEmbed{
				Title:       "Ready Check",
				Description: timeRemainingStr + "\n\n" + countStr + "\n\n" + playerStr,
				Color:       1752220,
			},
		},
	)
	if err != nil {
		return errors.Wrap(err, "Cannot edit ready check message.")
	}

	return nil
}
