package src

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	config "github.com/emnopal/go-telegram-bot/src/config"
	handler "github.com/emnopal/go-telegram-bot/src/handler"
	loader "github.com/emnopal/go-telegram-bot/src/loader"
)

func RunBot() {

	log.Println("Press Ctrl+C to exit.")

	botConf := config.GetBotConfig()

	botUpdateConf := config.GetUpdateBotConfig()

	bot := loader.BotInit(botConf)

	go handler.BotHandler(bot, botUpdateConf)

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT)

	<-interruptChan
	log.Println()
	log.Println("Received Ctrl+C. Exiting...")

}
