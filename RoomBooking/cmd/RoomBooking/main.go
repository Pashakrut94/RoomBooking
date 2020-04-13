package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Pashakrut94/cmd/RoomBooking/calendarAPI"
	"github.com/Pashakrut94/cmd/RoomBooking/configuration"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func main() {

	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := calendarAPI.GetClient(config)

	srv, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
	fmt.Println("Upcoming events:")
	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		for _, item := range events.Items {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}
			fmt.Printf("%v (%v)\n", item.Summary, date)
		}
	}

	addr := configuration.ViperConfigVariable("port")

	r := mux.NewRouter()
	r.HandleFunc("/", Handler)
	http.Handle("/", r)
	fmt.Printf("Server starts at %s\n", addr)
	http.ListenAndServe(addr, nil)

	fmt.Println("hello!")
}

func Handler(w http.ResponseWriter, q *http.Request) {
	fmt.Fprintf(w, "Congratz! You have a connection!")
}
