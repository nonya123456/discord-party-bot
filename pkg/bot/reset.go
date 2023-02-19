package bot

func (b *Bot) Reset() {
	b.Ready = make(map[string]ReadyType)
	b.CurrentTime = nil

	b.ResetTicker.Stop()
	b.UpdateEmbedTicker.Stop()
}
