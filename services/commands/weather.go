package commands

import (
	"Helper_Bot/services"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strings"
)

func init() {
	Register("/weather", WeatherCommand)
}

func WeatherCommand(ctx context.Context, b *bot.Bot, update *models.Update) {
	msg := update.Message.Text
	parts := strings.Fields(msg)

	if len(parts) < 2 {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Zəhmət olmasa ünvanı qeyd et. Məs: /hava Baku",
		})
		return
	}

	city := strings.ToUpper(parts[1])

	result := services.Weather(city)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   result,
	})
}
