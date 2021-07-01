package internal

import "reflect"

func IsEqual(actual, given interface{}) bool {
	if IsNil(actual) && IsNil(given) {
		return true
	}

	if IsNil(actual) || IsNil(given) {
		return false
	}

	if reflect.DeepEqual(actual, given) {
		return true
	}

	aValue := reflect.ValueOf(actual)
	bValue := reflect.ValueOf(given)
	return aValue == bValue
}
