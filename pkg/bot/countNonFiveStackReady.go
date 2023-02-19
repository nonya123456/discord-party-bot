package bot

func (b *Bot) CountNonFiveStackReady() int {
	total := 0
	for _, rt := range b.Ready {
		if rt == Ready {
			total += 1
		}
	}

	return total
}
