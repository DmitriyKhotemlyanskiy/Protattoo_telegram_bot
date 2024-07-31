package entities

import (
	//"telegram_bot/entities"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type KeyBoard struct{}

func (k KeyBoard) GetKeyBoard(buttons []Button) tgbotapi.InlineKeyboardMarkup {
	var newKeyboard [][]tgbotapi.InlineKeyboardButton
	for _, button := range buttons {
		btn := tgbotapi.NewInlineKeyboardButtonData(button.GetText(), button.GetData())
		row := tgbotapi.NewInlineKeyboardRow(btn)
		newKeyboard = append(newKeyboard, row)
	}
	return tgbotapi.NewInlineKeyboardMarkup(newKeyboard...)
}
