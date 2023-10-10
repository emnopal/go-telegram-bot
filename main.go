package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"
)

func InitBot(conf telebot.Settings) *telebot.Bot {
	bot, err := telebot.NewBot(conf)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	log.Printf("Authorized as @%s", bot.Me.Username)

	return bot
}

func HandleBot(bot *telebot.Bot) *telebot.Bot {

	// Handle the "/start" command
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

	// Handle incoming messages
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

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	token := os.Getenv("TELEGRAM_TOKEN")

	botConf := telebot.Settings{
		Token: token,
		Poller: &telebot.LongPoller{
			Timeout: 10 * time.Second,
		},
	}

	botInit := InitBot(botConf)

	bot := HandleBot(botInit)

	bot.Start()
}
