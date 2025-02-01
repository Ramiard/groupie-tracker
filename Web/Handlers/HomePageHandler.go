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

		if r.FormValue("filterBy-CreationDate-min") != "" && r.FormValue("filterBy-CreationDate-max") != "" {

			fmt.Println("FILTER LOG : minCreationDate = [", r.FormValue("filterBy-CreationDate-min"), "] AND maxCreationDate = [", r.FormValue("filterBy-CreationDate-max"), "]")

			// Check if the user sent an integer
			minCreationDate, isValid1 := Api.IsAnInteger("minCreationDate", r.FormValue("filterBy-CreationDate-min"), w)

			maxCreationDate, isValid2 := Api.IsAnInteger("maxCreationDate", r.FormValue("filterBy-CreationDate-max"), w)

			// Check if the user sent a valid range and apply the filter if all the entries are valid
			if isValid1 && isValid2 && Api.IsValidRange(minCreationDate, maxCreationDate, w) == true {

				// Apply the 'CreationDate' filter
				HomePageData.Groups = Api.FilterGroupsByCreationDate(minCreationDate, maxCreationDate, HomePageData.Groups)

			}
		}

		// --------------------------------------------------------------------------------------------------------- //

		// Checking the 'QtyOfMembers' filter

		if r.FormValue("filterBy-NumberOfMembers-min") != "" && r.FormValue("filterBy-NumberOfMembers-max") != "" {

			fmt.Println("FILTER LOG : minQtyOfMembers = [", r.FormValue("filterBy-NumberOfMembers-min"), "] AND maxQtyOfMembers = [", r.FormValue("filterBy-NumberOfMembers-max"), "]")

			// Check if the user sent an integer
			minQtyOfMembers, isValid1 := Api.IsAnInteger("minQtyOfMembers", r.FormValue("filterBy-NumberOfMembers-min"), w)

			maxQtyOfMembers, isValid2 := Api.IsAnInteger("maxQtyOfMembers", r.FormValue("filterBy-NumberOfMembers-max"), w)

			// Check if the user sent a valid range and apply the filter if all the entries are valid
			if isValid1 && isValid2 && Api.IsValidRange(minQtyOfMembers, maxQtyOfMembers, w) == true {
				// Apply the filter
				HomePageData.Groups = Api.FilterGroupsByQtyOfMembers(minQtyOfMembers, maxQtyOfMembers, HomePageData.Groups)
			}
		}

		// --------------------------------------------------------------------------------------------------------- //

		// Checking the 'Country' filter

		if r.FormValue("filterBy-Country") != "" {

			fmt.Println("FILTER LOG : Country = [", r.FormValue("filterBy-Country"), "]")

			// Check if the user sent a string
			if Api.IsAString("Country", r.FormValue("filterBy-Country"), w) == true {
				// Apply the filter
				HomePageData.Groups = Api.FilterGroupsByCountry(r.FormValue("filterBy-Country"), HomePageData.Groups)
			}

		}

		// --------------------------------------------------------------------------------------------------------- //

		// Run the template with the filtered data
		err = tmpl.ExecuteTemplate(w, "HomePage", HomePageData)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error while loading the template : %v", err), http.StatusInternalServerError)
		}
	}
}
