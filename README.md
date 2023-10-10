# Telegram Bot

Simple Telegram Bot using Golang

## How to run

1. Run: `go mod init`
2. Then, run: `go mod tidy`
3. Copy `.env.example` then rename it to `.env`
4. Go to [Here](https://t.me/BotFather) and type `/newbot`
5. Fill key of token to `TELEGRAM_TOKEN=` inside `.env` files
6. Run the bot with `go run main.go`

## Modifications

For modification, simply add or remove handler inside `handler.go`.

For advanced modification just change `config.go`, `loader.go` or `runbot.go` as you wish

## To Do

1. Add advanced handler (maybe games, downloader or feeds)
2. Add UnitTest
3. Add docker
