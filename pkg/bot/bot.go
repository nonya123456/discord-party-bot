package bot

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/nonya123456/discord-party-bot/internal/config"
	"github.com/pkg/errors"
)

type ReadyType string

const (
	Ready          ReadyType = "ready"
	ReadyFiveStack ReadyType = "ready-5-stack"
)

type Bot struct {
	ReadyCheckChannel   string
	ReadyChannel        string
	Session             *discordgo.Session
	Message             *discordgo.Message
	NotificationMessage *discordgo.Message
	Ready               map[string]ReadyType
	MaxTime             int64
	UpdateEmbedPeriod   int64
	CurrentTime         *int64
	ResetTicker         *time.Ticker
	UpdateEmbedTicker   *time.Ticker
}

func New(token string, config *config.Config) (*Bot, error) {
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, errors.Wrap(err, "error creating Discord session,")
	}

	err = s.Open()
	if err != nil {
		return nil, errors.Wrap(err, "error opening Discord session,")
	}

	bot := &Bot{
		ReadyCheckChannel:   config.ReadyCheckChannel,
		ReadyChannel:        config.ReadyChannel,
		Session:             s,
		NotificationMessage: nil,
		Ready:               make(map[string]ReadyType),
		MaxTime:             config.MaxTime,
		UpdateEmbedPeriod:   config.UpdateEmbedPeriod,
		CurrentTime:         nil,
		ResetTicker:         time.NewTicker(time.Duration(config.MaxTime) * time.Second),
		UpdateEmbedTicker:   time.NewTicker(time.Duration(config.UpdateEmbedPeriod) * time.Second),
	}

	return bot, nil
}
