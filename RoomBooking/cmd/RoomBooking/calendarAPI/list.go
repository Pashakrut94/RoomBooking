package calendarAPI

import (
	"fmt"
	"log"
	"time"
)

func HandleListEvents() error {
	srv, err := CreateService()
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
		return err
	}

	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List(CalendarId).ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
		return err
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
	return nil
}
