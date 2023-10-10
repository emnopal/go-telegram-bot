package main

import (
	runbot "github.com/emnopal/go-telegram-bot/src"
)

func main() {
	bot := runbot.RunBot()
	bot.Start()
}
