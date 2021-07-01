package is

import (
	"fmt"
	"github.com/scarabsoft/go-hamcrest/internal"
	"github.com/scarabsoft/go-hamcrest/matcher"
)

func Empty() matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(matcher.FailIfIsNil("actual", actual)).
				Add(func() matcher.MatchResult {
					length, err := internal.Length(actual)
					if err != nil {
						return matcher.Failed(err.Error())
					}

					if length != 0 {
						return matcher.Failed(fmt.Sprintf("want to be empty; got length %d", length))
					}

					return matcher.Matched()
				})
		},
	)
}

func NotEmpty() matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(matcher.FailIfIsNil("actual", actual)).
				Add(func() matcher.MatchResult {
					length, err := internal.Length(actual)
					if err != nil {
						return matcher.Failed(err.Error())
					}

					if length == 0 {
						return matcher.Failed(fmt.Sprintf("want not to be empty; got empty"))
					}

					return matcher.Matched()
				})
		},
	)
}
