package main

import (
	"context"
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	botToken, exist := os.LookupEnv("BOT_TOKEN")
	if !exist {
		log.Panic("Doesn't found bot token")
	}
	bot, err := tgbotapi.NewBotAPI(botToken)
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
