package main

import (
	"log"
	"os"

	"github.com/Pashakrut94/cmd/RoomBooking/telegramAPI"
)

var (
	localHost = os.Getenv("LOCALHOST")
)

func main() {
	bot, err := telegramAPI.BotConnection()
	if err != nil {
		log.Fatalf("unable to register bot %s:", err)
	}
	telegramAPI.GetUpdates(bot)

	// router := mux.NewRouter()

	// router.HandleFunc("/api/create", calendarAPI.CreateEvent())
	// router.HandleFunc("/api/list", calendarAPI.ListEvents())

	// http.Handle("/", router)
	// fmt.Printf("Server starts at %s\n", localHost)
	// http.ListenAndServe(localHost, nil)
}
