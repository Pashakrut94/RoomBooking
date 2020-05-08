package telegramAPI

import tgbotapi "gopkg.in/telegram-bot-api.v4"

var listMenu = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🔍List events", "DeleteEvent"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🔙 Back to main menu", "MainMenu"),
	),
)
