package handler

import (
	"log"

	"gopkg.in/telebot.v3"
)

func BotHandler(bot *telebot.Bot) *telebot.Bot {

	bot.Handle("/start", func(c telebot.Context) error {
		err := c.Send("Hello, I'm your bot! Send me a message.")
		if err != nil {
			log.Fatal(err)
		}
		return err
	})

	bot.Handle("/end", func(c telebot.Context) error {
		err := c.Send("Bye")
		if err != nil {
			log.Fatal(err)
		}
		return err
	})

	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		text := c.Message().Text
		log.Printf("User send: %s", text)
		reply := "You said: " + text
		err := c.Send(reply)
		if err != nil {
			log.Fatal(err)
		}
		return err
	})

	return bot
}
