package internal

import (
	"reflect"
)

func IsNil(actual interface{}) bool {
	if actual == nil {
		return true
	}
	switch reflect.TypeOf(actual).Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Slice:
		return reflect.ValueOf(actual).IsNil()
	default:
		return false
	}
}

