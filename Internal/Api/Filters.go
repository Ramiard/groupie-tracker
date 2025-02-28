package Api

import (
	"fmt"
	"net/http"
	"strconv"
)

// ApplyFilters will apply the all the filters sent by the user to the groupList
// and if a filter isn't 'complete' (the user didn't send all the entries needed)
// the filter will not be applied
func ApplyFilters(filters Filters, groupList []GroupInfos, w http.ResponseWriter) []GroupInfos {
	var filteredGroups []GroupInfos
	// Check for the 'CreationDate' filter
	if filters.IsCreationDateFilter == true {
		fmt.Println("FILTER LOG : minCreationDate = [", filters.MinCreationDate, "] AND maxCreationDate = [", filters.MaxCreationDate, "]")
		// Check if the user sent an integer
		minCreationDate, isValid1 := IsAnInteger("minCreationDate", filters.MinCreationDate, w)

		maxCreationDate, isValid2 := IsAnInteger("maxCreationDate", filters.MaxCreationDate, w)

		// Check if the user sent a valid range and apply the filter if all the entries are valid
		if isValid1 && isValid2 && IsValidRange(minCreationDate, maxCreationDate, w) == true {

			// Apply the 'CreationDate' filter
			filteredGroups = filterGroupsByCreationDate(minCreationDate, maxCreationDate, groupList)
		}
	}
	// Check for the 'QtyOfMembers' filter
	if filters.IsQtyOfMembersFilter == true {
		fmt.Println("FILTER LOG : QtyOfMembers = [", filters.QtyOfMembersList, "]")
		// Check if the user sent an integer list
		qtyOfMembers, isValid := IsAIntList("QtyOfMembers", filters.QtyOfMembersList, w)

		if isValid == true {
			// Apply the filter
			filteredGroups = filterGroupsByQtyOfMembers(qtyOfMembers, filteredGroups)
		}
	}

	// Check the 'FirstAlbumDate' filter
	if filters.IsFirstAlbumDateFilter == true {
		fmt.Println("FILTER LOG : minFirstAlbumDate = [", filters.MinFirstAlbumDate, "] AND maxFirstAlbumDate = [", filters.MaxFirstAlbumDate, "]")
		// Check if the user sent an integer
		minFirstAlbumDate, isValid1 := IsAnInteger("minFirstAlbumDate", filters.MinFirstAlbumDate, w)
		maxFirstAlbumDate, isValid2 := IsAnInteger("maxFirstAlbumDate", filters.MaxFirstAlbumDate, w)

		// Check if the user sent a valid range and apply the filter if all the entries are valid
		if isValid1 && isValid2 && IsValidRange(minFirstAlbumDate, maxFirstAlbumDate, w) == true {
			filteredGroups = filterGroupsByFirstAlbumDate(minFirstAlbumDate, maxFirstAlbumDate, filteredGroups)
		}
	}

	// Check the 'Country' filter
	if filters.IsCountryFilter == true {
		fmt.Println("FILTER LOG : Country = [", filters.CountryToFilter, "]")
		// Check if the user sent a string
		if IsAString("Country", filters.CountryToFilter, w) == true {
			filteredGroups = filterGroupsByCountry(filters.CountryToFilter, filteredGroups)
		}
	}
	return filteredGroups
}

func filterGroupsByCreationDate(minCreationDate int, maxCreationDate int, groups []GroupInfos) []GroupInfos {
	var filteredGroups []GroupInfos

	for _, group := range groups {
		if group.CreationDate >= minCreationDate && group.CreationDate <= maxCreationDate {
			filteredGroups = append(filteredGroups, group)
		}
	}
	return filteredGroups
}

func filterGroupsByQtyOfMembers(qtyList []int, groups []GroupInfos) []GroupInfos {
	var filteredGroups []GroupInfos

	for _, group := range groups {
		for _, qty := range qtyList {
			if group.QtyOfMembers == qty {
				filteredGroups = append(filteredGroups, group)
			}
		}
	}
	return filteredGroups
}

func filterGroupsByCountry(countryToFilter string, groupList []GroupInfos) []GroupInfos {
	var filteredGroups []GroupInfos

	// Check if there is the default value of the 'Country' filter
	// i added 'tous' in case we want to upgrade the wep app to support french
	if countryToFilter == "All Countries" || countryToFilter == "Tous" {
		return groupList
	}

	for _, group := range groupList {
		for _, country := range group.Relations.CountriesList {
			if country == countryToFilter {
				filteredGroups = append(filteredGroups, group)
				break
			}
		}
	}
	return filteredGroups
}

func filterGroupsByFirstAlbumDate(minFirstAlbumDate int, maxFirstAlbumDate int, groups []GroupInfos) []GroupInfos {
	var filteredGroups []GroupInfos

	for _, group := range groups {
		// Extract the year from the 'FirstAlbum' string and pass it to an integer
		// the 'firstAlbumDate' is in the format 'DD-MM-YYYY' so we need to extract the last 4 characters to get the year
		groupFirstAlbum, _ := strconv.Atoi(group.FirstAlbum[len(group.FirstAlbum)-4:])
		if groupFirstAlbum >= minFirstAlbumDate && groupFirstAlbum <= maxFirstAlbumDate {
			filteredGroups = append(filteredGroups, group)
		}
	}
	return filteredGroups
}
