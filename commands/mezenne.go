package commands

import (
	"Helper_Bot/logger"
	"Helper_Bot/services"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strings"
)

func init() {
	Register("/mezenne", MezenneCommand)
}

func MezenneCommand(ctx context.Context, b *bot.Bot, update *models.Update) {
	msg := update.Message.Text
	parts := strings.Fields(msg)

	if len(parts) < 3 {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Zəhmət olmasa valyutanı tam qeyd et. Məs: /mezenne USD AZN",
		})
		return
	}

	currency := strings.ToUpper(parts[1])
	currency2 := strings.ToUpper(parts[2])

	result := services.Mezenne(currency, currency2)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   result,
	})
	logger.Log("User: %s (ChatID: %d) command: %s", update.Message.From.FirstName, update.Message.Chat.ID, update.Message.Text)

}
