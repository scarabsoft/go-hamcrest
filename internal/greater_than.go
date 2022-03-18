package internal

import (
	"reflect"
	"time"
)

// IsGreaterThan only supports numeric, string and boolean values
func IsGreaterThan(actual, expected interface{}) bool {

	if reflect.TypeOf(actual) != reflect.TypeOf(expected) {
		panic(FormatComparingIncompatibleTypes(actual, expected))
	}

	actualValue := reflect.ValueOf(actual)
	expectedValue := reflect.ValueOf(expected)

	switch expected.(type) {
	case int, int8, int16, int32, int64:
		return actualValue.Int() > expectedValue.Int()
	case uint, uint8, uint16, uint32, uint64:
		return actualValue.Uint() > expectedValue.Uint()
	case float32, float64:
		return actualValue.Float() > expectedValue.Float()
	case string:
		return actualValue.String() > expectedValue.String()
	case bool:
		return convertBoolToInt(actualValue.Bool()) > convertBoolToInt(expectedValue.Bool())
	case time.Duration:
		return actual.(time.Duration).Microseconds() > expected.(time.Duration).Microseconds()
	}

	return false
}

// IsGreaterThanEqual only supports numeric, string and boolean values
func IsGreaterThanEqual(actual, expected interface{}) bool {
	return IsEqual(actual, expected) || IsGreaterThan(actual, expected)
}
