package bot

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

var bot *Bot
var exists = struct{}{}

type Bot struct {
	Session *discordgo.Session
	Message *discordgo.Message
	Ready   map[string]struct{}
	Count   int
}

func New(token string) (*Bot, error) {
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, errors.Wrap(err, "error creating Discord session,")
	}

	err = s.Open()
	if err != nil {
		return nil, errors.Wrap(err, "error opening Discord session,")
	}

	bot := &Bot{
		Session: s,
		Ready:   make(map[string]struct{}),
		Count:   0,
	}

	s.AddHandler(func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate,
	) {
		if i.MessageComponentData().CustomID == "ready" {
			bot.Ready[i.Member.User.ID] = exists

			s.ChannelMessageEditComplex(
				&discordgo.MessageEdit{
					Channel: bot.Message.ChannelID,
					ID:      bot.Message.ID,
					Embed: &discordgo.MessageEmbed{
						Title:       "Ready Check",
						Description: strconv.Itoa(len(bot.Ready)) + "/5",
					},
				},
			)

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
			})
		} else if bot.Ready[i.Member.User.ID] == exists {
			delete(bot.Ready, i.Member.User.ID)

			s.ChannelMessageEditComplex(
				&discordgo.MessageEdit{
					Channel: bot.Message.ChannelID,
					ID:      bot.Message.ID,
					Embed: &discordgo.MessageEmbed{
						Title:       "Ready Check",
						Description: strconv.Itoa(len(bot.Ready)) + "/5",
					},
				},
			)

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
			})
		}
	})

	return bot, nil
}
