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
