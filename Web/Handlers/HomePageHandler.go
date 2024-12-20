package Handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	// Load the template
	tmpl := template.Must(template.ParseFiles("Web/Templates/HomePage.gohtml"))

	// Run the template
	err := tmpl.ExecuteTemplate(w, "HomePage", nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while loading the template : %v", err), http.StatusInternalServerError)
	}
}
