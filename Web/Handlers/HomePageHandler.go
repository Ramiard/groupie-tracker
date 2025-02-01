package Handlers

import (
	"fmt"
	"groupie-tracker/Internal/Api"
	"html/template"
	"net/http"
	"strconv"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	// Load the template
	tmpl := template.Must(template.ParseFiles("Web/Templates/HomePage.gohtml"))

	var HomePageData Api.Data
	HomePageData.Groups = Api.GetAllGroups()
	HomePageData.Countries = Api.GetAllCountries(HomePageData.Groups)
	// Check if the user send a POST request containing a filter
	if r.Method != http.MethodPost {
		// Run the template without filtering
		err := tmpl.ExecuteTemplate(w, "HomePage", HomePageData)
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

		// Checking the 'CreationDate' filter
		fmt.Println("FILTER LOG : minCreationDate = [", r.FormValue("filterBy-CreationDate-min"), "] AND maxCreationDate = [", r.FormValue("filterBy-CreationDate-max"), "]")

		// Check if the user sent an integer
		minCreationDate, err := strconv.Atoi(r.FormValue("filterBy-CreationDate-min"))
		if err != nil {
			http.Error(w, "Error, the value of 'minCreationDate' that you sent is not an integer", http.StatusBadRequest)
			return
		}
		maxCreationDate, err := strconv.Atoi(r.FormValue("filterBy-CreationDate-max"))
		if err != nil {
			http.Error(w, "Error, the value of 'maxCreationDate' that you sent is not an integer", http.StatusBadRequest)
			return
		}

		// Check if the user sent a valid range
		if minCreationDate > maxCreationDate {
			http.Error(w, "Error, the minimum value is higher than the maximum value", http.StatusBadRequest)
			return
		}

		// Apply the 'CreationDate' filter
		HomePageData.Groups = Api.FilterGroupsByCreationDate(minCreationDate, maxCreationDate, HomePageData.Groups)

		// --------------------------------------------------------------------------------------------------------- //

		// Checking the 'QtyOfMembers' filter
		fmt.Println("FILTER LOG : minQtyOfMembers = [", r.FormValue("filterBy-NumberOfMembers-min"), "] AND maxQtyOfMembers = [", r.FormValue("filterBy-NumberOfMembers-max"), "]")

		// Check if the user sent an integer
		minQtyOfMembers, err := strconv.Atoi(r.FormValue("filterBy-NumberOfMembers-min"))
		if err != nil {
			http.Error(w, "Error, the value of 'minQtyOfMembers' that you sent is not an integer", http.StatusBadRequest)
			return
		}
		maxQtyOfMembers, err := strconv.Atoi(r.FormValue("filterBy-NumberOfMembers-max"))
		if err != nil {
			http.Error(w, "Error, the value of 'maxQtyOfMembers' that you sent is not an integer", http.StatusBadRequest)
			return
		}

		// Check if the user sent a valid range
		if minQtyOfMembers > maxQtyOfMembers {
			http.Error(w, "Error, the minimum value is higher than the maximum value", http.StatusBadRequest)
			return
		}

		// Apply the filter
		HomePageData.Groups = Api.FilterGroupsByQtyOfMembers(minQtyOfMembers, maxQtyOfMembers, HomePageData.Groups)

		// Run the template with the filtered data
		err = tmpl.ExecuteTemplate(w, "HomePage", HomePageData)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error while loading the template : %v", err), http.StatusInternalServerError)
		}

		// --------------------------------------------------------------------------------------------------------- //

		// Checking the 'Country' filter
		fmt.Println("FILTER LOG : Country = [", r.FormValue("filterBy-Country"), "]")

	}
}
