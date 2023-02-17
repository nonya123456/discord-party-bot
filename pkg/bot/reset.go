package bot

func (bot *Bot) Reset() {
	bot.Ready = make(map[string]struct{})
	bot.CurrentTime = nil

	bot.ResetTicker.Stop()
	bot.UpdateEmbedTicker.Stop()
}
