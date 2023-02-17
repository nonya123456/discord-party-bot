package bot

func (bot *Bot) Reset() {
	bot.Ready = make(map[string]struct{})
}
