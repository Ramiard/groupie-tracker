package Api

import (
	"net/http"
	"strconv"
)

func IsValidRange(min int, max int, w http.ResponseWriter) bool {
	if min > max {
		http.Error(w, "Error, the minimum value is higher than the maximum value", http.StatusBadRequest)
		return false
	}
	return true
}

func IsAnInteger(variableName string, value string, w http.ResponseWriter) (int, bool) {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		errorMsg := "Error, the value of '" + variableName + "' that you sent is not an integer"
		http.Error(w, errorMsg, http.StatusBadRequest)
		return 0, false
	}
	return intValue, true
}

func IsAString(variableName string, value string, w http.ResponseWriter) bool {
	for _, char := range value {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char == ' ') || (char == '-') {
			continue
		} else {
			errorMsg := "Error, the value of '" + variableName + "' that you sent contains an invalid character"
			http.Error(w, errorMsg, http.StatusBadRequest)
			return false
		}
	}
	return true
}

func IsAIntList(variableName string, qtyList []string, w http.ResponseWriter) ([]int, bool) {
	qtyOfMembers := []int{}
	isListValid := false
	for _, value := range qtyList {
		number, isValid := IsAnInteger("QtyOfMembers", value, w)
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
