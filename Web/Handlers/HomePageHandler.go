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

		if r.FormValue("filterBy-creationDate-min") != "" && r.FormValue("filterBy-creationDate-max") != "" {

			fmt.Println("FILTER LOG : minCreationDate = [", r.FormValue("filterBy-creationDate-min"), "] AND maxCreationDate = [", r.FormValue("filterBy-creationDate-max"), "]")

			// Check if the user sent an integer
			minCreationDate, isValid1 := Api.IsAnInteger("minCreationDate", r.FormValue("filterBy-creationDate-min"), w)

			maxCreationDate, isValid2 := Api.IsAnInteger("maxCreationDate", r.FormValue("filterBy-creationDate-max"), w)

			// Check if the user sent a valid range and apply the filter if all the entries are valid
			if isValid1 && isValid2 && Api.IsValidRange(minCreationDate, maxCreationDate, w) == true {

				// Apply the 'CreationDate' filter
				HomePageData.Groups = Api.FilterGroupsByCreationDate(minCreationDate, maxCreationDate, HomePageData.Groups)

			}
		}

		// --------------------------------------------------------------------------------------------------------- //

		// Checking the 'QtyOfMembers' filter
		if r.FormValue("filterBy-membersNumber") != "" {

			fmt.Println("FILTER LOG : QtyOfMembers = [", r.Form["filterBy-membersNumber"], "]")
			// Check if the user sent an integer list
			qtyOfMembers, isValid := Api.IsAIntList("QtyOfMembers", r.Form["filterBy-membersNumber"], w)

			if isValid == true {
				// Apply the filter
				HomePageData.Groups = Api.FilterGroupsByQtyOfMembers(qtyOfMembers, HomePageData.Groups)
			}

		}

		// --------------------------------------------------------------------------------------------------------- //

		// Checking the 'Country' filter

		if r.FormValue("filterBy-country") != "" {

			fmt.Println("FILTER LOG : Country = [", r.FormValue("filterBy-country"), "]")

			// Check if the user sent a string
			if Api.IsAString("Country", r.FormValue("filterBy-country"), w) == true {
				// Apply the filter
				HomePageData.Groups = Api.FilterGroupsByCountry(r.FormValue("filterBy-country"), HomePageData.Groups)
			}

		}

		// --------------------------------------------------------------------------------------------------------- //

		// Checking the 'First Album Date' filter
		if r.FormValue("filterBy-firstAlbumDate-min") != "" && r.FormValue("filterBy-firstAlbumDate-max") != "" {

			fmt.Println("FILTER LOG : minFirstAlbumDate = [", r.FormValue("filterBy-firstAlbumDate-min"), "] AND maxFirstAlbumDate = [", r.FormValue("filterBy-firstAlbumDate-max"), "]")

			// Check if the user sent an integer
			minFirstAlbumDate, valid1 := Api.IsAnInteger("minFirstAlbumDate", r.FormValue("filterBy-firstAlbumDate-min"), w)
			maxFirstAlbumDate, valid2 := Api.IsAnInteger("maxFirstAlbumDate", r.FormValue("filterBy-firstAlbumDate-max"), w)

			if valid1 && valid2 && Api.IsValidRange(minFirstAlbumDate, maxFirstAlbumDate, w) == true {
				// Apply the filter
				HomePageData.Groups = Api.FilterGroupsByFirstAlbumDate(minFirstAlbumDate, maxFirstAlbumDate, HomePageData.Groups)
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
