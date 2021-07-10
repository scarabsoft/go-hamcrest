package is

import (
	"fmt"
	"github.com/scarabsoft/go-hamcrest/internal"
	"github.com/scarabsoft/go-hamcrest/matcher"
	"reflect"
)

func EqualTo(given interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(matcher.MatchIfAllAreNil(actual, given)).
				Add(matcher.FailIfAnyIsNil("want value equal to %s; got %s", actual, given)).
				Add(matcher.FailIfNotSameType(actual, given)).
				Add(func() matcher.MatchResult {
					if !internal.IsEqual(actual, given) {
						return matcher.Failed(fmt.Sprintf(
							"want value equal to %s; got %s",
							internal.FormatTypeWithValue(actual),
							internal.FormatTypeWithValue(given),
						))
					}
					return matcher.Matched()
				})
		},
	)
}

func NotEqualTo(given interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(func() matcher.MatchResult {
					if internal.IsNil(actual) && internal.IsNil(given) {
						return matcher.Failed(fmt.Sprintf(
							"want value not equal to %s; got %s",
							internal.FormatTypeWithValue(actual),
							internal.FormatTypeWithValue(given),
						))
					}
					return matcher.Next()
				}).
				Add(func() matcher.MatchResult {
					if internal.IsNil(actual) || internal.IsNil(given) {
						return matcher.Matched()
					}
					return matcher.Next()
				}).
				Add(func() matcher.MatchResult {
					if reflect.ValueOf(actual).Type() != reflect.ValueOf(given).Type() {
						return matcher.Failed("comparing different types; " + internal.FormatTypes("!=", actual, given))
					}
					return matcher.Next()
				}).
				Add(func() matcher.MatchResult {
					if !internal.IsEqual(actual, given) {
						return matcher.Matched()
					}
					return matcher.Failed(fmt.Sprintf(
						"want value not equal to %s; got %s",
						internal.FormatTypeWithValue(actual),
						internal.FormatTypeWithValue(given),
					))
				})
		},
	)
}

func numericStringBoolKindsComparisonMatcher(given interface{}, fn func(actual, given interface{}) bool, format string) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(
					matcher.FailIfIsNil("actual", actual),
					matcher.FailIfIsNil("given", given),
					matcher.FailIfNotRestrictedType("actual", actual, internal.RestrictedToNumericStringBoolKinds),
					matcher.FailIfNotRestrictedType("given", given, internal.RestrictedToNumericStringBoolKinds),
					matcher.FailIfNotSameType(actual, given),
					func() matcher.MatchResult {
						if !fn(actual, given) {
							return matcher.Failed(fmt.Sprintf(
								format,
								internal.FormatTypeWithValue(given),
								internal.FormatTypeWithValue(actual),
							))
						}
						return matcher.Matched()
					})
		},
	)
}

func GreaterThan(given interface{}) matcher.Matcher {
	return numericStringBoolKindsComparisonMatcher(
		given, internal.IsGreaterThan, "want value greater than %s; got %s",
	)
}

func GreaterThanEqual(given interface{}) matcher.Matcher {
	return numericStringBoolKindsComparisonMatcher(
		given, internal.IsGreaterThanEqual, "want value greater than equal to %s; got %s",
	)
}

func LessThan(given interface{}) matcher.Matcher {
	return numericStringBoolKindsComparisonMatcher(
		given, internal.IsLessThan, "want value less than %s; got %s",
	)
}

func LessThanEqual(given interface{}) matcher.Matcher {
	return numericStringBoolKindsComparisonMatcher(
		given, internal.IsLessThanEqual, "want value less than equal to %s; got %s",
	)
}

