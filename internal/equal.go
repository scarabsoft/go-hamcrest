package internal

import "reflect"

func IsEqual(actual, expected interface{}) bool {
	if IsNil(actual) && IsNil(expected) {
		return true
	}

	if IsNil(actual) || IsNil(expected) {
		return false
	}

	if reflect.DeepEqual(actual, expected) {
		return true
	}

	aValue := reflect.ValueOf(actual)
	bValue := reflect.ValueOf(expected)
	return aValue == bValue
}
