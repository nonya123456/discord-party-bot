package bot

import "github.com/pkg/errors"

func (b *Bot) SendNotification() error {
	msg := "Please head over to " + "<#" + b.ReadyCheckChannel + ">" + " and let us know if you're ready for the upcoming match."

	_, err := b.Session.ChannelMessageSend(b.ReadyChannel, msg)
	if err != nil {
		return errors.Wrap(err, "Fail to send message")
	}

	return nil
}
