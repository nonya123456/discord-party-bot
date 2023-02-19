package bot

import "github.com/bwmarrin/discordgo"

func (b *Bot) HandlePlayNowButton(i *discordgo.InteractionCreate) {
	l := b.GetNonFiveStackReady()
	b.SendReadyEmbed(l)
}
