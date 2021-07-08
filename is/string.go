package is

import (
	"fmt"
	"github.com/scarabsoft/go-hamcrest/internal"
	"github.com/scarabsoft/go-hamcrest/matcher"
	"regexp"
)

func MatchingPattern(given interface{}) matcher.Matcher {
	return matchingCheck(given, func(a string, g string) bool {
		pattern := regexp.MustCompile(g)
		return pattern.MatchString(a)
	}, "want value to match %s; got %s")
}

func NotMatchingPattern(given interface{}) matcher.Matcher {
	return matchingCheck(given, func(a string, g string) bool {
		pattern := regexp.MustCompile(g)
		return !pattern.MatchString(a)
	}, "want value not to match %s; got %s")
}

func matchingCheck(given interface{}, fn func(string, string) bool, format string) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(matcher.FailIfIsNil("actual", actual)).
				Add(matcher.FailIfIsNil("given", given)).
				Add(matcher.FailIfRestrictedType("actual", actual, internal.RestrictedToStringKind)).
				Add(matcher.FailIfRestrictedType("given", given, internal.RestrictedToStringKind)).
				Add(func() matcher.MatchResult {
					actualString := actual.(string)
					givenString := given.(string)
					if !fn(actualString, givenString) {
						return matcher.Failed(
							fmt.Sprintf(format,
								givenString,
								actualString,
							),
						)
					}
					return matcher.Matched()
				})
		},
	)
}
