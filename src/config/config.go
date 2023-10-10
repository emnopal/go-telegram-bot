package config

import (
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

type BotConfig struct {
	Token string
	Debug bool
}

func GetBotConfig() BotConfig {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	token := os.Getenv("TELEGRAM_TOKEN")
	debug, err := strconv.ParseBool(os.Getenv("TELEGRAM_BOT_DEBUG"))

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	conf := BotConfig{
		Token: token,
		Debug: debug,
	}

	return conf
}

func GetUpdateBotConfig() tgbotapi.UpdateConfig {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	return updateConfig
}
