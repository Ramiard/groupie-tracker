package Server

import (
	"groupie-tracker/Web/Handlers"
	"net/http"
)

func Routes() {
	// Route par défaut
	http.HandleFunc("/", Handlers.HomePageHandler)

	// Reste des routes
	http.HandleFunc("/group", Handlers.GroupPageHandler)
	//http.HandleFunc("/search", Handlers.SearchHandler)
}
