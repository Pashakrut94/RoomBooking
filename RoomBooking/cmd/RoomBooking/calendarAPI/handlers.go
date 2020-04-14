package calendarAPI

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ListEvents() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		HandleListEvents()
	}
}

// Error handle
func CreateEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var event Event
		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
			log.Fatalf("error parsing request: %s", err)
			return
		}
		if err := HandleCreateEvent(event); err != nil {
			log.Fatalf("error creating new event %s", err)
			return
		}
		fmt.Fprintf(w, "Created '%s'\n", event.Description)

	}
}

// func DeleteEvent() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		HandleDeleteEvent()
// 	}
// }
