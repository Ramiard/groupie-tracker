package Handlers

import (
	"encoding/base64"
	"fmt"
	"groupie-tracker/Internal/Api"
	"html/template"
	"net/http"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Load the template
	tmpl := template.Must(template.ParseFiles("Web/Templates/SearchResults.gohtml"))

	// Prevent the webbrowser from loading the favicon, and sending 2 request to the website
	if r.URL.Path == "/favicon.ico" {
		return
	}

	// Check if the user send a POST request containing a search
	if r.Method != http.MethodPost {
		fmt.Println("SEARCH PAGE LOG: No POST used, redirecting to /HomePage")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var searchResultsData Api.Data
	// Get all the groups before making a search
	searchResultsData.AllGroups = Api.GetAllGroups()
	searchResultsData.Groups = searchResultsData.AllGroups

	var searchQuery string
	// If the user send a POST request with a search we make the search and stock the results in a cookie
	if r.FormValue("search") != "" {
		searchQuery = r.FormValue("search")
		fmt.Println("SEARCH PAGE LOG: Search query : '", searchQuery, "'")

		// Let's make the search results list
		searchResultsData.Groups = Api.SearchGroups(searchQuery, searchResultsData.Groups)
		// Update the filters 'min' and 'max' values according to the search results and the countries
		searchResultsData.Countries = Api.GetAllCountries(searchResultsData.Groups)
		Api.GetFiltersMinAndMax(searchResultsData.Groups, &searchResultsData)

		// Setting up a cookie to store the search query
		// Encode it in base64
		encodedQuery := base64.StdEncoding.EncodeToString([]byte(searchQuery))
		// Set the cookie
		http.SetCookie(w, &http.Cookie{
			Name:  "searchQuery",
			Value: encodedQuery,
			Path:  "/search",
			// Make it expire after 1 minute
			MaxAge: 60,
		})
		fmt.Println("SEARCH PAGE LOG: Search query cookie set")
	}

	// If there Is a POST request without a search, we check if there is a cookie containing the previous search results
	if r.FormValue("search") == "" {

		// Get the cookie
		searchCookie, err := r.Cookie("searchQuery")
		if err != nil {
			fmt.Println("SEARCH PAGE LOG: No search cookie found and no search, redirecting to /HomePage")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		fmt.Println("SEARCH PAGE LOG: Search cookie found")

		// Decode the cookie
		decodedSearchCookie, err := base64.StdEncoding.DecodeString(searchCookie.Value)
		if err != nil {
			http.Error(w, "Error while decoding the search query cookie", http.StatusInternalServerError)
			return
		}

		// Read the cookie and apply the search
		var previousSearchQuery = string(decodedSearchCookie)
		searchResultsData.Groups = Api.SearchGroups(previousSearchQuery, searchResultsData.Groups)
		// Setting up the filters 'min' and 'max' values according to the search results and the countries
		searchResultsData.Countries = Api.GetAllCountries(searchResultsData.Groups)
		Api.GetFiltersMinAndMax(searchResultsData.Groups, &searchResultsData)

		// Check if there is any filter applied
		// Get the form value
		err = r.ParseForm()
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
		searchResultsData.Groups = Api.ApplyFilters(filters, searchResultsData.AllGroups, w)

	}

	// Run the template
	err := tmpl.ExecuteTemplate(w, "SearchResults", searchResultsData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while loading the template : %v", err), http.StatusInternalServerError)
	}
}
