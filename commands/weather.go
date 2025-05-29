package commands

import (
	"Helper_Bot/logger"
	"Helper_Bot/services"
	"context"
	"strings"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/patrickmn/go-cache"
)

var (
	c = cache.New(5*time.Minute, 10*time.Minute)
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
	cacheKey := "city:" + city

	val, found := c.Get(cacheKey)
	var result string

	if found {
		result = val.(string)
	} else {
		result = services.Weather(city)
		c.Set(cacheKey, result, cache.DefaultExpiration)
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   result,
	})
	logger.Log("User: %s (ChatID: %d) command: %s", update.Message.From.FirstName, update.Message.Chat.ID, update.Message.Text)

}
