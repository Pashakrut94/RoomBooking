package telegramAPI

import tgbotapi "gopkg.in/telegram-bot-api.v4"

var deleteMenu = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("❌Delete event", "ListEvents"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🔙 Back to main menu", "MainMenu"),
	),
)
