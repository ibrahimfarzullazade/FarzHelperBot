package commands

import (
	"Helper_Bot/logger"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func init() {
	Register("/help", HelpCommand)
}
func HelpCommand(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Kömək üçün buradayam! 💡",
	})
	logger.Log("User: %s (ChatID: %d) command: %s", update.Message.From.FirstName, update.Message.Chat.ID, update.Message.Text)
}
