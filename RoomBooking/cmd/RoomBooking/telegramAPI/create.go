package telegramAPI

import tgbotapi "gopkg.in/telegram-bot-api.v4"

// Produce dates: slice of titles of the buttons for create event in such format ["Thursday 30.04"],
// produce callBackData: slice of callback data for slice dates in format ["2020-04-30"]
func generateDaysButtons() (dates, callBackData []string) {
	days := GenerateDays()
	for _, day := range days {
		date := day[0]
		callBackData = append(callBackData, day[1]) //итоговый р-т формата: "2020-04-14T11:00:00"
		dates = append(dates, date)
	}
	return
}

// Produce buttons for create event related to days of the week in a column format
func daysButtonsColumn() (daysMenu tgbotapi.InlineKeyboardMarkup) {
	var buttons [][]tgbotapi.InlineKeyboardButton
	backButton := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("<< Back to main menu", "MainMenu"),
	)
	dates, callBackData := generateDaysButtons()
	for i := 0; i < len(dates); i++ {
		button := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(dates[i], callBackData[i]),
		)
		buttons = append(buttons, button)
	}
	buttons = append(buttons, backButton)
	daysMenu = tgbotapi.NewInlineKeyboardMarkup(buttons...)
	return
}
