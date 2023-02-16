package bot

func (b *Bot) AddHandler(handler interface{}) {
	b.Session.AddHandler(handler)
}
