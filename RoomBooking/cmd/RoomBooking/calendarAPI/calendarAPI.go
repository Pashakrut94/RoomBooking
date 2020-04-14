package calendarAPI

// TODO: Make function for merge date and time in one field
type Event struct {
	Summary     string `json:"summary"`
	Location    string `json:"location"`
	Description string `json:"description"`
	// Date        string `json:"date"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
