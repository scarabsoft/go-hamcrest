package internal

import (
	"fmt"
)

func FormatTypeWithValue(v interface{}) string {
	if IsNil(v) {
		return "<nil>"
	}
	return fmt.Sprintf("%[1]T(%[1]v)", v)
}

func FormatTypes(op string, actual, expected interface{}) string {
	return fmt.Sprintf("%[1]T %s %[3]T", actual, op, expected)
}

func FormatComparingIncompatibleTypes(actual, expected interface{}) string {
	return "try to compare incompatible types: " + FormatTypes(",", actual, expected)
}

func FormatUnsupportedVariable(name string, kinds *RestrictedToKinds) string {
	return fmt.Sprintf("%s not one of %s", name, kinds)
}

func convertBoolToInt(b bool) int {
	var r = 0
	if b {
		r = 1
	}
	return r
}