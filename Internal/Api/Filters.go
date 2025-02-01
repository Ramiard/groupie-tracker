package Api

func FilterGroupsByCreationDate(minCreationDate int, maxCreationDate int, groups []GroupInfos) []GroupInfos {
	var filteredGroups []GroupInfos

	for _, group := range groups {
		if group.CreationDate >= minCreationDate && group.CreationDate <= maxCreationDate {
			filteredGroups = append(filteredGroups, group)
		}
	}
	return filteredGroups
}

func FilterGroupsByQtyOfMembers(minQtyOfMembers int, maxQtyOfMembers int, groups []GroupInfos) []GroupInfos {
	var filteredGroups []GroupInfos

	for _, group := range groups {
		if group.QtyOfMembers >= minQtyOfMembers && group.QtyOfMembers <= maxQtyOfMembers {
			filteredGroups = append(filteredGroups, group)
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
