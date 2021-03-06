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
	Add(fns ...Function) Chain

	iter() []Function
}

type matcherChainImpl struct {
	filterResult *MatchResult
	functions    []Function
}

func (c *matcherChainImpl) Add(fns ...Function) Chain {
	c.functions = append(c.functions, fns...)
	return c
}

func (c matcherChainImpl) iter() []Function {
	return c.functions
}

func exec(c Chain) MatchResult {
	for _, fn := range c.iter() {
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

func MatchIfAllAreNil(actual, expected interface{}) Function {
	return func() MatchResult {
		if internal.IsNil(actual) && internal.IsNil(expected) {
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

func FailIfAnyIsNil(format string, actual, expected interface{}) Function {
	return func() MatchResult {
		if internal.IsNil(actual) || internal.IsNil(expected) {
			return Failed(
				fmt.Sprintf(
					format,
					internal.FormatTypeWithValue(actual),
					internal.FormatTypeWithValue(expected),
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

func FailIfNotSameType(actual, expected interface{}) Function {
	return func() MatchResult {
		if internal.IsNil(actual) && internal.IsNil(expected) {
			return Next()
		}

		if (internal.IsNil(actual) || internal.IsNil(expected)) ||
			reflect.ValueOf(actual).Type() != reflect.ValueOf(expected).Type() {
			return Failed("comparing different types; " + internal.FormatTypes("!=", actual, expected))
		}
		return Next()
	}
}

// LoopAndApply applies a matcherFn to every element of expected
func LoopAndApply(matcherFn func(elem interface{}) Matcher, expected ...interface{}) Matcher {
	return New(
		func(actual interface{}, chain Chain) Chain {
			for _, g := range expected {
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
