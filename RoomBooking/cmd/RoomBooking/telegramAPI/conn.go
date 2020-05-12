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
	WebHookURL = "https://1d32ee2d.ngrok.io"
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
