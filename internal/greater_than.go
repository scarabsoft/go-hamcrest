package internal

import "reflect"

// IsGreaterThan only supports numeric, string and boolean values
func IsGreaterThan(actual, given interface{}) bool {

	if reflect.TypeOf(actual) != reflect.TypeOf(given) {
		panic(FormatComparingIncompatibleTypes(actual, given))
	}

	actualValue := reflect.ValueOf(actual)
	givenValue := reflect.ValueOf(given)

	switch given.(type) {
	case int, int8, int16, int32, int64:
		return actualValue.Int() > givenValue.Int()
	case uint, uint8, uint16, uint32, uint64:
		return actualValue.Uint() > givenValue.Uint()
	case float32, float64:
		return actualValue.Float() > givenValue.Float()
	case string:
		return actualValue.String() > givenValue.String()
	case bool:
		return convertBoolToInt(actualValue.Bool()) > convertBoolToInt(givenValue.Bool())
	}

	return false
}

// IsGreaterThanEqual only supports numeric, string and boolean values
func IsGreaterThanEqual(actual, given interface{}) bool {
	return IsEqual(actual, given) || IsGreaterThan(actual, given)
}
