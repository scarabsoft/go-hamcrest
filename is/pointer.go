package is

import (
	"fmt"
	"github.com/scarabsoft/go-hamcrest/internal"
	"github.com/scarabsoft/go-hamcrest/matcher"
	"reflect"
)

func PointingTo(expected interface{}) matcher.Matcher {
	return matcher.New(func(actual interface{}, chain matcher.Chain) matcher.Chain {
		return chain.
			Add(
				matcher.MatchIfAllAreNil(actual, expected),
				matcher.FailIfIsNil("actual", actual),
				matcher.FailIfIsNil("expected", expected),
				matcher.FailIfNotRestrictedType("actual", actual, internal.RestrictedToPointer),
				matcher.FailIfNotRestrictedType("expected", expected, internal.RestrictedToPointer),
				func() matcher.MatchResult {
					actualAddress := reflect.ValueOf(actual).Pointer()
					expectedAddress := reflect.ValueOf(expected).Pointer()

					if !internal.IsEqual(actualAddress, expectedAddress) {
						return matcher.Failed(fmt.Sprintf(
							"want pointer to point to %s; got %s",
							internal.FormatTypeWithValue(actual),
							internal.FormatTypeWithValue(expected),
						))
					}

					return matcher.Matched()
				})
	})
}

func NotPointingTo(expected interface{}) matcher.Matcher {
	return matcher.New(func(actual interface{}, chain matcher.Chain) matcher.Chain {
		return chain.
			Add(
				matcher.MatchIfAllAreNil(actual, expected),
				matcher.FailIfIsNil("actual", actual),
				matcher.FailIfIsNil("expected", expected),
				matcher.FailIfNotRestrictedType("actual", actual, internal.RestrictedToPointer),
				matcher.FailIfNotRestrictedType("expected", expected, internal.RestrictedToPointer),
				func() matcher.MatchResult {
					actualAddress := reflect.ValueOf(actual).Pointer()
					expectedAddress := reflect.ValueOf(expected).Pointer()

					if internal.IsEqual(actualAddress, expectedAddress) {
						return matcher.Failed(fmt.Sprintf(
							"want pointer to not point to %s; got %s",
							internal.FormatTypeWithValue(actual),
							internal.FormatTypeWithValue(expected),
						))
					}

					return matcher.Matched()
				})
	})
}