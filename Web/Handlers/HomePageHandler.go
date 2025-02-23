package Handlers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/Internal/Api"
	"html/template"
	"net/http"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	// Load the template
	tmpl := template.Must(template.ParseFiles("Web/Templates/HomePage.gohtml"))

	// Prevent the webbrowser from loading the favicon, and sending 2 request to the website
	if r.URL.Path == "/favicon.ico" {
		return
	}

	var HomePageData Api.Data
	HomePageData.Groups = Api.GetAllGroups()
	HomePageData.Countries = Api.GetAllCountries(HomePageData.Groups)
	Api.GetFiltersMinAndMax(HomePageData.Groups, &HomePageData)

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

		// Setting up the filters
		var filters Api.Filters
		if r.FormValue("filterBy-creationDate-min") != "" && r.FormValue("filterBy-creationDate-max") != "" {
			filters.IsCreationDateFilter = true
			filters.MinCreationDate = r.FormValue("filterBy-creationDate-min")
			filters.MaxCreationDate = r.FormValue("filterBy-creationDate-max")
		}
		if r.FormValue("filterBy-membersNumber") != "" {
			filters.IsQtyOfMembersFilter = true
			filters.QtyOfMembersList = r.Form["filterBy-membersNumber"]
		}
		if r.FormValue("filterBy-firstAlbumDate-min") != "" && r.FormValue("filterBy-firstAlbumDate-max") != "" {
			filters.IsFirstAlbumDateFilter = true
			filters.MinFirstAlbumDate = r.FormValue("filterBy-firstAlbumDate-min")
			filters.MaxFirstAlbumDate = r.FormValue("filterBy-firstAlbumDate-max")
		}
		if r.FormValue("filterBy-country") != "" {
			filters.IsCountryFilter = true
			filters.CountryToFilter = r.FormValue("filterBy-country")
		}
		// Apply the filters
		HomePageData.Groups = Api.ApplyFilters(filters, HomePageData.Groups, w)

		//// --------------------------------------------------------------------------------------------------------- //

		// JSON the data so it can be sent to the javascript
		jsonData, err := json.Marshal(HomePageData)
		if err != nil {
			http.Error(w, "Error while marshalling the data", http.StatusInternalServerError)
			return
		}

		// Run the template with the filtered data
		err = tmpl.ExecuteTemplate(w, "HomePage", string(jsonData))
		if err != nil {
			http.Error(w, fmt.Sprintf("Error while loading the template : %v", err), http.StatusInternalServerError)
		}
	}
}
