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
