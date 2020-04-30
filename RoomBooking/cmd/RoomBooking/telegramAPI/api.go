package telegramAPI

import (
	"fmt"
	"log"
	"net/http"
	"os"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

// Making connection through ngrok WebHookURL
var (
	BotPort    = os.Getenv("BOT_PORT")
	BotToken   = os.Getenv("BOT_TOKEN")
	WebHookURL = "https://464a87c6.ngrok.io"
)

// Register telegram bot with BotToken
func BotConnection() (bot *tgbotapi.BotAPI, err error) {
	bot, err = tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		log.Fatalf("unable to register new bot with this token: %v\n", err)
		return nil, err
	}
	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(WebHookURL))
	if err != nil {
		log.Fatalf("unable to set webhook %v: %v\n", WebHookURL, err)
		return nil, err
	}

	go http.ListenAndServe(BotPort, nil)
	fmt.Printf("listening on %v\n", BotPort)

	return bot, nil
}

// Makes channel updates, which implements client-server model
func GetUpdates(bot *tgbotapi.BotAPI) {
	updates := bot.ListenForWebhook("/")

	for update := range updates {
		if update.CallbackQuery != nil {
			bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data))
			switch update.CallbackQuery.Data {
			case "ListEvents":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Handler for list events")
				bot.Send(msg)
			case "CreateEvent":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Choose the day you want to book")
				daysMenu := daysButtonsColumn()
				msg.ReplyMarkup = daysMenu
				bot.Send(msg)
			case "DeleteEvent":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Handler for delete events")
				bot.Send(msg)
			case "MainMenu":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Main menu")
				msg.ReplyMarkup = mainMenu
				bot.Send(msg)
			default:
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Choose the time you want to book")
				msg.ReplyMarkup = timeMenu
				bot.Send(msg)
			}
		}
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			switch update.Message.Text {
			case "/menu":
				msg.ReplyMarkup = mainMenu
			}
			bot.Send(msg)
		}
	}
}
