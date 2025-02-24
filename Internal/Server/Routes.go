package Server

import (
	"groupie-tracker/Web/Handlers"
	"net/http"
)

func Routes() {
	// Default route
	http.HandleFunc("/", Handlers.HomePageHandler)
	// Others routes
	http.HandleFunc("/group", Handlers.GroupPageHandler)
	http.HandleFunc("/search", Handlers.SearchHandler)
}
