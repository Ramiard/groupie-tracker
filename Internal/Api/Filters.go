package Api

import (
	"strconv"
)

func FilterGroupsByCreationDate(creationDate string, groups []GroupInfos) []GroupInfos {
	var filteredGroups []GroupInfos

	for _, group := range groups {
		if strconv.Itoa(group.CreationDate) == creationDate {
			filteredGroups = append(filteredGroups, group)
		}
	}
	return filteredGroups
}
