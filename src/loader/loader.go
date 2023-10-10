package loader

import (
	"log"

	c "github.com/emnopal/go-telegram-bot/src/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func BotInit(conf c.BotConfig) *tgbotapi.BotAPI {

	bot, err := tgbotapi.NewBotAPI(conf.Token)

	bot.Debug = conf.Debug

	if err != nil {
		log.Fatal(err)
		return nil
	}

	log.Printf("Authorized as @%s", bot.Self.UserName)

	return bot
}
