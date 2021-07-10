package has

import (
	"fmt"
	"github.com/scarabsoft/go-hamcrest/internal"
	"github.com/scarabsoft/go-hamcrest/matcher"
	"strings"
)

func stringMatcher(actual interface{}, given string, fn func(actual, given string) bool, format string) func() matcher.MatchResult {
	return func() matcher.MatchResult {
		actualValue := actual.(string)
		if fn(actualValue, given) {
			return matcher.Failed(
				fmt.Sprintf(format, given, actual),
			)
		}
		return matcher.Matched()
	}
}

func Prefix(given string) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(matcher.FailIfIsNil("actual", actual)).
				Add(matcher.FailIfNotRestrictedType("actual", actual, internal.RestrictedToStringKind)).
				Add(stringMatcher(actual, given, func(a, g string) bool {
					return !strings.HasPrefix(a, g)
				}, "want %s to be prefix of %s; but was not"))
		},
	)
}

func NotPrefix(given string) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(matcher.FailIfIsNil("actual", actual)).
				Add(matcher.FailIfNotRestrictedType("actual", actual, internal.RestrictedToStringKind)).
				Add(stringMatcher(actual, given, func(a, g string) bool {
					return strings.HasPrefix(a, g)
				}, "want %s not to be prefix of %s; but was"))
		},
	)
}

func Suffix(given string) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(matcher.FailIfIsNil("actual", actual)).
				Add(matcher.FailIfNotRestrictedType("actual", actual, internal.RestrictedToStringKind)).
				Add(stringMatcher(actual, given, func(a, g string) bool {
					return !strings.HasSuffix(a, g)
				}, "want %s to be suffix of %s; but was not"))
		},
	)
}

func NotSuffix(given string) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(matcher.FailIfIsNil("actual", actual)).
				Add(matcher.FailIfNotRestrictedType("actual", actual, internal.RestrictedToStringKind)).
				Add(stringMatcher(actual, given, func(a, g string) bool {
					return strings.HasSuffix(a, g)
				}, "want %s not to be suffix of %s; but was"))
		},
	)
}
