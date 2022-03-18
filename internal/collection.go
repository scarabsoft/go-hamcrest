package internal

import (
	"errors"
	"reflect"
)

func Length(actual interface{}) (int64, error) {
	actualValue := reflect.ValueOf(actual)
	switch actualValue.Kind() {
	case reflect.Array, reflect.Slice, reflect.Chan, reflect.Map, reflect.String:
		return int64(actualValue.Len()), nil
	case reflect.Ptr:
		elem := actualValue.Elem()
		kind := elem.Kind()
		switch kind {
		case reflect.Array, reflect.Slice, reflect.Map, reflect.String:
			return int64(elem.Len()), nil
		}
		return -1, errors.New(FormatUnsupportedVariable("actual pointer", NewRestrictedToKind(
			reflect.Array, reflect.Map, reflect.Slice, reflect.String, reflect.Ptr),
		))
	}
	return -1, errors.New(FormatUnsupportedVariable("actual", NewRestrictedToKind(
		reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String, reflect.Ptr),
	))
}

func HasItem(actual, expected interface{}) (bool, error) {
	if expected == nil {
		actualValue := reflect.ValueOf(actual)
		for i := 0; i < actualValue.Len(); i++ {
			switch actualValue.Index(i).Kind() {
			case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer,
				reflect.Interface, reflect.Slice:
				if actualValue.Index(i).IsNil() {
					return true, nil
				}
			}
		}
		return false, nil
	}

	length, err := Length(actual)
	if err != nil {
		panic("was not able to resolve length of actual")
	}
	expectedValue, actualValue := reflect.ValueOf(expected), reflect.ValueOf(actual)

	switch actualValue.Kind() {
	case reflect.Ptr:
		elem := actualValue.Elem()
		for i := 0; i < int(length); i++ {
			v := elem.Index(i)
			if IsEqual(v.Interface(), expectedValue.Interface()) {
				return true, nil
			}
		}
	case reflect.Chan:
		store := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(actual).Elem()), 0, 0)

		for i := 0; i < int(length); i++ {
			v, ok := actualValue.Recv()
			if !ok {
				panic("something went wrong while consuming from channel")
			}
			store = reflect.Append(store, v)
		}

		matched := false
		for i := 0; i < int(length); i++ {
			v := store.Index(i)
			actualValue.Send(v)
			if IsEqual(v.Interface(), expectedValue.Interface()) {
				matched = true
			}
		}
		return matched, nil

	default:
		for i := 0; i < int(length); i++ {
			if IsEqual(actualValue.Index(i).Interface(), expectedValue.Interface()) {
				return true, nil
			}
		}
	}
	return false, nil
}
