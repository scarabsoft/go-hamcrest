package is

import (
	"fmt"
	"github.com/scarabsoft/go-hamcrest/internal"
	"github.com/scarabsoft/go-hamcrest/matcher"
	"reflect"
)

func EqualTo(expected interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(matcher.MatchIfAllAreNil(actual, expected)).
				Add(matcher.FailIfAnyIsNil("want value equal to %s; got %s", expected, actual)).
				Add(matcher.FailIfNotSameType(actual, expected)).
				Add(func() matcher.MatchResult {
					if !internal.IsEqual(actual, expected) {
						return matcher.Failed(fmt.Sprintf(
							"want value equal to %s; got %s",
							internal.FormatTypeWithValue(expected),
							internal.FormatTypeWithValue(actual),
						))
					}
					return matcher.Matched()
				})
		},
	)
}

func NotEqualTo(expected interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(func() matcher.MatchResult {
					if internal.IsNil(actual) && internal.IsNil(expected) {
						return matcher.Failed(fmt.Sprintf(
							"want value not equal to %s; got %s",
							internal.FormatTypeWithValue(expected),
							internal.FormatTypeWithValue(actual),
						))
					}
					return matcher.Next()
				}).
				Add(func() matcher.MatchResult {
					if internal.IsNil(actual) || internal.IsNil(expected) {
						return matcher.Matched()
					}
					return matcher.Next()
				}).
				Add(func() matcher.MatchResult {
					if reflect.ValueOf(actual).Type() != reflect.ValueOf(expected).Type() {
						return matcher.Failed("comparing different types; " + internal.FormatTypes("!=", actual, expected))
					}
					return matcher.Next()
				}).
				Add(func() matcher.MatchResult {
					if !internal.IsEqual(actual, expected) {
						return matcher.Matched()
					}
					return matcher.Failed(fmt.Sprintf(
						"want value not equal to %s; got %s",
						internal.FormatTypeWithValue(expected),
						internal.FormatTypeWithValue(actual),
					))
				})
		},
	)
}

func numericStringBoolKindsComparisonMatcher(expected interface{}, fn func(actual, expected interface{}) bool, format string) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(
					matcher.FailIfIsNil("actual", actual),
					matcher.FailIfIsNil("expected", expected),
					matcher.FailIfNotRestrictedType("actual", actual, internal.RestrictedToNumericStringBoolKinds),
					matcher.FailIfNotRestrictedType("expected", expected, internal.RestrictedToNumericStringBoolKinds),
					matcher.FailIfNotSameType(actual, expected),
					func() matcher.MatchResult {
						if !fn(actual, expected) {
							return matcher.Failed(fmt.Sprintf(
								format,
								internal.FormatTypeWithValue(expected),
								internal.FormatTypeWithValue(actual),
							))
						}
						return matcher.Matched()
					})
		},
	)
}

func GreaterThan(expected interface{}) matcher.Matcher {
	return numericStringBoolKindsComparisonMatcher(
		expected, internal.IsGreaterThan, "want value greater than %s; got %s",
	)
}

func GreaterThanEqual(expected interface{}) matcher.Matcher {
	return numericStringBoolKindsComparisonMatcher(
		expected, internal.IsGreaterThanEqual, "want value greater than equal to %s; got %s",
	)
}

func LessThan(expected interface{}) matcher.Matcher {
	return numericStringBoolKindsComparisonMatcher(
		expected, internal.IsLessThan, "want value less than %s; got %s",
	)
}

func LessThanEqual(expected interface{}) matcher.Matcher {
	return numericStringBoolKindsComparisonMatcher(
		expected, internal.IsLessThanEqual, "want value less than equal to %s; got %s",
	)
}

func numericRangeComparisonMatcher(expectedMin, expectedMax interface{}, fn func(actual, expectedMin, expectedMax interface{}) bool, format string) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(
					matcher.FailIfIsNil("actual", actual),
					matcher.FailIfIsNil("expectedMin", expectedMin),
					matcher.FailIfIsNil("expectedMax", expectedMax),
					matcher.FailIfNotRestrictedType("actual", actual, internal.RestrictedToNumericKinds),
					matcher.FailIfNotRestrictedType("expectedMin", expectedMin, internal.RestrictedToNumericKinds),
					matcher.FailIfNotRestrictedType("expectedMax", expectedMax, internal.RestrictedToNumericKinds),
					matcher.FailIfNotSameType(actual, expectedMin),
					matcher.FailIfNotSameType(actual, expectedMax),
					func() matcher.MatchResult {
						if internal.IsLessThanEqual(expectedMin, expectedMax) {
							return matcher.Next()
						}
						return matcher.Failed(fmt.Sprintf("expectedMin %s must be <= expectedMax %s",
							internal.FormatTypeWithValue(expectedMin),
							internal.FormatTypeWithValue(expectedMax),
						))
					},
					func() matcher.MatchResult {
						if !fn(actual, expectedMin, expectedMax) {
							return matcher.Failed(fmt.Sprintf(
								format,
								internal.FormatTypeWithValue(expectedMin),
								internal.FormatTypeWithValue(expectedMax),
								internal.FormatTypeWithValue(actual),
							))
						}
						return matcher.Matched()
					})
		},
	)
}

