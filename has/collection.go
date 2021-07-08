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
				Add(matcher.FailIfIsNil("length", length)).
				Add(matcher.FailIfRestrictedType("length", length, internal.RestrictedToInteger)).
				Add(matcher.FailIfIsNil("actual", actual)).
				Add(matcher.FailIfRestrictedType("actual", actual, internal.NewRestrictedToKind(
					reflect.Array, reflect.Chan, reflect.Map, reflect.String, reflect.Ptr)),
				).
				Add(func() matcher.MatchResult {
					var givenLength int64

					lenValue := reflect.ValueOf(length)
					switch lenValue.Kind() {
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						givenLength = lenValue.Int()
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
						givenLength = int64(lenValue.Uint())
					}

					actualLength, err := internal.Length(actual)
					if err != nil {
						return matcher.Failed(err.Error())
					}

					if actualLength == givenLength {
						return matcher.Matched()
					}

					return matcher.Failed(fmt.Sprintf("want length of %d; got %d", givenLength, actualLength))
				})
		},
	)
}

func Item(given interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(matcher.FailIfIsNil("actual", actual)).
				Add(func() matcher.MatchResult {
					result, err := internal.HasItem(actual, given)
					if err != nil {
						return matcher.Failed(err.Error())
					}

					if !result {
						return matcher.Failed(fmt.Sprintf(
							"want %s to be part of %s; but was not",
							internal.FormatTypeWithValue(given),
							internal.FormatTypeWithValue(actual),
						))
					}

					return matcher.Matched()
				})
		},
	)
}

func NotItem(given interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(matcher.FailIfIsNil("actual", actual)).
				Add(func() matcher.MatchResult {
					result, err := internal.HasItem(actual, given)
					if err != nil {
						return matcher.Failed(err.Error())
					}

					if result {
						return matcher.Failed(fmt.Sprintf(
							"want %s to not be part of %s; but was",
							internal.FormatTypeWithValue(given),
							internal.FormatTypeWithValue(actual),
						))
					}
					return matcher.Matched()
				})
		},
	)
}

func Items(given ...interface{}) matcher.Matcher {
	return matcher.LoopAndApply(Item, given...)
}

func NotItems(given ...interface{}) matcher.Matcher {
	return matcher.LoopAndApply(NotItem, given...)
}

func Key(given interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(matcher.FailIfIsNil("actual", actual)).
				Add(matcher.FailIfRestrictedType("actual", actual, internal.NewRestrictedToKind(reflect.Map))).
				Add(func() matcher.MatchResult {
					actualValue := reflect.ValueOf(actual)
					givenValue := reflect.ValueOf(given)

					if reflect.TypeOf(actual).Key() != reflect.TypeOf(given) {
						return matcher.Failed(fmt.Sprintf(
							"want %s to be key of %s; but was not",
							internal.FormatTypeWithValue(given),
							internal.FormatTypeWithValue(actual),
						))
					}

					valueVal := actualValue.MapIndex(givenValue)
					if valueVal.Kind().String() == "invalid" {
						return matcher.Failed(fmt.Sprintf(
							"want %s to be key of %s; but was not",
							internal.FormatTypeWithValue(given),
							internal.FormatTypeWithValue(actual),
						))
					}
					return matcher.Matched()
				})
		},
	)
}

func NotKey(given interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(matcher.FailIfIsNil("actual", actual)).
				Add(matcher.FailIfRestrictedType("actual", actual, internal.NewRestrictedToKind(reflect.Map))).
				Add(func() matcher.MatchResult {
					actualValue := reflect.ValueOf(actual)
					for _, key := range actualValue.MapKeys() {
						if internal.IsEqual(key.Interface(), given) {
							return matcher.Failed(fmt.Sprintf(
								"want %s not to be key of %s; but was",
								internal.FormatTypeWithValue(given),
								internal.FormatTypeWithValue(actual),
							))
						}
					}
					return matcher.Matched()
				})
		},
	)
}

func Keys(given ...interface{}) matcher.Matcher {
	return matcher.LoopAndApply(Key, given...)
}

func NotKeys(given ...interface{}) matcher.Matcher {
	return matcher.LoopAndApply(NotKey, given...)
}
