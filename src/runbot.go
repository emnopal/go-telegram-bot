package src

import (
	config "github.com/emnopal/go-telegram-bot/src/config"
	handler "github.com/emnopal/go-telegram-bot/src/handler"
	loader "github.com/emnopal/go-telegram-bot/src/loader"
	"gopkg.in/telebot.v3"
)

func RunBot() *telebot.Bot {

	botConf := config.BotConfig()

	botInit := loader.BotInit(botConf)

	bot := handler.BotHandler(botInit)

	return bot
}
