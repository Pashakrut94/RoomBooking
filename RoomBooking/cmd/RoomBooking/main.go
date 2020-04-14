package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Pashakrut94/cmd/RoomBooking/calendarAPI"
	"github.com/gorilla/mux"
)

var (
	localHost = os.Getenv("LOCALHOST")
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/create", calendarAPI.CreateEvent())
	router.HandleFunc("/api/list", calendarAPI.ListEvents())

	http.Handle("/", router)
	fmt.Printf("Server starts at %s\n", localHost)
	http.ListenAndServe(localHost, nil)
}
