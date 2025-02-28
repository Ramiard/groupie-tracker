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

	// Prevent the webbrowser from loading the favicon, and sending 2 request to the website
	if r.URL.Path == "/favicon.ico" {
		return
	}

	var homePageData Api.Data
	// Get the groups and countries
	homePageData.AllGroups = Api.GetAllGroups()
	homePageData.Groups = homePageData.AllGroups
	homePageData.Countries = Api.GetAllCountries(homePageData.Groups)
	// Get the filters 'min' and 'max' values
	Api.GetFiltersMinAndMax(homePageData.Groups, &homePageData)

	// Check if the user send a POST request containing a filter
	if r.Method != http.MethodPost {
		// Run the template without filtering
		err := tmpl.ExecuteTemplate(w, "HomePage", homePageData)
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

		homePageData.Groups = Api.ApplyFilters(filters, homePageData.Groups, w)
		Api.GetFiltersMinAndMax(homePageData.Groups, &homePageData)

		//// --------------------------------------------------------------------------------------------------------- //

		// Run the template with the filtered data
		err = tmpl.ExecuteTemplate(w, "HomePage", homePageData)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error while loading the template : %v", err), http.StatusInternalServerError)
		}
	}
}
