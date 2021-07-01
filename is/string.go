package is

import (
	"fmt"
	"github.com/scarabsoft/go-hamcrest/matcher"
	"regexp"
)

func MatchingPattern(pattern interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.Add(func() matcher.MatchResult {

				var a string
				switch tempActual := actual.(type) {
				case string:
					a = tempActual
				case fmt.Stringer:
					a = tempActual.String()
				default:
					return matcher.Failed("Unable to get string out of actual")
				}

				var p regexp.Regexp
				switch temPattern := pattern.(type) {
				case string:
					p = *regexp.MustCompile(temPattern)
				case regexp.Regexp:
					p = temPattern
				case *regexp.Regexp:
					p = *temPattern
				default:
					return matcher.Failed("Dont know how to handle pattern!")
				}

				if p.MatchString(a) {
					return matcher.Matched()
				}

				return matcher.Failed("did not match")
			})
		},
	)
}

func NotMatchingPattern(pattern interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			panic("Implement me")
		},
	)
}
