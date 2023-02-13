package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

type Bot struct {
	Session *discordgo.Session
	Message *discordgo.Message
	Count   int
}

func New(token string) (*Bot, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, errors.Wrap(err, "error creating Discord session,")
	}

	err = session.Open()
	if err != nil {
		return nil, errors.Wrap(err, "error opening Discord session,")
	}

	return &Bot{
		Session: session,
		Count:   0,
	}, nil
}
