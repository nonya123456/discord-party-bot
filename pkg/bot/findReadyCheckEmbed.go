package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

func (b *Bot) FindReadyCheckEmbedMessage() (*discordgo.Message, error) {
	if b.Message != nil {
		return b.Message, nil
	}

	messages, err := b.Session.ChannelMessages(b.ReadyCheckChannel, 1, "", "", "")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot get channel messages.")
	}

	if len(messages) == 0 {
		return nil, nil
	}

	return messages[0], nil
}
