package telegramAPI

import (
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

var mainMenu = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🔍List events", "ListEvents"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("📅Create event", "CreateEvent"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("❌Delete event", "DeleteEvent"),
	),
)

// "2020-04-14T11:00:00"
var timeMenu = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("9:00", "T9:00:00"),
		tgbotapi.NewInlineKeyboardButtonData("9:15", "T9:15:00"),
		tgbotapi.NewInlineKeyboardButtonData("9:30", "T9:30:00"),
		tgbotapi.NewInlineKeyboardButtonData("9:45", "T9:45:00"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("10:00", "T10:00:00"),
		tgbotapi.NewInlineKeyboardButtonData("10:15", "T10:15:00"),
		tgbotapi.NewInlineKeyboardButtonData("10:30", "T10:30:00"),
		tgbotapi.NewInlineKeyboardButtonData("10:45", "T10:45:00"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("11:00", "T11:00:00"),
		tgbotapi.NewInlineKeyboardButtonData("11:15", "T11:15:00"),
		tgbotapi.NewInlineKeyboardButtonData("11:30", "T11:30:00"),
		tgbotapi.NewInlineKeyboardButtonData("11:45", "T11:45:00"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("12:00", "T12:00:00"),
		tgbotapi.NewInlineKeyboardButtonData("12:15", "T12:15:00"),
		tgbotapi.NewInlineKeyboardButtonData("12:30", "T12:30:00"),
		tgbotapi.NewInlineKeyboardButtonData("12:45", "T12:45:00"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("13:00", "T13:00:00"),
		tgbotapi.NewInlineKeyboardButtonData("13:15", "T13:15:00"),
		tgbotapi.NewInlineKeyboardButtonData("13:30", "T13:30:00"),
		tgbotapi.NewInlineKeyboardButtonData("13:45", "T13:45:00"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("14:00", "T14:00:00"),
		tgbotapi.NewInlineKeyboardButtonData("14:15", "T14:15:00"),
		tgbotapi.NewInlineKeyboardButtonData("14:30", "T14:30:00"),
		tgbotapi.NewInlineKeyboardButtonData("14:45", "T14:45:00"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("15:00", "T15:00:00"),
		tgbotapi.NewInlineKeyboardButtonData("15:15", "T15:15:00"),
		tgbotapi.NewInlineKeyboardButtonData("15:30", "T15:30:00"),
		tgbotapi.NewInlineKeyboardButtonData("15:45", "T15:45:00"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("16:00", "T16:00:00"),
		tgbotapi.NewInlineKeyboardButtonData("16:15", "T16:15:00"),
		tgbotapi.NewInlineKeyboardButtonData("16:30", "T16:30:00"),
		tgbotapi.NewInlineKeyboardButtonData("16:45", "T16:45:00"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("17:00", "T17:00:00"),
		tgbotapi.NewInlineKeyboardButtonData("17:15", "T17:15:00"),
		tgbotapi.NewInlineKeyboardButtonData("17:30", "T17:30:00"),
		tgbotapi.NewInlineKeyboardButtonData("17:45", "T17:45:00"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("<< Back to choose day", "CreateEvent"),
	),
)
