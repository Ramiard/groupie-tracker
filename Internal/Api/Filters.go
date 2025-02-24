package Api

import (
	"fmt"
	"net/http"
	"strconv"
)

func ApplyFilters(filters Filters, groupList []GroupInfos, w http.ResponseWriter) []GroupInfos {
	var filteredGroups []GroupInfos
	// Apply the 'CreationDate' filter
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
	if filters.IsQtyOfMembersFilter == true {
		fmt.Println("FILTER LOG : QtyOfMembers = [", filters.QtyOfMembersList, "]")
		// Check if the user sent an integer list
		qtyOfMembers, isValid := IsAIntList("QtyOfMembers", filters.QtyOfMembersList, w)

		if isValid == true {
			// Apply the filter
			filteredGroups = filterGroupsByQtyOfMembers(qtyOfMembers, filteredGroups)
		}
	}

	// Apply the 'FirstAlbumDate' filter
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

	// Apply the 'Country' filter
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
		groupFirstAlbum, _ := strconv.Atoi(group.FirstAlbum[len(group.FirstAlbum)-4:])
		if groupFirstAlbum >= minFirstAlbumDate && groupFirstAlbum <= maxFirstAlbumDate {
			filteredGroups = append(filteredGroups, group)
		}
	}
	return filteredGroups
}