func numericRangeComparisonMatcher(givenMin, givenMax interface{}, fn func(actual, givenMin, givenMax interface{}) bool, format string) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(
					matcher.FailIfIsNil("actual", actual),
					matcher.FailIfIsNil("givenMin", givenMin),
					matcher.FailIfIsNil("givenMax", givenMax),
					matcher.FailIfNotRestrictedType("actual", actual, internal.RestrictedToNumericKinds),
					matcher.FailIfNotRestrictedType("givenMin", givenMin, internal.RestrictedToNumericKinds),
					matcher.FailIfNotRestrictedType("givenMax", givenMax, internal.RestrictedToNumericKinds),
					matcher.FailIfNotSameType(actual, givenMin),
					matcher.FailIfNotSameType(actual, givenMax),
					func() matcher.MatchResult {
						if internal.IsLessThanEqual(givenMin, givenMax) {
							return matcher.Next()
						}
						return matcher.Failed(fmt.Sprintf("givenMin %s must be <= givenMax %s",
							internal.FormatTypeWithValue(givenMin),
							internal.FormatTypeWithValue(givenMax),
						))
					},
					func() matcher.MatchResult {
						if !fn(actual, givenMin, givenMax) {
							return matcher.Failed(fmt.Sprintf(
								format,
								internal.FormatTypeWithValue(givenMin),
								internal.FormatTypeWithValue(givenMax),
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

func Between(givenMin, givenMax interface{}) matcher.Matcher {
	return numericRangeComparisonMatcher(
		givenMin, givenMax,
		func(a, gMi, gMax interface{}) bool {
			return internal.IsGreaterThan(a, gMi) && internal.IsLessThan(a, gMax)
		},
		"want value between %s and %s; got %s",
	)
}

func BetweenOrEqual(givenMin, givenMax interface{}) matcher.Matcher {
	return numericRangeComparisonMatcher(
		givenMin, givenMax,
		func(a, gMi, gMax interface{}) bool {
			return betweenOrEqual(a, gMi, gMax)
		},
		"want value between or equal to %s and %s; got %s",
	)
}

func NotBetween(givenMin, givenMax interface{}) matcher.Matcher {
	return numericRangeComparisonMatcher(
		givenMin, givenMax,
		func(a, gMi, gMax interface{}) bool {
			return !(internal.IsGreaterThan(a, gMi) && internal.IsLessThan(a, gMax))
		},
		"want value not between %s and %s; got %s",
	)
}

func NotBetweenOrEqual(givenMin, givenMax interface{}) matcher.Matcher {
	return numericRangeComparisonMatcher(
		givenMin, givenMax,
		func(a, gMi, gMax interface{}) bool {
			return !(internal.IsGreaterThanEqual(a, gMi) && internal.IsLessThanEqual(a, gMax))
		},
		"want value not between or equal to %s and %s; got %s",
	)
}

func CloseTo(given, error interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(
					matcher.FailIfIsNil("actual", actual),
					matcher.FailIfIsNil("given", given),
					matcher.FailIfIsNil("error", error),
					matcher.FailIfNotRestrictedType("actual", actual, internal.RestrictedToFloat),
					matcher.FailIfNotRestrictedType("given", given, internal.RestrictedToFloat),
					matcher.FailIfNotRestrictedType("error", error, internal.RestrictedToFloat),
					matcher.FailIfNotSameType(actual, given),
					matcher.FailIfNotSameType(actual, error),
					func() matcher.MatchResult {
						givenValue := given.(float64)
						errorValue := error.(float64)
						if !betweenOrEqual(actual, givenValue-errorValue, givenValue+errorValue) {
							return matcher.Failed(
								fmt.Sprintf("want value close to %s with error %s; got %s",
									internal.FormatTypeWithValue(given),
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

func NotCloseTo(given, error interface{}) matcher.Matcher {
	return matcher.New(
		func(actual interface{}, chain matcher.Chain) matcher.Chain {
			return chain.
				Add(
					matcher.FailIfIsNil("actual", actual),
					matcher.FailIfIsNil("given", given),
					matcher.FailIfIsNil("error", error),
					matcher.FailIfNotRestrictedType("actual", actual, internal.RestrictedToFloat),
					matcher.FailIfNotRestrictedType("given", given, internal.RestrictedToFloat),
					matcher.FailIfNotRestrictedType("error", error, internal.RestrictedToFloat),
					matcher.FailIfNotSameType(actual, given),
					matcher.FailIfNotSameType(actual, error),
					func() matcher.MatchResult {
						givenValue := given.(float64)
						errorValue := error.(float64)
						if betweenOrEqual(actual, givenValue-errorValue, givenValue+errorValue) {
							return matcher.Failed(
								fmt.Sprintf("want value not close to %s with error %s; got %s",
									internal.FormatTypeWithValue(given),
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
