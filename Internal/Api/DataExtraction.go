package Api

import (
	"encoding/json"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"io/ioutil"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

func GetAllGroups() []GroupInfos {
	// Getting all the groups present in the API

	// Creation of the struct that will contain all the groups
	var groups []GroupInfos

	// Sending a GET request to the API
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Print("LOG: Error while sending the GET request (", err, ")")
		return groups
	}

	// Reading the response
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print("LOG: Error while reading the response (", err, ")")
		return groups
	}

	// We need to Unmarshal the data and put it in a Struct to be able to use it
	var unmarshalledData []GroupInfos
	err = json.Unmarshal(responseData, &unmarshalledData)
	if err != nil {
		fmt.Print("LOG: Error while unmarshalling the data (", err, ")")
		return groups
	}

	// Now we need to parcours all the groups to get the relations and more informations
	// (countries list, qty of members, ...)
	for index, group := range unmarshalledData {
		group.Relations = getGroupRelations(fmt.Sprint(group.Id))
		group.Relations = setCountriesList(group.Relations)
		qtyOfMembers := len(group.Members)
		group.QtyOfMembers = qtyOfMembers
		unmarshalledData[index] = group
	}

	groups = unmarshalledData
	return groups
}

func GetGroupInfos(id string) GroupInfos {
	// Sending a GET request to the API
	var chosenGroupInfos GroupInfos
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)
	if err != nil {
		fmt.Print("LOG: Error while sending the GET request (", err, ")")
		return chosenGroupInfos
	}

	// Reading the response
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print("LOG: Error while reading the response (", err, ")")
		return chosenGroupInfos
	}

	// Now we need to Unmarshal the data and put it in a Struct to be able to use it
	var unmarshalledData GroupInfos
	err = json.Unmarshal(responseData, &unmarshalledData)
	if err != nil {
		fmt.Print("LOG: Error while unmarshalling the data (", err, ")")
		return chosenGroupInfos
	}

	// Checking if the group is present in the API
	if unmarshalledData.Name == "" {
		fmt.Println("LOG: Group not found")
		return chosenGroupInfos
	}

	chosenGroupInfos = unmarshalledData
	// Getting the relations of the group and the other informations
	chosenGroupInfos.QtyOfMembers = len(chosenGroupInfos.Members)
	chosenGroupInfos.Relations = getGroupRelations(id)
	chosenGroupInfos.Relations.Coordinates = getConcertsCoordinates(chosenGroupInfos.Relations.DatesLocations)
	chosenGroupInfos.Relations = setCountriesList(chosenGroupInfos.Relations)
	return chosenGroupInfos
}

func getGroupRelations(id string) Relation {
	// Sending a GET request to the API
	var relations Relation
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + id)
	if err != nil {
		fmt.Print("LOG: Error while sending the GET request (", err, ")")
		return relations
	}

	// Reading the response
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print("LOG: Error while reading the response (", err, ")")
		return relations
	}

	// Now we need to Unmarshal the data and put it in a Struct to be able to use it
	var unmarshalledData Relation
	err = json.Unmarshal(responseData, &unmarshalledData)
	if err != nil {
		fmt.Print("LOG: Error while unmarshalling the data (", err, ")")
		return relations
	}
	// Checking if the group is present in the API
	if unmarshalledData.Id == 0 {
		fmt.Println("LOG: Group not found")
		return relations
	}
	relations = unmarshalledData

	return relations
}

func getConcertsCoordinates(datesLocations map[string][]string) map[string][]float64 {
	// Now we need to get all coordinates of all concerts locations
	coordinates := make(map[string][]float64)
	for key, _ := range datesLocations {

		// Creation of all structures needed to get the coordinates from the Geocoding API
		// according to the geocoding api response's structure
		type geoResponse struct {
			Features []struct {
				Geometry struct {
					Coordinates []float64 `json:"coordinates"`
				} `json:"geometry"`
			} `json:"features"`
		}

		// Sending a GET request to the Geocoding API
		apiResponse, err := http.Get("https://api.mapbox.com/search/geocode/v6/forward?q=" + key + "&access_token=pk.eyJ1IjoicmFtaWFyZDEyIiwiYSI6ImNtNXBibmE1cTA4bWcybXNpaWg4cWgydDgifQ.eZ661gjAiBWtfYMxXvN9Hw")
		if err != nil {
			fmt.Print("LOG Geocoding: Error while sending the GET request to the Geocoding API (", err, ")")
			return coordinates
		}

		// Reading the response
		apiResponseData, err := ioutil.ReadAll(apiResponse.Body)
		if err != nil {
			fmt.Println("LOG Geocoding: Error while reading the response from the Geocoding API (", err, ")")
			return coordinates
		}

		// Now we need to Unmarshal the data and put it into our previous 'geoResponse' struct to be able to use it
		var unmarshalledGeoData geoResponse
		err = json.Unmarshal(apiResponseData, &unmarshalledGeoData)
		if err != nil {
			fmt.Println("LOG Geocoding: Error while unmarshalling the data from the Geocoding API (", err, ")")
			return coordinates
		}
		// Add to the map the coordinates of the locations
		coordinates[key] = unmarshalledGeoData.Features[0].Geometry.Coordinates
	}
	return coordinates
}

