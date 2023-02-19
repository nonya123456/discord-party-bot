package bot

import "github.com/bwmarrin/discordgo"

func (b *Bot) HandleReadyFiveStackButton(i *discordgo.InteractionCreate) {
	b.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
	})

	r, ok := b.Ready[i.Member.User.ID]
	if ok && r == ReadyFiveStack {
		return
	}

	if b.CurrentTime == nil {
		b.StartTicker()
	}

	b.Ready[i.Member.User.ID] = ReadyFiveStack

	if len(b.Ready) >= 5 {
		b.SendReadyEmbed()
		b.Reset()
	}

	b.UpdateReadyCheckEmbed()
}
