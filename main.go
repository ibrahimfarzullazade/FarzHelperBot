package main

import (
	"context"
	
	"log"
	"os"
	"os/signal"
	"strings"

	"Helper_Bot/commands"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("TOKEN")
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(defaultHandler),
	}
	for command, handler := range commands.Handlers {
		opts = append(opts, bot.WithMessageTextHandler(command, bot.MatchTypeExact, bot.HandlerFunc(handler)))
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		log.Fatal(err)
	}

	b.Start(ctx)
}

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	text := update.Message.Text
	if handler, ok := commands.Handlers[strings.Fields(text)[0]]; ok {
		handler(ctx, b, update)
		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Bilinməyən komanda. /help yaz!",
	})
}
