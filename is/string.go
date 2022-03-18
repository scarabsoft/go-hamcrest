package is

import (
	"fmt"
	"github.com/scarabsoft/go-hamcrest/internal"
	"github.com/scarabsoft/go-hamcrest/matcher"
	"regexp"
)

func MatchingPattern(expected interface{}) matcher.Matcher {
	return matchingCheck(expected, func(a string, g string) bool {
		pattern := regexp.MustCompile(g)
		return pattern.MatchString(a)
	}, "want value to match %s; got %s")
}

func NotMatchingPattern(expected interface{}) matcher.Matcher {
	return matchingCheck(expected, func(a string, g string) bool {
		pattern := regexp.MustCompile(g)
		return !pattern.MatchString(a)
	}, "want value not to match %s; got %s")
}

func matchingCheck(expected interface{}, fn func(string, string) bool, format string) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(
					matcher.FailIfIsNil("actual", actual),
					matcher.FailIfIsNil("expected", expected),
					matcher.FailIfNotRestrictedType("actual", actual, internal.RestrictedToStringKind),
					matcher.FailIfNotRestrictedType("expected", expected, internal.RestrictedToStringKind),
					func() matcher.MatchResult {
						actualString := actual.(string)
						expectedString := expected.(string)
						if !fn(actualString, expectedString) {
							return matcher.Failed(
								fmt.Sprintf(format,
									expectedString,
									actualString,
								),
							)
						}
						return matcher.Matched()
					})
		},
	)
}
