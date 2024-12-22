package Api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAllGroups() GroupList {
	// Getting all the groups present in the API
	var groups GroupList
	var groupsLeft bool = true
	var id int = 1

	for groupsLeft == true {
		// Sending a GET request to the API
		response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + fmt.Sprint(id))
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
		var UnmarshalledData GroupInfos
		err = json.Unmarshal(responseData, &UnmarshalledData)
		if err != nil {
			fmt.Print("LOG: Error while unmarshalling the data (", err, ")")
			return groups
		}

		// Checking if we got all the groups present in the API
		if UnmarshalledData.Name == "" {
			groupsLeft = false
			fmt.Println("LOG: All groups have been loaded")
			fmt.Println("LOG: Last group loaded :", id-1)
			return groups
		}
		var tempGroupInfos GroupInfos
		tempGroupInfos = UnmarshalledData
		tempGroupInfos.Relations = GetGroupRelations(fmt.Sprint(id))
		groups.List = append(groups.List, tempGroupInfos)
		id++
	}
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
	chosenGroupInfos.Relations = GetGroupRelations(id)
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
