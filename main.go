package main

import (
	"telegram_bot/config"
	"telegram_bot/handlers"
)

func main() {
	bot, updates := config.Bot{}.GetBot()
	handlers.HandleUpdates(bot, updates)
}
