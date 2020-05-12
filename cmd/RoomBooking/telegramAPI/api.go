package telegramAPI

import (
	"fmt"
	"log"

	create "github.com/Pashakrut94/cmd/RoomBooking/telegramAPI/create"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

type CreateEvent struct {
	UserID        int    `json:"user_id"`
	UserFirstName string `json:"user_first_name"`
	UserLastName  string `json:"user_last_name"`
	Date          string `json:"date"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
	Location      string `json:"location"`
	Summary       string `json:"summary"`

	// Description string `json:"description"` // Description add to struct of CalendarAPI: "AndersenBookingBot: made by %v, at %v", UserLastName, time.Now()
}

type UserMessage struct {
	LastID int
}

var (
	createEventMap map[int]*CreateEvent
	message        map[int]*UserMessage
)

func init() {
	createEventMap = make(map[int]*CreateEvent)
	message = make(map[int]*UserMessage)
}

// Makes channel updates, which implements client-server model
func GetUpdates(bot *tgbotapi.BotAPI) {
	updates := bot.ListenForWebhook("/")

	// userLastName := update.CallbackQuery.From.LastName // Фамилия

	for update := range updates {

		if update.Message != nil {
			cmdText := update.Message.Command()
			switch cmdText {
			case "menu":

				msgMap, ok := message[update.Message.From.ID]

				// TODO: Rework 1 case? Is it necessary?
				if ok && update.CallbackQuery != nil {
					msg := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, msgMap.LastID, "Main menu")
					msg.ReplyMarkup = &mainMenu
					ms, _ := bot.Send(msg)
					message[update.CallbackQuery.From.ID].LastID = ms.MessageID
					createEventMap[update.CallbackQuery.From.ID] = new(CreateEvent) // ??? Is it right to make new map in this case?
					log.Printf("messsage owner: %v, data: %v\n ", update.CallbackQuery.From.ID, update.CallbackQuery.Data)

				} else {
					message[update.Message.From.ID] = new(UserMessage)
					createEventMap[update.Message.From.ID] = new(CreateEvent)

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Main menu")
					msg.ReplyMarkup = mainMenu
					ms, _ := bot.Send(msg)
					message[update.Message.From.ID].LastID = ms.MessageID

					createEventMap[update.Message.From.ID].UserID = update.Message.From.ID
					createEventMap[update.Message.From.ID].UserFirstName = update.Message.From.FirstName
					createEventMap[update.Message.From.ID].UserLastName = update.Message.From.LastName

					log.Printf("messsage owner: %v, data: %v; createEventMap: %+v\n ", update.Message.From.ID, update.Message.Command(), *createEventMap[update.Message.From.ID])
				}
			}

		} else {

			if update.CallbackQuery != nil {

				log.Printf("CallBackQueryData = %v\n", update.CallbackQuery.Data)
				msgMap, ok := message[update.CallbackQuery.From.ID]

				if ok {
					switch update.CallbackQuery.Data {
					case "ListEvents":
						msg := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, msgMap.LastID, "List events handler")
						msg.ReplyMarkup = &listMenu
						ms, _ := bot.Send(msg)
						log.Printf("messsage owner: %v, data: %v\n ", update.CallbackQuery.From.ID, update.CallbackQuery.Data)
						message[update.CallbackQuery.From.ID].LastID = ms.MessageID
					case "CreateEvent":
						msg := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, msgMap.LastID, "Choose the day you want to book")
						daysMenu := create.DaysButtonsColumn()
						msg.ReplyMarkup = &daysMenu
						ms, _ := bot.Send(msg)

						log.Printf("messsage owner: %v, data: %v\n ", update.CallbackQuery.From.ID, update.CallbackQuery.Data)
						message[update.CallbackQuery.From.ID].LastID = ms.MessageID
					case "DeleteEvent":
						msg := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, msgMap.LastID, "Delete event handler")
						msg.ReplyMarkup = &deleteMenu
						ms, _ := bot.Send(msg)
						log.Printf("messsage owner: %v, data: %v\n ", update.CallbackQuery.From.ID, update.CallbackQuery.Data)
						message[update.CallbackQuery.From.ID].LastID = ms.MessageID
					case "MainMenu":
						msg := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, msgMap.LastID, "Main menu")
						msg.ReplyMarkup = &mainMenu
						ms, _ := bot.Send(msg)
						log.Printf("messsage owner: %v, data: %v\n ", update.CallbackQuery.From.ID, update.CallbackQuery.Data)
						message[update.CallbackQuery.From.ID].LastID = ms.MessageID
					case "Date": // data from create handler
						msg := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, msgMap.LastID, "Choose time you want to book")
						msg.ReplyMarkup = &timeMenu
						ms, _ := bot.Send(msg)

						fmt.Printf("Text from button: %+v\n", update.CallbackQuery.Message)

						log.Printf("messsage owner: %v, data: %v\n ", update.CallbackQuery.From.ID, update.CallbackQuery.Data)
						message[update.CallbackQuery.From.ID].LastID = ms.MessageID
					}

				} else {
					// Если пользователя нету в мапе то почистить весь диалог
					// Этот case о: если остановил сервер, а у клиента осталось меню
					// и он нажимает на кнопку(получает data), то в ответ он получит новое "Main menu"
					message[update.CallbackQuery.From.ID] = new(UserMessage)
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Main menu")
					msg.ReplyMarkup = mainMenu
					ms, _ := bot.Send(msg)
					log.Printf("messsage owner: %v, data: %v\n ", update.CallbackQuery.From.ID, update.CallbackQuery.Data)
					message[update.CallbackQuery.From.ID].LastID = ms.MessageID
				}
			}
		}
	}
}
