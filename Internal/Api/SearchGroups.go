package Api

import (
	"fmt"
	"strconv"
	"strings"
)

func SearchGroups(searchQuery string, groups []GroupInfos) []GroupInfos {
	var searchResults []GroupInfos
	for _, group := range groups {
		// Check if the search query is in the group name
		if strings.Contains(strings.ToLower(group.Name), strings.ToLower(searchQuery)) {
			fmt.Println("SEARCH LOG: Search query found in the", group.Name, "name")
			searchResults = append(searchResults, group)
			continue
		}
		// Check if the search query is in the group members
		isInMember := false
		for _, member := range group.Members {
			if strings.Contains(strings.ToLower(member), strings.ToLower(searchQuery)) {
				fmt.Println("SEARCH LOG: Search query found in the ", group.Name, "'s members")
				searchResults = append(searchResults, group)
				isInMember = true
				break
			}
		}
		// If the search query is found in the members, we skip the rest of the checks
		// if not, we check the others fields
		if isInMember == true {
			continue
		}
		// Check if the search query is in the group number of members
		if strings.Contains(strings.ToLower(strconv.Itoa(group.QtyOfMembers)), strings.ToLower(searchQuery)) {
			fmt.Println("SEARCH LOG: Search query found in the", group.Name, "number of members")
			searchResults = append(searchResults, group)
			continue
		}
		// Check if the search query is the group Creation Date
		if strings.Contains(strings.ToLower(strconv.Itoa(group.CreationDate)), strings.ToLower(searchQuery)) {
			fmt.Println("SEARCH LOG: Search query found in the ", group.Name, " creation date")
			searchResults = append(searchResults, group)
			continue
		}
		// Check if the search query contains the group first Album Date
		if strings.Contains(strings.ToLower(group.FirstAlbum), strings.ToLower(searchQuery)) {
			fmt.Println("SEARCH LOG: Search query found in the", group.Name, "first album date")
			searchResults = append(searchResults, group)
			continue
		}
		// Check if the search query contains the group relations, locations or dates
		isInRelation := false
		for key, value := range group.Relations.DatesLocations {
			if strings.Contains(strings.ToLower(key), strings.ToLower(searchQuery)) {
				fmt.Println("SEARCH LOG: Search query found in the", group.Name, "locations")
				searchResults = append(searchResults, group)
				isInRelation = true
				break
			}
			for _, dates := range value {
				if strings.Contains(strings.ToLower(dates), strings.ToLower(searchQuery)) {
					fmt.Println("SEARCH LOG: Search query found in the", group.Name, "dates")
					searchResults = append(searchResults, group)
					isInRelation = true
					break
				}
			}
		}
		if isInRelation == true {
			continue
		}
	}
	fmt.Println("SEARCH LOG: Number of results : ", len(searchResults))
	return searchResults
}
