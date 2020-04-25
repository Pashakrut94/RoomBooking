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
	WebHookURL = "https://def41599.ngrok.io"
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

		if update.Message.IsCommand() {
			cmdText := update.Message.Command()
			if cmdText == "menu" {
				menuMessage := tgbotapi.NewMessage(update.Message.Chat.ID, "Main menu")
				menuMessage.ReplyMarkup = mainMenu
				bot.Send(menuMessage)
			}
		} else {
			if update.Message.Text == mainMenu.Keyboard[0][0].Text { //list

			} else if update.Message.Text == mainMenu.Keyboard[0][1].Text { //create
				dayMessage := tgbotapi.NewMessage(update.Message.Chat.ID, "Choose the day you want to book")
				dayMessage.ReplyMarkup = daysMenu
				bot.Send(dayMessage)
			} else if update.Message.Text == mainMenu.Keyboard[0][2].Text { //delete

			}

		}

	}
}
