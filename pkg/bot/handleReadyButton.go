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
	}

	b.Ready[i.Member.User.ID] = Ready

	if len(b.Ready) >= 5 {
		b.SendReadyEmbed()
		b.Reset()
	}

	b.UpdateReadyCheckEmbed()
}
