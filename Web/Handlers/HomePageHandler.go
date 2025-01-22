package Handlers

import (
	"fmt"
	"groupie-tracker/Internal/Api"
	"html/template"
	"net/http"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	// Load the template
	tmpl := template.Must(template.ParseFiles("Web/Templates/HomePage.gohtml"))

	GroupList := Api.GetAllGroups()

	// Check if the user send a POST request containing a filter
	if r.Method != http.MethodPost {
		// Run the template without filtering
		err := tmpl.ExecuteTemplate(w, "HomePage", GroupList)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error while loading the template : %v", err), http.StatusInternalServerError)
		}
		return
	} else {

		// If the User send a POST request with a filter
		// Get the form value
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error while analysing the form", http.StatusBadRequest)
			return
		}

		if r.FormValue("filterBy-CreationDate") != "" {
			creationDate := r.FormValue("filterBy-CreationDate")
			// Apply the filter
			GroupList = Api.FilterGroupsByCreationDate(creationDate, GroupList)
		}

		// Run the template with the filtered data
		err = tmpl.ExecuteTemplate(w, "HomePage", GroupList)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error while loading the template : %v", err), http.StatusInternalServerError)
		}
	}
}
