package calendarAPI

import "log"

// TODO: srv.Events.Delete(): make implementation of eventID
func HandleDeleteEvent() error {
	srv, err := CreateService()
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
		return err
	}
	srv.Events.Delete("", "")
	return nil
}
