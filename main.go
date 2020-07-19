package main

import (
	"context"
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const token = ""

func main() {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	jwtBot := NewJwtBot(bot)
	ctx := context.Background()
	if err := jwtBot.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
