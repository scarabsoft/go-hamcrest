package internal

import (
	"fmt"
	"reflect"
	"time"
)

// IsLessThan only supports numeric, string and boolean values
func IsLessThan(actual, expected interface{}) bool {

	if reflect.TypeOf(actual) != reflect.TypeOf(expected) {
		panic(FormatComparingIncompatibleTypes(actual, expected))
	}

	actualValue := reflect.ValueOf(actual)
	expectedValue := reflect.ValueOf(expected)

	switch expected.(type) {
	case int, int8, int16, int32, int64:
		return actualValue.Int() < expectedValue.Int()
	case uint, uint8, uint16, uint32, uint64:
		return actualValue.Uint() < expectedValue.Uint()
	case float32, float64:
		return actualValue.Float() < expectedValue.Float()
	case string:
		return actualValue.String() < expectedValue.String()
	case bool:
		return convertBoolToInt(actualValue.Bool()) < convertBoolToInt(expectedValue.Bool())
	case time.Duration:
		return actual.(time.Duration).Microseconds() < expected.(time.Duration).Microseconds()
	default:
		fmt.Println("CALLED")
	}

	return false
}

// IsLessThanEqual only supports numeric, string and boolean values
func IsLessThanEqual(actual, expected interface{}) bool {
	return IsEqual(actual, expected) || IsLessThan(actual, expected)
}
