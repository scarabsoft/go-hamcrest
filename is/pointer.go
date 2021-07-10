package is

import (
	"fmt"
	"github.com/scarabsoft/go-hamcrest/internal"
	"github.com/scarabsoft/go-hamcrest/matcher"
	"reflect"
)

func PointingTo(given interface{}) matcher.Matcher {
	return matcher.New(func(actual interface{}, chain matcher.Chain) matcher.Chain {
		return chain.
			Add(
				matcher.MatchIfAllAreNil(actual, given),
				matcher.FailIfIsNil("actual", actual),
				matcher.FailIfIsNil("given", given),
				matcher.FailIfNotRestrictedType("actual", actual, internal.RestrictedToPointer),
				matcher.FailIfNotRestrictedType("given", given, internal.RestrictedToPointer),
				func() matcher.MatchResult {
					actualAddress := reflect.ValueOf(actual).Pointer()
					givenAddress := reflect.ValueOf(given).Pointer()

					if !internal.IsEqual(actualAddress, givenAddress) {
						return matcher.Failed(fmt.Sprintf(
							"want pointer to point to %s; got %s",
							internal.FormatTypeWithValue(actual),
							internal.FormatTypeWithValue(given),
						))
					}

					return matcher.Matched()
				})
	})
}
