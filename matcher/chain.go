package matcher

import (
	"fmt"
	"github.com/scarabsoft/go-hamcrest/internal"
	"reflect"
)

type Function func() MatchResult

type MatchResult struct {
	Matched      bool
	Cause        string
	continueNext bool
}

func Matched() MatchResult {
	return MatchResult{
		Matched:      true,
		Cause:        "",
		continueNext: false,
	}
}

func Failed(cause string) MatchResult {
	return MatchResult{
		Matched:      false,
		Cause:        cause,
		continueNext: false,
	}
}

func Next() MatchResult {
	return MatchResult{
		Matched:      true,
		Cause:        "",
		continueNext: true,
	}
}

func (c *MatchResult) String() string {
	return fmt.Sprintf("%t %s", c.Matched, c.Cause)
}

type Chain interface {
	Add(fn Function) Chain

	Exec() MatchResult
}

type matcherChainImpl struct {
	filterResult *MatchResult
	functions    []Function
}

func (c *matcherChainImpl) Add(fn Function) Chain {
	c.functions = append(c.functions, fn)
	return c
}

func (c *matcherChainImpl) Exec() MatchResult {
	for _, fn := range c.functions {
		filterResult := fn()
		if !filterResult.continueNext {
			return filterResult
		}
	}
	return Matched()
}

func newMatcherChain() Chain {
	return &matcherChainImpl{}
}

func MatchIfAllAreNil(actual, given interface{}) Function {
	return func() MatchResult {
		if internal.IsNil(actual) && internal.IsNil(given) {
			return Matched()
		}
		return Next()
	}
}

func FailIfIsNil(name string, value interface{}) Function {
	return func() MatchResult {
		if internal.IsNil(value) {
			return Failed(fmt.Sprintf("%s is <nil>", name))
		}
		return Next()
	}
}

func FailIfAnyIsNil(format string, actual, given interface{}) Function {
	return func() MatchResult {
		if internal.IsNil(actual) || internal.IsNil(given) {
			return Failed(
				fmt.Sprintf(
					format,
					internal.FormatTypeWithValue(actual),
					internal.FormatTypeWithValue(given),
				),
			)
		}
		return Next()
	}
}

func FailIfNotRestrictedType(name string, value interface{}, kinds *internal.RestrictedToKinds) Function {
	return func() MatchResult {
		if !kinds.Allowed(value) {
			return Failed(internal.FormatUnsupportedVariable(name, kinds))
		}
		return Next()
	}
}

func FailIfNotSameType(actual, given interface{}) Function {
	return func() MatchResult {
		if internal.IsNil(actual) && internal.IsNil(given) {
			return Next()
		}

		if (internal.IsNil(actual) || internal.IsNil(given)) ||
			reflect.ValueOf(actual).Type() != reflect.ValueOf(given).Type() {
			return Failed("comparing different types; " + internal.FormatTypes("!=", actual, given))
		}
		return Next()
	}
}

// applies a matcherFn to every element of given
func LoopAndApply(matcherFn func(elem interface{}) Matcher, given ...interface{}) Matcher {
	return New(
		func(actual interface{}, chain Chain) Chain {
			for _, g := range given {
				chain.Add(func() MatchResult {
					m := matcherFn(g)
					if m.Matches(actual) {
						return Next()
					}
					return Failed(m.Cause())
				})
			}
			// if none of the previous matcher failed - it was successful
			return chain.Add(func() MatchResult {
				return Matched()
			})
		},
	)
}
