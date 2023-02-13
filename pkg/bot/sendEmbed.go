package bot

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

func (bot *Bot) SendEmbed() error {
	embed := &discordgo.MessageEmbed{
		Title:       "Ready Check",
		Description: strconv.Itoa(bot.Count),
	}

	readyButton := discordgo.Button{
		CustomID: "ready",
		Label:    "Ready",
		Style:    discordgo.PrimaryButton,
	}

	notReadyButton := discordgo.Button{
		CustomID: "not-ready",
		Label:    "Not Ready",
		Style:    discordgo.SecondaryButton,
	}

	actionsRow := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{readyButton, notReadyButton},
	}

	message, err := bot.Session.ChannelMessageSendComplex("760902112028262452", &discordgo.MessageSend{
		Embeds:     []*discordgo.MessageEmbed{embed},
		Components: []discordgo.MessageComponent{actionsRow},
	})
	if err != nil {
		return errors.Wrap(err, "Fail to send message")
	}

	bot.Message = message
	return nil
}