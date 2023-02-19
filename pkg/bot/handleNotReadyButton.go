package bot

import "github.com/bwmarrin/discordgo"

func (b *Bot) HandleNotReadyButton(i *discordgo.InteractionCreate) {
	b.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
	})

	_, ok := b.Ready[i.Member.User.ID]
	if !ok {
		return
	}

	delete(b.Ready, i.Member.User.ID)

	if len(b.Ready) == 0 {
		b.Reset()
	}

	b.UpdateReadyCheckEmbed()
}
