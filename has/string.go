package has

import (
	"fmt"
	"github.com/scarabsoft/go-hamcrest/internal"
	"github.com/scarabsoft/go-hamcrest/matcher"
	"strings"
)

func stringMatcher(actual interface{}, expected string, fn func(actual, expected string) bool, format string) func() matcher.MatchResult {
	return func() matcher.MatchResult {
		actualValue := actual.(string)
		if fn(actualValue, expected) {
			return matcher.Failed(
				fmt.Sprintf(format, expected, actual),
			)
		}
		return matcher.Matched()
	}
}

func Prefix(expected string) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(
					matcher.FailIfIsNil("actual", actual),
					matcher.FailIfNotRestrictedType("actual", actual, internal.RestrictedToStringKind),
					stringMatcher(actual, expected, func(a, g string) bool {
						return !strings.HasPrefix(a, g)
					}, "want %s to be prefix of %s; but was not"))
		},
	)
}

func NotPrefix(expected string) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(
					matcher.FailIfIsNil("actual", actual),
					matcher.FailIfNotRestrictedType("actual", actual, internal.RestrictedToStringKind),
					stringMatcher(actual, expected, func(a, g string) bool {
						return strings.HasPrefix(a, g)
					}, "want %s not to be prefix of %s; but was"))
		},
	)
}

func Suffix(expected string) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(
					matcher.FailIfIsNil("actual", actual),
					matcher.FailIfNotRestrictedType("actual", actual, internal.RestrictedToStringKind),
					stringMatcher(actual, expected, func(a, g string) bool {
						return !strings.HasSuffix(a, g)
					}, "want %s to be suffix of %s; but was not"))
		},
	)
}

func NotSuffix(expected string) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(
					matcher.FailIfIsNil("actual", actual),
					matcher.FailIfNotRestrictedType("actual", actual, internal.RestrictedToStringKind),
					stringMatcher(actual, expected, func(a, g string) bool {
						return strings.HasSuffix(a, g)
					}, "want %s not to be suffix of %s; but was"))
		},
	)
}