func setCountriesList(relations Relation) Relation {
	// Let's stylise the location (Capitalizing the first letter of each word and deleting '_')
	// ("north_carolina-usa" become "North Carolina-Usa")
	groupRelation := relations
	for key, value := range groupRelation.DatesLocations {
		keyStylized := key
		keyStylized = strings.ReplaceAll(keyStylized, "_", " ")
		keyStylized = cases.Title(language.English).String(keyStylized)
		// Delete the old key and add the new one (because in go we cant change the key of a map)
		delete(groupRelation.DatesLocations, key)
		groupRelation.DatesLocations[keyStylized] = value
	}
	// Make the countries list
	var countriesList []string
	for key, _ := range groupRelation.DatesLocations {
		_, after, _ := strings.Cut(key, "-")
		// If 'countriesList' already contains the country, we don't add it
		if slices.Contains(countriesList, after) {
			continue
		}
		countriesList = append(countriesList, after)
	}
	groupRelation.CountriesList = countriesList
	return groupRelation
}

// GetAllCountries is a function that returns a list of all the countries present in the API
func GetAllCountries(groupList []GroupInfos) []string {

	var countriesList []string
	for _, group := range groupList {
		for _, country := range group.Relations.CountriesList {
			// If 'countriesList' already contains the country, we don't add it
			if slices.Contains(countriesList, country) {
				continue
			}
			countriesList = append(countriesList, country)
		}
	}
	slices.Sort(countriesList)
	return countriesList
}

// GetFiltersMinAndMax is a function that returns the min and max values of the API data
func GetFiltersMinAndMax(groupList []GroupInfos, data *Data) {
	// Getting the 'min' creation date

	// We need to set the 'min' creation date to a very high value to be able to compare it with the real values
	minCreationDate := 9999999
	for _, group := range groupList {
		if group.CreationDate < minCreationDate {
			minCreationDate = group.CreationDate
		}
	}
	// If the 'min' creation date is still the same, it means that there is no group in the API
	// so we set it to 0 to avoid to display '9999999' in the template
	if minCreationDate == 9999999 {
		minCreationDate = 0
	}
	data.MinCreationDate = minCreationDate

	// Getting the 'max' creation date
	maxCreationDate := 0
	for _, group := range groupList {
		if group.CreationDate > maxCreationDate {
			maxCreationDate = group.CreationDate
		}
	}
	data.MaxCreationDate = maxCreationDate

	// Getting the 'max' qty of members
	maxQtyOfMembers := 0
	for _, group := range groupList {
		if group.QtyOfMembers > maxQtyOfMembers {
			maxQtyOfMembers = group.QtyOfMembers
		}
	}

	// Making a list containing all the possible qty of members to use it in our go template
	qtyOfMemberList := []int{}
	for i := 0; i < maxQtyOfMembers; i++ {
		qtyOfMemberList = append(qtyOfMemberList, i+1)
	}
	if len(qtyOfMemberList) == 0 {
		qtyOfMemberList = append(qtyOfMemberList, 0)
	}
	data.QtyOfMemberList = qtyOfMemberList

	// Getting the 'min' first album date
	// We need to set the 'min' first album date to a very high value to be able to compare it with the real values
	minFirstAlbum := 9999999
	for _, group := range groupList {
		firstAlbum, _ := strconv.Atoi(group.FirstAlbum[len(group.FirstAlbum)-4:])
		if firstAlbum < minFirstAlbum {
			minFirstAlbum = firstAlbum
		}
	}
	// If the 'min' creation date is still the same, it means that there is no group in the API
	// so we set it to 0 to avoid to display '9999999' in the template
	if minFirstAlbum == 9999999 {
		minFirstAlbum = 0
	}
	data.MinFirstAlbum = minFirstAlbum

	// Getting the 'max' first album date
	maxFirstAlbum := 0
	for _, group := range groupList {
		firstAlbum, _ := strconv.Atoi(group.FirstAlbum[len(group.FirstAlbum)-4:])
		if firstAlbum > maxFirstAlbum {
			maxFirstAlbum = firstAlbum
		}
	}
	data.MaxFirstAlbum = maxFirstAlbum
}
