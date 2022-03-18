package has

import (
	"fmt"
	"github.com/scarabsoft/go-hamcrest/matcher"
	"reflect"
)

func SameTypeAs(expected interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(
					matcher.MatchIfAllAreNil(actual, expected),
					func() matcher.MatchResult {
						if reflect.TypeOf(actual) != reflect.TypeOf(expected) {
							return matcher.Failed(
								fmt.Sprintf("want same type as %[1]T; got %[2]T", actual, expected))
						}
						return matcher.Matched()
					})
		},
	)
}

func NotSameTypeAs(expected interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(func() matcher.MatchResult {
					if actual == nil && expected == nil {
						return matcher.Failed("want not same type as <nil>; got <nil>")
					}
					if reflect.TypeOf(actual) == reflect.TypeOf(expected) {
						return matcher.Failed(
							fmt.Sprintf("want not same type as %[1]T; got %[2]T", actual, expected))
					}
					return matcher.Matched()
				})
		},
	)
}
