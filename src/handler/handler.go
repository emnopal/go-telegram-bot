package handler

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func botHandlerSendText(bot *tgbotapi.BotAPI, msg *tgbotapi.Message, text string) {
	response := tgbotapi.NewMessage(msg.Chat.ID, text)
	bot.Send(response)
}

func botHandlerText(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	text := msg.Text
	log.Printf("User send: %s", text)

	switch text {
	case "/start":
		textToSend := "Hello, I'm your bot! Send me a message."
		botHandlerSendText(bot, msg, textToSend)
	case "/end":
		textToSend := "Bye"
		botHandlerSendText(bot, msg, textToSend)
	default:
		reply := "You said: " + text
		botHandlerSendText(bot, msg, reply)
	}
}

// TODO: Handle Audio, Photos, Videos, Sticker etc
func botHandlerMessage(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {

	if msg.Text == "" {
		log.Println("User send other than text which is not yet implemented")
		botHandlerSendText(bot, msg, "Still not implemented other than text")
		return
	}

	botHandlerText(bot, msg)
}

func BotHandler(bot *tgbotapi.BotAPI, config tgbotapi.UpdateConfig) {

	updates, err := bot.GetUpdatesChan(config)

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		botHandlerMessage(bot, update.Message)
	}
}
