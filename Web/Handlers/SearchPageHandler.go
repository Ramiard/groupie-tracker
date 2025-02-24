package Handlers

import (
	"encoding/base64"
	"encoding/json"
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

	var SearchResultsData Api.Data
	SearchResultsData.Groups = Api.GetAllGroups()

	var searchQuery string
	// If the user send a POST request with a search we make the search and stock the results in a cookie
	if r.FormValue("search") != "" {
		searchQuery = r.FormValue("search")
		fmt.Println("SEARCH PAGE LOG: Search query : '", searchQuery, "'")

		// Let's make the search results list
		SearchResultsData.SearchResults = Api.SearchGroups(searchQuery, SearchResultsData.Groups)
		// Update the filters 'min' and 'max' values according to the search results and the countries
		SearchResultsData.Countries = Api.GetAllCountries(SearchResultsData.SearchResults)
		Api.GetFiltersMinAndMax(SearchResultsData.SearchResults, &SearchResultsData)

		// Setting up a cookie to store the search query
		searchQueryJson, err := json.Marshal(searchQuery)
		if err != nil {
			http.Error(w, "Error while marshalling the search query", http.StatusInternalServerError)
			return
		}
		encodedQuery := base64.StdEncoding.EncodeToString(searchQueryJson)
		http.SetCookie(w, &http.Cookie{
			Name:   "searchQuery",
			Value:  encodedQuery,
			Path:   "/search",
			MaxAge: 60, // Expire after 1 minute
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

		// Unmarshal the cookie
		var previousSearchQuery string
		err = json.Unmarshal(decodedSearchCookie, &previousSearchQuery)
		if err != nil {
			http.Error(w, "Error while unmarshalling the search results cookie", http.StatusInternalServerError)
			return
		}
		SearchResultsData.SearchResults = Api.SearchGroups(previousSearchQuery, SearchResultsData.Groups)
		// Setting up the filters 'min' and 'max' values according to the search results and the countries
		SearchResultsData.Countries = Api.GetAllCountries(SearchResultsData.SearchResults)
		Api.GetFiltersMinAndMax(SearchResultsData.SearchResults, &SearchResultsData)

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
		SearchResultsData.SearchResults = Api.ApplyFilters(filters, SearchResultsData.SearchResults, w)

	}

	// Run the template
	err := tmpl.ExecuteTemplate(w, "SearchResults", SearchResultsData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while loading the template : %v", err), http.StatusInternalServerError)
	}
}
