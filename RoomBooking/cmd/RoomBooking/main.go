package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// addr := config.viperConfigVariable("port")

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
