package Api

import (
	"strconv"
)

func FilterGroupsByCreationDate(minCreationDate int, maxCreationDate int, groups []GroupInfos) []GroupInfos {
	var filteredGroups []GroupInfos

	for _, group := range groups {
		if group.CreationDate >= minCreationDate && group.CreationDate <= maxCreationDate {
			filteredGroups = append(filteredGroups, group)
		}
	}
	return filteredGroups
}

func FilterGroupsByQtyOfMembers(qtyList []int, groups []GroupInfos) []GroupInfos {
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

func FilterGroupsByCountry(countryToFilter string, groupList []GroupInfos) []GroupInfos {
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

func FilterGroupsByFirstAlbumDate(minFirstAlbumDate int, maxFirstAlbumDate int, groups []GroupInfos) []GroupInfos {
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
