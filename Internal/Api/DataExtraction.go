package Api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAllGroups() []GroupInfos {
	// Getting all the groups present in the API
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
	var UnmarshalledData []GroupInfos
	err = json.Unmarshal(responseData, &UnmarshalledData)
	if err != nil {
		fmt.Print("LOG: Error while unmarshalling the data (", err, ")")
		return groups
	}

	for index, group := range UnmarshalledData {
		group.Relations = GetGroupRelations(fmt.Sprint(group.Id))
		qtyOfMembers := len(group.Members)
		group.QtyOfMembers = qtyOfMembers
		UnmarshalledData[index] = group
	}

	groups = UnmarshalledData
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
	var UnmarshalledData GroupInfos
	err = json.Unmarshal(responseData, &UnmarshalledData)
	if err != nil {
		fmt.Print("LOG: Error while unmarshalling the data (", err, ")")
		return chosenGroupInfos
	}

	// Checking if the group is present in the API
	if UnmarshalledData.Name == "" {
		fmt.Println("LOG: Group not found")
		return chosenGroupInfos
	}

	chosenGroupInfos = UnmarshalledData
	chosenGroupInfos.QtyOfMembers = len(chosenGroupInfos.Members)
	chosenGroupInfos.Relations = GetGroupRelations(id)
	chosenGroupInfos.Relations.Coordinates = getConcertsCoordinates(chosenGroupInfos.Relations.DatesLocations)
	return chosenGroupInfos
}

func GetGroupRelations(id string) Relation {
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
	var UnmarshalledData Relation
	err = json.Unmarshal(responseData, &UnmarshalledData)
	if err != nil {
		fmt.Print("LOG: Error while unmarshalling the data (", err, ")")
		return relations
	}

	// Checking if the group is present in the API
	if UnmarshalledData.Id == 0 {
		fmt.Println("LOG: Group not found")
		return relations
	}
	relations = UnmarshalledData

	return relations
}

func getConcertsCoordinates(datesLocations map[string][]string) map[string][]float64 {
	// Now we need to get all coordinates of all concerts locations
	coordinates := make(map[string][]float64)
	for key, _ := range datesLocations {

		// Creation of all structures needed to get the coordinates from the Geocoding API
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
		var UnmarshalledGeoData geoResponse
		err = json.Unmarshal(apiResponseData, &UnmarshalledGeoData)
		if err != nil {
			fmt.Println("LOG Geocoding: Error while unmarshalling the data from the Geocoding API (", err, ")")
			return coordinates
		}
		// Add to the map the coordinates of the locations
		coordinates[key] = UnmarshalledGeoData.Features[0].Geometry.Coordinates
	}
	return coordinates
}
