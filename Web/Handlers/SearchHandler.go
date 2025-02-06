package Handlers

import (
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
	if r.Method != http.MethodPost || r.FormValue("search") == "" {
		fmt.Println("SEARCH LOG: No POST used or empty search, redirecting to /HomePage")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var SearchResultsData Api.Data
	SearchResultsData.Groups = Api.GetAllGroups()
	SearchResultsData.Countries = Api.GetAllCountries(SearchResultsData.Groups)
	Api.GetFiltersMinAndMax(SearchResultsData.Groups, &SearchResultsData)

	searchQuery := r.FormValue("search")
	fmt.Println("SEARCH LOG: Search query : '", searchQuery, "'")

	// Let's make the search results list
	SearchResultsData.SearchResults = Api.SearchGroups(searchQuery, SearchResultsData.Groups)

	// Run the template
	err := tmpl.ExecuteTemplate(w, "SearchResults", SearchResultsData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error while loading the template : %v", err), http.StatusInternalServerError)
	}
}
