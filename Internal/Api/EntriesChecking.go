package Api

import (
	"fmt"
	"net/http"
	"strconv"
)

func IsValidRange(min int, max int, w http.ResponseWriter) bool {
	// Checking if the min isn't higher than the max
	if min > max {
		fmt.Println("ENTRIES CHECK LOG: Error, the minimum value is higher than the maximum value")
		http.Error(w, "Error, the minimum value is higher than the maximum value", http.StatusBadRequest)
		return false
	}
	return true
}

func IsAnInteger(variableName string, value string, w http.ResponseWriter) (int, bool) {
	// Checking if the value is an integer
	// if Atoi returns an error, it means that the value isn't an integer
	intValue, err := strconv.Atoi(value)
	if err != nil {
		errorMsg := "Error, the value of '" + variableName + "' that you sent is not an integer"
		fmt.Println("ENTRIES CHECK LOG:", errorMsg)
		http.Error(w, errorMsg, http.StatusBadRequest)
		return 0, false
	}
	return intValue, true
}

func IsAString(variableName string, value string, w http.ResponseWriter) bool {
	// Checking if the value is a 'valid' string
	// a valid string is a string that contains only letters, spaces and '-' characters
	for _, char := range value {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char == ' ') || (char == '-') {
			continue
		} else {
			errorMsg := "Error, the value of '" + variableName + "' that you sent contains an invalid character"
			fmt.Println("ENTRIES CHECK LOG:", errorMsg)
			http.Error(w, errorMsg, http.StatusBadRequest)
			return false
		}
	}
	return true
}

func IsAIntList(variableName string, qtyList []string, w http.ResponseWriter) ([]int, bool) {
	// Checking if the list contains only integers
	qtyOfMembers := []int{}
	isListValid := false
	for _, value := range qtyList {
		number, isValid := IsAnInteger(variableName, value, w)
		if isValid != true {
			isListValid = false
			break
		} else {
			isListValid = true
			qtyOfMembers = append(qtyOfMembers, number)
		}
	}
	return qtyOfMembers, isListValid
}
