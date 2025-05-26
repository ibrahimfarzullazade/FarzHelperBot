package commands

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type CommandHandler func(ctx context.Context, b *bot.Bot, update *models.Update)

var Handlers = map[string]CommandHandler{}

func Register(command string, handler CommandHandler) {
	Handlers[command] = handler
}