func betweenOrEqual(a, gMi, gMax interface{}) bool {
	return internal.IsGreaterThanEqual(a, gMi) && internal.IsLessThanEqual(a, gMax)

}

func Between(expectedMin, expectedMax interface{}) matcher.Matcher {
	return numericRangeComparisonMatcher(
		expectedMin, expectedMax,
		func(a, gMi, gMax interface{}) bool {
			return internal.IsGreaterThan(a, gMi) && internal.IsLessThan(a, gMax)
		},
		"want value between %s and %s; got %s",
	)
}

func BetweenOrEqual(expectedMin, expectedMax interface{}) matcher.Matcher {
	return numericRangeComparisonMatcher(
		expectedMin, expectedMax,
		func(a, gMi, gMax interface{}) bool {
			return betweenOrEqual(a, gMi, gMax)
		},
		"want value between or equal to %s and %s; got %s",
	)
}

func NotBetween(expectedMin, expectedMax interface{}) matcher.Matcher {
	return numericRangeComparisonMatcher(
		expectedMin, expectedMax,
		func(a, gMi, gMax interface{}) bool {
			return !(internal.IsGreaterThan(a, gMi) && internal.IsLessThan(a, gMax))
		},
		"want value not between %s and %s; got %s",
	)
}

func NotBetweenOrEqual(expectedMin, expectedMax interface{}) matcher.Matcher {
	return numericRangeComparisonMatcher(
		expectedMin, expectedMax,
		func(a, gMi, gMax interface{}) bool {
			return !(internal.IsGreaterThanEqual(a, gMi) && internal.IsLessThanEqual(a, gMax))
		},
		"want value not between or equal to %s and %s; got %s",
	)
}

func CloseTo(expected, error interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(
					matcher.FailIfIsNil("actual", actual),
					matcher.FailIfIsNil("expected", expected),
					matcher.FailIfIsNil("error", error),
					matcher.FailIfNotRestrictedType("actual", actual, internal.RestrictedToFloat),
					matcher.FailIfNotRestrictedType("expected", expected, internal.RestrictedToFloat),
					matcher.FailIfNotRestrictedType("error", error, internal.RestrictedToFloat),
					matcher.FailIfNotSameType(actual, expected),
					matcher.FailIfNotSameType(actual, error),
					func() matcher.MatchResult {
						expectedValue := expected.(float64)
						errorValue := error.(float64)
						if !betweenOrEqual(actual, expectedValue-errorValue, expectedValue+errorValue) {
							return matcher.Failed(
								fmt.Sprintf("want value close to %s with error %s; got %s",
									internal.FormatTypeWithValue(expected),
									internal.FormatTypeWithValue(error),
									internal.FormatTypeWithValue(actual),
								),
							)
						}
						return matcher.Matched()
					})
		},
	)
}

func NotCloseTo(expected, error interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(
					matcher.FailIfIsNil("actual", actual),
					matcher.FailIfIsNil("expected", expected),
					matcher.FailIfIsNil("error", error),
					matcher.FailIfNotRestrictedType("actual", actual, internal.RestrictedToFloat),
					matcher.FailIfNotRestrictedType("expected", expected, internal.RestrictedToFloat),
					matcher.FailIfNotRestrictedType("error", error, internal.RestrictedToFloat),
					matcher.FailIfNotSameType(actual, expected),
					matcher.FailIfNotSameType(actual, error),
					func() matcher.MatchResult {
						expectedValue := expected.(float64)
						errorValue := error.(float64)
						if betweenOrEqual(actual, expectedValue-errorValue, expectedValue+errorValue) {
							return matcher.Failed(
								fmt.Sprintf("want value not close to %s with error %s; got %s",
									internal.FormatTypeWithValue(expected),
									internal.FormatTypeWithValue(error),
									internal.FormatTypeWithValue(actual),
								),
							)
						}
						return matcher.Matched()
					})
		},
	)
}
