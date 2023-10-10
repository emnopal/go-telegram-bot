package src

import (
	config "github.com/emnopal/go-telegram-bot/src/config"
	handler "github.com/emnopal/go-telegram-bot/src/handler"
	loader "github.com/emnopal/go-telegram-bot/src/loader"
)

func RunBot() {

	botConf := config.GetBotConfig()
	botUpdateConf := config.GetUpdateBotConfig()

	bot := loader.BotInit(botConf)

	go handler.BotHandler(bot, botUpdateConf)

	select {}
}
