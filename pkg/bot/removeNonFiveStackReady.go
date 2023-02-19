package bot

func (b *Bot) RemoveReadyByIds(l []string) {
	for _, userId := range l {
		delete(b.Ready, userId)
	}
}
