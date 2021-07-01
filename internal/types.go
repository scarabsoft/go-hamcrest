package internal

import (
	"reflect"
	"strings"
)

type RestrictedToKinds struct {
	allowedKinds []reflect.Kind
}

var RestrictedToPointer = NewRestrictedToKind(
	reflect.Ptr,
	reflect.UnsafePointer,
)

var RestrictedToBoolKind = NewRestrictedToKind(
	reflect.Bool,
)

var RestrictedToStringKind = NewRestrictedToKind(
	reflect.String,
)

var RestrictedToInteger = NewRestrictedToKind(
	reflect.Int,
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,

	reflect.Uint,
	reflect.Uint8,
	reflect.Uint16,
	reflect.Uint32,
	reflect.Uint64,
)

var RestrictedToFloat = NewRestrictedToKind(
	reflect.Float32,
	reflect.Float64,
)

var RestrictedToNumericKinds = merge(
	RestrictedToInteger,
	RestrictedToFloat,
)

var AllAreAllowed = NewRestrictedToKind()

var RestrictedToNumericStringBoolKinds = merge(
	RestrictedToBoolKind,
	RestrictedToStringKind,
	RestrictedToNumericKinds,
)

func merge(kinds ...*RestrictedToKinds) *RestrictedToKinds {
	ks := make([]reflect.Kind, 0)

	for _, kind := range kinds {
		for _, k := range kind.allowedKinds {
			ks = append(ks, k)
		}
	}

	return &RestrictedToKinds{allowedKinds: ks}
}

func NewRestrictedToKind(allowedKinds ...reflect.Kind) *RestrictedToKinds {
	ks := make([]reflect.Kind, len(allowedKinds))

	for _, kind := range allowedKinds {
		ks = append(ks, kind)
	}

	return &RestrictedToKinds{allowedKinds: allowedKinds}
}

func (s RestrictedToKinds) Allowed(value interface{}) bool {
	return s.AllowedKind(reflect.TypeOf(value).Kind())
}

func (s RestrictedToKinds) AllowedKind(kind reflect.Kind) bool {
	if len(s.allowedKinds) == 0 {
		return true
	}

	for _, k := range s.allowedKinds {
		if k == kind {
			return true
		}
	}
	return false
}

func (s RestrictedToKinds) String() string {
	strs := []string{}
	for _, kind := range s.allowedKinds {
		strs = append(strs, kind.String())
	}
	return "[" + strings.Join(strs, ",") + "]"
}
