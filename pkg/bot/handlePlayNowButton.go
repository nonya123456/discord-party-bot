package bot

import "github.com/bwmarrin/discordgo"

func (b *Bot) HandlePlayNowButton(i *discordgo.InteractionCreate) {
	b.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
	})

	l := b.GetNonFiveStackReady()
	b.SendReadyEmbed(l)
	b.RemoveReadyByIds(l)

	if len(b.Ready) == 0 {
		b.Reset()
	}

	b.UpdateReadyCheckEmbed()
}
