package is

import (
	"github.com/scarabsoft/go-hamcrest/internal"
	"github.com/scarabsoft/go-hamcrest/matcher"
)

func True() matcher.Matcher {
	return matcher.New(func(actual interface{}, chain matcher.Chain) matcher.Chain {
		return chain.Add(func() matcher.MatchResult {
			if internal.IsTrue(actual) {
				return matcher.Matched()
			}
			return matcher.Failed("want true; got false")
		})
	})
}

func False() matcher.Matcher {
	return matcher.New(func(actual interface{}, chain matcher.Chain) matcher.Chain {
		return chain.Add(func() matcher.MatchResult {
			if internal.IsFalse(actual) {
				return matcher.Matched()
			}
			return matcher.Failed("want false; got true")
		})
	})
}

func Ok() matcher.Matcher {
	return True()
}

func NotOk() matcher.Matcher {
	return False()
}
