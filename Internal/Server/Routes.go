package Server

import (
	"groupie-tracker/Web/Handlers"
	"net/http"
)

func Routes() {
	// Default route
	http.HandleFunc("/", Handlers.HomePageHandler)

	//
	http.HandleFunc("/group", Handlers.GroupPageHandler)
	http.HandleFunc("/search", Handlers.SearchHandler)
}
