package bot

func (bot *Bot) ResetReady() {
	bot.Ready = make(map[string]struct{})
}
