package bot

import "github.com/bwmarrin/discordgo"

func (b *Bot) HandleReadyButton(i *discordgo.InteractionCreate) {
	b.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
	})

	r, ok := b.Ready[i.Member.User.ID]
	if ok && r == Ready {
		return
	}

	if b.CurrentTime == nil {
		b.StartTicker()
		if !b.NotificationSent {
			err := b.SendNotification()
			if err != nil {
				panic(err)
			}
			b.NotificationSent = true
		}
	}

	b.Ready[i.Member.User.ID] = Ready

	if len(b.Ready) >= 5 {
		l := make([]string, len(b.Ready))
		i := 0
		for k := range b.Ready {
			l[i] = k
			i++
		}

		b.SendReadyEmbed(l)
		b.Reset()
	}

	b.UpdateReadyCheckEmbed()
}
