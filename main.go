package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("8303648766:AAGQ5jwPDxzefxnqn_cbKpvztSYVDAoGu18")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Text == "/start" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Здарова нолик ебаный)")
			bot.Send(msg)
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ты что-то насрал: "+update.Message.Text)
		bot.Send(msg)
	}
}
