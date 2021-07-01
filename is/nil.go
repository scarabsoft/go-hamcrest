package is

import (
	"fmt"
	"github.com/scarabsoft/go-hamcrest/internal"
	"github.com/scarabsoft/go-hamcrest/matcher"
)

func Nil() matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.Add(func() matcher.MatchResult {
				if !internal.IsNil(actual) {
					return matcher.Failed(fmt.Sprintf("want <nil>; got=%s", internal.FormatTypeWithValue(actual)))
				}
				return matcher.Matched()
			})
		})
}

func NotNil() matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.Add(func() matcher.MatchResult {
				if internal.IsNil(actual) {
					return matcher.Failed(fmt.Sprintf("want not <nil>; got=<nil>"))
				}
				return matcher.Matched()
			})
		})
}