package bot

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

func (bot *Bot) SendReadyCheckEmbed() (*discordgo.Message, error) {
	if bot.Message != nil {
		return nil, errors.New("Message is already there")
	}

	embed := &discordgo.MessageEmbed{
		Title:       "Ready Check (reset every 30 mins)",
		Description: strconv.Itoa(len(bot.Ready)) + "/5",
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

	message, err := bot.Session.ChannelMessageSendComplex(bot.ReadyCheckChannel, &discordgo.MessageSend{
		Embed:      embed,
		Components: []discordgo.MessageComponent{actionsRow},
	})
	if err != nil {
		return nil, errors.Wrap(err, "Fail to send message")
	}

	bot.Message = message
	return message, nil
}
