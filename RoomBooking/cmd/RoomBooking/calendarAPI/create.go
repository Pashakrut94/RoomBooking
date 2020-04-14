package calendarAPI

import (
	"fmt"
	"log"

	"google.golang.org/api/calendar/v3"
)

func HandleCreateEvent(evnt Event) error {

	srv, err := CreateService()
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
		return err
	}

	event := &calendar.Event{
		Summary:     evnt.Summary,
		Location:    evnt.Location,
		Description: evnt.Description,
		Start: &calendar.EventDateTime{
			DateTime: evnt.StartTime, // "2020-04-14T11:00:00"
			TimeZone: TimeZone,
		},
		End: &calendar.EventDateTime{
			DateTime: evnt.EndTime, //"2020-04-14T13:00:00",
			TimeZone: TimeZone,
		},
	}

	event, err = srv.Events.Insert(CalendarId, event).Do()
	if err != nil {
		log.Fatalf("Unable to create event. %v\n", err)
		return err
	}
	fmt.Printf("Event created: %s\n", event.HtmlLink)
	return nil
}
