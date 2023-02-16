package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

var bot *Bot

type Bot struct {
	ReadyCheckChannel string
	ReadyChannel      string
	Session           *discordgo.Session
	Message           *discordgo.Message
	Ready             map[string]struct{}
}

func New(token string, readyCheckChannel string, readyChannel string) (*Bot, error) {
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, errors.Wrap(err, "error creating Discord session,")
	}

	err = s.Open()
	if err != nil {
		return nil, errors.Wrap(err, "error opening Discord session,")
	}

	bot := &Bot{
		ReadyCheckChannel: readyCheckChannel,
		ReadyChannel:      readyChannel,
		Session:           s,
		Ready:             make(map[string]struct{}),
	}

	return bot, nil
}
