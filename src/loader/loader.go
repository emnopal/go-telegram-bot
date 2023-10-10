package loader

import (
	"log"

	"gopkg.in/telebot.v3"
)

func BotInit(conf telebot.Settings) *telebot.Bot {

	bot, err := telebot.NewBot(conf)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	log.Printf("Authorized as @%s", bot.Me.Username)

	return bot
}
