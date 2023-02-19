package bot

func (b *Bot) GetNonFiveStackReady() []string {
	l := []string{}
	for k, rt := range b.Ready {
		if rt == Ready {
			l = append(l, k)
		}
	}

	return l
}
