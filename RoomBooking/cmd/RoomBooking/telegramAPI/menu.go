package telegramAPI

import (
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

var mainMenu = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🔍List events", "ListEvents"),
		tgbotapi.NewInlineKeyboardButtonData("📅Create event", "CreateEvent"),
		tgbotapi.NewInlineKeyboardButtonData("❌Delete event", "DeleteEvent"),
	),
)
