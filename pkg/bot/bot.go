package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

var bot *Bot

type Bot struct {
	Channel string
	Session *discordgo.Session
	Message *discordgo.Message
	Ready   map[string]struct{}
	Count   int
}

func New(token string, channel string) (*Bot, error) {
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, errors.Wrap(err, "error creating Discord session,")
	}

	err = s.Open()
	if err != nil {
		return nil, errors.Wrap(err, "error opening Discord session,")
	}

	bot := &Bot{
		Channel: channel,
		Session: s,
		Ready:   make(map[string]struct{}),
		Count:   0,
	}

	return bot, nil
}
