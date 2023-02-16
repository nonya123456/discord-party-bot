package bot

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func (bot *Bot) UpdateReadyCheckEmbed() {
	bot.Session.ChannelMessageEditComplex(
		&discordgo.MessageEdit{
			Channel: bot.Message.ChannelID,
			ID:      bot.Message.ID,
			Embed: &discordgo.MessageEmbed{
				Title:       "Ready Check",
				Description: strconv.Itoa(len(bot.Ready)) + "/5",
			},
		},
	)
}
