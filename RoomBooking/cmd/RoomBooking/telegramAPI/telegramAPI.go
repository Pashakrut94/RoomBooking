package telegramAPI

import (
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

var mainMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🔍List events"),
		tgbotapi.NewKeyboardButton("📅Create event"),
		tgbotapi.NewKeyboardButton("❌Delete event"),
	),
)

var daysMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Monday"),
		tgbotapi.NewKeyboardButton("Tuesday"),
		tgbotapi.NewKeyboardButton("Wednesday"),
		tgbotapi.NewKeyboardButton("Thursday"),
		tgbotapi.NewKeyboardButton("Friday"),
	),
)

type CreateEventBot struct {
	Summary     string `json:"summary"`
	Location    string `json:"location"`
	Description string `json:"description"`
	// Date        string `json:"date"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
