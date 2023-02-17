package bot

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/nonya123456/discord-party-bot/internal/config"
	"github.com/pkg/errors"
)

var bot *Bot

type Bot struct {
	ReadyCheckChannel string
	ReadyChannel      string
	Session           *discordgo.Session
	Message           *discordgo.Message
	Ready             map[string]struct{}
	MaxTime           int64
	CurrentTime       *int64
	MainTicker        *time.Ticker
	UpdateEmbedTicker *time.Ticker
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
		ReadyCheckChannel: config.ReadyCheckChannel,
		ReadyChannel:      config.ReadyChannel,
		Session:           s,
		Ready:             make(map[string]struct{}),
		MaxTime:           config.MaxTime,
		CurrentTime:       nil,
		MainTicker:        time.NewTicker(time.Duration(config.MaxTime) * time.Second),
		UpdateEmbedTicker: time.NewTicker(time.Duration(config.UpdateEmbedPeriod) * time.Second),
	}

	go func() {
		for {
			select {
			case <-bot.MainTicker.C:
				fmt.Println("MAIN")
			case <-bot.UpdateEmbedTicker.C:
				fmt.Println("UPDATEEMBED")
			}
		}
	}()

	return bot, nil
}
