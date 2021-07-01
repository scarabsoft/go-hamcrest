package internal

import "reflect"

func IsTrue(actual interface{}) bool {
	//FIXME true case only for boolean?
	//FIXME maybe even panic?
	if reflect.TypeOf(actual).Kind() != reflect.Bool {
		return false
	}
	return actual == true
}

func IsFalse(actual interface{}) bool {
	return !IsTrue(actual)
}
