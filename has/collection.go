package has

import (
	"fmt"
	"github.com/scarabsoft/go-hamcrest/internal"
	"github.com/scarabsoft/go-hamcrest/matcher"
	"reflect"
)

func Length(length interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(
					matcher.FailIfIsNil("length", length),
					matcher.FailIfNotRestrictedType("length", length, internal.RestrictedToInteger),
					matcher.FailIfIsNil("actual", actual),
					matcher.FailIfNotRestrictedType("actual", actual, internal.NewRestrictedToKind(
						reflect.Array, reflect.Chan, reflect.Map, reflect.String, reflect.Slice, reflect.Ptr),
					),
					func() matcher.MatchResult {
						var expectedLength int64

						lenValue := reflect.ValueOf(length)
						switch lenValue.Kind() {
						case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
							expectedLength = lenValue.Int()
						case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
							expectedLength = int64(lenValue.Uint())
						}

						actualLength, err := internal.Length(actual)
						if err != nil {
							return matcher.Failed(err.Error())
						}

						if actualLength == expectedLength {
							return matcher.Matched()
						}

						return matcher.Failed(fmt.Sprintf("want length of %d; got %d", expectedLength, actualLength))
					})
		},
	)
}

func Item(expected interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(
					matcher.FailIfIsNil("actual", actual),
					func() matcher.MatchResult {
						result, err := internal.HasItem(actual, expected)
						if err != nil {
							return matcher.Failed(err.Error())
						}

						if !result {
							return matcher.Failed(fmt.Sprintf(
								"want %s to be part of %s; but was not",
								internal.FormatTypeWithValue(expected),
								internal.FormatTypeWithValue(actual),
							))
						}

						return matcher.Matched()
					})
		},
	)
}

func NotItem(expected interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(
					matcher.FailIfIsNil("actual", actual),
					func() matcher.MatchResult {
						result, err := internal.HasItem(actual, expected)
						if err != nil {
							return matcher.Failed(err.Error())
						}

						if result {
							return matcher.Failed(fmt.Sprintf(
								"want %s to not be part of %s; but was",
								internal.FormatTypeWithValue(expected),
								internal.FormatTypeWithValue(actual),
							))
						}
						return matcher.Matched()
					})
		},
	)
}

func Items(expected ...interface{}) matcher.Matcher {
	return matcher.LoopAndApply(Item, expected...)
}

func NotItems(expected ...interface{}) matcher.Matcher {
	return matcher.LoopAndApply(NotItem, expected...)
}

func Key(expected interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(
					matcher.FailIfIsNil("actual", actual),
					matcher.FailIfNotRestrictedType("actual", actual, internal.NewRestrictedToKind(reflect.Map)),
					func() matcher.MatchResult {
						actualValue := reflect.ValueOf(actual)
						expectedValue := reflect.ValueOf(expected)

						if reflect.TypeOf(actual).Key() != reflect.TypeOf(expected) {
							return matcher.Failed(fmt.Sprintf(
								"want %s to be key of %s; but was not",
								internal.FormatTypeWithValue(expected),
								internal.FormatTypeWithValue(actual),
							))
						}

						valueVal := actualValue.MapIndex(expectedValue)
						if valueVal.Kind().String() == "invalid" {
							return matcher.Failed(fmt.Sprintf(
								"want %s to be key of %s; but was not",
								internal.FormatTypeWithValue(expected),
								internal.FormatTypeWithValue(actual),
							))
						}
						return matcher.Matched()
					})
		},
	)
}

func NotKey(expected interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(
					matcher.FailIfIsNil("actual", actual),
					matcher.FailIfNotRestrictedType("actual", actual, internal.NewRestrictedToKind(reflect.Map)),
					func() matcher.MatchResult {
						actualValue := reflect.ValueOf(actual)
						for _, key := range actualValue.MapKeys() {
							if internal.IsEqual(key.Interface(), expected) {
								return matcher.Failed(fmt.Sprintf(
									"want %s not to be key of %s; but was",
									internal.FormatTypeWithValue(expected),
									internal.FormatTypeWithValue(actual),
								))
							}
						}
						return matcher.Matched()
					})
		},
	)
}

func Keys(expected ...interface{}) matcher.Matcher {
	return matcher.LoopAndApply(Key, expected...)
}

func NotKeys(expected ...interface{}) matcher.Matcher {
	return matcher.LoopAndApply(NotKey, expected...)
}
