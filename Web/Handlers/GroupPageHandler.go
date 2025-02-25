package Handlers

import (
	"fmt"
	"groupie-tracker/Internal/Api"
	"html/template"
	"net/http"
	"strconv"
)

func GroupPageHandler(w http.ResponseWriter, r *http.Request) {
	// Load the template
	tmpl := template.Must(template.ParseFiles("Web/Templates/GroupPage.gohtml"))

	// Get the form value
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error while analysing the form", http.StatusBadRequest)
		return
	}

	// Check if the user send a POST request containing an ID
	if r.Method != http.MethodPost {
		fmt.Println("LOG: No POST used redirecting to /HomePage")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Print("GROUP-PAGE LOG: Given ID : '", r.FormValue("id"), "' \n")

	// Get the group infos
	groupInfos := Api.GetGroupInfos(r.FormValue("id"))

	// Convert the form value 'string' to an 'int'
	_, err = strconv.Atoi(r.FormValue("id"))
	// Check if the Id is empty or if it's not a number
	if r.FormValue("id") == "" || groupInfos.Id == 0 || err != nil {
		fmt.Println("LOG: Group not found or Id is in a wrong format")
		http.Error(w, "Group not found.\nOr the given Id is in a wrong format", http.StatusBadRequest)
		return
	}

	// Run the template
	err = tmpl.ExecuteTemplate(w, "GroupPage", groupInfos)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while loading the template : %v", err), http.StatusInternalServerError)
	}

}
