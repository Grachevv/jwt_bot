package main

import (
	"context"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type JwtBot struct {
	tBot *tgbotapi.BotAPI
}

func NewJwtBot(tBot *tgbotapi.BotAPI) *JwtBot {
	return &JwtBot{tBot: tBot}
}

func (b *JwtBot) Start(ctx context.Context) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.tBot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		select {
		case <-ctx.Done():
			return err
		default:
		}
		if update.Message == nil {
			continue
		}
		if err := b.sendAnswer(update.Message); err != nil {
			log.Println(err)
		}
	}
	return nil
}

func (b *JwtBot) sendAnswer(message *tgbotapi.Message) error {
	decoded, err := DecodeJwt(message.Text)
	var text string
	if err != nil {
		text = fmt.Sprintln(err)
	} else {
		text = fmt.Sprintf("```\n%s\n```", decoded)
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	msg.ParseMode = tgbotapi.ModeMarkdown
	msg.ReplyToMessageID = message.MessageID
	if _, err := b.tBot.Send(msg); err != nil {
		return err
	}
	return nil
}
