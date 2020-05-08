package telegramAPI

import (
	"fmt"

	create "github.com/Pashakrut94/cmd/RoomBooking/telegramAPI/create"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

type CreateEvent struct {
	UserLastName string `json:"user_last_name"`
	Date         string `json:"date"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	Location     string `json:"location"`
	Summary      string `json:"summary"`

	// Description string `json:"description"` // Description add to struct of CalendarAPI: "AndersenBookingBot: made by %v, at %v", UserLastName, time.Now()
}

type UserMessage struct {
	LastID string
}

var (
	createEventMap map[int]*CreateEvent
	message        map[int]*UserMessage
)

// Makes channel updates, which implements client-server model
func GetUpdates(bot *tgbotapi.BotAPI) {
	updates := bot.ListenForWebhook("/")
	lastID := 0 //Сделать структуру message (у каждого свой lastID) Message{UserID: int, lastID: int}
	// userLastName := update.CallbackQuery.From.LastName // Фамилия
	// userID := update.CallbackQuery.From.UserID // Число
	for update := range updates {

		if update.Message != nil {
			cmdText := update.Message.Command()
			switch cmdText {
			case "menu":

				if lastID != 0 && update.CallbackQuery != nil {
					msg := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, lastID, "Main menu")
					msg.ReplyMarkup = &mainMenu
					ms, _ := bot.Send(msg)
					lastID = ms.MessageID
				} else {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Main menu")
					msg.ReplyMarkup = mainMenu
					ms, _ := bot.Send(msg)
					lastID = ms.MessageID
				}
			}

		} else {

			if update.CallbackQuery != nil {
				fmt.Println(update.CallbackQuery.Data)

				switch update.CallbackQuery.Data {

				case "ListEvents":
					msg := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, lastID, "List events handler")
					fmt.Println(userID)
					msg.ReplyMarkup = &listMenu
					ms, _ := bot.Send(msg)
					lastID = ms.MessageID
				case "CreateEvent":
					msg := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, lastID, "Choose the day you want to book")
					daysMenu := create.DaysButtonsColumn()
					msg.ReplyMarkup = &daysMenu
					ms, _ := bot.Send(msg)
					lastID = ms.MessageID
				case "DeleteEvent":
					msg := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, lastID, "Delete event handler")
					msg.ReplyMarkup = &deleteMenu
					ms, _ := bot.Send(msg)
					lastID = ms.MessageID
				case "MainMenu":
					msg := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, lastID, "Main menu")
					msg.ReplyMarkup = &mainMenu
					ms, _ := bot.Send(msg)
					lastID = ms.MessageID
				case "2020-05-08": // data from create handler
					msg := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, lastID, "Choose time you want to book")
					msg.ReplyMarkup = &timeMenu
					ms, _ := bot.Send(msg)
					lastID = ms.MessageID
				}

			}

		}

	}
}
