package bot

import "time"

func (bot *Bot) StartTicker() {
	bot.CurrentTime = new(int64)
	*bot.CurrentTime = bot.MaxTime

	bot.ResetTicker.Reset(time.Duration(bot.MaxTime) * time.Second)
	bot.UpdateEmbedTicker.Reset(time.Duration(bot.UpdateEmbedPeriod) * time.Second)
}
